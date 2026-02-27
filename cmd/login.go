package cmd

import (
	"fmt"
	"os"

	"github.com/Cloverhound/five9-cli/internal/appconfig"
	"github.com/Cloverhound/five9-cli/internal/auth"
	"github.com/Cloverhound/five9-cli/internal/config"
	"github.com/Cloverhound/five9-cli/internal/soap"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to Five9 and store credentials",
	Long:  "Prompts for username and password, validates against the Five9 API, and stores credentials in the OS keyring.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Prompt for username
		fmt.Print("Five9 Username: ")
		var username string
		fmt.Scanln(&username)
		if username == "" {
			return fmt.Errorf("username is required")
		}

		// Prompt for password (hidden)
		fmt.Print("Five9 Password: ")
		passBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			return fmt.Errorf("reading password: %w", err)
		}
		password := string(passBytes)
		if password == "" {
			return fmt.Errorf("password is required")
		}

		// Validate credentials with a simple API call
		config.SetUsername(username)
		config.SetPassword(password)

		fmt.Print("Validating credentials...")
		req := soap.NewRequest("getApiVersions")
		_, err = soap.Do(req)
		if err != nil {
			fmt.Println(" failed")
			return fmt.Errorf("authentication failed: %w", err)
		}
		fmt.Println(" ok")

		// Store in keyring
		if err := auth.SaveCredentials(username, password); err != nil {
			return fmt.Errorf("saving credentials: %w", err)
		}

		// Update config
		cfg, err := appconfig.Load()
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}
		cfg.SetDefaultUser(username)
		if err := cfg.Save(); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}

		fmt.Printf("Logged in as %s\n", username)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
