package gitea

import (
	"fmt"
	"strings"

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
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("GITEA_BASE_URL", ""),
				Description:  descriptions["base_url"],
				ValidateFunc: validateAPIURLVersion,
			},
		},
		// ResourcesMap: map[string]*schema.Resource{
		// 	"gitea_project": resourceGiteaProject(),
		// 	"gitea_user":    resourceGiteaUser(),
		// },
		// DataSourcesMap: map[string]*schema.Resource{
		// 	"gitea_projects": dataSourceGiteaProjects(),
		// 	"gitea_project":  dataSourceGiteaProject(),
		// 	"gitea_users":    dataSourceGiteaUsers(),
		// 	"gitea_user":     dataSourceGiteaUser(),
		// },
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

func validateAPIURLVersion(value interface{}, key string) (ws []string, es []error) {
	v := value.(string)
	if !strings.HasSuffix(v, "/api/v1") || !strings.HasSuffix(v, "/api/v1/") {
		es = append(es, fmt.Errorf("The Gitea provider is only compatible with version 1 of the API: %s", v))
	}
	return
}
