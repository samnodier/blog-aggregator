package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	ctx := context.Background()
	if len(cmd.args) > 0 {
		return errors.New("the login handler expects no arguments")
	}
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("error listing all users: %w", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
