package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	if len(cmd.args) > 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	err := s.db.Reset(ctx)
	if err != nil {
		return fmt.Errorf("error deleting table: %w", err)
	}
	fmt.Println("table was deleted successfully!")
	return nil
}
