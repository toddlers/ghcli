package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toddlers/ghcli/repos"
	"github.com/toddlers/ghcli/user"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search github",
	Long:  `search github for various information`,
	RunE:  search,
}

func search(cmd *cobra.Command, args []string) (err error) {
	if username != "" {
		userInfo := user.GetUser(username)
		fmt.Println(`Username:	`, userInfo.Login)
		fmt.Println(`Name:		`, userInfo.Name)
		fmt.Println(`Bio:		`, userInfo.Bio)
		fmt.Println(`Location:		`, userInfo.Location)
		fmt.Println(`Blog:		`, userInfo.Blog)
		fmt.Println(`Public Repos :		`, userInfo.PublicRepos)
		fmt.Println(`Public Gists :		`, userInfo.PublicGists)
		fmt.Println(`Followers : `, userInfo.Followers)
		fmt.Println(`Following : `, userInfo.Following)
		fmt.Println(`Last Updated :		`, userInfo.UpdatedAt)
		fmt.Println(`Created :		`, userInfo.CreatedAt)
		fmt.Println("")
		return nil
	} else {
		//docker+language:go&sort=stars&order=desc
		query := querystring + "language:" + language + "&sort=stars&order=desc"
		repos, err := repos.SearchRepos(query)
		if err != nil {
			return err
		}
		for _, repo := range repos.Items {
			fmt.Println(`Owner : `, repo.Owner.URL)
			fmt.Println(`Repo Name : `, repo.FullName)
			fmt.Println(`Repo URL : `, repo.HTMLURL)
			fmt.Println(`Description : `, repo.Description)
			fmt.Println(`Watchers : `, repo.Watchers)
			fmt.Println(`Forks : `, repo.Forks)
			fmt.Println(`Open Issues : `, repo.OpenIssuesCount)
			fmt.Println(`Created :		`, repo.CreatedAt)
			fmt.Println(`Last Updated :		`, repo.UpdatedAt)
			fmt.Println("============================================================")
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&username, "username", "u", "", `user's github handle`)
	searchCmd.Flags().StringVarP(&querystring, "query", "q", "", `string to search`)
	searchCmd.Flags().StringVarP(&language, "language", "l", "", `language project written in`)
}
