package cmd

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/toddlers/ghcli/gist"
	"github.com/toddlers/ghcli/templates"
	"github.com/toddlers/ghcli/utils"
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
		var report = template.Must(template.New("Gist Info").Funcs(template.FuncMap{"daysAgo": utils.DaysAgo}).Parse(templates.GistInfo))
		snippets, err := gist.GetGists(username)
		if err != nil {
			return err
		}
		if err := report.Execute(os.Stdout, snippets); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(gistCmd)
	gistCmd.Flags().StringVarP(&username, "username", "u", "", "github handle")
}
