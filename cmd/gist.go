package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toddlers/ghcli/gist"
	"github.com/toddlers/ghcli/models"
)

// gistCmd represents the gist command
var gistCmd = &cobra.Command{
	Use:   "gist",
	Short: "gist fetch",
	Long:  `show all public gists`,
	RunE:  getGists,
}

func getGists(cmd *cobra.Command, args []string) (err error) {
	if username != "" {
		snippets, err := gist.GetGists(username)
		if err != nil {
			return err
		}
		for _, snippet := range snippets {

			fmt.Println(`Description : `, snippet.Description)
			fmt.Println(`Github Handle : `, snippet.Owner.Login)
			var file models.File
			for _, file = range snippet.Files {
				fmt.Println(`Filename : `, file.Filename)
				fmt.Println(`Language : `, file.Language)
			}
			fmt.Println(`Created At : `, snippet.CreatedAt)
			fmt.Println("============================================================")
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(gistCmd)
	gistCmd.Flags().StringVarP(&username, "username", "u", "", "github handle")
}
