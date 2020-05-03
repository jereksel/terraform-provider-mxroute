package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/jereksel/terraform-provider-mxroute/mxroute"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mxroute.Provider,
	})
}
