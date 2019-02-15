package gitea

import (
	"log"
	"strconv"

	giteaapi "code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGiteaMilestone() *schema.Resource {
	return &schema.Resource{
		Create: resourceGiteaMilestoneCreate,
		Read:   resourceGiteaMilestoneRead,
		Update: resourceGiteaMilestoneUpdate,
		Delete: resourceGiteaMilestoneDelete,
		Schema: map[string]*schema.Schema{
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGiteaMilestoneSetToState(d *schema.ResourceData, milestone *giteaapi.Milestone) error {
	if err := d.Set("title", milestone.Title); err != nil {
		return err
	}
	if err := d.Set("description", milestone.Description); err != nil {
		return err
	}
	return nil
}

func resourceGiteaMilestoneCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	owner := d.Get("owner").(string)
	repository := d.Get("repository").(string)
	options := giteaapi.CreateMilestoneOption{
		Title:       d.Get("title").(string),
		Description: d.Get("description").(string),
	}

	log.Printf("[DEBUG] create milestone: %s %s %v", owner, repository, options)

	milestone, err := client.CreateMilestone(owner, repository, options)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] milestone created %v", milestone)
	d.SetId(strconv.FormatInt(milestone.ID, 10))
	return resourceGiteaMilestoneRead(d, meta)
}

func resourceGiteaMilestoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	log.Printf("[DEBUG] milestone informations: %s", d.Id())
	milestoneId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	owner := d.Get("owner").(string)
	repository := d.Get("repository").(string)
	log.Printf("[DEBUG] read milestone %q", milestoneId)

	milestone, err := client.GetMilestone(owner, repository, milestoneId)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] milestone find %v", milestone)
	return resourceGiteaMilestoneSetToState(d, milestone)
}

func resourceGiteaMilestoneUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*giteaapi.Client)
	// id := d.Get("id").(int64)
	// owner := d.Get("owner").(string)
	// name := d.Get("name").(string)
	// options := giteaapi.EditMilestoneOption{
	// 	Name:  d.Get("name").(string),
	// 	Color: d.Get("color").(string),
	// }
	// log.Printf("[DEBUG] edit gitea Milestone: %q %q", id, options)
	// Milestone, err := client.EditMilestone(owner, name, id, options)
	// if err != nil {
	// 	return err
	// }
	// return resourceGiteaMilestoneSetToState(d, Milestone)
	return nil
}

func resourceGiteaMilestoneDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	owner := d.Get("owner").(string)
	repository := d.Get("repository").(string)
	log.Printf("[DEBUG] delete milestone: %d %s %s", id, owner, repository)
	return client.DeleteMilestone(owner, repository, id)
}
