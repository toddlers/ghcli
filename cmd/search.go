package cmd

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/toddlers/ghcli/repos"
	"github.com/toddlers/ghcli/templates"
	"github.com/toddlers/ghcli/user"
	"github.com/toddlers/ghcli/utils"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search github",
	Long:  `search github for various information`,
}
var repoSearchCmd = &cobra.Command{
	Use:   "repo",
	Short: "search repo",
	Long:  `search repo on github for various information`,
	RunE:  repoSearch,
}
var userSearchCmd = &cobra.Command{
	Use:   "user",
	Short: "search user",
	Long:  `search user on github for various information`,
	RunE:  userSearch,
}

func repoSearch(cmd *cobra.Command, args []string) (err error) {
	//docker+language:go&sort=stars&order=desc
	query := querystring + "language:" + language + "&sort=stars&order=desc"
	repos, err := repos.SearchRepos(query)
	if err != nil {
		return err
	}
	var report = template.Must(template.New("Repo Info").Funcs(template.FuncMap{"daysAgo": utils.DaysAgo}).Parse(templates.RepoInfo))
	if err := report.Execute(os.Stdout, repos); err != nil {
		log.Fatal(err)
	}
	return nil
}

func userSearch(cmd *cobra.Command, args []string) (err error) {
	if username != "" {
		userInfo := user.GetUser(username)

		var report = template.Must(template.New("User Info").Funcs(template.FuncMap{"daysAgo": utils.DaysAgo}).Parse(templates.UserInfo))
		if err := report.Execute(os.Stdout, userInfo); err != nil {
			log.Fatal(err)
		}
		return nil
	}
	return nil
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.AddCommand(repoSearchCmd)
	searchCmd.AddCommand(userSearchCmd)
	userSearchCmd.Flags().StringVarP(&username, "username", "u", "", `user's github handle`)
	repoSearchCmd.Flags().StringVarP(&querystring, "query", "q", "", `string to search`)
	repoSearchCmd.Flags().StringVarP(&language, "language", "l", "", `language project written in`)
}
