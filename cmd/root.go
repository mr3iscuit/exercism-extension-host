package cmd

import (
	"log"
	"os"

	"github.com/mr3iscuit/exercism-extension-host/router"
	native "github.com/rickypc/native-messaging-host"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "extension-host",
	Short: "A native messaging host for browser extensions",
	Long:  `A native messaging host for browser extensions. Use --help for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		host := (&native.Host{
			AppName: "com.exercism.extension",
		}).Init()

		r := router.NewRouter()

		r.On("text", func(data *native.H) (*native.H, error) {
			text, ok := (*data)["text"].(string)
			if !ok {
				return &native.H{
					"error": "invalid text data",
				}, nil
			}
			return &native.H{
				"text": text,
			}, nil
		})

		r.On("ide", func(data *native.H) (*native.H, error) {
			return &native.H{}, nil
		})

		request := &native.H{}
		if err := host.OnMessage(os.Stdin, request); err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}

		response, err := r.HandleMessage(request)
		if err != nil {
			log.Fatalf("Failed to handle message: %v", err)
		}

		if err := host.PostMessage(os.Stdout, response); err != nil {
			log.Fatalf("Failed to send response: %v", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for extension-host")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}
