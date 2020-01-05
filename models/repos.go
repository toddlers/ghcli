package models

import "time"

type ReposSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Repo
}

type Owner struct {
	URL   string `json:"url"`
	Login string `json:"login"`
}

type Repo struct {
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Private         bool   `json:"private"`
	Owner           *Owner
	HTMLURL         string    `json:"html_url"`
	Description     string    `json:"description"`
	Fork            bool      `json:"fork"`
	URL             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
	GitURL          string    `json:"git_url"`
	SSHURL          string    `json:"ssh_url"`
	CloneURL        string    `json:"clone_url"`
	Homepage        string    `json:"homepage"`
	Size            int       `json:"size"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Language        string    `json:"language"`
	HasIssues       bool      `json:"has_issues"`
	HasProjects     bool      `json:"has_projects"`
	HasDownloads    bool      `json:"has_downloads"`
	HasWiki         bool      `json:"has_wiki"`
	HasPages        bool      `json:"has_pages"`
	ForksCount      int       `json:"forks_count"`
	Archived        bool      `json:"archived"`
	Disabled        bool      `json:"disabled"`
	OpenIssuesCount int       `json:"open_issues_count"`
	Forks           int       `json:"forks"`
	OpenIssues      int       `json:"open_issues"`
	Watchers        int       `json:"watchers"`
	DefaultBranch   string    `json:"default_branch"`
	Score           float64   `json:"score"`
}
