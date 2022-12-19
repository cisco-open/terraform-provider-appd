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

var resourceConnectionConfigurationAWSTest = map[string]interface{}{
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
			"valid":           []interface{}{"us-east-1", "us-west-1"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"1le4sgpplf", "wtusot7s3z", "k59tedt7kx", "b1j3dbmw8j", "w9k4djpfyf", "p8yfk11huy", "pj2315wwa9", "oe9leltrpu", "2pqrg8dy6n", "7b3qd072xy", "waj52s6wkv", "mch4zxjleo", "6p62x14gka", "cfgfkv56zl", "pdemvpw8l1"},
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
				"valid":           []interface{}{"ebs", "ec2", "ecs", "elb", "rds"},
				"invalid":         []interface{}{10, 12.43},
				"multiple_valids": []interface{}{"ebs", "ec2", "ecs", "elb", "rds"},
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

func TestAccAppdynamicscloudConnectionConfigurationAWS_Basic(t *testing.T) {
	var connectionConfigurationAWS_default cloudConnectionApi.ConfigurationDetail
	var connectionConfigurationAWS_updated cloudConnectionApi.ConfigurationDetail
	resourceName := "appdynamicscloud_connection_configuration_aws.test"

	rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAWSDestroy,
		Steps: append([]resource.TestStep{
			{
				Config:      CreateAccConnectionConfigurationAWSWithoutDisplayName(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config: CreateAccConnectionConfigurationAWSConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, &connectionConfigurationAWS_default),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),

					resource.TestCheckResourceAttr(resourceName, "description", ""),

					resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				),
			},
			{
				Config: CreateAccConnectionConfigurationAWSConfigWithOptional(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, &connectionConfigurationAWS_updated),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.regions.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "details.0.regions.0", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.regions.1", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.polling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.excluded_keys.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"))),

					resource.TestCheckResourceAttr(resourceName, "details.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.name", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.excluded_keys.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"))),

					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))),

					testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(&connectionConfigurationAWS_default, &connectionConfigurationAWS_updated),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"details_service_default"},
			},
			// {
			// 	Config: CreateAccConnectionConfigurationAWSConfig(rName),
			// },
		}, generateStepForUpdatedRequiredAttrConnectionConfigurationAWS(rName, resourceName, &connectionConfigurationAWS_default, &connectionConfigurationAWS_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionConfigurationAWS_Update(t *testing.T) {
	var connectionConfigurationAWS_default cloudConnectionApi.ConfigurationDetail
	var connectionConfigurationAWS_updated cloudConnectionApi.ConfigurationDetail
	resourceName := "appdynamicscloud_connection_configuration_aws.test"
	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAWSDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionConfigurationAWSConfigWithOptional(rName),
				Check:  testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, &connectionConfigurationAWS_default),
			},
		}, generateStepForUpdatedAttrConnectionConfigurationAWS(rName, resourceName, &connectionConfigurationAWS_default, &connectionConfigurationAWS_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionConfigurationAWS_NegativeCases(t *testing.T) {
	resourceName := "appdynamicscloud_connection_configuration_aws.test"

	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAWSDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionConfigurationAWSConfigWithOptional(rName),
			},
		}, generateNegativeStepsConnectionConfigurationAWS(rName, resourceName)...),
	})
}

func TestAccAppdynamicscloudConnectionConfigurationAWS_MultipleCreateDelete(t *testing.T) {

	// [TODO]: Add makeTestVariable() to utils.go file
	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionConfigurationAWSDestroy,
		Steps: []resource.TestStep{
			{
				Config: CreateAccConnectionConfigurationAWSMultipleConfig(rName),
			},
		},
	})
}

func CreateAccConnectionConfigurationAWSWithoutDisplayName(rName string) string {
	resource := fmt.Sprintf(`
				resource  "appdynamicscloud_connection_configuration_aws" "test" {

									description = "%v"

									details {
    
									                        
                                        regions = ["%v","%v"]
                        

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
			`, searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAWSConfig(rName string) string {
	resource := fmt.Sprintf(`
		resource  "appdynamicscloud_connection_configuration_aws" "test" {
			display_name = "%v"
		}
	`, rName)
	return resource
}

func CreateAccConnectionConfigurationAWSConfigWithOptional(rName string) string {

	resource := fmt.Sprintf(`
		resource  "appdynamicscloud_connection_configuration_aws" "test" {

						display_name = "%v"

						description = "%v"

                        details {
    
                                                
                            regions = ["%v","%v"]
            

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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}

func generateStepForUpdatedRequiredAttrConnectionConfigurationAWS(rName string, resourceName string, connectionConfigurationAWS_default, connectionConfigurationAWS_updated *cloudConnectionApi.ConfigurationDetail) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	value := searchInObject(resourceConnectionConfigurationAWSTest, "display_name.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionConfigurationAWSUpdateRequiredDisplayName(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
			resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
		),
	})
	return testSteps
}
func CreateAccConnectionConfigurationAWSUpdateRequiredDisplayName(rName string) string {
	var resource string
	value := searchInObject(resourceConnectionConfigurationAWSTest, "display_name.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {
							
							display_name = "%v"

							description = "%v"

							details {
    
							                        
                            regions = ["%v","%v"]
            

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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAWSUpdatedAttrDescription(rName string, value interface{}) string {
	var resource string

	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {
							display_name = "%v"

							description = "%v"

							details {

                            regions = ["%v","%v"]

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
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsRegions(rName string, value interface{}) string {
	var resource string
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v"]


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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsPollingInterval(rName string, value interface{}) string {

	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]


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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		value,
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsPollingUnit(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]


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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}

// func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsImportTagsEnabled(rName string, value interface{}) string {
// 	resource := fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_configuration_aws" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							details {

// 						    regions = ["%v", "%v"]

//

// 						    polling {

//                                 interval = %v

//                                 unit = "%v"

//                               }

// 						        import_tags {

// 						            enabled = "%v"

// 						            excluded_keys = ["%v", "%v"]

// 						          }

//                                     tag_filter = "%v"

// 						            services {

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

//								}
//			}
//		`, rName,
//			searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
//
//
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
//			value,
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
//			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
//		return resource
//	}
func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsImportTagsExcludedKeys(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]


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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsTagFilter(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]


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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		value,
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}

// func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesImportTagsEnabled(rName string, value interface{}) string {
// 	resource := fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_configuration_aws" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							details {

// 						    regions = ["%v", "%v"]

//

// 						    polling {

//                                 interval = %v

//                                 unit = "%v"

//                               }

// 						        import_tags {

//                                     enabled = "%v"

//                                     excluded_keys = ["%v","%v"]

//                                   }

//                                     tag_filter = "%v"

// 						            services {

//                                         name = "%v"

// 						                import_tags {

// 						                    enabled = "%v"

// 						                    excluded_keys = ["%v", "%v"]

// 						                  }

//                                             tag_filter = "%v"

// 						                    polling {

//                                                 interval = %v

//                                                 unit = "%v"

//                                               }

// 						                      }

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
//
//
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.services.polling.unit.valid.0"))
// 	return resource
// }

func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesImportTagsExcludedKeys(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionConfigurationAWS(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionConfigurationAWSConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]


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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesTagFilter(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]


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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		value,
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesPollingInterval(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]


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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		value,
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesPollingUnit(rName string, value interface{}) string {
	resource := fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test" {

							display_name = "%v"

							description = "%v"

							details {

						    regions = ["%v", "%v"]


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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		value)
	return resource
}

func generateStepForUpdatedAttrConnectionConfigurationAWS(rName string, resourceName string, connectionConfigurationAWS_default, connectionConfigurationAWS_updated *cloudConnectionApi.ConfigurationDetail) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = searchInObject(resourceConnectionConfigurationAWSTest, "description.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDescription(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "description", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsRegions(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.regions.0", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsPollingInterval(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.polling.0.interval", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsPollingUnit(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.polling.0.unit", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}

	// valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsImportTagsEnabled(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.enabled", v),
	// 			testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
	// 		),
	// 	})
	// }
	valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsImportTagsExcludedKeys(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.import_tags.0.excluded_keys.0", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsTagFilter(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.tag_filter", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}
	// valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesImportTagsEnabled(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.enabled", v),
	// 			testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
	// 		),
	// 	})
	// }
	valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesImportTagsExcludedKeys(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.services.0.import_tags.0.excluded_keys.0", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesTagFilter(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.services.0.tag_filter", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesPollingInterval(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.0.interval", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}
	valid = searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesPollingUnit(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(resourceName, connectionConfigurationAWS_updated),
				resource.TestCheckResourceAttr(resourceName, "details.0.services.0.polling.0.unit", v),
				testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS_default, connectionConfigurationAWS_updated),
			),
		})
	}

	return testSteps
}

func generateNegativeStepsConnectionConfigurationAWS(rName string, resourceName string) []resource.TestStep {
	//Use Update Config Function with false value
	testSteps := make([]resource.TestStep, 0, 1)
	var invalid []interface{}
	invalid = searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionConfigurationAWSUpdatedAttrDetailsPollingUnit(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}

	invalid = searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionConfigurationAWSUpdatedAttrDetailsServicesPollingUnit(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}

	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionConfigurationAWSConfigWithOptional(rName),
	})
	return testSteps
}

func CreateAccConnectionConfigurationAWSMultipleConfig(rName string) string {
	var resource string
	multipleValues := searchInObject(resourceConnectionConfigurationAWSTest, "display_name.multiple_valids").([]interface{})
	for i, val := range multipleValues {
		resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_configuration_aws" "test%d" {

							display_name = "%v"

							description = "%v"

							details {

                            regions = ["%v","%v"]

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
			searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),

			searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
			searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
	}
	return resource
}

func testAccCheckAppdynamicscloudConnectionConfigurationAWSExists(name string, connectionConfigurationAWS *cloudConnectionApi.ConfigurationDetail) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Connection Configuration AWS %s not found", name)
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
			return fmt.Errorf("Cloud Connection Configuration AWS %s not found", rs.Primary.ID)
		}
		*connectionConfigurationAWS = *resp
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionConfigurationAWSDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(config)
	myCtx, _, apiClient := initializeCloudConnectionClient(config)
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "appdynamicscloud_connection_configuration_aws" {
			_, _, err := apiClient.ConfigurationsApi.GetConfiguration(myCtx, rs.Primary.ID).Execute()
			if err == nil {
				return fmt.Errorf("Cloud Connection Configuration AWS %s Still exists", rs.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckAppdynamicscloudConnectionConfigurationAWSIdEqual(connectionConfigurationAWS1, connectionConfigurationAWS2 *cloudConnectionApi.ConfigurationDetail) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		if connectionConfigurationAWS1.Id != connectionConfigurationAWS2.Id {
			return fmt.Errorf("Connection Configuration AWS IDs are not equal")
		}
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionConfigurationAWSIdNotEqual(connectionConfigurationAWS1, connectionConfigurationAWS2 *cloudConnectionApi.ConfigurationDetail) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionConfigurationAWS1.Id == connectionConfigurationAWS2.Id {
			return fmt.Errorf("Connection Configuration AWS IDs are equal")
		}
		return nil
	}
}

func getParentConnectionConfigurationAWS(rName string) []string {
	t := []string{}
	t = append(t, connectionConfigurationAWSBlock(rName))
	return t
}

func connectionConfigurationAWSBlock(rName string) string {
	return fmt.Sprintf(`
		resource  "appdynamicscloud_connection_configuration_aws" "test" {

						display_name = "%v"


						description = "%v"


                        details {
    
                                                
                            regions = ["%v","%v"]
            

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
		searchInObject(resourceConnectionConfigurationAWSTest, "description.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.regions.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.polling.unit.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.name.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionConfigurationAWSTest, "details.services.polling.unit.valid.0"))
}

// To eliminate duplicate resource block from slice of resource blocks
func createConnectionConfigurationAWSConfig(configSlice []string) string {
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
