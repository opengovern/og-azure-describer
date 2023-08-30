package main

import (
	"github.com/kaytu-io/kaytu-azure-describer/steampipe-plugin-azuread/azuread"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: azuread.Plugin})
}
