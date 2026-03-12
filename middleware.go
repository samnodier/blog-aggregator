package main

import (
	"context"
	"fmt"

	"github.com/samnodier/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()
		loggedInUser := s.cfg.CurrentUserName
		currentUser, err := s.db.GetUser(ctx, loggedInUser)
		if err != nil {
			return fmt.Errorf("error fetching user from database: %w", err)
		}
		return handler(s, cmd, currentUser)
	}
}
