package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the Chrome native messaging host",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Read the manifest file
		manifestData, err := os.ReadFile("chrome-host.json")
		if err != nil {
			return fmt.Errorf("failed to read manifest file: %w", err)
		}

		// Parse the manifest
		var manifest map[string]interface{}
		if err := json.Unmarshal(manifestData, &manifest); err != nil {
			return fmt.Errorf("failed to parse manifest file: %w", err)
		}

		// Get the absolute path of the current executable
		exePath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("failed to get executable path: %w", err)
		}
		exePath, err = filepath.Abs(exePath)
		if err != nil {
			return fmt.Errorf("failed to get absolute path: %w", err)
		}

		// Update the path in the manifest
		manifest["path"] = exePath

		// Create the NativeMessagingHosts directory if it doesn't exist
		hostsDir := filepath.Join(os.Getenv("HOME"), ".config", "google-chrome", "NativeMessagingHosts")
		if err := os.MkdirAll(hostsDir, 0755); err != nil {
			return fmt.Errorf("failed to create NativeMessagingHosts directory: %w", err)
		}

		// Write the updated manifest
		manifestPath := filepath.Join(hostsDir, "com.biscuit.extensions.exercism.json")
		updatedData, err := json.MarshalIndent(manifest, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal updated manifest: %w", err)
		}

		if err := os.WriteFile(manifestPath, updatedData, 0644); err != nil {
			return fmt.Errorf("failed to write manifest file: %w", err)
		}

		fmt.Printf("Chrome native messaging host installed at %s\n", manifestPath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
