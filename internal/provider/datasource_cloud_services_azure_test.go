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

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppdynamicscloudServicesAzureDataSource_Basic(t *testing.T) {
	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
	randomValue := makeTestVariable(acctest.RandString(5))
	rName := makeTestVariable(acctest.RandString(5))
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: CreateAccServicesAzureDataSourceConfig(rName),
			},
			{
				Config:      CreateAccServicesAzureUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue, rName),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},
			{
				Config: CreateAccServicesAzureDataSourceConfig(rName),
			},
		},
	})
}

func CreateAccServicesAzureDataSourceConfig(rName string) string {
	resource := fmt.Sprintln(`
	data "appdynamicscloud_services_azure" "test" {
	}
	`)
	return resource
}
func CreateAccServicesAzureUpdatedConfigDataSourceRandomAttr(key, value, rName string) string {
	resource := fmt.Sprintf(`
	data "appdynamicscloud_services_azure" "test" {
		%s = "%s"
	}
	`, key, value)
	return resource
}
