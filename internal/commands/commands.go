package commands

import (
	"fmt"
)

type Command struct {
	Name        string
	Description string
	Callback    func(*Config) error
	Args        []string
}

type Config struct {
	Cmds map[string]Command
}

func NewConfig(cmds map[string]Command) *Config {
	return &Config{
		Cmds: cmds,
	}
}

func HelpCallBack(cfg *Config) error {
	fmt.Println("A CLI tool for controlling expenses")
	fmt.Println("Usage: moneydoc [command] <args>")
	fmt.Println("Supported Commands:")
	for key, command := range cfg.Cmds {
		fmt.Printf("\t-%s: %s\n", key, command.Description)
	}
	return nil
}

var Help = Command{
	Name:        "help",
	Description: "Help about any command",
	Callback:    HelpCallBack,
}

func (cmd Command) Run(cfg *Config) error {
	err := cmd.Callback(cfg)
	if err != nil {
		return err
	}
	return nil
}
