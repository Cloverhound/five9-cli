package cmd

import (
	"fmt"

	"github.com/Cloverhound/five9-cli/internal/appconfig"
	"github.com/Cloverhound/five9-cli/internal/auth"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication status and management",
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current authentication status",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := appconfig.Load()
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		if cfg.DefaultUser == "" {
			fmt.Println("Not logged in")
			return nil
		}

		_, err = auth.LoadCredentials(cfg.DefaultUser)
		if err != nil {
			fmt.Printf("Default user: %s (credentials not found in keyring)\n", cfg.DefaultUser)
			return nil
		}

		fmt.Printf("Logged in as: %s\n", cfg.DefaultUser)
		return nil
	},
}

func init() {
	authCmd.AddCommand(authStatusCmd)
	rootCmd.AddCommand(authCmd)
}
