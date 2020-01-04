package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/toddlers/ghcli/config"
)

var (
	username    string
	querystring string
	language    string
	Tag         bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github cli",
	Short: "search github",
	Long:  "search github for various things like repos, users etc.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
