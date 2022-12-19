package provider

// import (
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
// )

// #   query="fetch id: id, name: attributes(service.name), cpm: metrics(apm:response_time) {source, timestamp, min, max} from entities(apm:service)[attributes(service.namespace) = 'Levitate'].out.to(apm:service_instance) since -3h"
// func TestAccAppdynamicscloudQueryDataSource_Basic(t *testing.T) {
// 	resourceName := "appdynamicscloud_query.test"
// 	dataSourceName := "data.appdynamicscloud_query.test"
// 	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
// 	randomValue := makeTestVariable(acctest.RandString(5))
// 	rName := makeTestVariable(acctest.RandString(5))
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudQueryDestroy,
// 		Steps: append([]resource.TestStep{

// 			{
// 				Config:      CreateAccQueryDataSourceWithoutQuery(rName),
// 				ExpectError: regexp.MustCompile(`Missing required argument`),
// 			},

// 			{
// 				Config: CreateAccQueryDataSourceConfig(rName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttrPair(dataSourceName, "query", resourceName, "query"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "reponse", resourceName, "reponse"),
// 				),
// 			},
// 			{
// 				Config:      CreateAccQueryUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue, rName),
// 				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
// 			},
// 			{
// 				Config:      CreateAccQueryDataSourceWithInvalidQuery(rName),
// 				ExpectError: regexp.MustCompile(""), // `(.)+ Object may not exists`
// 			},
// 		}, generateStepForDataSourceUpdatedOptionalAttrQuery(rName, dataSourceName, resourceName)...),
// 	})
// }
// func CreateAccQueryDataSourceWithoutQuery(rName string) string {
// 	resource := CreateAccQueryConfigWithOptional(rName)
// 	resource += fmt.Sprintf(`
// 			data "appdynamicscloud_query" "test" {
// 			}
// 			`)
// 	return resource
// }
// func CreateAccQueryDataSourceConfig(rName string) string {
// 	resource := CreateAccQueryConfigWithOptional(rName)
// 	resource += fmt.Sprintf(`
// 	data "appdynamicscloud_query" "test" {

// 					query = appdynamicscloud_query.test.query
// 	}
// 	`)
// 	return resource
// }
// func CreateAccQueryUpdatedConfigDataSourceRandomAttr(key, value, rName string) string {
// 	resource := CreateAccQueryConfigWithOptional(rName)
// 	resource += fmt.Sprintf(`
// 	data "appdynamicscloud_query" "test" {

// 					query = appdynamicscloud_query.test.query
// 		%s = "%s"
// 	}
// 	`, key, value)
// 	return resource
// }

// func CreateAccQueryDataSourceWithInvalidQuery(rName string) string {
// 	var resource string
// 	parentResources := getParentQuery(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createQueryConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			data "appdynamicscloud_query" "test" {

// 							query = "%v"

// 							reponse = "%v"
// 		}
// 	`, "abcd",
// 		searchInObject(resourceQueryTest, "reponse.valid.0"))
// 	return resource
// }

// func generateStepForDataSourceUpdatedOptionalAttrQuery(rName, dataSourceName, resourceName string) []resource.TestStep {
// 	testSteps := make([]resource.TestStep, 0, 1)
// 	var valid interface{}
// 	valid = searchInObject(resourceQueryTest, "reponse.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccQueryDataSourceUpdatedOptionalAttrReponse(rName, valid),
// 		Check: resource.ComposeTestCheckFunc(
// 			resource.TestCheckResourceAttrPair(dataSourceName, "reponse", resourceName, "reponse"),
// 		),
// 	})
// 	return testSteps
// }

// func CreateAccQueryDataSourceUpdatedOptionalAttrReponse(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentQuery(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createQueryConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_query" "test" {

// 							query = "%v"

// 							reponse = "%v"
// 			}
// 			`, searchInObject(resourceQueryTest, "query.valid.0"),
// 		value)
// 	resource += fmt.Sprintf(`
// 		data "appdynamicscloud_query" "test" {

// 						query = appdynamicscloud_query.test.query
// 		}
// 		`)
// 	return resource
// }
