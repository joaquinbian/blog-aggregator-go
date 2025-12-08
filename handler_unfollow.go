package main

import (
	"context"
	"fmt"

	"github.com/joaquinbian/blog-aggregator-go/internal/database"
)

func handlerUnfollow(state *State, cmd Command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <feed url>\n", cmd.Name)
	}

	feedUrl := cmd.Args[0]

	feed, err := state.db.GetFeedBYUrl(context.Background(), feedUrl)

	if err != nil {
		return fmt.Errorf("error getting feed: %v\n", err)
	}

	_, err = state.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})

	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v\n", err)
	}

	fmt.Println("unfollowed feed:")
	printFeed(feed)
	return nil
}
