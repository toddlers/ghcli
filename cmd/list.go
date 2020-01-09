package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/toddlers/ghcli/gist"
	"github.com/toddlers/ghcli/templates"
	"github.com/toddlers/ghcli/user"
	"github.com/toddlers/ghcli/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list-stars",
	Short: "list starred repositories",
	Long:  `list starred repositories and gists by an authenticated user`,
}
var repoListCmd = &cobra.Command{
	Use:   "repos",
	Short: "list starred repositories",
	Long:  `list starred repositories by an authenticated user`,
	RunE:  repoListStars,
}
var gistsListCmd = &cobra.Command{
	Use:   "gists",
	Short: "list starred gists",
	Long:  `list starred gists by an authenticated user`,
	RunE:  gistListStars,
}

func repoListStars(cmd *cobra.Command, args []string) (err error) {
	if username != "" {
		fmt.Println("listing starred repositories")
		repos, err := user.GetStarredRepos(username)
		if err != nil {
			return err
		}
		var report = template.Must(template.New("Repo Info").Funcs(template.FuncMap{"daysAgo": utils.DaysAgo}).Parse(templates.RepoInfo))
		if err := report.Execute(os.Stdout, repos); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Username is mandatory flag for listing the starred repositories")
	}
	return nil
}

func gistListStars(cmd *cobra.Command, args []string) (err error) {
	fmt.Println("listing starred gists")
	var report = template.Must(template.New("Gist Info").Funcs(template.FuncMap{"daysAgo": utils.DaysAgo}).Parse(templates.GistInfo))
	snippets, err := gist.GetGists(username)
	if err != nil {
		return err
	}
	if err := report.Execute(os.Stdout, snippets); err != nil {
		log.Fatal(err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(repoListCmd)
	listCmd.AddCommand(gistsListCmd)
	repoListCmd.Flags().StringVarP(&username, "username", "u", "", `user's github handle`)
	gistsListCmd.Flags().StringVarP(&username, "username", "u", "", `user's github handle`)
}
