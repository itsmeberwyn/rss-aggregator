package helper

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	V1Model "github.com/itsmeberwyn/rss-service/pkg/http/model"
	"github.com/itsmeberwyn/rss-service/pkg/http/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

func StartScraping(conn *pgxpool.Pool, concurrency int, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		feeds, err := repository.FetchFeed(conn, context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feeds ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrape(conn, wg, feed)
		}
		wg.Wait()
	}
}

func scrape(conn *pgxpool.Pool, wg *sync.WaitGroup, feed V1Model.FeedModel) {
	defer wg.Done()

	err := repository.UpdateLastFetchedFeed(conn, context.Background(), feed.Id.String())
	if err != nil {
		log.Println("Error updating feed: ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed: ", err)
		return
	}

	for _, item := range rssFeed.Entry {
		published_at, err := time.Parse(time.RFC1123Z, item.Published)
		if err != nil {
			published_at = time.Now()
		}

		post := V1Model.PostModel{
			Title:        item.Title,
			Published_at: published_at,
			Url:          item.Link.Key,
			Feed_id:      feed.Id,
		}
		post, err = repository.CreatePost(conn, context.Background(), post)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				continue
			}
			log.Println("Error creating post: ", err)
		}
	}
}

func urlToFeed(url string) (V1Model.RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return V1Model.RSSFeed{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return V1Model.RSSFeed{}, err
	}

	rssFeed := V1Model.RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return V1Model.RSSFeed{}, err
	}
	return rssFeed, nil
}
