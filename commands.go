package main

import (
	"gator/internal/config"
	"errors"
	"fmt"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	registry map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.registry[cmd.name]
	if !ok {
		return errors.New("key not found")
	}
	err := value(s, cmd)
	if err != nil {
		return fmt.Errorf("failed the command: %w", err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registry[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username required")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to set username: %w", err)
	}
	fmt.Println("username set")
	return nil
}