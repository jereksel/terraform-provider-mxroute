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
	emailUsername1 := "username1"
	emailPassword1 := "password1"

	emailUsername2 := "username2"
	emailPassword2 := "password2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMxRouteDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMxRouteEmailAccountConfigDomainWithEmail(domainName, emailUsername1, emailPassword1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMxRouteDomainContainsEmail(domainName, emailUsername1),
					testAccCheckMxRouteDomainDoesNotContainEmail(domainName, emailUsername2),
					testAccCheckMxRouteEmailPasswordIsCorrect(domainName, emailUsername1, emailPassword1),
					testAccCheckMxRouteEmailPasswordIsIncorrect(domainName, emailUsername1, emailPassword2),
				),
			},
			//Change password
			{
				Config: testAccCheckMxRouteEmailAccountConfigDomainWithEmail(domainName, emailUsername1, emailPassword2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMxRouteDomainContainsEmail(domainName, emailUsername1),
					testAccCheckMxRouteDomainDoesNotContainEmail(domainName, emailUsername2),
					testAccCheckMxRouteEmailPasswordIsCorrect(domainName, emailUsername1, emailPassword2),
					testAccCheckMxRouteEmailPasswordIsIncorrect(domainName, emailUsername1, emailPassword1),
				),
			},
			//Change username
			{
				Config: testAccCheckMxRouteEmailAccountConfigDomainWithEmail(domainName, emailUsername2, emailPassword2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMxRouteDomainContainsEmail(domainName, emailUsername2),
					testAccCheckMxRouteDomainDoesNotContainEmail(domainName, emailUsername1),
					testAccCheckMxRouteEmailPasswordIsCorrect(domainName, emailUsername2, emailPassword2),
					testAccCheckMxRouteEmailPasswordIsIncorrect(domainName, emailUsername2, emailPassword1),
				),
			},
			//Change username and password and the same time
			{
				Config: testAccCheckMxRouteEmailAccountConfigDomainWithEmail(domainName, emailUsername1, emailPassword1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMxRouteDomainContainsEmail(domainName, emailUsername1),
					testAccCheckMxRouteDomainDoesNotContainEmail(domainName, emailUsername2),
					testAccCheckMxRouteEmailPasswordIsCorrect(domainName, emailUsername1, emailPassword1),
					testAccCheckMxRouteEmailPasswordIsIncorrect(domainName, emailUsername1, emailPassword2),
				),
			},
			{
				Config: testAccCheckMxRouteEmailAccountConfigDomainWithoutEmail(domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMxRouteDomainDoesNotContainAnyEmails(domainName),
				),
			},
		},
	})

}

func testAccCheckMxRouteDomainContainsEmail(domainName string, emailUsername string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		client := testAccProvider.Meta().(config)

		emails, err := api.GetEmailAccounts(client.Username, client.Password, domainName)
		if err != nil {
			return err
		}

		if len(emails) == 0 {
			return fmt.Errorf("domain '%s' does not contain emails", domainName)
		}
		if len(emails) > 1 {
			return fmt.Errorf("domain '%s' contains more than one email '%v'", domainName, emails)
		}
		if emails[0] != emailUsername {
			return fmt.Errorf("domain '%s' should have only email '%s', but has '%s'", domainName, emails[0], emailUsername)
		}

		return nil
	}
}

func testAccCheckMxRouteDomainDoesNotContainEmail(domainName string, emailUsername string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		client := testAccProvider.Meta().(config)

		emails, err := api.GetEmailAccounts(client.Username, client.Password, domainName)
		if err != nil {
			return err
		}

		if len(emails) == 0 {
			return fmt.Errorf("domain '%s' does not contain emails", domainName)
		}
		if len(emails) > 1 {
			return fmt.Errorf("domain '%s' contains more than one email '%v'", domainName, emails)
		}
		if emails[0] == emailUsername {
			return fmt.Errorf("domain '%s' should NOT have email '%s', but it has", domainName, emailUsername)
		}

		return nil
	}
}

func testAccCheckMxRouteDomainDoesNotContainAnyEmails(domainName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		client := testAccProvider.Meta().(config)

		emails, err := api.GetEmailAccounts(client.Username, client.Password, domainName)
		if err != nil {
			return err
		}

		if len(emails) > 0 {
			return fmt.Errorf("domain '%s' contains emails '%v'", domainName, emails)
		}

		return nil
	}

}

func testAccCheckMxRouteEmailPasswordIsCorrect(domainName, emailUsername, emailPassword string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		client := testAccProvider.Meta().(config)

		passwordIsCorrect, err := api.CheckPasswordIsCorrect(client.Username, client.Password, domainName, emailUsername, emailPassword)
		if err != nil {
			return err
		}

		if !*passwordIsCorrect {
			return fmt.Errorf("password '%s' is incorrect, but should be correct", emailPassword)
		}

		return nil
	}
}

// In theory only Correct would be sufficient, but I also want to test if api.CheckPasswordIsCorrect works correctly
func testAccCheckMxRouteEmailPasswordIsIncorrect(domainName, emailUsername, emailPassword string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		client := testAccProvider.Meta().(config)

		passwordIsCorrect, err := api.CheckPasswordIsCorrect(client.Username, client.Password, domainName, emailUsername, emailPassword)
		if err != nil {
			return err
		}

		if *passwordIsCorrect {
			return fmt.Errorf("password '%s' is correct, but should be incorrect", emailPassword)
		}

		return nil
	}
}

func testAccCheckMxRouteEmailAccountConfigDomainWithEmail(domainName, emailUsername, emailPassword string) string {
	return fmt.Sprintf(`
		
		resource "mxroute_domain" "foobar" {
			name = "%s"
		}
	
		resource "mxroute_email_account" "email" {
			domain = mxroute_domain.foobar.name
			username = "%s"
			password = "%s"
		}

`, domainName, emailUsername, emailPassword)
}

func testAccCheckMxRouteEmailAccountConfigDomainWithoutEmail(domainName string) string {
	return fmt.Sprintf(`
		
		resource "mxroute_domain" "foobar" {
			name = "%s"
		}

`, domainName)
}
