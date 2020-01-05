package config

type GistConfig struct {
	FileName    string
	AccessToken string
	GistID      string
	Public      bool
	AutoSync    bool
}

var Gc GistConfig

func LoadConfig() {
	Gc.Public = false
}
