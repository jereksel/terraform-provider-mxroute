package mxroute

import (
	"github.com/jereksel/terraform-provider-mxroute/api"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainCreate,
		Read:   resourceDomainRead,
		Delete: resourceDomainDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dkim": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDomainCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(config)
	domainName := d.Get("name").(string)
	if err := api.CreateDomain(config.Username, config.Password, domainName); err != nil {
		return err
	}
	return resourceDomainRead(d, m)
}

func resourceDomainRead(d *schema.ResourceData, m interface{}) error {
	config := m.(config)
	domainName := d.Get("name").(string)
	allDomains, err := api.GetAllDomains(config.Username, config.Password)
	if err != nil {
		return err
	}
	exists := false
	for _, domain := range allDomains {
		if domain == domainName {
			exists = true
		}
	}
	if exists {
		d.SetId(domainName)
	} else {
		d.SetId("")
	}
	return nil
}

func resourceDomainDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(config)
	domainName := d.Get("name").(string)
	return api.RemoveDomain(config.Username, config.Password, domainName)
}
