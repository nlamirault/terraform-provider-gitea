package gitea

// import (
// 	"fmt"
// 	"testing"

// 	giteaapi "code.gitea.io/sdk/gitea"
// 	"github.com/hashicorp/terraform/helper/resource"
// 	"github.com/hashicorp/terraform/terraform"
// )

// var testAccGiteaRepositoryConfig = fmt.Sprintf(`
// resource "gitea_repository" "test" {
//     owner = "test"
// 	name = "resourcetest"
// 	description = "Just a test"
// }`)

// func TestAccGiteaRepository_basic(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccGiteaRepositoryDestroy,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: testAccGiteaRepositoryConfig,
// 				Check: resource.ComposeTestCheckFunc(
// 					testCheckGiteaRepositoryExists("gitea_repository.test", t),
// 				),
// 			},
// 		},
// 	})
// }

// func testCheckGiteaRepositoryExists(n string, t *testing.T) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		client := testAccProvider.Meta().(*giteaapi.Client)

// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", n)
// 		}

// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("No ID is set")
// 		}

// 		username := rs.Primary.Attributes["username"]
// 		if username == "" {
// 			return fmt.Errorf("No Username is set")
// 		}

// 		name := rs.Primary.Attributes["name"]
// 		if name == "" {
// 			return fmt.Errorf("No name is set")
// 		}

// 		_, err := client.GetRepo(username, name)
// 		return err
// 	}
// }

// func testAccGiteaRepositoryDestroy(s *terraform.State) error {
// 	client := testAccProvider.Meta().(*giteaapi.Client)

// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "gitea_repository" {
// 			continue
// 		}

// 		username := rs.Primary.Attributes["username"]
// 		name := rs.Primary.Attributes["name"]
// 		_, err := client.GetRepo(username, name)

// 		if err == nil {
// 			return fmt.Errorf("Repository %s/%s still exists", username, name)
// 		}
// 	}

// 	return nil
// }
