package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samnodier/gator/internal/database"
)

func parsePubDate(pubDate string) (time.Time, error) {
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
		"Mon, 02 Jan 2006 15:04:05 GMT",
	}
	for _, format := range formats {
		t, err := time.Parse(format, pubDate)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("could not parse date: %s", pubDate)
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error fetching next feed to fetch: %w", err)
	}
	now := time.Now().UTC()
	_, err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  now,
			Valid: true,
		},
		UpdatedAt: now,
		ID:        nextFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}
	feed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed %s: %w", nextFeed.Name, err)
	}
	for _, feedItem := range feed.Channel.Item {
		now := time.Now().UTC()
		publishedAt := sql.NullTime{}
		if t, err := parsePubDate((feedItem.PubDate)); err == nil {
			publishedAt = sql.NullTime{Time: t, Valid: true}
		}
		post, err := s.db.CreatePost(ctx, database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Title:     feedItem.Title,
			Url:       feedItem.Link,
			Description: sql.NullString{
				String: feedItem.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			fmt.Printf("error creating post: %v\n", err)
		} else {
			fmt.Printf("inserted post: %s\n", post.Title)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("agg command takes one argument, the time_between_reqs [1s, 1m, 1h]")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing the duration: %w", err)
	}
	if timeBetweenRequests < time.Second {
		return errors.New("time between requests must be at least 1 second")
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("error scraping feeds: %v\n", err)
		}
		<-ticker.C
	}
	return nil
}
