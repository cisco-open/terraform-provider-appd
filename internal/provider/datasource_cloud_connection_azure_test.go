package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppdynamicscloudConnectionAzureDataSource_Basic(t *testing.T) {
	resourceName := "appdynamicscloud_connection_azure.test"
	dataSourceName := "data.appdynamicscloud_connection_azure.test"
	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
	randomValue := makeTestVariable(acctest.RandString(5))
	rName := makeTestVariable(acctest.RandString(5))
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config:      CreateAccConnectionAzureDataSourceWithoutConnectionId(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config: CreateAccConnectionAzureDataSourceConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					// resource.TestCheckResourceAttrPair(dataSourceName, "connection_id", resourceName, "connection_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "state", resourceName, "state"),
					resource.TestCheckResourceAttrPair(dataSourceName, "configuration_id", resourceName, "configuration_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "details.#", resourceName, "details.#"),
				),
			},
			{
				Config:      CreateAccConnectionAzureUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue, rName),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},
			{
				Config:      CreateAccConnectionAzureDataSourceWithInvalidConnectionId(rName),
				ExpectError: regexp.MustCompile("expected (.)+ to be a valid UUID, got (.)+"),
			},
			{
				Config: CreateAccConnectionAzureDataSourceConfig(rName),
			},
		},
	})
}
func CreateAccConnectionAzureDataSourceWithoutConnectionId(rName string) string {
	resource := CreateAccConnectionAzureConfigWithOptional(rName)
	resource += fmt.Sprintf(`
			data "appdynamicscloud_connection_azure" "test" {
			}
			`)
	return resource
}
func CreateAccConnectionAzureDataSourceConfig(rName string) string {
	resource := CreateAccConnectionAzureConfigWithOptional(rName)
	resource += fmt.Sprintf(`
	data "appdynamicscloud_connection_azure" "test" {
		connection_id = appdynamicscloud_connection_azure.test.id
	}
	`)
	return resource
}
func CreateAccConnectionAzureUpdatedConfigDataSourceRandomAttr(key, value, rName string) string {
	resource := CreateAccConnectionAzureConfigWithOptional(rName)
	resource += fmt.Sprintf(`
	data "appdynamicscloud_connection_azure" "test" {
		connection_id = appdynamicscloud_connection_azure.test.id
		%s = "%s"
	}
	`, key, value)
	return resource
}

func CreateAccConnectionAzureDataSourceWithInvalidConnectionId(rName string) string {
	var resource string
	resource += fmt.Sprintf(`
	data "appdynamicscloud_connection_azure" "test" {
		connection_id = "%v"
	}
	`, "abcd")
	return resource
}
