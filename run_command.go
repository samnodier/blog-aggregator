package main

import (
	"fmt"
)

func (c *commands) run(s *state, cmd command) error {
	command, exists := c.handlers[cmd.name]
	if !exists {
		return fmt.Errorf("command not found: %s", cmd.name)
	}
	err := command(s, cmd)
	if err != nil {
		return fmt.Errorf("error running command: %w", err)
	}
	return nil
}
