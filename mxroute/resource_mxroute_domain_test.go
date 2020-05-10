package mxroute

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/jereksel/terraform-provider-mxroute/api"
)

const TestDomainPrefix = "terraform-test-"

func init() {
	resource.AddTestSweepers("domain", &resource.Sweeper{
		Name: "domain",
		F: func(region string) error {

			config, err := sharedClient()
			if err != nil {
				return err
			}

			allDomains, err := api.GetAllDomains(config.Username, config.Password)
			if err != nil {
				return err
			}

			for _, domain := range allDomains {

				if strings.HasPrefix(domain, TestDomainPrefix) {
					if err := api.RemoveDomain(config.Username, config.Password, domain); err != nil {
						return err
					}
				}

			}

			return nil
		},
	})
}

func TestAccMxRouteDomain_basic(t *testing.T) {
	domainName := TestDomainPrefix + generateRandomResourceName() + ".email"
	resourceName := "mxroute_domain.foobar"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMxRouteDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMxRouteDomainConfig(domainName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "dkim", regexp.MustCompile(`.+`)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     domainName,
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "dkim", regexp.MustCompile(`.+`)),
				),
			},
		},
	})
}

func testAccCheckMxRouteDomainDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(config)

	allDomains, err := api.GetAllDomains(config.Username, config.Password)
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mxroute_domain" {
			continue
		}

		resourceDomain := rs.Primary.Attributes["name"]

		exists := false
		for _, domain := range allDomains {
			if domain == resourceDomain {
				exists = true
			}
		}

		if exists {
			return fmt.Errorf("domain '%s' still exists", resourceDomain)
		}

	}

	return nil
}

func testAccCheckMxRouteDomainConfig(domainName string) string {
	return fmt.Sprintf(`

		resource "mxroute_domain" "foobar" {
			name = "%s"
		}


`, domainName)
}
