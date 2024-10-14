package main

import (
	"github.com/opengovern/og-azure-describer/steampipe-plugin-azure/azure"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: azure.Plugin})
}
