package main

import (
	"blog-aggregator-go/internal/database"
	"context"
	"fmt"
)

func handlerFollowing(state *State, cmd Command) error {

	feeds, err := state.db.GetFeedFollowsForUser(context.Background(), state.cfg.Current_user_name)

	if err != nil {
		return fmt.Errorf("error getting feed follows for user: %v\n", err)
	}

	printFeedFollows(feeds)
	return nil
}

func printFeedFollows(feeds []database.GetFeedFollowsForUserRow) {

	for _, f := range feeds {
		fmt.Print("--------------------------------------------------\n\n")
		fmt.Printf("* ID:            %s\n", f.ID)
		fmt.Printf("* Created:       %v\n", f.CreatedAt)
		fmt.Printf("* Updated:       %v\n", f.UpdatedAt)
		fmt.Printf("* User Name:          %s\n", f.UserName)
		fmt.Printf("* Feed Name:           %s\n", f.FeedName)
	}
}
