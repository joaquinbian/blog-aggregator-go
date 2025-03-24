package commands

import "fmt"

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {

	command, ok := c.Commands[cmd.Name]

	if !ok {
		return fmt.Errorf("The %s command does not exists!", cmd.Name)
	}

	err := command(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
