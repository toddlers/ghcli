package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/toddlers/ghcli/templates"
	"github.com/toddlers/ghcli/user"
	"github.com/toddlers/ghcli/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list starred repos",
	Long:  `list starred repos by the user`,
	RunE:  listStars,
}

func listStars(cmd *cobra.Command, args []string) (err error) {
	status, _ := cmd.Flags().GetBool("Gists")
	if status {
		fmt.Println("listing starred gists")
	}
	repos, err := user.GetStars(username)
	if err != nil {
		return err
	}
	var report = template.Must(template.New("Repo Info").Funcs(template.FuncMap{"daysAgo": utils.DaysAgo}).Parse(templates.RepoInfo))
	if err := report.Execute(os.Stdout, repos); err != nil {
		log.Fatal(err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&username, "username", "u", "", `user's github handle`)
	listCmd.Flags().BoolP("Gists", "g", false, "list starred gists")
}
