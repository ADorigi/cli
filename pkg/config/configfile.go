package config

import (
	"encoding/json"
	"io"
	"os"
	"path"
)

func CreateConfigFile(config *Configuration) error {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configFileFolder := path.Join(homeDir, ".checkctl")

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

func ReadConfigFile() (*Configuration, error) {
	config := &Configuration{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configFilePath := path.Join(homeDir, ".checkctl", "config.json")

	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	data, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
