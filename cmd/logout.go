package cmd

import (
	"fmt"

	"github.com/Cloverhound/five9-cli/internal/appconfig"
	"github.com/Cloverhound/five9-cli/internal/auth"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout [username]",
	Short: "Log out and remove stored credentials",
	Long:  "Removes stored credentials from the OS keyring. Without arguments, removes the default user.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := appconfig.Load()
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		username := ""
		if len(args) > 0 {
			username = args[0]
		} else {
			username = cfg.DefaultUser
		}

		if username == "" {
			return fmt.Errorf("no user specified and no default user — nothing to log out")
		}

		if err := auth.DeleteCredentials(username); err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Warning: could not remove credentials for %s: %v\n", username, err)
		}

		if cfg.DefaultUser == username {
			cfg.DefaultUser = ""
			if err := cfg.Save(); err != nil {
				return fmt.Errorf("saving config: %w", err)
			}
		}

		fmt.Printf("Logged out %s\n", username)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
