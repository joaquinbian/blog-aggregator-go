package main

import (
	"context"
	"fmt"
)

func handlerFetchFeed(state *State, cmd Command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return fmt.Errorf("error al obtener feed: %v", err)
	}

	fmt.Printf("%v\n", feed)

	return nil
}
