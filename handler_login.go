package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	ctx := context.Background()
	if len(cmd.args) == 0 || len(cmd.args) > 1 {
		return errors.New("the login handler expects a single argument, the username")
	}
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error writing user: %w", err)
	}
	registeredUser, err := s.db.GetUser(ctx, username)
	if err != nil {
		return fmt.Errorf("error reading database: %w", err)
	}
	fmt.Printf("user %s was set successfully!\n", username)
	fmt.Println(registeredUser)
	return nil
}
