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
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// "fullname": {
			// 	Type:     schema.TypeString,
			// 	Required: true,
			// },
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

func resourceGiteaRepositorySetToState(d *schema.ResourceData, repo *giteaapi.Repository) error {
	if err := d.Set("owner", repo.Owner.UserName); err != nil {
		return err
	}
	if err := d.Set("name", repo.Name); err != nil {
		return err
	}
	if err := d.Set("description", repo.Description); err != nil {
		return err
	}
	return nil
}

func resourceGiteaRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	owner := d.Get("owner").(string)
	options := giteaapi.CreateRepoOption{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Private:     d.Get("is_private").(bool),
		AutoInit:    d.Get("auto_init").(bool),
		Gitignores:  d.Get("gitignores").(string),
		License:     d.Get("license").(string),
		Readme:      d.Get("readme").(string),
	}

	log.Printf("[DEBUG] create repository %s", options.Name)

	repository, err := client.AdminCreateRepo(owner, options)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Repository created: %v", repository)
	d.SetId(fmt.Sprintf("%d", repository.ID))
	return nil
}

func resourceGiteaRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	owner := d.Get("owner").(string)
	name := d.Get("name").(string)
	log.Printf("[DEBUG] read repository %q %s %s", d.Id(), owner, name)
	repo, err := client.GetRepo(owner, name)
	if err != nil {
		return fmt.Errorf("unable to retrieve repository %s %s", owner, name)
	}
	log.Printf("[DEBUG] repository find: %v", repo)
	return resourceGiteaRepositorySetToState(d, repo)
}

func resourceGiteaRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGiteaRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	owner := d.Get("owner").(string)
	name := d.Get("name").(string)
	log.Printf("[DEBUG] delete repository: %s %s", owner, name)
	return client.DeleteRepo(owner, name)
}
