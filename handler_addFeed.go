package main

import (
	"blog-aggregator-go/internal/database"
	"blog-aggregator-go/internal/utils"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(state *State, cmd Command) error {

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <web name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	if utils.ValidateUrl(name) {
		return fmt.Errorf("error: url must be 2nd parameter\nfe: addFeed <webname> <url>")
	}

	if !utils.ValidateUrl(url) {
		return fmt.Errorf("error: invalid url")
	}

	currentUser := state.cfg.Current_user_name

	user, err := state.db.GetUser(context.Background(), currentUser)

	if err != nil {
		return fmt.Errorf("error: error al obtener el usuario actual")
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	f, err := state.db.CreateFeed(context.Background(), feed)

	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = state.db.CreateFeedFollow(context.Background(), feedFollow)

	if err != nil {
		return fmt.Errorf("error registering feed follow for current user: %v", err)
	}

	printFeed(f)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
