package mxroute

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/jereksel/terraform-provider-mxroute/api"
	"testing"
)

func TestAccMxRouteEmailAccount_basic(t *testing.T) {

	domainName := TestDomainPrefix + generateRandomResourceName() + ".email"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMxRouteDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMxRouteEmailAccountConfigDomainWithEmail(domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMxRouteDomainContainsEmail(domainName),
				),
			},
			{
				Config: testAccCheckMxRouteEmailAccountConfigDomainWithoutEmail(domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMxRouteDomainDoesNotContainEmail(domainName),
				),
			},
		},
	})

}

func testAccCheckMxRouteDomainContainsEmail(domainName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		client := testAccProvider.Meta().(config)

		emails, err := api.GetEmailAccounts(client.Username, client.Password, domainName)
		if err != nil {
			return err
		}

		if len(emails) == 0 {
			return fmt.Errorf("domain %s does not contain emails", domainName)
		}

		return nil
	}
}

func testAccCheckMxRouteDomainDoesNotContainEmail(domainName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		client := testAccProvider.Meta().(config)

		emails, err := api.GetEmailAccounts(client.Username, client.Password, domainName)
		if err != nil {
			return err
		}

		if len(emails) > 0 {
			return fmt.Errorf("domain %s contains emails %v", domainName, emails)
		}

		return nil
	}

}

func testAccCheckMxRouteEmailAccountConfigDomainWithEmail(domainName string) string {
	return fmt.Sprintf(`
		
		resource "mxroute_domain" "foobar" {
			name = "%s"
		}
	
		resource "mxroute_email_account" "email" {
			domain = mxroute_domain.foobar.name
			username = "email"
			password = "password"
		}

`, domainName)
}

func testAccCheckMxRouteEmailAccountConfigDomainWithoutEmail(domainName string) string {
	return fmt.Sprintf(`
		
		resource "mxroute_domain" "foobar" {
			name = "%s"
		}

`, domainName)
}
