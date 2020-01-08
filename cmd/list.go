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
	Long: `list starred repositories by an authenticated user
	        by default, for listing starred gists please use
					the option "-g"`,
	RunE: listStars,
}

func listStars(cmd *cobra.Command, args []string) (err error) {
	status, _ := cmd.Flags().GetBool("Gists")
	if status {
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

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&username, "username", "u", "", `user's github handle`)
	listCmd.Flags().BoolP("Gists", "g", false, "list starred gists")
}
