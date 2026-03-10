package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

type state struct {
	config *Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmdsMap map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	username := cmd.args[1]
	err := s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("Error writing user: %w", err)
	}
	fmt.Printf("User %s was set successfully!\n")
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("Error getting home directory")
	}
	return filepath.Join(home, configFileName), nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(cfgPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	c := Config{}
	configPath, err := getConfigFilePath()
	if err != nil {
		return c, fmt.Errorf("Error: %w", err)
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return c, fmt.Errorf("Error Reading Config: %w", err)
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return c, fmt.Errorf("Error: %w", err)
	}
	return c, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}
