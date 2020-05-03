package mxroute

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"mxroute": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func generateRandomResourceName() string {
	return acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("MXROUTE_USERNAME"); v == "" {
		t.Fatal("MXROUTE_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("MXROUTE_PASSWORD"); v == "" {
		t.Fatal("MXROUTE_PASSWORD must be set for acceptance tests")
	}
}
