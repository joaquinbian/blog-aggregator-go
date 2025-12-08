# Gator - RSS Feed Aggregator

A command-line RSS feed aggregator built with Go and PostgreSQL.

## Prerequisites

Before running this program, you'll need to have the following installed on your system:

### 1. PostgreSQL
You need PostgreSQL installed and running on your machine.

**Installation:**
- **macOS**: `brew install postgresql`
- **Ubuntu/Debian**: `sudo apt-get install postgresql`
- **Windows**: Download from [postgresql.org](https://www.postgresql.org/download/)

After installation, make sure PostgreSQL is running:
```bash
# Check if PostgreSQL is running
psql --version
```

Create a database for the application:
```bash
createdb gator
```

### 2. Go
You need Go 1.25 or higher installed.

**Installation:**
- Download from [go.dev/dl](https://go.dev/dl/)
- Or use a package manager:
  - **macOS**: `brew install go`
  - **Ubuntu/Debian**: `sudo apt-get install golang`

Verify installation:
```bash
go version
```

## Installation

Install the `gator` CLI tool using `go install`:

```bash
go install github.com/joaquinbian/blog-aggregator-go/cmd/gator@latest
```

This will install the `gator` binary to your `$GOPATH/bin` directory (typically `~/go/bin`).

This will install the `gator` binary to your `$GOPATH/bin` directory (typically `~/go/bin`).

Make sure your `$GOPATH/bin` is in your system's PATH:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Configuration

Before running the program, you need to set up a configuration file.

Create a file named `.gatorconfig.json` in your home directory:

```bash
cd ~
touch .gatorconfig.json
```

Add the following content to `.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username` and `password` with your PostgreSQL credentials.

## Usage

Once installed and configured, you can use the `gator` CLI with the following commands:

### User Management

**Register a new user:**
```bash
gator register <your_name>
```

**Login:**
```bash
gator login <your_name>
```

**List all users:**
```bash
gator users
```

**Reset database (delete all users):**
```bash
gator reset
```
⚠️ Warning: This will delete all users and their associated data from the database.

### Feed Management

**Add a new RSS feed:**
```bash
gator addfeed <feed_name> <feed_url>
```
Example:
```bash
gator addfeed "Golang Blog" https://go.dev/blog/feed.atom
```

**List all feeds:**
```bash
gator feeds
```

**Follow a feed:**
```bash
gator follow <feed_url>
```

**See your followed feeds:**
```bash
gator following
```

**Unfollow a feed:**
```bash
gator unfollow <feed_url>
```

### Feed Aggregation

**Start the aggregator (fetches feeds at regular intervals):**
```bash
gator agg <time_interval>
```
Examples:
```bash
gator agg 30s   # Fetch every 30 seconds
gator agg 5m    # Fetch every 5 minutes
gator agg 1h    # Fetch every hour
```

This command will continuously fetch RSS feeds in the background based on the time interval you specify. It automatically stores all posts from your followed feeds into the database.

### Browse Posts

**View posts from your followed feeds:**
```bash
gator browse <limit>
```
Examples:
```bash
gator browse       # View 2 most recent posts (default)
gator browse 10    # View 10 most recent posts
gator browse 50    # View 50 most recent posts
```

This command displays posts from all the feeds you're following, including the post title, description, feed name, and link.

## Example Workflow

```bash
# 1. Register yourself
gator register alice

# 2. Add some RSS feeds
gator addfeed "Tech Crunch" https://techcrunch.com/feed/
gator addfeed "Hacker News" https://news.ycombinator.com/rss

# 3. Follow a feed
gator follow https://techcrunch.com/feed/

# 4. Check which feeds you're following
gator following

# 5. Start aggregating feeds (fetches every minute)
# Run this in a separate terminal window
gator agg 1m

# 6. Browse the collected posts
gator browse 5
```

## Database Schema

The application uses four main tables:

- **users**: Stores user information
- **feeds**: Stores RSS feed information
- **feed_follows**: Tracks which users follow which feeds
- **posts**: Stores individual posts/articles from RSS feeds

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.