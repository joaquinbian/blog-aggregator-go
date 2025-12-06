package main

import (
	"blog-aggregator-go/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFeedFollow(state *State, cmd Command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := state.db.GetFeedBYUrl(context.Background(), url)

	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	user, err := state.db.GetUser(context.Background(), state.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("error getting current user: %v", err)
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

	printFeedFollow(feedFollow)
	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("* ID:            %s\n", feedFollow.ID)
	fmt.Printf("* Created:       %v\n", feedFollow.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feedFollow.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feedFollow.UserName)
	fmt.Printf("* Feed Name:           %s\n", feedFollow.FeedName)
}
