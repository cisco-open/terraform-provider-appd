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
