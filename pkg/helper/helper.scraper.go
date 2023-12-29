package helper

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
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
	rssFeed, err := urlToFeed(feed.Url)
	fmt.Println(feed.Url)
	if err != nil {
		log.Println("Error fetching feed: ", err)
		return
	}

	// update last fetched field in feed

	// sanitize the feed and save to post table

	fmt.Println(rssFeed.Entry[0].Title)
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
