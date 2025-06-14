package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the Chrome native messaging host",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the absolute path of the current executable
		execPath, err := os.Executable()
		if err != nil {
			log.Fatalf("Failed to get executable path: %v", err)
		}
		absPath, err := filepath.Abs(execPath)
		if err != nil {
			log.Fatalf("Failed to get absolute path: %v", err)
		}

		// TODO: fix manifest file
		manifest := map[string]interface{}{
			"name":            "com.google.chrome.biscuit.exercism",
			"description":     "Chrome Native Messaging API Example Host",
			"path":            absPath,
			"type":            "stdio",
			"allowed_origins": []string{"chrome-extension://knldjmfmopnpolahpmmgbagdohdnhkik/"},
		}

		// Marshal the manifest to JSON
		manifestJSON, err := json.MarshalIndent(manifest, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal manifest: %v", err)
		}

		// Define the user-specific target directory and file path
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get user home directory: %v", err)
		}
		targetDir := filepath.Join(homeDir, ".config", "google-chrome", "NativeMessagingHosts")
		targetFile := filepath.Join(targetDir, "com.google.chrome.example.echo.json")

		// Ensure the target directory exists
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			log.Fatalf("Failed to create directory %s: %v", targetDir, err)
		}

		// Write the manifest file
		if err := os.WriteFile(targetFile, manifestJSON, 0644); err != nil {
			log.Fatalf("Failed to write manifest file: %v", err)
		}

		log.Printf("Chrome native messaging host installed successfully at %s", targetFile)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
