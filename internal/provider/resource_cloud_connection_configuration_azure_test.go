package provider

import (
	"fmt"
	"regexp"
	"testing"

	cloudConnectionApi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var resourceConnectionConfigurationAzureTest = map[string]interface{}{
	"display_name": map[string]interface{}{
		"valid":           []interface{}{"gzruzybnxs", "37nri64m0y", "iec4uszee1", "p046m4nkta"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"xwhw1oyi6i", "wgkdbkjm91", "gg6du8kdjr"},
	},

	"description": map[string]interface{}{
		"valid":           []interface{}{"l6pirpl9ja", "kz8yho0kiy", "vde9dgkual", "zjtio8m82l"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"1nqz9fbrf1", "phx7abhmg0", "wn0jl3kk9i", "2l8epqwfqh", "j2l2hcv6p1", "tusgk21gr2", "29smli4dyc", "98k5jvbu01", "ovkyaeze67", "9tweg4k350", "5givjlmmce", "bqqjkcy0t6", "d36paaymso", "9djznf7bk5", "y2qgeno36r"},
	},

	"details": map[string]interface{}{
		"regions": map[string]interface{}{
			"valid":           []interface{}{"eastus", "westus"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"1le4sgpplf", "wtusot7s3z", "k59tedt7kx", "b1j3dbmw8j", "w9k4djpfyf", "p8yfk11huy", "pj2315wwa9", "oe9leltrpu", "2pqrg8dy6n", "7b3qd072xy", "waj52s6wkv", "mch4zxjleo", "6p62x14gka", "cfgfkv56zl", "pdemvpw8l1"},
		},

		"resource_groups": map[string]interface{}{
			"valid":           []interface{}{"resourceGroup1", "resourceGroup2", "resourceGroup3"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"aluicyyh53", "yeimb26ddv", "hwgt0et47j", "g673j39erk", "gnzw38906u", "tkv7lu5sn0", "hcdmicmzra", "q3h5befzo1", "i6lfco9qy3", "m8rljrtj3z", "7zvyt6ywmf", "1yat3mdd4p", "yb1hbuy9wy", "lw9tos3h03", "v06b8a3mjp"},
		},

		"polling": map[string]interface{}{
			"interval": map[string]interface{}{
				"valid":           []interface{}{5},
				"invalid":         []interface{}{"random", 10.023},
				"multiple_valids": []interface{}{5},
			},

			"unit": map[string]interface{}{
				"valid":           []interface{}{"minute"},
				"invalid":         []interface{}{"nn0t1edqsy"},
				"multiple_valids": []interface{}{"minute"},
			},
		},

		"import_tags": map[string]interface{}{
			"enabled": map[string]interface{}{
				"valid":           []interface{}{true, false},
				"invalid":         []interface{}{"random", 10},
				"multiple_valids": []interface{}{true, false},
			},

			"excluded_keys": map[string]interface{}{
				"valid":           []interface{}{"key1", "key2", "key3", "key4"},
				"invalid":         []interface{}{10, 12.43},
				"multiple_valids": []interface{}{"key1", "key2", "key3", "key4"},
			},
		},

		"tag_filter": map[string]interface{}{
			"valid":           []interface{}{"tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)", "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)", "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)"},
		},

		"services": map[string]interface{}{
			"name": map[string]interface{}{
				"valid":           []interface{}{"vm", "disk", "postgresql", "lb", "sql"},
				"invalid":         []interface{}{10, 12.43},
				"multiple_valids": []interface{}{"vm", "disk", "postgresql", "lb", "sql"},
			},

			"import_tags": map[string]interface{}{
				"enabled": map[string]interface{}{
					"valid":           []interface{}{true, false},
					"invalid":         []interface{}{"random", 10},
					"multiple_valids": []interface{}{true, false},
				},

				"excluded_keys": map[string]interface{}{
					"valid":           []interface{}{"key1", "key2", "key3", "key4"},
					"invalid":         []interface{}{10, 12.43},
					"multiple_valids": []interface{}{"key1", "key2", "key3", "key4"},
				},
			},

			"tag_filter": map[string]interface{}{
				"valid":           []interface{}{"tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)", "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)"},
				"invalid":         []interface{}{10, 12.43},
				"multiple_valids": []interface{}{"tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)", "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)"},
			},

			"polling": map[string]interface{}{
				"interval": map[string]interface{}{
					"valid":           []interface{}{5},
					"invalid":         []interface{}{"random", 10.023},
					"multiple_valids": []interface{}{5},
				},

				"unit": map[string]interface{}{
					"valid":           []interface{}{"minute"},
					"invalid":         []interface{}{"nn0t1edqsy"},
					"multiple_valids": []interface{}{"minute"},
				},
			},
		},
	},
}

func TestAccAppdynamicscloudConnectionConfigurationAzure_Basic(t *testing.T) {
	var connectionConfigurationAzure_default cloudConnectionApi.ConfigurationDetail
	var connectionConfigurationAzure_updated cloudConnectionApi.ConfigurationDetail
	resourceName := "appdynamicscloud_connection_configuration_azure.test"

	rName := makeTestVariable(acctest.RandString(5))
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAzureDestroy,
		Steps: append([]resource.TestStep{
			{
				Config:      CreateAccConnectionConfigurationAzureWithoutDisplayName(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config: CreateAccConnectionConfigurationAzureConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, &connectionConfigurationAzure_default),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),

					resource.TestCheckResourceAttr(resourceName, "description", ""),

					resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				),
			},
			{
				Config: CreateAccConnectionConfigurationAzureConfigWithOptional(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, &connectionConfigurationAzure_updated),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.regions.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "details.0.regions.0", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.regions.1", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.resource_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "details.0.resource_groups.0", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.resource_groups.1", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.polling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.excluded_keys.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"))),

					resource.TestCheckResourceAttr(resourceName, "details.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.name", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.excluded_keys.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"))),

					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))),

					testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(&connectionConfigurationAzure_default, &connectionConfigurationAzure_updated),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"details_service_default"},
			},
		}, generateStepForUpdatedRequiredAttrConnectionConfigurationAzure(rName, resourceName, &connectionConfigurationAzure_default, &connectionConfigurationAzure_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionConfigurationAzure_Update(t *testing.T) {
	var connectionConfigurationAzure_default cloudConnectionApi.ConfigurationDetail
	var connectionConfigurationAzure_updated cloudConnectionApi.ConfigurationDetail
	resourceName := "appdynamicscloud_connection_configuration_azure.test"
	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAzureDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionConfigurationAzureConfigWithOptional(rName),
				Check:  testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, &connectionConfigurationAzure_default),
			},
		}, generateStepForUpdatedAttrConnectionConfigurationAzure(rName, resourceName, &connectionConfigurationAzure_default, &connectionConfigurationAzure_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionConfigurationAzure_NegativeCases(t *testing.T) {
	resourceName := "appdynamicscloud_connection_configuration_azure.test"

	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAzureDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionConfigurationAzureConfigWithOptional(rName),
			},
		}, generateNegativeStepsConnectionConfigurationAzure(rName, resourceName)...),
	})
}

func TestAccAppdynamicscloudConnectionConfigurationAzure_MultipleCreateDelete(t *testing.T) {

	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: CreateAccConnectionConfigurationAzureMultipleConfig(rName),
			},
		},
	})
}

func CreateAccConnectionConfigurationAzureWithoutDisplayName(rName string) string {
	resource := fmt.Sprintf(`
				resource  "appdynamicscloud_connection_configuration_azure" "test" {

									description = "%v"

									details {
    
									                        
                                        regions = ["%v","%v"]
                        
                                        resource_groups = ["%v","%v"]

                                        polling {
                                                    
                                            interval = %v
                        
                                            unit = "%v"

                                          }

                                            import_tags {
                                                    
                                                enabled = "%v"
                        
                                                excluded_keys = ["%v","%v"]

                                              }
                        
                                                tag_filter = "%v"

                                                services {
                                                    
                                                    name = "%v"

                                                    import_tags {
                                                    
                                                        enabled = "%v"
                        
                                                        excluded_keys = ["%v","%v"]

                                                      }
                        
                                                        tag_filter = "%v"

                                                        polling {
                                                    
                                                            interval = %v
                        
                                                            unit = "%v"

                                                          }

                                                          }

									}
				}
			`, searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAzureConfig(rName string) string {
	resource := fmt.Sprintf(`
		resource  "appdynamicscloud_connection_configuration_azure" "test" {
			display_name = "%v"
		}
	`, rName)
	return resource
}

func CreateAccConnectionConfigurationAzureConfigWithOptional(rName string) string {

	resource := fmt.Sprintf(`
		resource  "appdynamicscloud_connection_configuration_azure" "test" {

						display_name = "%v"

						description = "%v"

                        details {
    
                                                
                            regions = ["%v","%v"]
                        
                            resource_groups = ["%v","%v"]

                            polling {
                                                    
                                interval = %v
                        
                                unit = "%v"

                              }

                                import_tags {
                                                    
                                    enabled = "%v"
                        
                                    excluded_keys = ["%v","%v"]

                                  }
                        
                                    tag_filter = "%v"

                                    services {
                                                    
                                        name = "%v"

                                        import_tags {
                                                    
                                            enabled = "%v"
                        
                                            excluded_keys = ["%v","%v"]

                                          }
                        
                                            tag_filter = "%v"

                                            polling {
                                                    
                                                interval = %v
                        
                                                unit = "%v"

                                              }

                                              }

                        }
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}

func generateStepForUpdatedRequiredAttrConnectionConfigurationAzure(rName string, resourceName string, connectionConfigurationAzure_default, connectionConfigurationAzure_updated *cloudConnectionApi.ConfigurationDetail) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	value := searchInObject(resourceConnectionConfigurationAzureTest, "display_name.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionConfigurationAzureUpdateRequiredDisplayName(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
			resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
		),
	})
	return testSteps
}
func CreateAccConnectionConfigurationAzureUpdateRequiredDisplayName(rName string) string {
	var resource string
	value := searchInObject(resourceConnectionConfigurationAzureTest, "display_name.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {
							
							display_name = "%v"

							description = "%v"

							details {
    
							                        
                            regions = ["%v","%v"]
                        
                            resource_groups = ["%v","%v"]

                            polling {
                                                    
                                interval = %v
                        
                                unit = "%v"

                              }

                                import_tags {
                                                    
                                    enabled = "%v"
                        
                                    excluded_keys = ["%v","%v"]

                                  }
                        
                                    tag_filter = "%v"

                                    services {
                                                    
                                        name = "%v"

                                        import_tags {
                                                    
                                            enabled = "%v"
                        
                                            excluded_keys = ["%v","%v"]

                                          }
                        
                                            tag_filter = "%v"

                                            polling {
                                                    
                                                interval = %v
                        
                                                unit = "%v"

                                              }

                                              }

							}
			}
		`, value,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAzureUpdatedAttrDescription(rName string, value interface{}) string {
	var resource string

	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {
							display_name = "%v"

							description = "%v"

							details {

                            regions = ["%v","%v"]

                            resource_groups = ["%v","%v"]

                            polling {

                                interval = %v

                                unit = "%v"

                              }

                                import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

                                    services {

                                        name = "%v"

                                        import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

                                            polling {

                                                interval = %v

                                                unit = "%v"

                                              }

                                              }

							}
		}
	`, rName,
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsRegions(rName string, value interface{}) string {
	var resource string
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v"]

						    resource_groups = ["%v", "%v"]

						    polling {

                                interval = %v

                                unit = "%v"

                              }

						        import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

						            services {

                                        name = "%v"

                                        import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

                                            polling {

                                                interval = %v

                                                unit = "%v"

                                              }

                                              }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsResourceGroups(rName string, value interface{}) string {

	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]

						    resource_groups = ["%v"]

						    polling {

                                interval = %v

                                unit = "%v"

                              }

						        import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

						            services {

                                        name = "%v"

                                        import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

                                            polling {

                                                interval = %v

                                                unit = "%v"

                                              }

                                              }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsPollingInterval(rName string, value interface{}) string {

	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]

						    resource_groups = ["%v", "%v"]

						    polling {

						        interval = %v

                                unit = "%v"

						      }

						        import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

						            services {

                                        name = "%v"

                                        import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

                                            polling {

                                                interval = %v

                                                unit = "%v"

                                              }

                                              }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsPollingUnit(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]

						    resource_groups = ["%v", "%v"]

						    polling {

                                interval = %v

						        unit = "%v"

						      }

						        import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

						            services {

                                        name = "%v"

                                        import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

                                            polling {

                                                interval = %v

                                                unit = "%v"

                                              }

                                              }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsImportTagsExcludedKeys(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]

						    resource_groups = ["%v", "%v"]

						    polling {

                                interval = %v

                                unit = "%v"

                              }

						        import_tags {

                                    enabled = "%v"

						            excluded_keys = ["%v"]

						          }

                                    tag_filter = "%v"

						            services {

                                        name = "%v"

                                        import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

                                            polling {

                                                interval = %v

                                                unit = "%v"

                                              }

                                              }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsTagFilter(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]

						    resource_groups = ["%v", "%v"]

						    polling {

                                interval = %v

                                unit = "%v"

                              }

						        import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

						            tag_filter = "%v"

						            services {

                                        name = "%v"

                                        import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

                                            polling {

                                                interval = %v

                                                unit = "%v"

                                              }

                                              }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsServicesImportTagsExcludedKeys(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionConfigurationAzure(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionConfigurationAzureConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]

						    resource_groups = ["%v", "%v"]

						    polling {

                                interval = %v

                                unit = "%v"

                              }

						        import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

						            services {

                                        name = "%v"

						                import_tags {

                                            enabled = "%v"

						                    excluded_keys = ["%v"]

						                  }

                                            tag_filter = "%v"

						                    polling {

                                                interval = %v

                                                unit = "%v"

                                              }

						                      }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsServicesTagFilter(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]

						    resource_groups = ["%v", "%v"]

						    polling {

                                interval = %v

                                unit = "%v"

                              }

						        import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

						            services {

                                        name = "%v"

						                import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

						                    tag_filter = "%v"

						                    polling {

                                                interval = %v

                                                unit = "%v"

                                              }

						                      }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsServicesPollingInterval(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]

						    resource_groups = ["%v", "%v"]

						    polling {

                                interval = %v

                                unit = "%v"

                              }

						        import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

						            services {

                                        name = "%v"

						                import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

						                    polling {

						                        interval = %v

                                                unit = "%v"

						                      }

						                      }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionConfigurationAzureUpdatedAttrDetailsServicesPollingUnit(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]

						    resource_groups = ["%v", "%v"]

						    polling {

                                interval = %v

                                unit = "%v"

                              }

						        import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

						            services {

                                        name = "%v"

						                import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

						                    polling {

                                                interval = %v

						                        unit = "%v"

						                      }

						                      }

							}
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		value)
	return resource
}

func generateStepForUpdatedAttrConnectionConfigurationAzure(rName string, resourceName string, connectionConfigurationAzure_default, connectionConfigurationAzure_updated *cloudConnectionApi.ConfigurationDetail) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = searchInObject(resourceConnectionConfigurationAzureTest, "description.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDescription(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "description", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsRegions(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.regions.0", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsResourceGroups(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.resource_groups.0", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsPollingInterval(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.polling.0.interval", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsPollingUnit(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.polling.0.unit", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsImportTagsExcludedKeys(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.excluded_keys.0", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsTagFilter(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.tag_filter", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsServicesImportTagsExcludedKeys(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.excluded_keys.0", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsServicesTagFilter(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.services.0.tag_filter", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsServicesPollingInterval(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.0.interval", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAzureUpdatedAttrDetailsServicesPollingUnit(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(resourceName, connectionConfigurationAzure_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.0.unit", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure_default, connectionConfigurationAzure_updated),
			),
		})
	}

	return testSteps
}

func generateNegativeStepsConnectionConfigurationAzure(rName string, resourceName string) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var invalid []interface{}
	invalid = searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionConfigurationAzureUpdatedAttrDetailsPollingUnit(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}

	invalid = searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionConfigurationAzureUpdatedAttrDetailsServicesPollingUnit(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}

	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionConfigurationAzureConfigWithOptional(rName),
	})
	return testSteps
}

func CreateAccConnectionConfigurationAzureMultipleConfig(rName string) string {
	var resource string
	multipleValues := searchInObject(resourceConnectionConfigurationAzureTest, "display_name.multiple_valids").([]interface{})
	for i, val := range multipleValues {
		resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_azure" "test%d" {

							display_name = "%v"

							description = "%v"

							details {

                            regions = ["%v","%v"]

                            resource_groups = ["%v","%v"]

                            polling {

                                interval = %v

                                unit = "%v"

                              }

                                import_tags {

                                    enabled = "%v"

                                    excluded_keys = ["%v","%v"]

                                  }

                                    tag_filter = "%v"

                                    services {

                                        name = "%v"

                                        import_tags {

                                            enabled = "%v"

                                            excluded_keys = ["%v","%v"]

                                          }

                                            tag_filter = "%v"

                                            polling {

                                                interval = %v

                                                unit = "%v"

                                              }

                                              }

							}
			}
		`, i, val,
			searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
			searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
	}
	return resource
}

func testAccCheckAppdynamicscloudConnectionConfigurationAzureExists(name string, connectionConfigurationAzure *cloudConnectionApi.ConfigurationDetail) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Connection Configuration Azure %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No configuration id was set")
		}

		config := testAccProvider.Meta().(config)
		myCtx, _, apiClient := initializeCloudConnectionClient(config)

		resp, _, err := apiClient.ConfigurationsApi.GetConfiguration(myCtx, rs.Primary.ID).Execute()
		if err != nil {
			return err
		}

		if resp.GetId() != rs.Primary.ID {
			return fmt.Errorf("Cloud Connection Configuration Azure %s not found", rs.Primary.ID)
		}
		*connectionConfigurationAzure = *resp
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionConfigurationAzureDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(config)
	myCtx, _, apiClient := initializeCloudConnectionClient(config)
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "appdynamicscloud_connection_configuration_azure" {
			_, _, err := apiClient.ConfigurationsApi.GetConfiguration(myCtx, rs.Primary.ID).Execute()
			if err == nil {
				return fmt.Errorf("Cloud Connection Configuration Azure %s Still exists", rs.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckAppdynamicscloudConnectionConfigurationAzureIdEqual(connectionConfigurationAzure1, connectionConfigurationAzure2 *cloudConnectionApi.ConfigurationDetail) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		if connectionConfigurationAzure1.Id != connectionConfigurationAzure2.Id {
			return fmt.Errorf("Connection Configuration Azure IDs are not equal")
		}
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionConfigurationAzureIdNotEqual(connectionConfigurationAzure1, connectionConfigurationAzure2 *cloudConnectionApi.ConfigurationDetail) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionConfigurationAzure1.Id == connectionConfigurationAzure2.Id {
			return fmt.Errorf("Connection Configuration Azure IDs are equal")
		}
		return nil
	}
}

func getParentConnectionConfigurationAzure(rName string) []string {
	t := []string{}
	t = append(t, connectionConfigurationAzureBlock(rName))
	return t
}

func connectionConfigurationAzureBlock(rName string) string {
	return fmt.Sprintf(`
		resource  "appdynamicscloud_connection_configuration_azure" "test" {

						display_name = "%v"


						description = "%v"


                        details {
    
                                                
                            regions = ["%v","%v"]
                        
                            resource_groups = ["%v","%v"]

                            polling {
                                                    
                                interval = %v
                        
                                unit = "%v"

                              }

                                import_tags {
                                                    
                                    enabled = "%v"
                        
                                    excluded_keys = ["%v","%v"]

                                  }
                        
                                    tag_filter = "%v"

                                    services {
                                                    
                                        name = "%v"

                                        import_tags {
                                                    
                                            enabled = "%v"
                        
                                            excluded_keys = ["%v","%v"]

                                          }
                        
                                            tag_filter = "%v"

                                            polling {
                                                    
                                                interval = %v
                        
                                                unit = "%v"

                                              }

                                              }

                        }
		}
	`, rName,
		searchInObject(resourceConnectionConfigurationAzureTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.resource_groups.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAzureTest, "details.services.polling.unit.valid.0"))
}

// To eliminate duplicate resource block from slice of resource blocks
func createConnectionConfigurationAzureConfig(configSlice []string) string {
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
