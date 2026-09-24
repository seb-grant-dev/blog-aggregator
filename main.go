package main

import (
	"fmt"
	"os"
	"github.com/seb-grant-dev/blog-aggregator/internal/config"
	"github.com/seb-grant-dev/blog-aggregator/commands"
	"github.com/seb-grant-dev/blog-aggregator/state"
)

func main() {

	cfg := config.Read()
	appState := state.State{
		Config: cfg,
	}

	// Register the command registry
	commandRegistry := commands.NewCommandRegistry()
	commandRegistry.Register("login",commands.HandlerLogin)



	// Get the command name and arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("There should be at least 1 argument passed to this program. You passed %d\n",len(args)-1)
		os.Exit(1)
	}

	// Parse command and args from cmd line
	cmd := commands.NewCommand(args[1],args[2:]...)


	// Run the actual command
	err := commandRegistry.Run(&appState,cmd)
	if err != nil {
		fmt.Printf("Error: %s\n",err)
		os.Exit(1)
	}
}
