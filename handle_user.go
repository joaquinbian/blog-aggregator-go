package main

import "fmt"

func handlerLogin(state *State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("error: no args provided for the %v command", cmd.Name)
	} else if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	err := state.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("username changed to %v\n", name)
	return nil
}
