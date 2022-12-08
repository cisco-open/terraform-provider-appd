package main

import (
	"flag"

	"github.com/aniketk-crest/terraform-provider-appdynamicscloud/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	// opts := &plugin.ServeOpts{
	// 	ProviderFunc: func() *schema.Provider {
	// 		return appdynamicscloud.Provider()
	// 	},
	// }

	// if debugMode {
	// 	err := plugin.Debug(context.Background(), "github.com/aniketk-crest/terraform-provider-appdynamicscloud", opts)
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	return
	// }

	opts := &plugin.ServeOpts{
		Debug: debugMode,

		ProviderAddr: "github.com/aniketk-crest/terraform-provider-appdynamicscloud",

		ProviderFunc: provider.New(),
	}

	plugin.Serve(opts)
}
