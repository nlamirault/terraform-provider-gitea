package gitea

import (
	"fmt"
	"testing"

	giteaapi "code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var testAccGiteaUserConfig = fmt.Sprintf(`
resource "gitea_user" "testuser" {
	login = "johndoe"
    password = "pass"
    username = "johndoe"
    fullname = "John Doe"
    email = "john.doe@gitea.io"
}
`)

func TestAccGiteaUser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGiteaUserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGiteaUserConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckGiteaUserExists("gitea_user.testuser", t),
				),
			},
		},
	})
}

func testCheckGiteaUserExists(n string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*giteaapi.Client)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		username := rs.Primary.Attributes["username"]

		if username == "" {
			return fmt.Errorf("No Username is set")
		}

		_, err := client.GetUserInfo(username)
		return err
	}
}

func testAccGiteaUserDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*giteaapi.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gitea_user" {
			continue
		}

		username := rs.Primary.Attributes["username"]
		_, err := client.GetUserInfo(username)

		if err == nil {
			return fmt.Errorf("User %s still exists", username)
		}
	}

	return nil
}
