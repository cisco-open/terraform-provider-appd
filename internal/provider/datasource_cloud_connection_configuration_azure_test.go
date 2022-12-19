package provider

// import (
// 	"fmt"
// 	"regexp"
// 	"testing"

// 	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
// 	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
// )

// func TestAccAppdynamicscloudConnectionConfigurationAzureDataSource_Basic(t *testing.T) {
// 	resourceName := "appdynamicscloud_connection_configuration_azure.test"
// 	dataSourceName := "data.appdynamicscloud_connection_configuration_azure.test"
// 	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
// 	randomValue := makeTestVariable(acctest.RandString(5))
// 	rName := makeTestVariable(acctest.RandString(5))
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAzureDestroy,
// 		Steps: append([]resource.TestStep{

// 			{
// 				Config:      CreateAccConnectionConfigurationAzureDataSourceWithoutConfigurationId(rName),
// 				ExpectError: regexp.MustCompile(`Missing required argument`),
// 			},

// 			{
// 				Config: CreateAccConnectionConfigurationAzureDataSourceConfig(rName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttrPair(dataSourceName, "configuration_id", resourceName, "configuration_id"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "details.#", resourceName, "details.#"),
// 				),
// 			},
// 			{
// 				Config:      CreateAccConnectionConfigurationAzureUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue, rName),
// 				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
// 			},
// 			{
// 				Config:      CreateAccConnectionConfigurationAzureDataSourceWithInvalidConfigurationId(rName),
// 				ExpectError: regexp.MustCompile(""), // `(.)+ Object may not exists`
// 			},
// 		}, generateStepForDataSourceUpdatedOptionalAttrConnectionConfigurationAzure(rName, dataSourceName, resourceName)...),
// 	})
// }
// func CreateAccConnectionConfigurationAzureDataSourceWithoutConfigurationId(rName string) string {
// 	resource := CreateAccConnectionConfigurationAzureConfigWithOptional(rName)
// 	resource += fmt.Sprintf(`
// 			data "appdynamicscloud_connection_configuration_azure" "test" {
// 			}
// 			`)
// 	return resource
// }
// func CreateAccConnectionConfigurationAzureDataSourceConfig(rName string) string {
// 	resource := CreateAccConnectionConfigurationAzureConfigWithOptional(rName)
// 	resource += fmt.Sprintf(`
// 	data "appdynamicscloud_connection_configuration_azure" "test" {

// 					configuration_id = appdynamicscloud_connection_configuration_azure.test.configuration_id
// 	}
// 	`)
// 	return resource
// }
// func CreateAccConnectionConfigurationAzureUpdatedConfigDataSourceRandomAttr(key, value, rName string) string {
// 	resource := CreateAccConnectionConfigurationAzureConfigWithOptional(rName)
// 	resource += fmt.Sprintf(`
// 	data "appdynamicscloud_connection_configuration_azure" "test" {

// 					configuration_id = appdynamicscloud_connection_configuration_azure.test.configuration_id
// 		%s = "%s"
// 	}
// 	`, key, value)
// 	return resource
// }

// func CreateAccConnectionConfigurationAzureDataSourceWithInvalidConfigurationId(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionConfigurationAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionConfigurationAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			data "appdynamicscloud_connection_configuration_azure" "test" {
							
// 							configuration_id = "%v"


// 							display_name = "%v"

// 							description = "%v"

// 						details {
    
						                        
//                             regions = ["%v","%v"]
                        
//                             resource_groups = ["%v","%v"]

//                             polling {
                                                    
//                                 interval = %v
                        
//                                 unit = "%v"

//                               }

//                                 import_tags {
                                                    
//                                     enabled = "%v"
                        
//                                     excluded_keys = ["%v","%v"]

//                                   }
                        
//                                     tag_filter = "%v"

//                                     services {
                                                    
//                                         name = "%v"

//                                         import_tags {
                                                    
//                                             enabled = "%v"
                        
//                                             excluded_keys = ["%v","%v"]

//                                           }
                        
//                                             tag_filter = "%v"

//                                             polling {
                                                    
//                                                 interval = %v
                        
//                                                 unit = "%v"

//                                               }

//                                               }

// 						}
// 		}
// 	`, "abcd",
// 		searchInObject(resourceConnectionConfigurationAzureTest, "display_name.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
// 	return resource
// }

// func generateStepForDataSourceUpdatedOptionalAttrConnectionConfigurationAzure(rName, dataSourceName, resourceName string) []resource.TestStep {
// 	testSteps := make([]resource.TestStep, 0, 1)
// 	var valid interface{}
// 	valid = searchInObject(resourceConnectionConfigurationAzureTest, "display_name.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionConfigurationAzureDataSourceUpdatedOptionalAttrDisplayName(rName, valid),
// 		Check: resource.ComposeTestCheckFunc(
// 			resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
// 		),
// 	})
// 	valid = searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionConfigurationAzureDataSourceUpdatedOptionalAttrDescription(rName, valid),
// 		Check: resource.ComposeTestCheckFunc(
// 			resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
// 		),
// 	})
// 	return testSteps
// }

// func CreateAccConnectionConfigurationAzureDataSourceUpdatedOptionalAttrDisplayName(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionConfigurationAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionConfigurationAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_configuration_azure" "test" {

// 							configuration_id = "%v"
							
// 							display_name = "%v"

// 							description = "%v"

// 							details {
    
							                        
//                             regions = ["%v","%v"]
                        
//                             resource_groups = ["%v","%v"]

//                             polling {
                                                    
//                                 interval = %v
                        
//                                 unit = "%v"

//                               }

//                                 import_tags {
                                                    
//                                     enabled = "%v"
                        
//                                     excluded_keys = ["%v","%v"]

//                                   }
                        
//                                     tag_filter = "%v"

//                                     services {
                                                    
//                                         name = "%v"

//                                         import_tags {
                                                    
//                                             enabled = "%v"
                        
//                                             excluded_keys = ["%v","%v"]

//                                           }
                        
//                                             tag_filter = "%v"

//                                             polling {
                                                    
//                                                 interval = %v
                        
//                                                 unit = "%v"

//                                               }

//                                               }

// 							}
// 			}
// 			`, searchInObject(resourceConnectionConfigurationAzureTest, "configuration_id.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
// 	resource += fmt.Sprintf(`
// 		data "appdynamicscloud_connection_configuration_azure" "test" {

// 						configuration_id = appdynamicscloud_connection_configuration_azure.test.configuration_id
// 		}
// 		`)
// 	return resource
// }
// func CreateAccConnectionConfigurationAzureDataSourceUpdatedOptionalAttrDescription(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionConfigurationAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionConfigurationAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_configuration_azure" "test" {

// 							configuration_id = "%v"

// 							display_name = "%v"
							
// 							description = "%v"

// 							details {
    
							                        
//                             regions = ["%v","%v"]
                        
//                             resource_groups = ["%v","%v"]

//                             polling {
                                                    
//                                 interval = %v
                        
//                                 unit = "%v"

//                               }

//                                 import_tags {
                                                    
//                                     enabled = "%v"
                        
//                                     excluded_keys = ["%v","%v"]

//                                   }
                        
//                                     tag_filter = "%v"

//                                     services {
                                                    
//                                         name = "%v"

//                                         import_tags {
                                                    
//                                             enabled = "%v"
                        
//                                             excluded_keys = ["%v","%v"]

//                                           }
                        
//                                             tag_filter = "%v"

//                                             polling {
                                                    
//                                                 interval = %v
                        
//                                                 unit = "%v"

//                                               }

//                                               }

// 							}
// 			}
// 			`, searchInObject(resourceConnectionConfigurationAzureTest, "configuration_id.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "display_name.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
// 	resource += fmt.Sprintf(`
// 		data "appdynamicscloud_connection_configuration_azure" "test" {

// 						configuration_id = appdynamicscloud_connection_configuration_azure.test.configuration_id
// 		}
// 		`)
// 	return resource
// }
