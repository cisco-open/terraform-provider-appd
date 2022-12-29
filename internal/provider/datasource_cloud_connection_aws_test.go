package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppdynamicscloudConnectionAwsDataSource_Basic(t *testing.T) {
	resourceName := "appdynamicscloud_connection_aws.test"
	dataSourceName := "data.appdynamicscloud_connection_aws.test"
	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
	randomValue := makeTestVariable(acctest.RandString(5))
	rName := makeTestVariable(acctest.RandString(5))
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsAccessDestroy,
		Steps: []resource.TestStep{

			{
				Config:      CreateAccConnectionAwsDataSourceWithoutConnectionId(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config: CreateAccConnectionAwsDataSourceConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "state", resourceName, "state"),
					resource.TestCheckResourceAttrPair(dataSourceName, "configuration_id", resourceName, "configuration_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "connection_details.#", resourceName, "connection_details.#"),
					resource.TestCheckResourceAttrPair(dataSourceName, "configuration_details.#", resourceName, "configuration_details.#"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_details", "configuration_details_service_default", "state"},
			},
			{
				Config:      CreateAccConnectionAwsUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue, rName),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},
			{
				Config:      CreateAccConnectionAwsDataSourceWithInvalidConnectionId(rName),
				ExpectError: regexp.MustCompile("Not Found Error"), // `(.)+ Object may not exists`
			},
			{
				Config: CreateAccConnectionAwsDataSourceConfig(rName),
			},
		},
	})
}
func CreateAccConnectionAwsDataSourceWithoutConnectionId(rName string) string {
	resource := CreateAccConnectionAwsAccessConfigWithOptional(rName)
	resource += fmt.Sprintf(`
			data "appdynamicscloud_connection_aws" "test" {
			}
			`)
	return resource
}
func CreateAccConnectionAwsDataSourceConfig(rName string) string {
	resource := CreateAccConnectionAwsAccessConfigWithOptional(rName)
	resource += fmt.Sprintf(`
	data "appdynamicscloud_connection_aws" "test" {
		connection_id = appdynamicscloud_connection_aws.test.id
	}
	`)
	return resource
}
func CreateAccConnectionAwsUpdatedConfigDataSourceRandomAttr(key, value, rName string) string {
	resource := CreateAccConnectionAwsAccessConfigWithOptional(rName)
	resource += fmt.Sprintf(`
	data "appdynamicscloud_connection_aws" "test" {
		connection_id = appdynamicscloud_connection_aws.test.id
		%s = "%s"
	}
	`, key, value)
	return resource
}

func CreateAccConnectionAwsDataSourceWithInvalidConnectionId(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsAccess(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsAccessConfig(parentResources)
	resource += fmt.Sprintf(`
		data "appdynamicscloud_connection_aws" "test" {					
				connection_id = "%v"
		}
	`, "123e4567-e89b-12d3-a456-426614174000")
	return resource
}
