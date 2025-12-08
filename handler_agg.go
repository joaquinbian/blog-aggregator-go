package main

import (
	"blog-aggregator-go/internal/database"
	"blog-aggregator-go/internal/utils"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFetchFeed(state *State, cmd Command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <1s, 1m, 1h..>", cmd.Name)
	}

	time_between_reqs := cmd.Args[0]

	duration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("error parsing duration: %v\n", err)
	}

	fmt.Printf("collecting feeds every %s!\n", time_between_reqs)

	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		scrapeFeeds(state)
	}

}

func scrapeFeeds(state *State) error {
	nextFeed, err := state.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return fmt.Errorf("error getting feed: %v\n", err)
	}

	_, err = state.db.MarkFeedFetched(context.Background(), nextFeed.ID)

	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %v\n", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)

	if err != nil {
		return fmt.Errorf("error fetching feed: %v\n", err)
	}

	err = savePostToDB(state.db, feed, nextFeed.ID)

	fmt.Printf("error de saveposts %v\n", err)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	fmt.Println()
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Printf("fetching posts from %s...\n", feed.Channel.Title)
	fmt.Printf("found %d post for %s feed!\n", len(feed.Channel.Item), feed.Channel.Title)
	return nil
}

func savePostToDB(db *database.Queries, feed *RSSFeed, idFeed uuid.UUID) error {
	for _, item := range feed.Channel.Item {
		pubDate, err := utils.ParsePubDate(item.PubDate)
		if err != nil {
			pubDate = time.Now()
		}

		postData := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: pubDate,
			FeedID:      idFeed,
		}

		_, err = db.CreatePost(context.Background(), postData)

		if err != nil {
			fmt.Printf("error inserting post %v: %v\n", item.Title, err)
			continue
		}
	}
	return nil

}
