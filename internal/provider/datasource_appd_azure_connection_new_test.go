package provider

// import (
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
// )

// func TestAccAppdynamicscloudConnectionAzureDataSource_Basic(t *testing.T) {
// 	resourceName := "appdynamicscloud_connection_azure.test"
// 	dataSourceName := "data.appdynamicscloud_connection_azure.test"
// 	randomParameter := acctest.RandStringFromCharSet(5, "abcdefghijklmnopqrstuvwxyz")
// 	randomValue := makeTestVariable(acctest.RandString(5))
// 	rName := makeTestVariable(acctest.RandString(5))
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAzureDestroy,
// 		Steps: append([]resource.TestStep{

// 			{
// 				Config:      CreateAccConnectionAzureDataSourceWithoutConnectionId(rName),
// 				ExpectError: regexp.MustCompile(`Missing required argument`),
// 			},

// 			{
// 				Config: CreateAccConnectionAzureDataSourceConfig(rName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttrPair(dataSourceName, "connection_id", resourceName, "connection_id"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "state", resourceName, "state"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "configuration_id", resourceName, "configuration_id"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "connection_details.#", resourceName, "connection_details.#"),
// 					resource.TestCheckResourceAttrPair(dataSourceName, "configuration_details.#", resourceName, "configuration_details.#"),
// 				),
// 			},
// 			{
// 				Config:      CreateAccConnectionAzureUpdatedConfigDataSourceRandomAttr(randomParameter, randomValue, rName),
// 				ExpectError: regexp.MustCompile(`An argument named (.)+ is not expected here.`),
// 			},
// 			{
// 				Config:      CreateAccConnectionAzureDataSourceWithInvalidConnectionId(rName),
// 				ExpectError: regexp.MustCompile(""), // `(.)+ Object may not exists`
// 			},
// 		}, generateStepForDataSourceUpdatedOptionalAttrConnectionAzure(rName, dataSourceName, resourceName)...),
// 	})
// }
// func CreateAccConnectionAzureDataSourceWithoutConnectionId(rName string) string {
// 	resource := CreateAccConnectionAzureConfigWithOptional(rName)
// 	resource += fmt.Sprintf(`
// 			data "appdynamicscloud_connection_azure" "test" {
// 			}
// 			`)
// 	return resource
// }
// func CreateAccConnectionAzureDataSourceConfig(rName string) string {
// 	resource := CreateAccConnectionAzureConfigWithOptional(rName)
// 	resource += fmt.Sprintf(`
// 	data "appdynamicscloud_connection_azure" "test" {

// 					connection_id = appdynamicscloud_connection_azure.test.connection_id
// 	}
// 	`)
// 	return resource
// }
// func CreateAccConnectionAzureUpdatedConfigDataSourceRandomAttr(key, value, rName string) string {
// 	resource := CreateAccConnectionAzureConfigWithOptional(rName)
// 	resource += fmt.Sprintf(`
// 	data "appdynamicscloud_connection_azure" "test" {

// 					connection_id = appdynamicscloud_connection_azure.test.connection_id
// 		%s = "%s"
// 	}
// 	`, key, value)
// 	return resource
// }

// func CreateAccConnectionAzureDataSourceWithInvalidConnectionId(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			data "appdynamicscloud_connection_azure" "test" {
							
// 							connection_id = "%v"


// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							configuration_id = "%v"

// 						connection_details {
    
						                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 						}

// 						configuration_details {
    
						                        
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
// 		searchInObject(resourceConnectionAzureTest, "display_name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.valid.0"))
// 	return resource
// }

// func generateStepForDataSourceUpdatedOptionalAttrConnectionAzure(rName, dataSourceName, resourceName string) []resource.TestStep {
// 	testSteps := make([]resource.TestStep, 0, 1)
// 	var valid interface{}
// 	valid = searchInObject(resourceConnectionAzureTest, "display_name.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAzureDataSourceUpdatedOptionalAttrDisplayName(rName, valid),
// 		Check: resource.ComposeTestCheckFunc(
// 			resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
// 		),
// 	})
// 	valid = searchInObject(resourceConnectionAzureTest, "description.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAzureDataSourceUpdatedOptionalAttrDescription(rName, valid),
// 		Check: resource.ComposeTestCheckFunc(
// 			resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
// 		),
// 	})
// 	valid = searchInObject(resourceConnectionAzureTest, "state.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAzureDataSourceUpdatedOptionalAttrState(rName, valid),
// 		Check: resource.ComposeTestCheckFunc(
// 			resource.TestCheckResourceAttrPair(dataSourceName, "state", resourceName, "state"),
// 		),
// 	})
// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_id.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAzureDataSourceUpdatedOptionalAttrConfigurationId(rName, valid),
// 		Check: resource.ComposeTestCheckFunc(
// 			resource.TestCheckResourceAttrPair(dataSourceName, "configuration_id", resourceName, "configuration_id"),
// 		),
// 	})
// 	return testSteps
// }

// func CreateAccConnectionAzureDataSourceUpdatedOptionalAttrDisplayName(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							connection_id = "%v"
							
// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							configuration_id = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							                        
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
// 			`, searchInObject(resourceConnectionAzureTest, "connection_id.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.valid.0"))
// 	resource += fmt.Sprintf(`
// 		data "appdynamicscloud_connection_azure" "test" {

// 						connection_id = appdynamicscloud_connection_azure.test.connection_id
// 		}
// 		`)
// 	return resource
// }
// func CreateAccConnectionAzureDataSourceUpdatedOptionalAttrDescription(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							connection_id = "%v"

// 							display_name = "%v"
							
// 							description = "%v"

// 							state = "%v"

// 							configuration_id = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							                        
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
// 			`, searchInObject(resourceConnectionAzureTest, "connection_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "display_name.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.valid.0"))
// 	resource += fmt.Sprintf(`
// 		data "appdynamicscloud_connection_azure" "test" {

// 						connection_id = appdynamicscloud_connection_azure.test.connection_id
// 		}
// 		`)
// 	return resource
// }
// func CreateAccConnectionAzureDataSourceUpdatedOptionalAttrState(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							connection_id = "%v"

// 							display_name = "%v"

// 							description = "%v"
							
// 							state = "%v"

// 							configuration_id = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							                        
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
// 			`, searchInObject(resourceConnectionAzureTest, "connection_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "display_name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "configuration_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.valid.0"))
// 	resource += fmt.Sprintf(`
// 		data "appdynamicscloud_connection_azure" "test" {

// 						connection_id = appdynamicscloud_connection_azure.test.connection_id
// 		}
// 		`)
// 	return resource
// }
// func CreateAccConnectionAzureDataSourceUpdatedOptionalAttrConfigurationId(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							connection_id = "%v"

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"
							
// 							configuration_id = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							                        
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
// 			`, searchInObject(resourceConnectionAzureTest, "connection_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "display_name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.valid.0"))
// 	resource += fmt.Sprintf(`
// 		data "appdynamicscloud_connection_azure" "test" {

// 						connection_id = appdynamicscloud_connection_azure.test.connection_id
// 		}
// 		`)
// 	return resource
// }
