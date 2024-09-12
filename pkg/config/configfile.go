package config

import (
	"encoding/json"
	"os"
	"path"
)

func CreateConfigFile(config *Configuration) error {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configFileFolder := path.Join(homeDir, ".cli")

	err = os.Mkdir(configFileFolder, os.ModePerm)
	if err != nil {
		return err
	}

	configFilePath := path.Join(configFileFolder, "config.json")

	js, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, js, os.ModePerm)
	if err != nil {
		return err
	}

	return nil

}

// func ReadConfigFile() (Configuration, error) {
// 	config :=
// 	return
// }
