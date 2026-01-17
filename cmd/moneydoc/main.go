package main

import (
	"database/sql"
	"fmt"
	"github.com/luism2302/moneydoc/internal/commands"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

const (
	dbFile = "moneydoc.db"
)

func main() {
	//Open conn
	dataDir, err := getLocationDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	db, err := sql.Open("sqlite", filepath.Join(dataDir, dbFile))
	if err != nil {
		fmt.Printf("Couldnt open moneydoc.db: %s", err.Error())
	}
	defer db.Close()

	//Initialize commands
	config := commands.NewConfig(supportedCommands)

	//Execute commands
	cliInput := os.Args[1:]
	if len(cliInput) == 0 {
		config.Cmds["help"].Run(config)
		os.Exit(0)
	}
	cliCommand := cliInput[0]
	calledCommand, ok := config.Cmds[cliCommand]
	if !ok {
		fmt.Printf("Unsupported command: %s\n", cliCommand)
		os.Exit(1)
	}
	if err := calledCommand.Run(config); err != nil {
		fmt.Printf("Couldnt run %s command: %s", calledCommand.Name, err.Error())
		os.Exit(1)
	}
}
