package commands

import (
	"fmt"
	"github.com/seb-grant-dev/blog-aggregator/state"
)

type command struct {
	name string
	arguments []string
}

func NewCommand(name string, args ...string) command {
	return command{
		name: name,
		arguments: args,
	}
}


type Commands struct {
	registry map[string]func(*state.State,command) error
}

func NewCommandRegistry() *Commands {
	registry := Commands {
		registry: map[string]func(*state.State,command) error{},
	}

	return &registry
}

func (c *Commands) Run(s *state.State, cmd command) error {
	command, ok := c.registry[cmd.name]
	if ok {
		err := command(s,cmd)
		return err
	}
	return fmt.Errorf("%s command not registered",cmd.name)
}

func (c *Commands) Register(name string, f func(*state.State,command) error) {
	c.registry[name] = f
}
