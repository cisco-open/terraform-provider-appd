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

// const connectionAwsRoleSelfRequiredCount = 2

var awsProvider = map[string]resource.ExternalProvider{
	"aws": {
		Source: "hashicorp/aws",
	},
}

var resourceConnectionAwsRoleTest = map[string]interface{}{
	"display_name": map[string]interface{}{
		"valid":           []interface{}{"mf29ikrvuq", "9gihbn0pxs", "4a98kcdsoe", "80e1pxmdgt"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"dnx8qbdcwc", "ull5dyttec", "nj0zb0dq6h", "6fdwm6m1yn", "goghukna96", "c5vgax4len", "2m4wq4xxr1", "8rq857h4op", "zloufg82ka", "3h1qvjg9q4", "dbb9baoy7m", "fi28rvmzrf", "23vnv8qokp", "nq6tqe9pms", "f7esp3bvn8"},
	},

	"description": map[string]interface{}{
		"valid":           []interface{}{"a4azty49iw", "2e10vrjady", "h5zyu013gw", "9t2z4sgyfd"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"tjl9ithmen", "w7mc5wonx6", "6ejzhbd0bs", "2018j431yq", "9apw3rbx7m", "62bmkn0oxq", "o93crtl3m2", "unht0g3tqz", "50aynd1ltp", "d43uivb409", "aqujh09yo5", "b1anqkem7c", "h618adqt7b", "y7leai85xi", "a8an6sa0hi"},
	},

	"state": map[string]interface{}{
		"valid":           []interface{}{"ACTIVE","INACTIVE"},
		"invalid":         []interface{}{"gv2xnyi3xn"},
		"multiple_valids": []interface{}{"ACTIVE","INACTIVE"},
	},

	"connection_details": map[string]interface{}{
		"access_type": map[string]interface{}{
			"valid":           []interface{}{"role_delegation"},
			"invalid":         []interface{}{"rbnigdc04i"},
			"multiple_valids": []interface{}{"role_delegation"},
		},

		"account_id": map[string]interface{}{
			"valid":           []interface{}{"860850072464", "860850072464"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"860850072464", "860850072464"},
		},
	},

	"configuration_details": map[string]interface{}{
		"regions": map[string]interface{}{
			"valid":           []interface{}{"us-east-1", "us-west-1"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"us-east-1", "us-west-1"},
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
				"valid":           []interface{}{"j0c8yg2w9u", "08rw88zrgx", "ypg30vq7ya", "bbmn5sr18j"},
				"invalid":         []interface{}{10, 12.43},
				"multiple_valids": []interface{}{"q1l5yw0dwm", "qomyte9mgd", "nea1hrghlg", "7m1kwkgbpg", "i4x2trfu1y", "kw1px3pvr7", "wdc3wip8j5", "q0ms78lway", "s3oofcpdsv", "kud9r0c6nt", "bv74ohbrhe", "0llxadx6v2", "vfr0z4ltkt", "jgmnz221vm", "1hslpndgld"},
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
					"valid":           []interface{}{"j0c8yg2w9u", "08rw88zrgx", "ypg30vq7ya", "bbmn5sr18j"},
					"invalid":         []interface{}{10, 12.43},
					"multiple_valids": []interface{}{"q1l5yw0dwm", "qomyte9mgd", "nea1hrghlg", "7m1kwkgbpg", "i4x2trfu1y", "kw1px3pvr7", "wdc3wip8j5", "q0ms78lway", "s3oofcpdsv", "kud9r0c6nt", "bv74ohbrhe", "0llxadx6v2", "vfr0z4ltkt", "jgmnz221vm", "1hslpndgld"},
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

func TestAccAppdynamicscloudConnectionAwsRole_Basic(t *testing.T) {
	var connectionAwsRole_default cloudconnectionapi.ConnectionResponse
	var connectionAwsRole_updated cloudconnectionapi.ConnectionResponse
	resourceName := "appdynamicscloud_connection_aws.test"

	rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		ExternalProviders: awsProvider,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsRoleDestroy,
		Steps: append([]resource.TestStep{
			{
				Config:      CreateAccConnectionAwsRoleWithoutDisplayName(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccConnectionAwsRoleWithoutConnectionDetails(rName),
				ExpectError: regexp.MustCompile(`Insufficient connection_details blocks`),
			},
			{
				Config: CreateAccConnectionAwsRoleConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, &connectionAwsRole_default),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),

					resource.TestCheckResourceAttr(resourceName, "description", ""),

					// resource.TestCheckResourceAttr(resourceName, "state", ""),

					resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_type", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.account_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				),
			},
			{
				Config: CreateAccConnectionAwsRoleConfigWithOptional(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, &connectionAwsRole_updated),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"))),
					// resource.TestCheckResourceAttr(resourceName, "state", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_type", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.account_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.name", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))),

					testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(&connectionAwsRole_default, &connectionAwsRole_updated),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_details", "configuration_details_service_default"},
			},
			{
				Config: CreateAccConnectionAwsRoleConfigWithOptional(rName),
			},
		}, generateStepForUpdatedRequiredAttrConnectionAwsRole(rName, resourceName, &connectionAwsRole_default, &connectionAwsRole_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionAwsRole_Update(t *testing.T) {
	var connectionAwsRole_default cloudconnectionapi.ConnectionResponse
	var connectionAwsRole_updated cloudconnectionapi.ConnectionResponse
	resourceName := "appdynamicscloud_connection_aws.test"
	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		ExternalProviders: awsProvider,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsRoleDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionAwsRoleConfig(rName),
				Check:  testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, &connectionAwsRole_default),
			},
		}, generateStepForUpdatedAttrConnectionAwsRole(rName, resourceName, &connectionAwsRole_default, &connectionAwsRole_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionAwsRole_NegativeCases(t *testing.T) {
	resourceName := "appdynamicscloud_connection_aws.test"

	// [TODO]: Add makeTestVariable() to utils.go file
	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		ExternalProviders: awsProvider,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsRoleDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionAwsRoleConfig(rName),
			},
		}, generateNegativeStepsConnectionAwsRole(rName, resourceName)...),
	})
}

// func TestAccAppdynamicscloudConnectionAwsRole_MultipleCreateDelete(t *testing.T) {

// 	// [TODO]: Add makeTestVariable() to utils.go file
// 	rName := makeTestVariable(acctest.RandString(5))

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsRoleDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: CreateAccConnectionAwsRoleMultipleConfig(rName),
// 			},
// 		},
// 	})
// }

func CreateAccConnectionAwsRoleWithoutDisplayName(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += role_attachment
	resource += fmt.Sprintf(`
				resource  "appdynamicscloud_connection_aws" "test" {

									description = "%v"

									state = "%v"

									connection_details {
    
									                        
                                        access_type = "%v"
                        
                                        account_id = "%v"

									}

									configuration_details {
    
									                        
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
			`, searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsRoleWithoutConnectionDetails(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += role_attachment
	resource += fmt.Sprintf(`
				resource  "appdynamicscloud_connection_aws" "test" {

									display_name = "%v"

									description = "%v"

									state = "%v"

									configuration_details {
    
									                        
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
			`, searchInObject(resourceConnectionAwsRoleTest, "display_name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsRoleConfig(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += role_attachment
	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws" "test" {


							display_name = "%v"

							connection_details {
    
							 

						          access_type = "%v"
 

						          account_id = "%v"

							}
		}
	`, rName,
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"))
	return resource
}

func CreateAccConnectionAwsRoleConfigWithOptional(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]

	resource += createConnectionAwsRoleConfig(parentResources)
	resource += role_attachment
	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws" "test" {

						display_name = "%v"

						description = "%v"

						state = "%v"

                        connection_details {
    
                                                
                            access_type = "%v"
                        
                            account_id = "%v"

                        }

                        configuration_details {
    
                                                
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func generateStepForUpdatedRequiredAttrConnectionAwsRole(rName string, resourceName string, connectionAwsRole_default, connectionAwsRole_updated *cloudconnectionapi.ConnectionResponse) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var value interface{}
	value = searchInObject(resourceConnectionAwsRoleTest, "display_name.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAwsRoleUpdateRequiredDisplayName(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
			resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
		),
	})
	// value = searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.1")
	// testSteps = append(testSteps, resource.TestStep{
	// 	Config: CreateAccConnectionAwsRoleUpdateRequiredConnectionDetailsAccessType(rName),
	// 	Check: resource.ComposeTestCheckFunc(
	// 		testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
	// 		resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_type", fmt.Sprintf("%v", value)),
	// 		testAccCheckAppdynamicscloudConnectionAwsRoleIdNotEqual(connectionAwsRole_default, connectionAwsRole_updated),
	// 	),
	// })
	// value = searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.1")
	// testSteps = append(testSteps, resource.TestStep{
	// 	Config: CreateAccConnectionAwsRoleUpdateRequiredConnectionDetailsAccountId(rName),
	// 	Check: resource.ComposeTestCheckFunc(
	// 		testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
	// 		resource.TestCheckResourceAttr(resourceName, "connection_details.0.account_id", fmt.Sprintf("%v", value)),
	// 		testAccCheckAppdynamicscloudConnectionAwsRoleIdNotEqual(connectionAwsRole_default, connectionAwsRole_updated),
	// 	),
	// })

	return testSteps
}
func CreateAccConnectionAwsRoleUpdateRequiredDisplayName(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	value := searchInObject(resourceConnectionAwsRoleTest, "display_name.valid.1")
	resource += role_attachment
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {
							
							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							                        
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

// func CreateAccConnectionAwsRoleUpdateRequiredConnectionDetailsAccessType(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsRoleConfig(parentResources)
// 	value := searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.1")
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_aws" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {

// 						    access_type = "%v"

//                             account_id = "%v"

// 							}

// 							configuration_details {

//                             regions = ["%v","%v"]

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
// 		`, searchInObject(resourceConnectionAwsRoleTest, "display_name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAwsRoleUpdateRequiredConnectionDetailsAccountId(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsRoleConfig(parentResources)
// 	value := searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.1")
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_aws" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {

//                             access_type = "%v"

// 						    account_id = "%v"

// 							}

// 							configuration_details {

//                             regions = ["%v","%v"]

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
// 		`, searchInObject(resourceConnectionAwsRoleTest, "display_name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
// 	return resource
// }

func CreateAccConnectionAwsRoleUpdatedAttrDescription(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"
							
							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							                        
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
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsRoleUpdatedAttrState(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"
							
							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							                        
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsRegions(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsPollingInterval(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsRoleConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_aws" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {

//                             access_type = "%v"

//                             account_id = "%v"

// 							}

// 							configuration_details {

// 						    regions = ["%v", "%v"]

// 						    polling {

// 						        interval = %v

//                                 unit = "%v"

// 						      }

// 						        import_tags {

//                                     enabled = "%v"

//                                     excluded_keys = ["%v","%v"]

//                                   }

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
//			searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
//			value,
//			searchInObject(resourceConnectionAwsRoleTest, "polling.unit.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
//		return resource
//	}
func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsPollingUnit(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsImportTagsEnabled(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {

                            access_type = "%v"

                            account_id = "%v"

							}

							configuration_details {

						    regions = ["%v", "%v"]

						    polling {

                                interval = %v

                                unit = "%v"

                              }

						        import_tags {

						            enabled = "%v"

						            excluded_keys = ["%v", "%v"]

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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsImportTagsExcludedKeys(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsTagFilter(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesImportTagsEnabled(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {

                            access_type = "%v"

                            account_id = "%v"

							}

							configuration_details {

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

						                    excluded_keys = ["%v", "%v"]

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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesImportTagsExcludedKeys(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesTagFilter(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesPollingInterval(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsRoleConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_aws" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {

//                             access_type = "%v"

//                             account_id = "%v"

// 							}

// 							configuration_details {

// 						    regions = ["%v", "%v"]

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

//                                             enabled = "%v"

//                                             excluded_keys = ["%v","%v"]

//                                           }

//                                             tag_filter = "%v"

// 						                    polling {

// 						                        interval = %v

//                                                 unit = "%v"

// 						                      }

// 						                      }

//								}
//			}
//		`, rName,
//			searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
//			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "services.name.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.enabled.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.excluded_keys.valid.0"),
//			searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.excluded_keys.valid.1"),
//			searchInObject(resourceConnectionAwsRoleTest, "services.tag_filter.valid.0"),
//			value,
//			searchInObject(resourceConnectionAwsRoleTest, "polling.unit.valid.0"))
//		return resource
//	}
func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesPollingUnit(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsRole(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsRoleConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            account_id = "%v"

							}

							configuration_details {
    
							
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		value)
	return resource
}

func generateStepForUpdatedAttrConnectionAwsRole(rName string, resourceName string, connectionAwsRole_default, connectionAwsRole_updated *cloudconnectionapi.ConnectionResponse) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = searchInObject(resourceConnectionAwsRoleTest, "description.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsRoleUpdatedAttrDescription(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
				resource.TestCheckResourceAttr(resourceName, "description", v),
				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
			),
		})
	}
	// valid = searchInObject(resourceConnectionAwsRoleTest, "state.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionAwsRoleUpdatedAttrState(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "state", v),
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
	// 		),
	// 	})
	// }
	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsRegions(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.0", v),
				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
			),
		})
	}
	// valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsPollingInterval(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.interval", v),
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
	// 		),
	// 	})
	// }
	// valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsPollingUnit(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.unit", v),
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
	// 		),
	// 	})
	// }

	// valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsImportTagsEnabled(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.enabled", v),
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
	// 		),
	// 	})
	// }
	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsImportTagsExcludedKeys(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.0", v),
				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsTagFilter(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.tag_filter", v),
				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
			),
		})
	}
	// valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesImportTagsEnabled(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.enabled", v),
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
	// 		),
	// 	})
	// }
	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesImportTagsExcludedKeys(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.0", v),
				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesTagFilter(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.tag_filter", v),
				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
			),
		})
	}
	// valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesPollingInterval(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.interval", v),
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
	// 		),
	// 	})
	// }
	// valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid").([]interface{})
	// for _, value := range valid {
	// 	v := fmt.Sprintf("%v", value)
	// 	testSteps = append(testSteps, resource.TestStep{
	// 		Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesPollingUnit(rName, value),
	// 		Check: resource.ComposeTestCheckFunc(
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
	// 			resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.unit", v),
	// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
	// 		),
	// 	})
	// }

	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAwsRoleConfigWithOptional(rName),
	})

	return testSteps
}

func generateNegativeStepsConnectionAwsRole(rName string, resourceName string) []resource.TestStep {
	//Use Update Config Function with false value
	testSteps := make([]resource.TestStep, 0, 1)
	var invalid []interface{}
	invalid = searchInObject(resourceConnectionAwsRoleTest, "state.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAwsRoleUpdatedAttrState(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}
	invalid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsPollingUnit(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}

	invalid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesPollingUnit(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}

	///
	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsRegions(rName, "invalid_region"),
		ExpectError: regexp.MustCompile("'details.regions' value must be from the list of AWS regions"),
	})

	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsTagFilter(rName, "invalid_tag_filter"),
		ExpectError: regexp.MustCompile("Invalid Input Error"),
	})

	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesTagFilter(rName, "invalid_tag_filter"),
		ExpectError: regexp.MustCompile("Invalid Input Error"),
	})

	invalid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid").([]interface{})
	value := fmt.Sprintf("%v", invalid[1])
	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsImportTagsEnabled(rName, value),
		ExpectError: regexp.MustCompile("Invalid Input Error"),
	})

	invalid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid").([]interface{})
	value = fmt.Sprintf("%v", invalid[1])
	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesImportTagsEnabled(rName, value),
		ExpectError: regexp.MustCompile("Invalid Input Error"),
	})

	invalid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.invalid").([]interface{})
	for _, value := range invalid {
		value = fmt.Sprintf("%v", invalid)
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesWithInvalidName(rName, value),
			ExpectError: regexp.MustCompile("Invalid Input Error"),
		})
	}
	// testSteps = append(testSteps, resource.TestStep{
	// 	Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsWithoutServices(rName),
	// 	ExpectError: regexp.MustCompile("At Least one services is required while updating, services cannot be updated as empty."),
	// })
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAwsRoleConfig(rName),
	})
	///
	return testSteps
}

func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesWithInvalidName(rName string, value interface{}) string {
	var resource string
	parentResources := getParentConnectionAwsAccess(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsAccessConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {

							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"

							account_id = "%v"

							}

							configuration_details {
    
							
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		value,
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsWithoutServices(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsAccess(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsAccessConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_aws" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             access_type = "%v"
		
// 							account_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    polling {
                                                
//                                 interval = %v
                        
//                                 unit = "%v"

//                               }

// 						        import_tags {
                                                
//                                     enabled = "%v"
                        
//                                     excluded_keys = ["%v","%v"]

//                                   }
                        
//                                     tag_filter = "%v"

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"))
// 	return resource
// }

// func CreateAccConnectionAwsRoleMultipleConfig(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsRoleConfig(parentResources)
// 	multipleValues := searchInObject(resourceConnectionAwsRoleTest, "display_name.multiple_valids").([]interface{})
// 	for i, val := range multipleValues {
// 		resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_aws" "test%d" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {

//                             access_type = "%v"

//                             account_id = "%v"

// 							}

// 							configuration_details {

//                             regions = ["%v","%v"]

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
// 		`, i, val,
// 			searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
// 			searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
// 	}
// 	return resource
// }

func testAccCheckAppdynamicscloudConnectionAwsRoleExists(name string, connectionAwsRole *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Connection AWS %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AWS connection id was set")
		}

		config := testAccProvider.Meta().(config)
		myctx, _, apiClient := initializeCloudConnectionClient(config)

		resp, _, err := apiClient.ConnectionsApi.GetConnection(myctx, rs.Primary.ID).Execute()
		if err != nil {
			return err
		}

		if resp.GetId() != rs.Primary.ID {
			return fmt.Errorf("Cloud Connection AWS %s not found", rs.Primary.ID)
		}
		*connectionAwsRole = *resp
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionAwsRoleDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(config)
	myctx, _, apiClient := initializeCloudConnectionClient(config)
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "appdynamicscloud_connection_aws" {
			_, _, err := apiClient.ConnectionsApi.GetConnection(myctx, rs.Primary.ID).Execute()
			if err == nil {
				return fmt.Errorf("Cloud Connection AWS %s Still exists", rs.Primary.ID)
			} else {
				return nil
			}
		}
	}
	return nil
}

func testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole1, connectionAwsRole2 *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionAwsRole1.Id != connectionAwsRole2.Id {
			return fmt.Errorf("ConnectionAwsRole IDs are not equal")
		}
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionAwsRoleIdNotEqual(connectionAwsRole1, connectionAwsRole2 *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionAwsRole1.Id == connectionAwsRole2.Id {
			return fmt.Errorf("ConnectionAwsRole IDs are equal")
		}
		return nil
	}
}

func getParentConnectionAwsRole(rName string) []string {
	t := []string{}
	t = append(t, connectionAwsRoleBlock(rName))
	return t
}

func connectionAwsRoleBlock(rName string) string {
	return fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws" "test" {

						display_name = "%v"


						description = "%v"


						state = "%v"


                        connection_details {
    
                                                
                            access_type = "%v"
                        
                            account_id = "%v"

                        }

                        configuration_details {
    
                                                
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
		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
}

// To eliminate duplicate resource block from slice of resource blocks
func createConnectionAwsRoleConfig(configSlice []string) string {
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

var role_attachment string = fmt.Sprintf(`

  resource "appdynamicscloud_connection_aws_role_attachment" "attachment" {
    connection_id = appdynamicscloud_connection_aws.test.id
    role_name = aws_iam_role.role.name
}

resource "aws_iam_role" "role" {
  name = "test-role-2"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
          "AWS": "arn:aws:iam::${appdynamicscloud_connection_aws.test.connection_details[0].appdynamics_aws_account_id}:root"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
          "StringEquals": {
              "sts:ExternalId": "${appdynamicscloud_connection_aws.test.connection_details[0].external_id}"
          }
      }
    }
  ]
}
EOF
}

resource "aws_iam_policy" "policy" {
  name        = "test-policy-2"
  description = "A test policy 2"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "elasticloadbalancing:DescribeLoadBalancers",
        "ec2:DescribeInstances",
        "cloudwatch:GetMetricData",
        "ec2:DescribeVpcs",
        "ec2:DescribeRegions",
        "ec2:DescribeVolumes",
        "elasticloadbalancing:DescribeTargetHealth",
        "rds:DescribeDBInstances",
        "elasticloadbalancing:DescribeTargetGroups",
        "ec2:DescribeSubnets",
        "cloudwatch:ListMetrics",
        "rds:DescribeDBClusters",
        "tag:GetResources",
        "ecs:ListClusters",
        "ecs:DescribeClusters",
        "ecs:ListServices",
        "ecs:DescribeServices",
        "ecs:ListTasks",
        "ecs:DescribeTasks",
        "ecs:DescribeContainerInstances",
        "ecs:ListTaskDefinitions"
      ],
      "Resource": "*",
      "Effect": "Allow",
      "Sid": "AllowMonitoring"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = aws_iam_role.role.name
  policy_arn = aws_iam_policy.policy.arn
}
`)
