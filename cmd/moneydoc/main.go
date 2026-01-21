package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/luism2302/moneydoc/internal/commands"
	_ "modernc.org/sqlite"
)

const (
	dbFile = "moneydoc.db"
)

func main() {
	//Open conn
	dataDir, err := getLocationDB()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	db, err := sql.Open("sqlite", filepath.Join(dataDir, dbFile))
	if err != nil {
		log.Fatalf("Couldnt open connection to moneydoc.db: %s", err.Error())
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("Couldnt connect to moneydoc.db: %s", err.Error())
	}

	//Initialize commands
	config := commands.NewConfig(supportedCommands, db)

	//Execute commands
	if len(os.Args) <= 1 {
		config.Cmds["help"].Run(config, os.Args)
		os.Exit(0)
	}
	cliInput := os.Args[1:]
	commandArgs := cliInput[1:]
	commandName := cliInput[0]

	calledCommand, ok := config.Cmds[commandName]
	if !ok {
		log.Fatalf("Unsupported command: %s\n", commandName)
	}
	if err := calledCommand.Run(config, commandArgs); err != nil {
		log.Fatalf("Couldnt run %s command: %s", calledCommand.Name, err.Error())
	}
}
