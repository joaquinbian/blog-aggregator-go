package main

import (
	"context"
	"fmt"

	"github.com/joaquinbian/blog-aggregator-go/internal/database"
)

func loginMiddleware(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(state *State, cmd Command) error {
		currentUser := state.cfg.Current_user_name

		user, err := state.db.GetUser(context.Background(), currentUser)

		if err != nil {
			return fmt.Errorf("error getting current user: %v", err)
		}

		return handler(state, cmd, user)
	}
}
