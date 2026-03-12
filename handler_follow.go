package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samnodier/blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("follow command takes one argument, the feed url")
	}
	ctx := context.Background()
	feed, err := s.db.GetFeedByUrl(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("error fetching feed from database: %w", err)
	}
	feedFollowId := uuid.New()
	userId := user.ID
	feedId := feed.ID
	now := time.Now().UTC()
	insertedFeedFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        feedFollowId,
		UserID:    userId,
		FeedID:    feedId,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}
	fmt.Println(insertedFeedFollow.FeedName)
	fmt.Println(insertedFeedFollow.UserName)
	return nil
}
