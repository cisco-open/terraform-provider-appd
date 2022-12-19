package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppdynamicscloudConnectionConfigurationAzureDataSource_Basic(t *testing.T) {
	resourceName := "appdynamicscloud_connection_configuration_azure.test"
	dataSourceName := "data.appdynamicscloud_connection_configuration_azure.test"
	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
	randomValue := makeTestVariable(acctest.RandString(5))
	rName := makeTestVariable(acctest.RandString(5))
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAzureDestroy,
		Steps: []resource.TestStep{

			{
				Config:      CreateAccConnectionConfigurationAzureDataSourceWithoutConfigurationId(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config: CreateAccConnectionConfigurationAzureDataSourceConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					// resource.TestCheckResourceAttrPair(dataSourceName, "configuration_id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "details.#", resourceName, "details.#"),
				),
			},
			{
				Config:      CreateAccConnectionConfigurationAzureUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue, rName),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},
			{
				Config:      CreateAccConnectionConfigurationAzureDataSourceWithInvalidConfigurationId(rName),
				ExpectError: regexp.MustCompile(""), // `(.)+ Object may not exists`
			},
			{
				Config: CreateAccConnectionConfigurationAzureDataSourceConfig(rName),
			},
		},
	})
}
func CreateAccConnectionConfigurationAzureDataSourceWithoutConfigurationId(rName string) string {
	resource := CreateAccConnectionConfigurationAzureConfigWithOptional(rName)
	resource += fmt.Sprintf(`
		data "appdynamicscloud_connection_configuration_azure" "test" {
		}
	`)
	return resource
}
func CreateAccConnectionConfigurationAzureDataSourceConfig(rName string) string {
	resource := CreateAccConnectionConfigurationAzureConfigWithOptional(rName)
	resource += fmt.Sprintf(`
	data "appdynamicscloud_connection_configuration_azure" "test" {

					configuration_id = appdynamicscloud_connection_configuration_azure.test.id
	}
	`)
	return resource
}
func CreateAccConnectionConfigurationAzureUpdatedConfigDataSourceRandomAttr(key, value, rName string) string {
	resource := CreateAccConnectionConfigurationAzureConfigWithOptional(rName)
	resource += fmt.Sprintf(`
	data "appdynamicscloud_connection_configuration_azure" "test" {

		configuration_id = appdynamicscloud_connection_configuration_azure.test.id
		%s = "%s"
	}
	`, key, value)
	return resource
}

func CreateAccConnectionConfigurationAzureDataSourceWithInvalidConfigurationId(rName string) string {
	var resource string
	// parentResources := getParentConnectionConfigurationAzure(rName)
	// parentResources = parentResources[:len(parentResources)-1]
	// resource += createConnectionConfigurationAzureConfig(parentResources)
	resource += fmt.Sprintf(`
		data "appdynamicscloud_connection_configuration_azure" "test" {					
			configuration_id = "%v"
		}
	`, "123e4567-e89b-12d3-a456-426614174000")
	return resource
}
