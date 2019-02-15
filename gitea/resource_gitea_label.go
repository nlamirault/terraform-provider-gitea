package gitea

import (
	"log"
	"strconv"

	giteaapi "code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGiteaLabel() *schema.Resource {
	return &schema.Resource{
		Create: resourceGiteaLabelCreate,
		Read:   resourceGiteaLabelRead,
		Update: resourceGiteaLabelUpdate,
		Delete: resourceGiteaLabelDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"color": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGiteaLabelSetToState(d *schema.ResourceData, label *giteaapi.Label) error {
	if err := d.Set("name", label.Name); err != nil {
		return err
	}
	if err := d.Set("color", label.Color); err != nil {
		return err
	}
	return nil
}

func resourceGiteaLabelCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	owner := d.Get("owner").(string)
	repository := d.Get("repository").(string)
	options := giteaapi.CreateLabelOption{
		Name:  d.Get("name").(string),
		Color: d.Get("color").(string),
	}

	log.Printf("[DEBUG] create label: %s %s %v", owner, repository, options)

	label, err := client.CreateLabel(owner, repository, options)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] label created %v", label)
	d.SetId(strconv.FormatInt(label.ID, 10))
	return resourceGiteaLabelRead(d, meta)
}

func resourceGiteaLabelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	log.Printf("[DEBUG] Label informations: %s", d.Id())
	labelId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	owner := d.Get("owner").(string)
	repository := d.Get("repository").(string)
	log.Printf("[DEBUG] read label %q", labelId)

	label, err := client.GetRepoLabel(owner, repository, labelId)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] label find %v", label)
	return resourceGiteaLabelSetToState(d, label)
}

func resourceGiteaLabelUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*giteaapi.Client)
	// id := d.Get("id").(int64)
	// owner := d.Get("owner").(string)
	// name := d.Get("name").(string)
	// options := giteaapi.EditLabelOption{
	// 	Name:  d.Get("name").(string),
	// 	Color: d.Get("color").(string),
	// }
	// log.Printf("[DEBUG] edit gitea label: %q %q", id, options)
	// label, err := client.EditLabel(owner, name, id, options)
	// if err != nil {
	// 	return err
	// }
	// return resourceGiteaLabelSetToState(d, label)
	return nil
}

func resourceGiteaLabelDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*giteaapi.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	owner := d.Get("owner").(string)
	repository := d.Get("repository").(string)
	log.Printf("[DEBUG] delete label: %d %s %s", id, owner, repository)
	return client.DeleteLabel(owner, repository, id)
}
