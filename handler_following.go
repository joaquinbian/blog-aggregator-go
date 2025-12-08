package main

import (
	"context"
	"fmt"

	"github.com/joaquinbian/blog-aggregator-go/internal/database"
)

func handlerFollowing(state *State, cmd Command, user database.User) error {

	feeds, err := state.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("error getting feed follows for user: %v\n", err)
	}

	fmt.Printf("Feed followed by user %s:\n", state.cfg.Current_user_name)
	for _, f := range feeds {
		fmt.Printf("* Feed Name:           %s\n", f.FeedName)
	}
	return nil
}
