package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joaquinbian/blog-aggregator-go/internal/database"
)

func handlerFeedFollow(state *State, cmd Command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>\n", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := state.db.GetFeedBYUrl(context.Background(), url)

	if err != nil {
		return fmt.Errorf("error getting feed: %v\n", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	feedFollow, err := state.db.CreateFeedFollow(context.Background(), params)

	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	fmt.Println("feed follow created!")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func printFeedFollow(userName string, feedName string) {
	fmt.Printf("* Name:          %s\n", userName)
	fmt.Printf("* Feed Name:           %s\n", feedName)
}
