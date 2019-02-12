package gitea

import (
	"fmt"
	"log"
	"strings"

	giteaapi "code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGiteaUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGiteaUserRead,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fullname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"avatar_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGiteaUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)

	log.Printf("[INFO] Reading Gitea user")

	userName := strings.ToLower(d.Get("username").(string))
	user, err := client.GetUserInfo(userName)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d", user.ID))
	d.Set("username", user.UserName)
	d.Set("fullname", user.FullName)
	d.Set("email", user.Email)
	d.Set("avatar_url", user.AvatarURL)
	return nil
}
