package mxroute

//Based on https://github.com/terraform-providers/terraform-provider-cloudflare/blob/7df730622f105a6737afedd2c24453c855b32322/cloudflare/cloudflare_sweeper_test.go

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// sharedClient returns a common MXRoute client setup needed for the
// sweeper functions.
func sharedClient() (config, error) {
	username := os.Getenv("MXROUTE_USERNAME")
	password := os.Getenv("MXROUTE_PASSWORD")

	return config{Username: username, Password: password}, nil

}
