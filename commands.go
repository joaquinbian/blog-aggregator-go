package main

import (
	"fmt"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	commands map[string]func(*State, Command) error
}

func (c *Commands) Run(state *State, cmd Command) error {
	command, ok := c.commands[cmd.Name]
	if !ok {
		return fmt.Errorf("error: command %v does not exists", cmd.Name)
	}

	return command(state, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) error {
	_, ok := c.commands[name]

	if ok {
		return fmt.Errorf("error: command %v already exists", name)
	}

	c.commands[name] = f
	return nil
}

func HandlerLogin(state *State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("error: no args provided for the %v command", cmd.Name)
	}
	name := cmd.Args[0]
	err := state.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("username changed to %v\n", name)
	return nil
}
