package templates

const (
	UserInfo = `Username: {{.Login}}
Name: {{.Name}}
Bio:   {{.Bio}}
Location:  {{.Location }}
Age:    {{.CreatedAt | daysAgo}} days
Blog: {{.Blog}}
Freshness: {{.UpdatedAt | daysAgo }} days
Public Repos:  {{.PublicRepos}}
Public Gists:  {{.PublicGists}}
Followers:  {{.Followers}}
Following: {{.Following}}
`

	GistInfo = `
{{range .}}----------------------------------------
Description : {{.Description}}
Github Handle: {{.Owner.Login}}
Gist ID: {{.ID}}
{{range .Files}}
Filename:   {{.Filename}}
Language:  {{.Language}}
{{end}}
{{end}}`

	RepoInfo = `
{{range .Items}}-------------------------------------------------------------------------
Owner :  {{.Owner.URL}}
Repo Name : {{.FullName}}
Repo URL :  {{.HTMLURL}}
Description :  {{.Description}}
Forks :  {{.ForksCount}}
Open Issues :  {{.OpenIssuesCount}}
Created :           {{.CreatedAt}}
Last Updated :              {{.UpdatedAt}}
{{end}}`
)
