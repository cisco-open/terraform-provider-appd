package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppdynamicscloudQueryDataSource_Basic(t *testing.T) {
	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
	randomValue := makeTestVariable(acctest.RandString(5))
	rName := makeTestVariable(acctest.RandString(5))
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{

			{
				Config:      CreateAccQueryDataSourceWithoutQuery(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},

			{
				Config: CreateAccQueryDataSourceConfig(rName),
			},
			{
				Config:      CreateAccQueryUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue, rName),
				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
			},
			{
				Config:      CreateAccQueryDataSourceWithInvalidQuery(rName),
				ExpectError: regexp.MustCompile("Query compilation failure!"),
			},
			{
				Config: CreateAccQueryDataSourceConfig(rName),
			},
		},
	})
}

func CreateAccQueryDataSourceWithoutQuery(rName string) string {
	resource := fmt.Sprintln(`
	data "appdynamicscloud_query" "test" {
	}
	`)
	return resource
}
func CreateAccQueryDataSourceConfig(rName string) string {
	resource := fmt.Sprintln(`
	data "appdynamicscloud_query" "test" {
		query = "fetch id: id, name: attributes(service.name), cpm: metrics(apm:response_time) {source, timestamp, min, max} from entities(apm:service)[attributes(service.namespace) = 'Levitate'].out.to(apm:service_instance) since -3h"
	}
	`)
	return resource
}
func CreateAccQueryUpdatedConfigDataSourceRandomAttr(key, value, rName string) string {
	resource := fmt.Sprintf(`
	data "appdynamicscloud_query" "test" {
		query = "fetch id: id, name: attributes(service.name), cpm: metrics(apm:response_time) {source, timestamp, min, max} from entities(apm:service)[attributes(service.namespace) = 'Levitate'].out.to(apm:service_instance) since -3h"
		%s = "%s"
	}
	`, key, value)
	return resource
}

func CreateAccQueryDataSourceWithInvalidQuery(rName string) string {
	resource := fmt.Sprintf(`
		data "appdynamicscloud_query" "test" {
				query = "%v"
		}
	`, "invalid_query")
	return resource
}
