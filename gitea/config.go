package gitea

import (
	"log"

	"code.gitea.io/sdk/gitea"
)

// Config is per-provider, specifies where to connect to Gitea
type Config struct {
	Token   string
	BaseURL string
}

// Client returns a *gitea.Client to interact with the configured Gitea instance
func (c *Config) Client() interface{} {
	log.Printf("[DEBUG] Create client using configuration : %v", c)
	return gitea.NewClient(c.BaseURL, c.Token)
}
