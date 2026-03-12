package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	if len(cmd.args) > 0 {
		return errors.New("the reset handler expects no argument")
	}
	err := s.db.Reset(ctx)
	if err != nil {
		return fmt.Errorf("error deleting users table: %w", err)
	}
	fmt.Println("users were deleted successfully!")
	return nil
}
