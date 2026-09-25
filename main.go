package main

import _ "github.com/lib/pq"

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/seb-grant-dev/blog-aggregator/internal/config"
	"github.com/seb-grant-dev/blog-aggregator/internal/database"
	"github.com/seb-grant-dev/blog-aggregator/commands"
	"github.com/seb-grant-dev/blog-aggregator/state"
	"github.com/seb-grant-dev/blog-aggregator/internal/middleware"
)

func main() {


	cfg := config.Read()
	appState := state.State{
		Config: &cfg,
	}

	db, err := sql.Open("postgres",cfg.DBUrl)
	if err != nil {
		fmt.Printf("Error: Could not connect to database: %s\n",err)
	}
	dbQueries := database.New(db)
	appState.DB = dbQueries


	// Register the command registry
	commandRegistry := commands.NewCommandRegistry()
	commandRegistry.Register("login",commands.HandlerLogin)
	commandRegistry.Register("register",commands.HandlerRegister)
	commandRegistry.Register("reset",commands.HandlerReset)
	commandRegistry.Register("users",commands.HandlerUsers)
	commandRegistry.Register("agg",commands.HandlerAgg)
	commandRegistry.Register("addfeed",middleware.MiddlewareLoggedIn(commands.HandlerAddFeed))
	commandRegistry.Register("feeds",commands.HandlerFeeds)
	commandRegistry.Register("follow",middleware.MiddlewareLoggedIn(commands.HandlerFollow))
	commandRegistry.Register("following",middleware.MiddlewareLoggedIn(commands.HandlerFollowing))
	commandRegistry.Register("unfollow",middleware.MiddlewareLoggedIn(commands.HandlerUnfollow))
	commandRegistry.Register("browse",middleware.MiddlewareLoggedIn(commands.HandlerBrowse))


	// Get the command name and arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("There should be at least 1 argument passed to this program. You passed %d\n",len(args)-1)
		os.Exit(1)
	}

	// Parse command and args from cmd line
	cmd := commands.NewCommand(args[1],args[2:]...)


	// Run the actual command
	err = commandRegistry.Run(&appState,cmd)
	if err != nil {
		fmt.Printf("Error: %s\n",err)
		os.Exit(1)
	}
}
