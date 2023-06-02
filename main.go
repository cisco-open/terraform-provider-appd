// Copyright 2023 Cisco Systems, Inc.
//
// Licensed under the MPL License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.mozilla.org/en-US/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"

	"github.com/cisco-open/appd-cloud-terraform/internal/provider"
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
	// 	err := plugin.Debug(context.Background(), "github.com/cisco-open/appd-cloud-terraform", opts)
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	return
	// }

	opts := &plugin.ServeOpts{
		Debug: debugMode,

		ProviderAddr: "registry.terraform.io/appdynamics/appdynamicscloud",

		ProviderFunc: provider.New(),
	}

	plugin.Serve(opts)
}
