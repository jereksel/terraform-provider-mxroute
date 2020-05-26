package mxroute

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MXROUTE_USERNAME", nil),
				Description: "Username for DirectAdmin (NOT MXROUTE) account",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MXROUTE_PASSWORD", nil),
				Description: "Password/Login Key for DirectAdmin (NOT MXROUTE) account",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mxroute_domain":        resourceDomain(),
			"mxroute_email_account": resourceEmailAccount(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			username := d.Get("username").(string)
			password := d.Get("password").(string)

			config := config{
				Username: username,
				Password: password,
			}

			return config, nil
		},
	}
}
