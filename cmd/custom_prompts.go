package cmd

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Cloverhound/five9-cli/internal/soap"
	"github.com/spf13/cobra"
)

func init() {
	registerPromptsUpload()
	registerPromptsReplace()
}

func registerPromptsUpload() {
	var filePath string
	var name string
	var description string

	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload a WAV file as a new prompt",
		Long: `Upload a WAV file as a new prompt using addPromptWavInline.

The file is base64-encoded and sent inline in the SOAP body.
If --name is omitted, the filename without extension is used.

Examples:
  five9 prompts upload --file greeting.wav
  five9 prompts upload --file greeting.wav --name "Main Greeting"
  five9 prompts upload --file greeting.wav --name "Main Greeting" --description "IVR main menu"
  five9 prompts upload --file greeting.wav --dry-run
  five9 prompts upload --file greeting.wav --debug`,
		RunE: func(c *cobra.Command, args []string) error {
			wavData, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("reading WAV file: %w", err)
			}

			if name == "" {
				name = strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
			}

			encoded := base64.StdEncoding.EncodeToString(wavData)

			req := soap.NewRequest("addPromptWavInline")

			promptChildren := []soap.Param{
				{Name: "name", Value: name},
			}
			if description != "" {
				promptChildren = append(promptChildren, soap.Param{Name: "description", Value: description})
			}
			req.SetComplexParam("prompt", promptChildren)
			req.SetParam("wavFile", encoded)

			_, err = soap.Do(req)
			if err != nil {
				return err
			}

			fmt.Fprintf(os.Stderr, "Prompt %q uploaded successfully\n", name)
			return nil
		},
	}

	cmd.Flags().StringVar(&filePath, "file", "", "Path to WAV file to upload")
	cmd.MarkFlagRequired("file")
	cmd.Flags().StringVar(&name, "name", "", "Prompt name (defaults to filename without extension)")
	cmd.Flags().StringVar(&description, "description", "", "Prompt description")

	PromptsCmd.AddCommand(cmd)
}

func registerPromptsReplace() {
	var filePath string
	var name string

	cmd := &cobra.Command{
		Use:   "replace",
		Short: "Replace the WAV file on an existing prompt",
		Long: `Replace the audio file of an existing prompt using modifyPromptWavInline.

The file is base64-encoded and sent inline in the SOAP body.

Examples:
  five9 prompts replace --file greeting.wav --name "Main Greeting"
  five9 prompts replace --file greeting.wav --name "Main Greeting" --dry-run
  five9 prompts replace --file greeting.wav --name "Main Greeting" --debug`,
		RunE: func(c *cobra.Command, args []string) error {
			wavData, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("reading WAV file: %w", err)
			}

			encoded := base64.StdEncoding.EncodeToString(wavData)

			req := soap.NewRequest("modifyPromptWavInline")
			req.SetComplexParam("prompt", []soap.Param{
				{Name: "name", Value: name},
			})
			req.SetParam("wavFile", encoded)

			_, err = soap.Do(req)
			if err != nil {
				return err
			}

			fmt.Fprintf(os.Stderr, "Prompt %q replaced successfully\n", name)
			return nil
		},
	}

	cmd.Flags().StringVar(&filePath, "file", "", "Path to WAV file to upload")
	cmd.MarkFlagRequired("file")
	cmd.Flags().StringVar(&name, "name", "", "Name of existing prompt to replace")
	cmd.MarkFlagRequired("name")

	PromptsCmd.AddCommand(cmd)
}
