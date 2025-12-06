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

	if len(cmd.Args) < 2 {
		return fmt.Errorf("error: expected 2 arguments\nfe: addFeed <web name> <url>")
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
	state.db.CreateFeed(context.Background(), feed)

	return nil
}
