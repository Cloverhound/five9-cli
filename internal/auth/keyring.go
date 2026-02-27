package auth

import (
	"encoding/json"
	"fmt"

	"github.com/zalando/go-keyring"
)

const serviceName = "five9-cli"

type StoredCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SaveCredentials(username, password string) error {
	creds := StoredCredentials{Username: username, Password: password}
	data, err := json.Marshal(creds)
	if err != nil {
		return fmt.Errorf("marshaling credentials: %w", err)
	}
	return keyring.Set(serviceName, username, string(data))
}

func LoadCredentials(username string) (*StoredCredentials, error) {
	data, err := keyring.Get(serviceName, username)
	if err != nil {
		return nil, fmt.Errorf("loading credentials for %s: %w", username, err)
	}
	var creds StoredCredentials
	if err := json.Unmarshal([]byte(data), &creds); err != nil {
		return nil, fmt.Errorf("parsing stored credentials: %w", err)
	}
	return &creds, nil
}

func DeleteCredentials(username string) error {
	err := keyring.Delete(serviceName, username)
	if err == keyring.ErrNotFound {
		return nil
	}
	return err
}
