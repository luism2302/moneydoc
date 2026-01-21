package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/luism2302/moneydoc/internal/commands"
)

const (
	dataDir = "MoneyDocData"
)

var supportedCommands = map[string]commands.Command{
	"help":     commands.Help,
	"register": commands.Register,
}

func getLocationDB() (dataDirLocation string, err error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("Couldnt get users config directory: %w", err)
	}
	dataDirLocation = filepath.Join(userConfigDir, dataDir)
	if _, err := os.Stat(dataDirLocation); os.IsNotExist(err) {
		err := createDataDir(userConfigDir)
		if err != nil {
			return "", fmt.Errorf("Couldnt create %s directory inside users config directory: %w", dataDir, err)
		}
	}
	return dataDirLocation, nil
}

func createDataDir(userConfigDir string) error {
	err := os.Mkdir(filepath.Join(userConfigDir, dataDir), 0755)
	return err
}
