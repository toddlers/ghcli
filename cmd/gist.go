package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/toddlers/ghcli/gist"
	"github.com/toddlers/ghcli/templates"
	"github.com/toddlers/ghcli/utils"
)

// gistCmd represents the gist command
var gistCmd = &cobra.Command{
	Use:   "gist",
	Short: "gist list",
	Long:  `list all public gists`,
	RunE:  listGists,
}

var gistCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create gist",
	Long:  `create a gist`,
	RunE:  createGists,
}

var gistDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download gist",
	Long:  `download a public gist`,
	RunE:  downloadGist,
}

func listGists(cmd *cobra.Command, args []string) (err error) {
	if len(username) > 0 {
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

func createGists(cmd *cobra.Command, args []string) (err error) {
	var command string
	var description string
	var tags []string
	if len(args) > 0 {
		command = strings.Join(args, " ")
		fmt.Fprintf(color.Output, "%s %s\n", color.YellowString("Command>"), command)
	} else {
		command, err = utils.Scan(color.YellowString("Command>"))
		if err != nil {
			return err
		}
	}

	description, err = utils.Scan(color.GreenString("Description> "))
	if err != nil {
		return err
	}
	if Tag {
		var t string
		if t, err = utils.Scan("Tag> "); err != nil {
			return err
		}
		tags = strings.Fields(t)
	}
	newSnippet := gist.SnippetInfo{
		Description: description,
		Command:     command,
		Tag:         tags,
	}
	body, err := json.Marshal(newSnippet)
	if err != nil {
		return err
	}
	err = gist.GistUpload(string(body))
	if err != nil {
		return err
	}
	fmt.Println(tags)
	fmt.Println("Create Gists called")
	return nil
}

func downloadGist(cmd *cobra.Command, args []string) (err error) {
	if Gid != "" {
		err := gist.GistDownload(Gid)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	gistCmd.Flags().StringVarP(&username, "username", "u", "", "github handle")
	gistDownloadCmd.Flags().StringVarP(&Gid, "gid", "i", "", "gist id to download")
	gistCreateCmd.Flags().BoolVarP(&Tag, "tag", "t", false,
		`Display tag prompt (delimiter: space)`)
	gistCmd.AddCommand(gistCreateCmd)
	gistCmd.AddCommand(gistDownloadCmd)
	rootCmd.AddCommand(gistCmd)
}
