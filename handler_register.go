package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samnodier/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 || len(cmd.args) > 1 {
		return errors.New("the register handler expects a single argument, the name")
	}
	ctx := context.Background()
	userId := uuid.New()
	username := strings.ToLower(cmd.args[0])
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()
	insertedUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        userId,
		Name:      username,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		return fmt.Errorf("error writing to database: %w", err)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error setting current user: %w", err)
	}
	fmt.Printf("user %s, created successfully!\n", insertedUser.Name)
	return nil
}
