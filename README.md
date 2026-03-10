# Gator 🐊

A CLI RSS feed aggregator built in Go. Because aggreGATOR.

## What it does

- Add RSS feeds from across the internet to be collected
- Store collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View post summaries in the terminal with links to the full articles

## Requirements

- Go
- PostgreSQL

## Usage

```bash
# Add a feed
gator addfeed "Feed Name" https://example.com/feed.xml

# Follow a feed
gator follow https://example.com/feed.xml

# View aggregated posts
gator browse
```
