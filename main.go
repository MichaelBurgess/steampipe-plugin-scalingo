package main

import (
	"github.com/francois2metz/steampipe-plugin-scalingo/scalingo"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: scalingo.Plugin})
}
