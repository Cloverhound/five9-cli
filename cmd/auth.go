package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Cloverhound/five9-cli/internal/appconfig"
	"github.com/Cloverhound/five9-cli/internal/auth"
	"github.com/Cloverhound/five9-cli/internal/localconfig"
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

		// Show folder default if set
		if cwd, err := os.Getwd(); err == nil {
			if lcfg, err := localconfig.Load(cwd); err == nil && lcfg != nil && lcfg.User != "" {
				fmt.Printf("Folder default: %s (for this directory)\n", lcfg.User)
			}
		}

		return nil
	},
}

var authListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all authenticated users",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := appconfig.Load()
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		if len(cfg.KnownUsers) == 0 {
			fmt.Println("No authenticated users. Run: five9 login")
			return nil
		}

		for _, username := range cfg.KnownUsers {
			status := "ok"
			if _, err := auth.LoadCredentials(username); err != nil {
				status = "no credentials"
			}

			marker := ""
			if username == cfg.DefaultUser {
				marker = " (default)"
			}

			fmt.Printf("  %s  [%s]%s\n", username, status, marker)
		}

		return nil
	},
}

var authSwitchCmd = &cobra.Command{
	Use:   "switch <username>",
	Short: "Switch the default user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]

		cfg, err := appconfig.Load()
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		// Check user is known
		found := false
		for _, u := range cfg.KnownUsers {
			if u == username {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("user %s not found — run: five9 login", username)
		}

		cfg.SetDefaultUser(username)
		if err := cfg.Save(); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}

		fmt.Printf("Default user set to %s\n", username)
		return nil
	},
}

var authSetFolderDefaultCmd = &cobra.Command{
	Use:   "set-folder-default <username>",
	Short: "Set the default user for the current folder",
	Long:  "Associates a Five9 user with the current working directory. When running commands from this folder, this user's credentials will be used automatically.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]

		// Verify credentials exist in keyring
		if _, err := auth.LoadCredentials(username); err != nil {
			return fmt.Errorf("no credentials found for %s — run 'five9 login' first", username)
		}

		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getting current directory: %w", err)
		}

		if err := localconfig.Save(cwd, username); err != nil {
			return fmt.Errorf("saving folder default: %w", err)
		}

		fmt.Printf("Set folder default to %s for %s\n", username, cwd)
		return nil
	},
}

var authClearFolderDefaultCmd = &cobra.Command{
	Use:   "clear-folder-default",
	Short: "Remove the folder default user for the current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getting current directory: %w", err)
		}

		cfg, err := localconfig.Load(cwd)
		if err != nil || cfg == nil {
			fmt.Println("No folder default set for this directory.")
			return nil
		}

		// Remove the config file and directory if empty
		configPath := filepath.Join(cwd, ".five9-cli", "config.json")
		if err := os.Remove(configPath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("removing folder config: %w", err)
		}

		configDir := filepath.Join(cwd, ".five9-cli")
		os.Remove(configDir) // ignore error — dir may not be empty

		fmt.Printf("Cleared folder default (was %s)\n", cfg.User)
		return nil
	},
}

func init() {
	authCmd.AddCommand(authStatusCmd)
	authCmd.AddCommand(authListCmd)
	authCmd.AddCommand(authSwitchCmd)
	authCmd.AddCommand(authSetFolderDefaultCmd)
	authCmd.AddCommand(authClearFolderDefaultCmd)
	rootCmd.AddCommand(authCmd)
}
