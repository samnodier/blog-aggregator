package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Error writing user: %w", err)
	}
	fmt.Printf("User %s was set successfully!\n", username)
	return nil
}
