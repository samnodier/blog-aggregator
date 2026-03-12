package main

import (
	"context"
	"fmt"
	"github.com/samnodier/gator/internal/database"
	"strconv"
	"strings"
)

func stripHTML(s string) string {
	var result strings.Builder
	inTag := false
	for _, r := range s {
		if r == '<' {
			inTag = true
		} else if r == '>' {
			inTag = false
		} else if !inTag {
			result.WriteRune(r)
		}
	}
	return strings.TrimSpace(result.String())
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		if num, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = num
		} else {
			fmt.Println("can't parse the limit argument")
		}
	}
	if len(cmd.args) > 1 {
		return fmt.Errorf("browse takes one argument: limit (default 2)")
	}
	ctx := context.Background()
	posts, err := s.db.GetUserPosts(ctx, database.GetUserPostsParams{
		ID:    user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}
	for i, post := range posts {
		fmt.Printf("%d. %s\n", i+1, post.Title)
		fmt.Printf("   %s\n", post.Url)
		if post.PublishedAt.Valid {
			fmt.Printf("   Published: %s\n", post.PublishedAt.Time.Format("2006-01-02"))
		}
		if post.Description.Valid {
			fmt.Printf("   %s\n", stripHTML(post.Description.String))
		}
		fmt.Println()
	}
	return nil
}
