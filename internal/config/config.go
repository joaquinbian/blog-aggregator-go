package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func write(cfg Config) error {
	dir, err := getConfigFilePath()

	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)

	if err != nil {
		return err
	}

	err = os.WriteFile(dir, data, 0644)

	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("error al leer home directory: %v", err)
	}

	return fmt.Sprintf("%s/%s", dir, configFileName), nil
}

func Read() (Config, error) {
	dir, err := getConfigFilePath()

	if err != nil {
		return Config{}, fmt.Errorf("error al leer home directory: %v", err)
	}

	data, err := os.ReadFile(dir)

	if err != nil {
		return Config{}, fmt.Errorf("error al leer archivo de configuracion: %v", err)
	}

	var cfg = Config{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("error al leer archivo de configuracion: %v", err)
	}

	return cfg, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.Current_user_name = name
	err := write(*cfg)
	if err != nil {
		fmt.Printf("en el if del error set user")
		return fmt.Errorf("error al escribir config file: %v", err)
	}
	return nil
}
