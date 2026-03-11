package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/samnodier/blog-aggregator/internal/database"
	"time"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the register handler expects a single argument, the name")
	}
	ctx := context.Background()
	username := cmd.args[0]
	userId := uuid.New()
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()
	insertedUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        userId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("error writing to database: %w", err)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error writing user: %w", err)
	}
	fmt.Printf("user %s, created successfully!\n", username)
	fmt.Println(insertedUser)
	return nil
}
