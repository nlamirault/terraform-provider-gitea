package gitea

import (
	"fmt"
	"log"

	giteaapi "code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGiteaOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceGiteaOrganizationCreate,
		Read:   resourceGiteaOrganizationRead,
		Update: resourceGiteaOrganizationUpdate,
		Delete: resourceGiteaOrganizationDelete,
		Schema: map[string]*schema.Schema{
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fullname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"website": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceGiteaOrganizationSetToState(d *schema.ResourceData, org *giteaapi.Organization) error {
	if err := d.Set("username", org.UserName); err != nil {
		return err
	}
	if err := d.Set("fullname", org.UserName); err != nil {
		return err
	}
	if err := d.Set("description", org.Description); err != nil {
		return err
	}
	if err := d.Set("website", org.Website); err != nil {
		return err
	}
	return nil
}

func resourceGiteaOrganizationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	owner := d.Get("owner").(string)
	options := giteaapi.CreateOrgOption{
		UserName:    d.Get("username").(string),
		FullName:    d.Get("fullname").(string),
		Description: d.Get("description").(string),
		Website:     d.Get("website").(string),
	}

	log.Printf("[DEBUG] create user %q", options.UserName)

	org, err := client.AdminCreateOrg(owner, options)
	if err != nil {
		return fmt.Errorf("unable to create organization: %v", err)
	}
	log.Printf("[DEBUG] organization created: %v", org)
	d.SetId(fmt.Sprintf("%d", org.ID))
	return resourceGiteaUserRead(d, meta)
}

func resourceGiteaOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	username := d.Get("username").(string)
	log.Printf("[DEBUG] read organization %q %s", d.Id(), username)
	org, err := client.GetOrg(username)
	if err != nil {
		return fmt.Errorf("unable to retrieve organization %s", username)
	}
	log.Printf("[DEBUG] organization find: %v", org)
	return resourceGiteaOrganizationSetToState(d, org)

}

func resourceGiteaOrganizationUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGiteaOrganizationDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*giteaapi.Client)
	// log.Printf("[DEBUG] delete organization %s", d.Id())
	// return client.AdminDeleteOrganization(d.Get("username").(string))
	return nil
}
