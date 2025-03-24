package main

import "fmt"

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	commands map[string]func(*state, Command) error
}

func (c *Commands) Register(name string, f func(*state, Command) error) {
	c.commands[name] = f
}

func (c *Commands) Run(s *state, cmd Command) error {

	command, ok := c.commands[cmd.Name]

	if !ok {
		return fmt.Errorf("The %s command does not exists!", cmd.Name)
	}

	return command(s, cmd)
}
