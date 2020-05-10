package mxroute

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/jereksel/terraform-provider-mxroute/api"
)

func resourceEmailAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailAccountCreate,
		Read:   resourceEmailAccountRead,
		Update: resourceEmailAccountUpdate,
		Delete: resourceEmailAccountDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
		},
	}
}

func resourceEmailAccountCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(config)
	domainName := d.Get("domain").(string)
	emailUsername := d.Get("username").(string)
	emailPassword := d.Get("password").(string)

	if err := api.CreateEmailAccount(config.Username, config.Password, domainName, emailUsername, emailPassword); err != nil {
		return err
	}

	return resourceEmailAccountRead(d, m)
}

func resourceEmailAccountRead(d *schema.ResourceData, m interface{}) error {
	config := m.(config)
	domainName := d.Get("domain").(string)
	emailUsername := d.Get("username").(string)

	emailExists, err := api.DoesEmailAccountExists(config.Username, config.Password, domainName, emailUsername)
	if err != nil {
		return err
	}

	if *emailExists {
		d.SetId(emailUsername)
	} else {
		d.SetId("")
	}

	return nil

}

func resourceEmailAccountUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(config)
	domainName := d.Get("domain").(string)

	//We can change password while changing username without additional code
	if (d.HasChange("username") && d.HasChange("password")) || d.HasChange("username") {
		iOldUsername, iNewUsername := d.GetChange("username")

		oldUsername := iOldUsername.(string)
		newUsername := iNewUsername.(string)

		password := d.Get("password").(string)

		if err := api.ChangeEmailAccountUsername(config.Username, config.Password, domainName, password, oldUsername, newUsername); err != nil {
			return err
		}

	} else if d.HasChange("password") {
		username := d.Get("username").(string)
		newPassword := d.Get("password").(string)

		if err := api.ChangeEmailAccountPassword(config.Username, config.Password, domainName, username, newPassword); err != nil {
			return err
		}
	}

	return resourceEmailAccountRead(d, m)
}

func resourceEmailAccountDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(config)
	domainName := d.Get("domain").(string)
	emailUsername := d.Get("username").(string)
	return api.RemoveEmailAccount(config.Username, config.Password, domainName, emailUsername)
}

func resourceEmailAccountImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	return []*schema.ResourceData{d}, nil
}
