package main

import (
	"fmt"
	"github.com/luism2302/moneydoc/internal/commands"
	"os"
)

func main() {
	supportedCommands := map[string]commands.Command{
		"help": commands.Help,
	}
	config := commands.NewConfig(supportedCommands)

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

	err := calledCommand.Run(config)
	if err != nil {
		fmt.Errorf("Couldnt run %s command: %w", calledCommand.Name, err)
		os.Exit(1)
	}

}
