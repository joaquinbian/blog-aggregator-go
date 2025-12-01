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
