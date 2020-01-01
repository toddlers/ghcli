package models

import "time"

type User struct {
	Login       string      `json:"login"`
	HTMLURL     string      `json:"html_url"`
	Name        string      `json:"name"`
	Company     string      `json:"company,omitempty"`
	Blog        string      `json:"blog"`
	Location    string      `json:"location"`
	Email       string      `json:"email,omitempty"`
	Hireable    interface{} `json:"hireable"`
	Bio         string      `json:"bio,omitempty"`
	PublicRepos int         `json:"public_repos"`
	PublicGists int         `json:"public_gists"`
	Followers   int         `json:"followers"`
	Following   int         `json:"following"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
