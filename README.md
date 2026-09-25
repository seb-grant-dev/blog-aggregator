# Gator

A simple CLI application for subscribing to and reading RSS feeds and posts.

## Requirements
To run the application locally, you will need to have [Postgres](https://www.postgresql.org/) and [Go](https://go.dev/) installed.

## Installation
1. Clone this repository
2. **cd** into the folder (`cd blog-aggregator`)
3. Install the dependencies: `go install`
4. Create the config file in your home directory (`~/.gatorconfig.json`)
```
{
  "db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

## Running

Run the `gator` commands as per below:

`gator register`

`gator addfeed "Hacker News" "https://news.ycombinator.com/rss"`

`gator agg 10m`

`gater browse 5`

## Commands

| command   | params            | description                                                                            |
| --------- | ----------------- | -------------------------------------------------------------------------------------- |
| register  | username          | Registers a user account                                                               |
| login     | username          | Logs in an existing user                                                               |
| addfeed   | title url         | Registers a feed to track for the logged in user                                       |
| follow    | url               | Follows a feed already registered by another user                                      |
| following |                   | Lists the feeds that the logged in user is following                                   |
| agg       | <frequency>       | Aggregates all feeds, polling them based on the passed frequency (i.e. 20s, 2m, 3h...) |
| browse    | <limit default=2> | Lists the latest posts, up to <limit>                                                  |
| reset     |                   | Completely resets the database. WARNING: This will immediately wipe your database!     |

```
