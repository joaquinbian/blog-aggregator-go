package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const ConfigFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	usrHomeDir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Error reading user home directory %w", err)
	}

	return filepath.Join(usrHomeDir, ConfigFileName), nil
}

func write(cfg Config) error {

	fileName, err := getConfigFilePath()

	if err != nil {
		return fmt.Errorf("There was an error writing to the config file: %w", err)
	}

	newFile, err := json.Marshal(cfg)

	if err != nil {
		return fmt.Errorf("There was an error marshalling the config file: %w", err)
	}

	if err := os.WriteFile(fileName, newFile, os.ModeAppend.Perm()); err != nil {
		return fmt.Errorf("There was an error writing the config file: %w", err)
	}

	return nil
}

func Read() (Config, error) {

	fileName, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(fileName)

	var config Config

	if err := json.Unmarshal(file, &config); err != nil {
		return Config{}, fmt.Errorf("Error reading user home directory %w", err)
	}

	return config, nil

}

func (c *Config) SetUser(username string) error {
	c.Current_user_name = username
	err := write(*c)

	if err != nil {
		return fmt.Errorf("There was parsing the config file: %w", err)
	}

	return nil

}
