package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/toddlers/ghcli/config"
	"github.com/toddlers/ghcli/user"
)

var (
	username    string
	querystring string
	language    string
	Tag         bool
	Gid         string
)

var rootCmd = &cobra.Command{
	Use:   "github cli",
	Short: "search github",
	Long:  "search github for various things like repos, users etc.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			message, err := user.GetZen()
			if err != nil {
				return err
			}
			fmt.Printf("\nProgramming Zen: %s\n\n", string(message))
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	config.LoadConfig()
}
