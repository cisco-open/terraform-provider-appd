package provider

import (
	"fmt"
	"regexp"
	"testing"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// const connectionAzureSelfRequiredCount = 2

var resourceConnectionAzureTest = map[string]interface{}{
	"display_name": map[string]interface{}{
		"valid":           []interface{}{"dt4vo4d4xz", "a0frbw35sw", "2rpqehuvuw", "mk859jgg17"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"5sjip2nath", "fz08dyshuj", "oev7l0hynv", "bxwjyfizwx", "uvi5ig8g12", "bbs5lmbh92", "pmnslf45bv", "syiz1i6s4l", "3s8j1nrcu1", "io16nqssxw", "l1i10qj5jl", "jdmebz0a4k", "r1lmz309d6", "hn0z8657k9", "fme9wtf1tf"},
	},

	"description": map[string]interface{}{
		"valid":           []interface{}{"1ayeizxvla", "l8nryel1f2", "dx56l00zee", "0ut8odktre"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"z3y6936fah", "2e6bb56vm6", "ro89r2poac", "yp09cg5kau", "q019suf5hv", "vhf7hnh89z", "wv84ofmdgd", "2vnxmb5ak6", "g2cyj9e6ic", "u9m6mpapgu", "8xvwjs361y", "ud01szdlme", "ie0eh8cwyy", "wsi5o83kwj", "lr3glbspd3"},
	},

	"state": map[string]interface{}{
		"valid":           []interface{}{"ACTIVE", "INACTIVE"},
		"invalid":         []interface{}{"xjp98xn8xb"},
		"multiple_valids": []interface{}{"ACTIVE", "INACTIVE"},
	},

	"configuration_id": map[string]interface{}{
		"valid":           []interface{}{"f3db65e4-ce5a-4d7e-a140-255bf017d87f", "f3db65e4-ce5a-4d7e-a140-255bf017d87f"},
		"invalid":         []interface{}{"abcddefsdf", "asdasdsafr"},
		"multiple_valids": []interface{}{"f3db65e4-ce5a-4d7e-a140-255bf017d87f", "f3db65e4-ce5a-4d7e-a140-255bf017d87f"},
	},

	"details": map[string]interface{}{
		"client_id": map[string]interface{}{
			"valid":           []interface{}{"adb90c29-204d-43d9-987c-ab406d9199cc", "adb90c29-204d-43d9-987c-ab406d9199cc"},
			"invalid":         []interface{}{"ACTIVE", "INACTIVE"},
			"multiple_valids": []interface{}{"adb90c29-204d-43d9-987c-ab406d9199cc", "adb90c29-204d-43d9-987c-ab406d9199cc"},
		},

		"client_secret": map[string]interface{}{
			"valid":           []interface{}{"ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx", "ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx"},
			"invalid":         []interface{}{"ACTIVE", "INACTIVE"},
			"multiple_valids": []interface{}{"ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx", "ZmD8Q~q87S6hbvMVFh5hKS3y8dLA9C1Xfc5jLbjx"},
		},

		"tenant_id": map[string]interface{}{
			"valid":           []interface{}{"f3db65e4-ce5a-4d7e-a140-255bf017d87f", "f3db65e4-ce5a-4d7e-a140-255bf017d87f"},
			"invalid":         []interface{}{"ACTIVE", "INACTIVE"},
			"multiple_valids": []interface{}{"f3db65e4-ce5a-4d7e-a140-255bf017d87f", "f3db65e4-ce5a-4d7e-a140-255bf017d87f"},
		},

		"subscription_id": map[string]interface{}{
			"valid":           []interface{}{"fca41da2-4908-49e2-b0cb-d3d2080fc5be", "fca41da2-4908-49e2-b0cb-d3d2080fc5be"},
			"invalid":         []interface{}{"ACTIVE", "INACTIVE"},
			"multiple_valids": []interface{}{"fca41da2-4908-49e2-b0cb-d3d2080fc5be", "fca41da2-4908-49e2-b0cb-d3d2080fc5be"},
		},
	},
}

func TestAccAppdynamicscloudConnectionAzure_Basic(t *testing.T) {
	var connectionAzure_default cloudconnectionapi.ConnectionResponse
	var connectionAzure_updated cloudconnectionapi.ConnectionResponse
	resourceName := "appdynamicscloud_connection_azure.test"

	rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAzureDestroy,
		Steps: append([]resource.TestStep{
			{
				Config:      CreateAccConnectionAzureWithoutDisplayName(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccConnectionAzureWithoutDetails(rName),
				ExpectError: regexp.MustCompile(`Insufficient details blocks`),
			},
			{
				Config: CreateAccConnectionAzureConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, &connectionAzure_default),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),

					resource.TestCheckResourceAttr(resourceName, "description", ""),

					// resource.TestCheckResourceAttr(resourceName, "state", ""),

					resource.TestCheckResourceAttr(resourceName, "configuration_id", ""),

					resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.client_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.client_secret", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.tenant_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.subscription_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))),
				),
			},
			{
				Config: CreateAccConnectionAzureConfigWithOptional(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, &connectionAzure_updated),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "description.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "state", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "state.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.client_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.client_secret", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.tenant_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.subscription_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))),
					testAccCheckAppdynamicscloudConnectionAzureIdEqual(&connectionAzure_default, &connectionAzure_updated),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"details"},
			},
		}, generateStepForUpdatedRequiredAttrConnectionAzure(rName, resourceName, &connectionAzure_default, &connectionAzure_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionAzure_Update(t *testing.T) {
	var connectionAzure_default cloudconnectionapi.ConnectionResponse
	var connectionAzure_updated cloudconnectionapi.ConnectionResponse
	resourceName := "appdynamicscloud_connection_azure.test"
	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAzureDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionAzureConfig(rName),
				Check:  testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, &connectionAzure_default),
			},
		}, generateStepForUpdatedAttrConnectionAzure(rName, resourceName, &connectionAzure_default, &connectionAzure_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionAzure_NegativeCases(t *testing.T) {
	resourceName := "appdynamicscloud_connection_azure.test"

	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAzureDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionAzureConfig(rName),
			},
		}, generateNegativeStepsConnectionAzure(rName, resourceName)...),
	})
}

// func TestAccAppdynamicscloudConnectionAzure_MultipleCreateDelete(t *testing.T) {

// 	// [TODO]: Add makeTestVariable() to utils.go file
// 	rName := makeTestVariable(acctest.RandString(5))

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAzureDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: CreateAccConnectionAzureMultipleConfig(rName),
// 			},
// 		},
// 	})
// }

func CreateAccConnectionAzureWithoutDisplayName(rName string) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "appdynamicscloud_connection_azure" "test" {
									description = "%v"
									state = "%v"
									configuration_id = appdynamicscloud_connection_configuration_azure.test.id
									details {
    
									                        
                                        client_id = "%v"
                        
                                        client_secret = "%v"
                        
                                        tenant_id = "%v"
                        
                                        subscription_id = "%v"
									}
				}
			`, searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}
func CreateAccConnectionAzureWithoutDetails(rName string) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "appdynamicscloud_connection_azure" "test" {
									display_name = "%v"
									description = "%v"
									state = "%v"
									configuration_id = appdynamicscloud_connection_configuration_azure.test.id
				}
			`, searchInObject(resourceConnectionAzureTest, "display_name.valid.0"),
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"))
	return resource
}

func CreateAccConnectionAzureConfig(rName string) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_azure" "test" {
							display_name = "%v"
							details {
    
							 
						          client_id = "%v"
 
						          client_secret = "%v"
 
						          tenant_id = "%v"
 
						          subscription_id = "%v"
							}
		}
	`, rName,
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}

func CreateAccConnectionAzureConfigWithOptional(rName string) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]

	resource += createConnectionAzureConfig(parentResources)

	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_azure" "test" {
						display_name = "%v"
						description = "%v"
						state = "%v"
						configuration_id = appdynamicscloud_connection_configuration_azure.test.id
                        details {
    
                                                
                            client_id = "%v"
                        
                            client_secret = "%v"
                        
                            tenant_id = "%v"
                        
                            subscription_id = "%v"
                        }
		}
	`, rName,
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}

func generateStepForUpdatedRequiredAttrConnectionAzure(rName string, resourceName string, connectionAzure_default, connectionAzure_updated *cloudconnectionapi.ConnectionResponse) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var value interface{}
	value = searchInObject(resourceConnectionAzureTest, "display_name.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAzureUpdateRequiredDisplayName(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
			resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
		),
	})
	value = searchInObject(resourceConnectionAzureTest, "details.client_id.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAzureUpdateRequiredDetailsClientId(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
			resource.TestCheckResourceAttr(resourceName, "details.0.client_id", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
		),
	})
	value = searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAzureUpdateRequiredDetailsClientSecret(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
			resource.TestCheckResourceAttr(resourceName, "details.0.client_secret", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
		),
	})
	// value = searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.1")
	// testSteps = append(testSteps, resource.TestStep{
	// 	Config: CreateAccConnectionAzureUpdateRequiredDetailsTenantId(rName),
	// 	Check: resource.ComposeTestCheckFunc(
	// 		testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
	// 		resource.TestCheckResourceAttr(resourceName, "details.0.tenant_id", fmt.Sprintf("%v", value)),
	// 		testAccCheckAppdynamicscloudConnectionAzureIdNotEqual(connectionAzure_default, connectionAzure_updated),
	// 	),
	// })
	// value = searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.1")
	// testSteps = append(testSteps, resource.TestStep{
	// 	Config: CreateAccConnectionAzureUpdateRequiredDetailsSubscriptionId(rName),
	// 	Check: resource.ComposeTestCheckFunc(
	// 		testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
	// 		resource.TestCheckResourceAttr(resourceName, "details.0.subscription_id", fmt.Sprintf("%v", value)),
	// 		testAccCheckAppdynamicscloudConnectionAzureIdNotEqual(connectionAzure_default, connectionAzure_updated),
	// 	),
	// })

	return testSteps
}
func CreateAccConnectionAzureUpdateRequiredDisplayName(rName string) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	value := searchInObject(resourceConnectionAzureTest, "display_name.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_azure" "test" {
							
							display_name = "%v"
							description = "%v"
							state = "%v"
							configuration_id = appdynamicscloud_connection_configuration_azure.test.id
							details {
    
							                        
                            client_id = "%v"
                        
                            client_secret = "%v"
                        
                            tenant_id = "%v"
                        
                            subscription_id = "%v"
							}
			}
		`, value,
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}
func CreateAccConnectionAzureUpdateRequiredDetailsClientId(rName string) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	value := searchInObject(resourceConnectionAzureTest, "details.client_id.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_azure" "test" {
							display_name = "%v"
							description = "%v"
							state = "%v"
							configuration_id = appdynamicscloud_connection_configuration_azure.test.id
							details {
    
												
						    client_id = "%v"
                        
                            client_secret = "%v"
                        
                            tenant_id = "%v"
                        
                            subscription_id = "%v"
							}
			}
		`, searchInObject(resourceConnectionAzureTest, "display_name.valid.0"),
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		value,
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}
func CreateAccConnectionAzureUpdateRequiredDetailsClientSecret(rName string) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	value := searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_azure" "test" {
							display_name = "%v"
							description = "%v"
							state = "%v"
							configuration_id = appdynamicscloud_connection_configuration_azure.test.id
							details {
    
							                        
                            client_id = "%v"
					
						    client_secret = "%v"
                        
                            tenant_id = "%v"
                        
                            subscription_id = "%v"
							}
							depends_on = ["appdynamicscloud_connection_configuration_azure.test"]
			}
		`, searchInObject(resourceConnectionAzureTest, "display_name.valid.0"),
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		value,
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}
func CreateAccConnectionAzureUpdateRequiredDetailsTenantId(rName string) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	value := searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_azure" "test" {
							display_name = "%v"
							description = "%v"
							state = "%v"
							configuration_id = appdynamicscloud_connection_configuration_azure.test.id
							details {
    
							                        
                            client_id = "%v"
                        
                            client_secret = "%v"
					
						    tenant_id = "%v"
                        
                            subscription_id = "%v"
							}
			}
		`, searchInObject(resourceConnectionAzureTest, "display_name.valid.0"),
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		value,
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}
func CreateAccConnectionAzureUpdateRequiredDetailsSubscriptionId(rName string) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	value := searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_azure" "test" {
							display_name = "%v"
							description = "%v"
							state = "%v"
							configuration_id = appdynamicscloud_connection_configuration_azure.test.id
							details {
    
							                        
                            client_id = "%v"
                        
                            client_secret = "%v"
                        
                            tenant_id = "%v"
					
						    subscription_id = "%v"
							}
			}
		`, searchInObject(resourceConnectionAzureTest, "display_name.valid.0"),
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		value)
	return resource
}

func CreateAccConnectionAzureUpdatedAttrDescription(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_azure" "test" {
							display_name = "%v"
							description = "%v"
							state = "%v"
							configuration_id = appdynamicscloud_connection_configuration_azure.test.id
							details {
                            client_id = "%v"
                            client_secret = "%v"
                            tenant_id = "%v"
                            subscription_id = "%v"
							}
		}
	`, rName,
		value,
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}
func CreateAccConnectionAzureUpdatedAttrState(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_azure" "test" {
							display_name = "%v"
							description = "%v"
							state = "%v"
							configuration_id = appdynamicscloud_connection_configuration_azure.test.id
							details {
                            client_id = "%v"
                            client_secret = "%v"
                            tenant_id = "%v"
                            subscription_id = "%v"
							}
		}
	`, rName,
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		value,
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}
func CreateAccConnectionAzureUpdatedAttrConfigurationId(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAzureConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_azure" "test" {
							display_name = "%v"
							description = "%v"
							state = "%v"
							configuration_id = "%v"
							details {
                            client_id = "%v"
                            client_secret = "%v"
                            tenant_id = "%v"
                            subscription_id = "%v"
							}
		}
	`, rName,
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		value,
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
	return resource
}

func generateStepForUpdatedAttrConnectionAzure(rName string, resourceName string, connectionAzure_default, connectionAzure_updated *cloudconnectionapi.ConnectionResponse) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = searchInObject(resourceConnectionAzureTest, "description.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAzureUpdatedAttrDescription(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "description", v),
				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
			),
		})
	}
	// valid = searchInObject(resourceConnectionAzureTest, "state.valid").([]interface{})
	// for _, value := range valid {
	// 	// v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionAzureUpdatedAttrState(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),

	// 			resource.TestCheckResourceAttrWith("appdynamicscloud_connection_azure.test", "state", func(s string) error {
	// 				states := []string{"INACTIVE", "ACTIVE", "PENDING CONFIGURATION", "INCOMPLETE", "CONFIGURED", "INSUFFICIENT LICENSE", "ERROR", "WARNING", "CRITICAL"}
	// 				for _, val := range states {
	// 					if val == s {
	// 						return nil
	// 					}
	// 				}
	// 				return fmt.Errorf("Invalid State Value Obtained: %v", s)
	// 			}),
	// 			testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
	// 		),
	// 	})
	// }
	// valid = searchInObject(resourceConnectionAzureTest, "configuration_id.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionAzureUpdatedAttrConfigurationId(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "configuration_id", v),
	// 			testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
	// 		),
	// 	})
	// }
	return testSteps
}

func generateNegativeStepsConnectionAzure(rName string, resourceName string) []resource.TestStep {
	//Use Update Config Function with false value
	testSteps := make([]resource.TestStep, 0, 1)
	var invalid []interface{}
	invalid = searchInObject(resourceConnectionAzureTest, "state.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAzureUpdatedAttrState(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}
	invalid = searchInObject(resourceConnectionAzureTest, "configuration_id.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAzureUpdatedAttrConfigurationId(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["IsUUID"]),
		})
	}
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAzureConfig(rName),
	})
	return testSteps
}

// func CreateAccConnectionAzureMultipleConfig(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	multipleValues := searchInObject(resourceConnectionAzureTest, "display_name.multiple_valids").([]interface{})
// 	for i, val := range multipleValues {
// 		resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test%d" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							configuration_id = appdynamicscloud_connection_configuration_azure.test.id

// 							details {

//                             client_id = "%v"

//                             client_secret = "%v"

//                             tenant_id = "%v"

//                             subscription_id = "%v"

// 							}
// 			}
// 		`, i, val,
// 			searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
// 	}
// 	return resource
// }

func testAccCheckAppdynamicscloudConnectionAzureExists(name string, connectionAzure *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Connection Azure %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No azure connection id was set")
		}

		config := testAccProvider.Meta().(config)
		myctx, _, apiClient := initializeCloudConnectionClient(config)

		resp, _, err := apiClient.ConnectionsApi.GetConnection(myctx, rs.Primary.ID).Execute()
		if err != nil {
			return err
		}

		if resp.GetId() != rs.Primary.ID {
			return fmt.Errorf("Cloud Connection Azure %s not found", rs.Primary.ID)
		}
		*connectionAzure = *resp
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionAzureDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(config)
	myctx, _, apiClient := initializeCloudConnectionClient(config)
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "appdynamicscloud_connection_azure" {
			_, _, err := apiClient.ConnectionsApi.GetConnection(myctx, rs.Primary.ID).Execute()
			if err == nil {
				return fmt.Errorf("Cloud Connection Azure %s Still exists", rs.Primary.ID)
			} else {
				return nil
			}
		}
	}
	return nil
}

func testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure1, connectionAzure2 *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionAzure1.Id != connectionAzure2.Id {
			return fmt.Errorf("Connection Azure IDs are not equal")
		}
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionAzureIdNotEqual(connectionAzure1, connectionAzure2 *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionAzure1.Id == connectionAzure2.Id {
			return fmt.Errorf("Connection Azure IDs are equal")
		}
		return nil
	}
}

func getParentConnectionAzure(rName string) []string {
	t := []string{}
	t = append(t, getParentConnectionConfigurationAzure(rName)...)
	t = append(t, connectionAzureBlock(rName))
	return t
}

func connectionAzureBlock(rName string) string {
	return fmt.Sprintf(`
		resource  "appdynamicscloud_connection_azure" "test" {
						display_name = "%v"
						description = "%v"
						state = "%v"
						configuration_id = appdynamicscloud_connection_configuration_azure.test.id
                        details {
    
                            client_id = "%v"
                        
                            client_secret = "%v"
                        
                            tenant_id = "%v"
                        
                            subscription_id = "%v"
                        }
		}
	`, rName,
		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.client_secret.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.tenant_id.valid.0"),
		searchInObject(resourceConnectionAzureTest, "details.subscription_id.valid.0"))
}

// To eliminate duplicate resource block from slice of resource blocks
func createConnectionAzureConfig(configSlice []string) string {
	keys := make(map[string]bool)
	str := ""

	for _, entry := range configSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			str += entry
		}
	}

	return str
}
