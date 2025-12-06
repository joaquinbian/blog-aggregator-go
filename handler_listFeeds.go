package main

import (
	"blog-aggregator-go/internal/database"
	"context"
	"fmt"
)

func handlerListFeeds(state *State, cmd Command) error {

	feeds, err := state.db.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	printFeeds(feeds)
	return nil
}

func printFeeds(feeds []database.GetFeedsRow) {

	for _, f := range feeds {
		fmt.Print("--------------------------------------------------\n\n")
		fmt.Printf("* ID:            %s\n", f.ID)
		fmt.Printf("* Created:       %v\n", f.CreatedAt)
		fmt.Printf("* Updated:       %v\n", f.UpdatedAt)
		fmt.Printf("* Name:          %s\n", f.Name)
		fmt.Printf("* URL:           %s\n", f.Url)
		fmt.Printf("* UserName:        %s\n", f.Name_2)
	}
}
