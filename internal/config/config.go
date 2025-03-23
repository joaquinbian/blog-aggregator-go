package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const ConfigFileName = "gatorconfig.json"

func Read() (Config, error) {

	usrHomeDir, err := os.UserHomeDir()

	if err != nil {
		return Config{}, fmt.Errorf("Error reading user home directory %w", err)
	}

	file, err := os.ReadFile(fmt.Sprint(usrHomeDir, "/", ConfigFileName))

	var config Config

	if err := json.Unmarshal(file, &config); err != nil {
		return Config{}, fmt.Errorf("Error reading user home directory %w", err)
	}

	return config, nil

}
