package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppdynamicscloudAccessClientAppDataSource_Basic(t *testing.T) {

	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
	randomValue := makeTestVariable(acctest.RandString(5))

	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{

			{
				Config:      CreateAccAccessClientAppDataSourceWithoutClientId(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config: CreateAccAccessClientAppDataSourceConfig(rName),
			},
			{
				Config:      CreateAccAccessClientAppUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue, rName),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},
			{
				Config:      CreateAccAccessClientAppDataSourceWithInvalidClientId(rName),
				ExpectError: regexp.MustCompile("Bad Request"),
			},
			{
				Config: CreateAccAccessClientAppDataSourceConfig(rName),
			},
		},
	})
}

func CreateAccAccessClientAppDataSourceWithoutClientId(rName string) string {
	resource := fmt.Sprintln(`
		data "appdynamicscloud_access_client_app" "test" {}
	`)
	return resource
}

func CreateAccAccessClientAppDataSourceConfig(rName string) string {
	resource := fmt.Sprintln(`
	resource "appdynamicscloud_access_client_app" "example" {
		display_name = "tfacc_acp"
		description = "orchestrated by terraform during acceptance tests"
		auth_type = "client_secret_basic" 
	}

	data "appdynamicscloud_access_client_app" "test" {
		client_id = appdynamicscloud_access_client_app.example.id
	}
	`)
	return resource
}
func CreateAccAccessClientAppUpdatedConfigDataSourceRandomAttr(key, value, rName string) string {
	resource := fmt.Sprintf(`
	resource "appdynamicscloud_access_client_app" "example" {
		display_name = "tfacc_acp"
		description = "orchestrated by terraform during acceptance tests"
		auth_type = "client_secret_basic" 
	}
	
	data "appdynamicscloud_access_client_app" "test" {
		client_id = appdynamicscloud_access_client_app.example.id
		%s = "%s"
	}
	`, key, value)

	return resource
}

func CreateAccAccessClientAppDataSourceWithInvalidClientId(rName string) string {
	resource := fmt.Sprintf(`
		data "appdynamicscloud_access_client_app" "test" {
			client_id = "%v"
		}
	`, "abc123")

	return resource
}
