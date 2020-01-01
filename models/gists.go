package models

import "time"

type File struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Language string `json:"language"`
	RawURL   string `json:"raw_url"`
	Size     int    `json:"size"`
}
type Gist struct {
	URL         string    `json:"url"`
	ForksURL    string    `json:"forks_url"`
	CommitsURL  string    `json:"commits_url"`
	GitPullURL  string    `json:"git_pull_url"`
	GitPushURL  string    `json:"git_push_url"`
	HTMLURL     string    `json:"html_url"`
	Public      bool      `json:"public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Comments    int       `json:"comments"`
	CommentsURL string    `json:"comments_url"`
	Owner       *Owner
	Files       map[string]File
}

type Snippets struct {
	Snippets []SnippetInfo
}

type SnippetInfo struct {
	Description string
	Command     string
	Tag         []string
	Output      string
}
