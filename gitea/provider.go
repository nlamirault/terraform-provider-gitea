package gitea

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	ENV_GITEA_BASE_URL = "GITEA_BASE_URL"
	ENV_GITEA_TOKEN    = "GITEA_TOKEN"
	descriptions       map[string]string
)

func init() {
	descriptions = map[string]string{
		"token":    "The token used to connect to Gitea.",
		"base_url": "The Gitea Base API URL",
	}
}

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(ENV_GITEA_TOKEN, nil),
				Description: descriptions["token"],
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(ENV_GITEA_BASE_URL, ""),
				Description: descriptions["base_url"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"gitea_organization": resourceGiteaOrganization(),
			"gitea_user":         resourceGiteaUser(),
			"gitea_repository":   resourceGiteaRepository(),
			"gitea_label":        resourceGiteaLabel(),
			"gitea_milestone":    resourceGiteaMilestone(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"gitea_user":       dataSourceGiteaUser(),
			"gitea_repository": dataSourceGiteaRepository(),
		},
		ConfigureFunc: providerConfigure,
	}

}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   d.Get("token").(string),
		BaseURL: d.Get("base_url").(string),
	}
	return config.Client(), nil
}
