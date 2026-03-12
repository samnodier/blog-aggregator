# Gator 🐊

A CLI RSS feed aggregator built in Go. Because aggreGATOR.

## What it does

- Add RSS feeds from across the internet to be collected
- Store collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View post summaries in the terminal with links to the full articles

## Requirements

- Go 1.22 or higher 
- PostgreSQL

## Installation

``` bash
go install github.com/samnodier/gator@latest
```

## Configuration

Create a config file at the following locations depending on your OS:

- Linux/macOS: `~/.gatorconfig.json`
- Windows: `C:\Users\<YourUsername>\.gatorconfig.json`

With the contents of your connection string:
``` json
{
    "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

- Replace `username`, `password`, and `gator` with your PostgreSQL credentials and database name.
- macOS(no password, your username): `postgres://wagslane:@localhost:5432/gator`

## Database Setup

Run the migrations using goose:

``` bash
goose -dir sql/schema postgres "your_connection_string" up
```

You don't need `?sslmode=disable` in the terminal connection string

## Commands

Register a new user:
```bash
gator register 
```

Login as a user:
```bash
gator login 
```

Add a feed:
```bash
gator addfeed "Feed Name" "https://feed-url.com/rss"
```

Follow a feed:
```bash
gator follow "https://feed-url.com/rss"
```

Start aggregating:
```bash
gator agg 1m
```

Browse posts:
```bash
gator browse 10
```
