package commands

import (
	"fmt"
	"github.com/seb-grant-dev/blog-aggregator/state"
)

type Command struct {
	name string
	arguments []string
}

func NewCommand(name string, args ...string) Command {
	return Command{
		name: name,
		arguments: args,
	}
}


type Commands struct {
	registry map[string]func(*state.State,Command) error
}

func NewCommandRegistry() *Commands {
	registry := Commands {
		registry: map[string]func(*state.State,Command) error{},
	}

	return &registry
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	command, ok := c.registry[cmd.name]
	if ok {
		err := command(s,cmd)
		return err
	}
	return fmt.Errorf("%s Command not registered",cmd.name)
}

func (c *Commands) Register(name string, f func(*state.State,Command) error) {
	c.registry[name] = f
}
