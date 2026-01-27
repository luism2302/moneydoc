package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"database/sql"

	"github.com/luism2302/moneydoc/internal/database/sqlc"
	"modernc.org/sqlite"
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
		return errors.New("usage: register <username> <email>")
	}
	username := args[0]
	email := args[1]

	newUserParams := sqlc.CreateNewUserParams{
		Username: username,
		Email:    email,
	}

	createdUser, err := cfg.Queries.CreateNewUser(context.Background(), newUserParams)
	var sqliteError *sqlite.Error
	if err != nil {
		if errors.As(err, &sqliteError) {
			failColumn := strings.Split(strings.Split(sqliteError.Error(), ".")[1], " ")[0]
			return fmt.Errorf("%s already exists in users table", failColumn)
		}
	}

	fmt.Printf("Created user: %s with email: %s", createdUser.Username, createdUser.Email)
	return nil
}

func ResetCallback(cfg *Config, args []string) error {
	if err := cfg.Queries.DeleteAllUsers(context.Background()); err != nil {
		return errors.New("error deleting users from database")
	}
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

var Reset = Command{
	Name:        "reset",
	Description: "deletes every register from every table in the database",
	Callback:    ResetCallback,
}
