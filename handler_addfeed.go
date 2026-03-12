package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samnodier/blog-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("the register handler expects a two arguments, name and url")
	}
	ctx := context.Background()
	feedId := uuid.New()
	feedName := cmd.args[0]
	feedURL := cmd.args[1]
	userId := user.ID
	now := time.Now().UTC()
	insertedFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        feedId,
		Name:      feedName,
		Url:       feedURL,
		UserID:    userId,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return fmt.Errorf("error writing to feeds database: %w", err)
	}
	now = time.Now().UTC()
	insertedFeedFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    userId,
		FeedID:    insertedFeed.ID,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}
	fmt.Printf("inserted feed: %s | %s\n", insertedFeed.ID, insertedFeed.Name)
	fmt.Printf("inserted feedfollow: %s\n", insertedFeedFollow.FeedName)
	return nil
}
