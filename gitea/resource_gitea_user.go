package gitea

import (
	"fmt"
	"log"

	giteaapi "code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGiteaUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceGiteaUserCreate,
		Read:   resourceGiteaUserRead,
		Update: resourceGiteaUserUpdate,
		Delete: resourceGiteaUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"gitea_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"avatar_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_admin": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceGiteaUserSetToState(d *schema.ResourceData, user *giteaapi.User) error {
	if err := d.Set("username", user.UserName); err != nil {
		return err
	}
	if err := d.Set("full_name", user.FullName); err != nil {
		return err
	}
	return nil
}

func resourceGiteaUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	options := giteaapi.CreateUserOption{
		Email:      d.Get("email").(string),
		FullName:   d.Get("full_name").(string),
		LoginName:  d.Get("login").(string),
		Password:   d.Get("password").(string),
		SendNotify: false,
		Username:   d.Get("username").(string),
	}

	log.Printf("[DEBUG] create gitea user %q", options.Username)

	user, err := client.AdminCreateUser(options)
	if err != nil {
		return fmt.Errorf("unable to create user: %v", err)
	}

	if d.Get("is_admin").(bool) {
		return resourceGiteaUserUpdate(d, meta)
	}

	return resourceGiteaUserRead(d, user)
}

func resourceGiteaUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	log.Printf("[DEBUG] read gitea user %s", d.Id())
	username := d.Get("username").(string)
	user, err := client.GetUserInfo(username)
	if err != nil {
		return fmt.Errorf("unable to retrieve user %s", username)
	}

	return resourceGiteaUserSetToState(d, user)

}

func resourceGiteaUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	log.Printf("[DEBUG] update gitea user %s", d.Id())
	isAdmin := d.Get("is_admin").(bool)
	username := d.Get("username").(string)
	edit := giteaapi.EditUserOption{
		Admin:     &isAdmin,
		Email:     d.Get("email").(string),
		FullName:  d.Get("full_name").(string),
		LoginName: d.Get("login").(string),
		Password:  d.Get("password").(string),
	}

	err := client.AdminEditUser(username, edit)
	if err != nil {
		return fmt.Errorf("unable to edit user: %s", username)
	}
	return nil
}

func resourceGiteaUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	log.Printf("[DEBUG] delete gitea user %s", d.Id())
	return client.AdminDeleteUser(d.Get("username").(string))
}
