package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/samnodier/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return errors.New("following command takes no arguments")
	}
	ctx := context.Background()
	userFeeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to retrieve user feeds: %w", err)
	}
	for _, feed := range userFeeds {
		fmt.Println(feed.FeedName)
	}
	return nil
}
