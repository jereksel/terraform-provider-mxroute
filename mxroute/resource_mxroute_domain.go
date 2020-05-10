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

		Importer: &schema.ResourceImporter{
			State: resourceDomainImport,
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
	doesDomainExists, err := api.DoesDomainExist(config.Username, config.Password, domainName)
	if err != nil {
		return err
	}
	if *doesDomainExists {
		dkim, err := api.GetDomainDkim(config.Username, config.Password, domainName)
		if err != nil {
			return err
		}
		d.SetId(domainName)
		if err := d.Set("dkim", dkim); err != nil {
			return err
		}
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

func resourceDomainImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	if err := d.Set("name", d.Id()); err != nil {
		return nil, err
	}

	if err := resourceDomainRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
