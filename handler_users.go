package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	ctx := context.Background()
	if len(cmd.args) > 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	err := s.db.Reset(ctx)
	return nil
}
