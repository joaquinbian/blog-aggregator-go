package main

import "fmt"

func HandlerLogin(state *state, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Not enough arguments! Username expected!\n")
	}

	if err := state.cfg.SetUser(cmd.Args[0]); err != nil {
		return fmt.Errorf("There was an error setting the username: %w\n")

	}

	fmt.Print("Username set correctly!\n")
	return nil
}
