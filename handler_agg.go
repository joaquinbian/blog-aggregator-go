package main

import (
	"context"
	"fmt"
	"time"
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

	printFeedItems(feed)
	return nil
}

func printFeedItems(feed *RSSFeed) {

	fmt.Printf("post from %s:\n", feed.Channel.Title)
	fmt.Println("-------------------------------------------------")
	for _, item := range feed.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}

	fmt.Println()
	fmt.Printf("found %d post for %s feed!\n", len(feed.Channel.Item), feed.Channel.Title)
	fmt.Println("-------------------------------------------------")
	fmt.Println()

}
