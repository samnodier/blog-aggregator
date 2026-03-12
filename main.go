package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/samnodier/blog-aggregator/internal/config"
	"github.com/samnodier/blog-aggregator/internal/database"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func main() {
	cfg, err := config.Read()
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("Error opening the database: %w", err)
	}
	dbQueries := database.New(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := &state{
		cfg: &cfg,
		db:  dbQueries,
	}
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))

	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments provided")
		os.Exit(1)
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	if err := cmds.run(s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
