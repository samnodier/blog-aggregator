package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/samnodier/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("follow command takes one argument, the feed url")
	}
	ctx := context.Background()
	userId := user.ID
	url := cmd.args[0]
	err := s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: userId,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}
	fmt.Printf("feed follow %s deleted successfully\n", url)
	return nil
}
