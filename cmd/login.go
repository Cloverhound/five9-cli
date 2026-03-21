package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Cloverhound/five9-cli/internal/appconfig"
	"github.com/Cloverhound/five9-cli/internal/auth"
	"github.com/Cloverhound/five9-cli/internal/config"
	"github.com/Cloverhound/five9-cli/internal/localconfig"
	"github.com/Cloverhound/five9-cli/internal/soap"
	"github.com/charmbracelet/huh"
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
		reader := bufio.NewReader(os.Stdin)
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
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

		// Update global config
		cfg, err := appconfig.Load()
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}
		cfg.SetDefaultUser(username)
		cfg.AddUser(username)
		if err := cfg.Save(); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}

		fmt.Printf("Logged in as %s\n\n", username)

		// Offer to associate this user with the current folder
		cwd, err := os.Getwd()
		if err == nil {
			if err := promptFolderAssociation(username, cwd); err != nil {
				return err
			}
		}

		return nil
	},
}

func promptFolderAssociation(username, dir string) error {
	folderName := filepath.Base(dir)

	var choice string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Associate this user with the current folder?").
				Description(
					fmt.Sprintf(
						"This saves \"%s\" as the default user for %s/.\n"+
							"Useful when different folders connect to different Five9 tenants,\n"+
							"so the right credentials are used automatically.",
						username, folderName,
					),
				).
				Options(
					huh.NewOption(fmt.Sprintf("Yes — use \"%s\" whenever I'm in %s/", username, folderName), "yes"),
					huh.NewOption("No — don't set folder default", "no"),
				).
				Value(&choice),
		),
	)

	if err := form.Run(); err != nil {
		return nil // user cancelled, not an error
	}

	if choice == "yes" {
		if err := localconfig.Save(dir, username); err != nil {
			return fmt.Errorf("saving local config: %w", err)
		}
		fmt.Printf("Saved to %s/.five9-cli/config.json\n", folderName)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
