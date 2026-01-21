package commands

import (
	"context"
	"fmt"

	"database/sql"

	"github.com/luism2302/moneydoc/internal/database/sqlc"
)

type Command struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

func (cmd Command) Run(cfg *Config, args []string) error {
	err := cmd.Callback(cfg, args)
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	Cmds    map[string]Command
	Queries *sqlc.Queries
}

func NewConfig(cmds map[string]Command, db *sql.DB) *Config {
	return &Config{
		Cmds:    cmds,
		Queries: sqlc.New(db),
	}
}

func HelpCallback(cfg *Config, args []string) error {
	fmt.Println("A CLI tool for controlling expenses")
	fmt.Println("Usage: moneydoc [command] <args>")
	fmt.Println("Supported Commands:")
	for key, command := range cfg.Cmds {
		fmt.Printf("\t-%s: %s\n", key, command.Description)
	}
	return nil
}

func RegisterCallback(cfg *Config, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: register <username> <email>")
	}
	username := args[0]
	email := args[1]

	newUserParams := sqlc.CreateNewUserParams{
		Username: username,
		Email:    email,
	}

	createdUser, err := cfg.Queries.CreateNewUser(context.Background(), newUserParams)

	if err != nil {
		return fmt.Errorf("Error creating new user: %s with email: %s. %w", username, email, err)
	}

	fmt.Printf("Created user: %s with email: %s", createdUser.Username, createdUser.Email)
	return nil

}

var Help = Command{
	Name:        "help",
	Description: "Help about any command",
	Callback:    HelpCallback,
}

var Register = Command{
	Name:        "register",
	Description: "registers a new user. usage: <username> <email>",
	Callback:    RegisterCallback,
}
