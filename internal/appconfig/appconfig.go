package appconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultUser string   `json:"default_user"`
	KnownUsers  []string `json:"known_users,omitempty"`
	Endpoint    string   `json:"endpoint,omitempty"`

	path string
}

func configDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".five9-cli")
}

func configPath() string {
	return filepath.Join(configDir(), "config.json")
}

func Load() (*Config, error) {
	cfg := &Config{path: configPath()}

	data, err := os.ReadFile(cfg.path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, fmt.Errorf("reading config: %w", err)
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}
	return cfg, nil
}

func (c *Config) Save() error {
	dir := filepath.Dir(c.path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("creating config dir: %w", err)
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling config: %w", err)
	}

	if err := os.WriteFile(c.path, data, 0600); err != nil {
		return fmt.Errorf("writing config: %w", err)
	}
	return nil
}

func (c *Config) SetDefaultUser(username string) {
	c.DefaultUser = username
}

// AddUser adds a username to the known users list if not already present.
func (c *Config) AddUser(username string) {
	for _, u := range c.KnownUsers {
		if u == username {
			return
		}
	}
	c.KnownUsers = append(c.KnownUsers, username)
}

// RemoveUser removes a username from the known users list and clears
// the default user if it matches.
func (c *Config) RemoveUser(username string) {
	filtered := c.KnownUsers[:0]
	for _, u := range c.KnownUsers {
		if u != username {
			filtered = append(filtered, u)
		}
	}
	c.KnownUsers = filtered
	if c.DefaultUser == username {
		c.DefaultUser = ""
	}
}
