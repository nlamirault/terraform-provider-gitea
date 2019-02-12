package gitea

import (
	"code.gitea.io/sdk/gitea"
)

// Config is per-provider, specifies where to connect to Gitea
type Config struct {
	Token   string
	BaseURL string
}

// Client returns a *gitea.Client to interact with the configured Gitea instance
func (c *Config) Client() interface{} {
	return gitea.NewClient(c.BaseURL, c.Token)
}
