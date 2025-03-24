package main

import (
	"blog-agregator-go/internal/commands"
	"blog-agregator-go/internal/config"
	"fmt"
	"os"
)

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

	state := commands.State{
		Cfg: &cfg,
	}

	cmds := commands.Commands{
		Commands: make(map[string]func(*commands.State, commands.Command) error),
	}

	cmds.Register("login", commands.HandlerLogin)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print("No enough arguments! Command expected!\n")
		os.Exit(1)
	}

	cmdName := args[0]
	if err := cmds.Run(&state, commands.Command{Name: cmdName, Args: args[1:]}); err != nil {
		fmt.Printf("there was an error running the %s command:\n%s\n", cmdName, err)
		os.Exit(1)
	}

}
