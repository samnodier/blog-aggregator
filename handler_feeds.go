package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()
	if len(cmd.args) > 0 {
		return errors.New("feeds handler takes no arguments")
	}
	feeds, err := s.db.ListFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error fetching feeds from database: %w", err)
	}
	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		user, err := s.db.GetUserById(ctx, feed.UserID)
		if err != nil {
			return fmt.Errorf("error fetching feeds from database: %w", err)
		}
		fmt.Println(user.Name)
	}
	return nil
}
