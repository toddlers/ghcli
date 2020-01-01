package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/chzyer/readline"
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

var gistCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create gist",
	Long:  `create a gist`,
	RunE:  createGists,
}

func scan(message string) (string, error) {
	tempFile := "/tmp/tempgist.tmp"
	l, err := readline.NewEx(&readline.Config{
		Prompt:            message,
		HistoryFile:       tempFile,
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
	})

	if err != nil {
		return "", err
	}

	defer l.Close()

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		return line, nil
	}
	return "", errors.New("Canceled")
}

func createGists(cmd *cobra.Command, args []string) (err error) {
	var tags []string
	if Tag {
		var t string
		if t, err = scan("Tag> "); err != nil {
			return err
		}
		tags = strings.Fields(t)
	}
	fmt.Println(tags)
	fmt.Println("Craete Gists called")
	return nil
}

func init() {
	gistCmd.Flags().StringVarP(&username, "username", "u", "", "github handle")
	gistCreateCmd.Flags().BoolVarP(&Tag, "tag", "t", false,
		`Display tag prompt (delimiter: space)`)
	//gistCreateCmd.Flags().StringVarP(&tag, "tag", "t", "", "tag for the gist")
	gistCmd.AddCommand(gistCreateCmd)
	rootCmd.AddCommand(gistCmd)
}
