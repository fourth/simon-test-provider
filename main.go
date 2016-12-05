package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SJJ_USER", nil),
				Description: "The user name",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"sjj_test": resourceTestResource(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func resourceTestResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceTestUpdate,
		Update: resourceTestUpdate,
		Read:   resourceTestRead,
		Delete: resourceTestDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceTestUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] write")
	err := ioutil.WriteFile(d.Get("name").(string), []byte(d.Get("content").(string)), 0600)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))
	d.Set("size", len(d.Get("content").(string)))
	return nil
}

func resourceTestRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] read")
	b, err := ioutil.ReadFile(d.Get("name").(string))
	if err != nil {
		return err
	}
	d.Set("content", string(b))
	return nil
}

func resourceTestDelete(d *schema.ResourceData, meta interface{}) error {
	return os.Remove(d.Get("name").(string))
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &config{
		user: d.Get("user").(string),
	}, nil
}

type config struct {
	user string
}

func main() {
	log.Println("hi")
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider,
	})
}
