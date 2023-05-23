// Copyright 2023 Cisco Systems, Inc.
//
// Licensed under the MPL License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.mozilla.org/en-US/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var resourceConnectionAwsAccessTest = map[string]interface{}{
	"display_name": map[string]interface{}{
		"valid":           []interface{}{"9ttp44xwtj", "23ri0trg54", "g7mkxrwkp6", "l9bcj880vi"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"x9ejdsemr6", "olje9mungr", "3zuh6b148v", "ltpblw71es", "sg2hr2ry74", "xc6mlp4ai9", "ky6cu8c8fw", "hfrtvu5fru", "ye3bfgj0j4", "qz6dem8wy8", "c7a3ony8cu", "4lkd1rc93a", "688hfs0wrc", "t6qce2d0kj", "s4upbvx9gu"},
	},

	"description": map[string]interface{}{
		"valid":           []interface{}{"TestAcc_ynnrw0nsyg", "TestAcc_j64nfvmm51", "TestAcc_ebg8hlszur", "TestAcc_u53fbt13o8"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"TestAcc_69ya5udp7l", "TestAcc_7xzzj3km8z", "TestAcc_fvjiruf0ae", "TestAcc_bp5hnkb2ow", "TestAcc_g7130ksf12", "TestAcc_sn5dmxhicz", "TestAcc_vrv4ynzt6e", "TestAcc_u0s8dzpwz9", "TestAcc_cmw95cf3kv", "TestAcc_98fqosmbds", "TestAcc_e3hkdpkams", "TestAcc_nhjsujp85r", "TestAcc_4rke4hgxru", "TestAcc_plns3k8p5u", "TestAcc_na39rgngyn"},
	},

	"state": map[string]interface{}{
		"valid":           []interface{}{"ACTIVE", "INACTIVE"},
		"invalid":         []interface{}{"t9rbvkt7hy"},
		"multiple_valids": []interface{}{"ACTIVE", "INACTIVE"},
	},

	"connection_details": map[string]interface{}{
		"access_type": map[string]interface{}{
			"valid":           []interface{}{"access_key"},
			"invalid":         []interface{}{"k3ta9kzd5c"},
			"multiple_valids": []interface{}{"access_key"},
		},

		"access_key_id": map[string]interface{}{
			"valid":           []interface{}{os.Getenv("TEST_AWS_ACCESS_KEY_ID")},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{os.Getenv("TEST_AWS_ACCESS_KEY_ID")},
		},

		"secret_access_key": map[string]interface{}{
			"valid":           []interface{}{os.Getenv("TEST_AWS_SECRET_ACCESS_KEY")},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{os.Getenv("TEST_AWS_SECRET_ACCESS_KEY")},
		},
	},

	"configuration_details": map[string]interface{}{
		"regions": map[string]interface{}{
			"valid":           []interface{}{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "eu-west-1"},
			"invalid":         []interface{}{10, 12.43},
			"multiple_valids": []interface{}{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "eu-west-1"},
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
				"invalid":         []interface{}{10, "", "random_service"},
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

func TestAccAppdynamicscloudConnectionAwsAccess_Basic(t *testing.T) {
	var connectionAwsAccess_default cloudconnectionapi.ConnectionResponse
	var connectionAwsAccess_updated cloudconnectionapi.ConnectionResponse
	resourceName := "appdynamicscloud_connection_aws.test"

	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsAccessDestroy,
		Steps: append([]resource.TestStep{
			{
				Config:      CreateAccConnectionAwsAccessWithoutDisplayName(rName),
				ExpectError: regexp.MustCompile(`Missing required argument`),
			},
			{
				Config:      CreateAccConnectionAwsAccessWithoutConnectionDetails(rName),
				ExpectError: regexp.MustCompile(`Insufficient connection_details blocks`),
			},
			{
				Config: CreateAccConnectionAwsAccessConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, &connectionAwsAccess_default),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),

					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_type", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_key_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.secret_access_key", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				),
			},
			{
				Config: CreateAccConnectionAwsAccessConfigWithOptional(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, &connectionAwsAccess_updated),
					resource.TestCheckResourceAttr(resourceName, "display_name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_type", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_key_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "connection_details.0.secret_access_key", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.name", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"))),

					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))),

					testAccCheckAppdynamicscloudConnectionAwsAccessIdEqual(&connectionAwsAccess_default, &connectionAwsAccess_updated),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_details", "configuration_details_service_default"},
			},
			{
				Config: CreateAccConnectionAwsAccessConfigWithOptional(rName),
			},
		}, generateStepForUpdatedRequiredAttrConnectionAwsAccess(rName, resourceName, &connectionAwsAccess_default, &connectionAwsAccess_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionAwsAccess_Update(t *testing.T) {
	var connectionAwsAccess_default cloudconnectionapi.ConnectionResponse
	var connectionAwsAccess_updated cloudconnectionapi.ConnectionResponse
	resourceName := "appdynamicscloud_connection_aws.test"
	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsAccessDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionAwsAccessConfig(rName),
				Check:  testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, &connectionAwsAccess_default),
			},
		}, generateStepForUpdatedAttrConnectionAwsAccess(rName, resourceName, &connectionAwsAccess_default, &connectionAwsAccess_updated)...),
	})
}

func TestAccAppdynamicscloudConnectionAwsAccess_NegativeCases(t *testing.T) {
	resourceName := "appdynamicscloud_connection_aws.test"
	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsAccessDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccConnectionAwsAccessConfigWithOptional(rName),
			},
		}, generateNegativeStepsConnectionAwsAccess(rName, resourceName)...),
	})
}

func CreateAccConnectionAwsAccessWithoutDisplayName(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsAccess(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsAccessConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "appdynamicscloud_connection_aws" "test" {

									description = "%v"

									state = "%v"

									connection_details {
    
									                        
                                        access_type = "%v"
                        
                                        access_key_id = "%v"
                        
                                        secret_access_key = "%v"

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
			`, searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsAccessWithoutConnectionDetails(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsAccess(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsAccessConfig(parentResources)
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
			`, searchInObject(resourceConnectionAwsAccessTest, "display_name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsAccessConfig(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsAccess(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsAccessConfig(parentResources)
	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws" "test" {


							display_name = "%v"

							connection_details {
    
							 

						          access_type = "%v"
 

						          access_key_id = "%v"
 

						          secret_access_key = "%v"

							}
		}
	`, rName,
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"))
	return resource
}

func CreateAccConnectionAwsAccessConfigWithOptional(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsAccess(rName)
	parentResources = parentResources[:len(parentResources)-1]

	resource += createConnectionAwsAccessConfig(parentResources)

	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws" "test" {

						display_name = "%v"

						description = "%v"

						state = "%v"

                        connection_details {
    
                                                
                            access_type = "%v"
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func generateStepForUpdatedRequiredAttrConnectionAwsAccess(rName string, resourceName string, connectionAwsAccess_default, connectionAwsAccess_updated *cloudconnectionapi.ConnectionResponse) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	value := searchInObject(resourceConnectionAwsAccessTest, "display_name.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAwsAccessUpdateRequiredDisplayName(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, connectionAwsAccess_updated),
			resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudConnectionAwsAccessIdEqual(connectionAwsAccess_default, connectionAwsAccess_updated),
		),
	})
	return testSteps
}
func CreateAccConnectionAwsAccessUpdateRequiredDisplayName(rName string) string {
	var resource string
	parentResources := getParentConnectionAwsAccess(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createConnectionAwsAccessConfig(parentResources)
	value := searchInObject(resourceConnectionAwsAccessTest, "display_name.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_connection_aws" "test" {
							
							display_name = "%v"

							description = "%v"

							state = "%v"

							connection_details {
    
							                        
                            access_type = "%v"
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsAccessUpdatedAttrDescription(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsAccessUpdatedAttrState(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsRegions(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsPollingUnit(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsImportTagsEnabled(rName string, value interface{}) string {
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

                            access_key_id = "%v"

                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsImportTagsExcludedKeys(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsTagFilter(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesImportTagsEnabled(rName string, value interface{}) string {
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

                            access_key_id = "%v"

                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesImportTagsExcludedKeys(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}
func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesTagFilter(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesPollingUnit(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		value)
	return resource
}

func generateStepForUpdatedAttrConnectionAwsAccess(rName string, resourceName string, connectionAwsAccess_default, connectionAwsAccess_updated *cloudconnectionapi.ConnectionResponse) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = searchInObject(resourceConnectionAwsAccessTest, "description.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsAccessUpdatedAttrDescription(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, connectionAwsAccess_updated),
				resource.TestCheckResourceAttr(resourceName, "description", v),
				testAccCheckAppdynamicscloudConnectionAwsAccessIdEqual(connectionAwsAccess_default, connectionAwsAccess_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsRegions(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, connectionAwsAccess_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.0", v),
				testAccCheckAppdynamicscloudConnectionAwsAccessIdEqual(connectionAwsAccess_default, connectionAwsAccess_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsImportTagsExcludedKeys(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, connectionAwsAccess_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.0", v),
				testAccCheckAppdynamicscloudConnectionAwsAccessIdEqual(connectionAwsAccess_default, connectionAwsAccess_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsTagFilter(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, connectionAwsAccess_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.tag_filter", v),
				testAccCheckAppdynamicscloudConnectionAwsAccessIdEqual(connectionAwsAccess_default, connectionAwsAccess_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesImportTagsExcludedKeys(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, connectionAwsAccess_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.0", v),
				testAccCheckAppdynamicscloudConnectionAwsAccessIdEqual(connectionAwsAccess_default, connectionAwsAccess_updated),
			),
		})
	}

	valid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesTagFilter(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudConnectionAwsAccessExists(resourceName, connectionAwsAccess_updated),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.tag_filter", v),
				testAccCheckAppdynamicscloudConnectionAwsAccessIdEqual(connectionAwsAccess_default, connectionAwsAccess_updated),
			),
		})
	}

	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAwsAccessConfigWithOptional(rName),
	})

	return testSteps
}

func generateNegativeStepsConnectionAwsAccess(rName string, resourceName string) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var invalid []interface{}
	invalid = searchInObject(resourceConnectionAwsAccessTest, "state.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAwsAccessUpdatedAttrState(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}
	invalid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsPollingUnit(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}

	invalid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesPollingUnit(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}

	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsRegions(rName, "invalid_region"),
		ExpectError: regexp.MustCompile("'details.regions' value must be from the list of AWS regions"),
	})

	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsTagFilter(rName, "invalid_tag_filter"),
		ExpectError: regexp.MustCompile("Invalid Input Error"),
	})

	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesTagFilter(rName, "invalid_tag_filter"),
		ExpectError: regexp.MustCompile("Invalid Input Error"),
	})

	invalid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid").([]interface{})
	value := fmt.Sprintf("%v", invalid[1])
	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsImportTagsEnabled(rName, value),
		ExpectError: regexp.MustCompile("Invalid Input Error"),
	})

	invalid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid").([]interface{})
	value = fmt.Sprintf("%v", invalid[1])
	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesImportTagsEnabled(rName, value),
		ExpectError: regexp.MustCompile("Invalid Input Error"),
	})

	invalid = searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.invalid").([]interface{})
	for _, value := range invalid {
		value = fmt.Sprintf("%v", invalid)
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesWithInvalidName(rName, value),
			ExpectError: regexp.MustCompile("Invalid Input Error"),
		})
	}
	testSteps = append(testSteps, resource.TestStep{
		Config:      CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsWithoutServices(rName),
		ExpectError: regexp.MustCompile(serviceEmptyErrorMsg),
	})
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccConnectionAwsAccessConfigWithOptional(rName),
	})
	return testSteps
}

func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsServicesWithInvalidName(rName string, value interface{}) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		value,
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
	return resource
}

func CreateAccConnectionAwsAccessUpdatedAttrConfigurationDetailsWithoutServices(rName string) string {
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
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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

							}
		}
	`, rName,
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"))
	return resource
}

func testAccCheckAppdynamicscloudConnectionAwsAccessExists(name string, connectionAWS *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
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
		*connectionAWS = *resp
		return nil
	}
}

func testAccCheckAppdynamicscloudConnectionAwsAccessDestroy(s *terraform.State) error {
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

func testAccCheckAppdynamicscloudConnectionAwsAccessIdEqual(connectionAwsAccess1, connectionAwsAccess2 *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionAwsAccess1.Id != connectionAwsAccess2.Id {
			return fmt.Errorf("ConnectionAwsAccess IDs are not equal")
		}
		return nil
	}
}

//lint:ignore U1000 might come in handy in the future
func testAccCheckAppdynamicscloudConnectionAwsAccessIdNotEqual(connectionAwsAccess1, connectionAwsAccess2 *cloudconnectionapi.ConnectionResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if connectionAwsAccess1.Id == connectionAwsAccess2.Id {
			return fmt.Errorf("ConnectionAwsAccess IDs are equal")
		}
		return nil
	}
}

func getParentConnectionAwsAccess(rName string) []string {
	t := []string{}
	t = append(t, connectionAwsAccessBlock(rName))
	return t
}

func connectionAwsAccessBlock(rName string) string {
	return fmt.Sprintf(`
		resource  "appdynamicscloud_connection_aws" "test" {

						display_name = "%v"


						description = "%v"


						state = "%v"


                        connection_details {
    
                                                
                            access_type = "%v"
                        
                            access_key_id = "%v"
                        
                            secret_access_key = "%v"

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
		searchInObject(resourceConnectionAwsAccessTest, "description.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "state.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_type.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.access_key_id.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "connection_details.secret_access_key.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.regions.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.polling.unit.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.name.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.enabled.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.tag_filter.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.interval.valid.0"),
		searchInObject(resourceConnectionAwsAccessTest, "configuration_details.services.polling.unit.valid.0"))
}

// To eliminate duplicate resource block from slice of resource blocks
func createConnectionAwsAccessConfig(configSlice []string) string {
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
