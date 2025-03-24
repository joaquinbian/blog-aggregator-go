package main

import (
	"blog-agregator-go/internal/config"
	"fmt"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Unexpected error:", r)
		}
	}()

	cfg, err := config.Read()

	if err != nil {
		fmt.Printf("there was an error: %w", err)
	}

	appState := &state{
		cfg: &cfg,
	}

	cmds := Commands{
		commands: make(map[string]func(*state, Command) error),
	}

	cmds.Register("login", HandlerLogin)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print("No enough arguments! Command expected!\n")
		os.Exit(1)
	}

	cmdName := args[0]
	if err := cmds.Run(appState, Command{Name: cmdName, Args: args[1:]}); err != nil {
		fmt.Printf("there was an error running the %s command:\n%s\n", cmdName, err)
		os.Exit(1)
	}

}
