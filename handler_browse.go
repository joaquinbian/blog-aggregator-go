package main

import (
	"blog-aggregator-go/internal/database"
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(state *State, cmd Command, user database.User) error {
	var limit int = 2
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <limit>\n", cmd.Name)
	} else if len(cmd.Args) == 1 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("usage: %s <number>\n", cmd.Name)
		}
		limit = n
	}
	params := database.GetPostsForUserParams{
		Limit:  int32(limit),
		UserID: user.ID,
	}
	posts, err := state.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error fetching posts for user: %v\n", err)
	}

	printPosts(posts)
	return nil
}

func printPosts(posts []database.GetPostsForUserRow) {
	for _, item := range posts {
		fmt.Println("---------------------------------------")
		fmt.Printf("* Feed: %s\n", item.FeedName)
		fmt.Printf("* Post title: %s\n", item.Title)
		fmt.Printf("* Description: %s\n", item.Description.String)
		fmt.Printf("* Link: %s\n", item.Url)
		fmt.Println()
	}
}
