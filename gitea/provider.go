package gitea

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var descriptions map[string]string

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
				DefaultFunc: schema.EnvDefaultFunc("GITEA_TOKEN", nil),
				Description: descriptions["token"],
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITEA_BASE_URL", ""),
				Description: descriptions["base_url"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"gitea_user": resourceGiteaUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"gitea_user": dataSourceGiteaUser(),
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
