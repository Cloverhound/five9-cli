package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Cloverhound/five9-cli/internal/appconfig"
	"github.com/Cloverhound/five9-cli/internal/auth"
	"github.com/Cloverhound/five9-cli/internal/config"
	"github.com/Cloverhound/five9-cli/internal/output"
	"github.com/Cloverhound/five9-cli/internal/soap"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "five9",
	Short:        "Five9 CLI — manage Five9 Configuration Web Services",
	Long:         `A command-line interface for the Five9 Configuration Web Services API.`,
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Debug
		debug, _ := cmd.Flags().GetBool("debug")
		config.SetDebug(debug)

		// Dry-run
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		config.SetDryRun(dryRun)

		// Output format
		format, _ := cmd.Flags().GetString("output")
		output.SetFormat(format)

		// Endpoint
		endpoint, _ := cmd.Flags().GetString("endpoint")
		config.SetEndpoint(endpoint)

		// Skip auth for certain commands
		if skipAuth(cmd) {
			return nil
		}

		// Resolve credentials: flags → env → keyring
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if username == "" {
			username = os.Getenv("FIVE9_USERNAME")
		}
		if password == "" {
			password = os.Getenv("FIVE9_PASSWORD")
		}

		// Try keyring if still missing
		if username == "" || password == "" {
			userFlag, _ := cmd.Flags().GetString("user")
			if userFlag == "" {
				userFlag = os.Getenv("FIVE9_USER")
			}

			if userFlag == "" {
				cfg, err := appconfig.Load()
				if err == nil && cfg.DefaultUser != "" {
					userFlag = cfg.DefaultUser
				}
			}

			if userFlag != "" {
				creds, err := auth.LoadCredentials(userFlag)
				if err == nil {
					if username == "" {
						username = creds.Username
					}
					if password == "" {
						password = creds.Password
					}
				}
			}
		}

		if username == "" || password == "" {
			return fmt.Errorf("no credentials found — run 'five9 login' or set FIVE9_USERNAME/FIVE9_PASSWORD")
		}

		config.SetUsername(username)
		config.SetPassword(password)

		// Load endpoint from app config if not set via flag/env
		if endpoint == "" {
			cfg, err := appconfig.Load()
			if err == nil && cfg.Endpoint != "" {
				config.SetEndpoint(cfg.Endpoint)
			}
		}

		return nil
	},
}

func Execute() error {
	err := rootCmd.Execute()
	if errors.Is(err, soap.ErrDryRun) {
		return nil
	}
	return err
}

func init() {
	rootCmd.PersistentFlags().String("username", "", "Five9 username (overrides keyring)")
	rootCmd.PersistentFlags().String("password", "", "Five9 password (overrides keyring)")
	rootCmd.PersistentFlags().String("output", "json", "Output format: json, table, csv, raw")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging of SOAP requests")
	rootCmd.PersistentFlags().Bool("dry-run", false, "Print write requests without executing them")
	rootCmd.PersistentFlags().String("user", "", "Use a specific authenticated user (username)")
	rootCmd.PersistentFlags().String("endpoint", "", "Override Five9 API endpoint URL")

	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.PrintErrf("Error: %s\n\n", err)
		cmd.PrintErr(cmd.UsageString())
		cmd.Root().SilenceErrors = true
		return err
	})
}

func skipAuth(cmd *cobra.Command) bool {
	name := cmd.Name()
	for c := cmd; c != nil; c = c.Parent() {
		switch c.Name() {
		case "login", "logout", "auth", "version", "help", "five9":
			if c.Name() == "five9" {
				continue
			}
			return true
		}
	}
	if name == "help" || name == "five9" {
		return true
	}
	// --sample flag prints sample JSON without needing auth
	if f := cmd.Flags().Lookup("sample"); f != nil && f.Value.String() == "true" {
		return true
	}
	return false
}
