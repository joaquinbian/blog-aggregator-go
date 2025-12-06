package main

import (
	"context"
	"fmt"
)

func handleReset(state *State, cmd Command) error {

	err := state.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error eliminando usuarios: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
