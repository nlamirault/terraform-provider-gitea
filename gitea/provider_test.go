package gitea

import (
	"os"
	"testing"

	giteaapi "code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testAccGiteaClient *giteaapi.Client

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"gitea": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv(ENV_GITEA_BASE_URL); v == "" {
		t.Fatalf("%s must be set for acceptance tests", ENV_GITEA_BASE_URL)
	}
	if v := os.Getenv(ENV_GITEA_TOKEN); v == "" {
		t.Fatalf("%s must be set for acceptance tests", ENV_GITEA_TOKEN)
	}
	if testAccGiteaClient == nil {
		config := Config{
			BaseURL: os.Getenv(ENV_GITEA_BASE_URL),
			Token:   os.Getenv(ENV_GITEA_TOKEN),
		}

		testAccGiteaClient = config.Client().(*giteaapi.Client)
	}
}
