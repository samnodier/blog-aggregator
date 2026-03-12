package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samnodier/blog-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	ctx := context.Background()
	currentUserName := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(ctx, currentUserName)
	if err != nil {
		return fmt.Errorf("error fetching user from database: %w", err)
	}
	if len(cmd.args) != 2 {
		return errors.New("the register handler expects a two arguments, name and url")
	}
	feedId := uuid.New()
	feedName := cmd.args[0]
	feedURL := cmd.args[1]
	userId := currentUser.ID
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()
	insertedFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        feedId,
		Name:      feedName,
		Url:       feedURL,
		UserID:    userId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		return fmt.Errorf("error writing to feeds database: %w", err)
	}
	fmt.Println(insertedFeed)
	return nil
}
