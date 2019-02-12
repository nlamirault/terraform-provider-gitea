package gitea

import (
	"fmt"
	"log"

	giteaapi "code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGiteaRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceGiteaRepositoryCreate,
		Read:   resourceGiteaRepositoryRead,
		Update: resourceGiteaRepositoryUpdate,
		Delete: resourceGiteaRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auto_init": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"gitignores": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"license": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"readme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGiteaRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	username := d.Get("username").(string)
	options := giteaapi.CreateRepoOption{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Private:     d.Get("is_private").(bool),
		AutoInit:    d.Get("auto_init").(bool),
		Gitignores:  d.Get("gitignores").(string),
		License:     d.Get("license").(string),
		Readme:      d.Get("readme").(string),
	}

	log.Printf("[DEBUG] create gitea repository %q", options.Name)

	repository, err := client.AdminCreateRepo(username, options)
	if err != nil {
		return err
	}

	id := fmt.Sprintf("%d", repository.ID)
	d.SetId(id)
	return nil
}

func resourceGiteaRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGiteaRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGiteaRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	username := d.Get("username").(string)
	name := d.Get("name").(string)
	log.Printf("[DEBUG] delete gitea repository: %s %s", username, name)
	return client.DeleteRepo(username, name)
}
