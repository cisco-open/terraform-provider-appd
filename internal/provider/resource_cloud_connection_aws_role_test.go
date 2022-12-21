package provider

// const connectionAwsRoleSelfRequiredCount = 2

// var resourceConnectionAwsRoleTest = map[string]interface{}{
// 	"display_name": map[string]interface{}{
// 		"valid":           []interface{}{"mf29ikrvuq", "9gihbn0pxs", "4a98kcdsoe", "80e1pxmdgt"},
// 		"invalid":         []interface{}{10, 12.43},
// 		"multiple_valids": []interface{}{"dnx8qbdcwc", "ull5dyttec", "nj0zb0dq6h", "6fdwm6m1yn", "goghukna96", "c5vgax4len", "2m4wq4xxr1", "8rq857h4op", "zloufg82ka", "3h1qvjg9q4", "dbb9baoy7m", "fi28rvmzrf", "23vnv8qokp", "nq6tqe9pms", "f7esp3bvn8"},
// 	},

// 	"description": map[string]interface{}{
// 		"valid":           []interface{}{"a4azty49iw", "2e10vrjady", "h5zyu013gw", "9t2z4sgyfd"},
// 		"invalid":         []interface{}{10, 12.43},
// 		"multiple_valids": []interface{}{"tjl9ithmen", "w7mc5wonx6", "6ejzhbd0bs", "2018j431yq", "9apw3rbx7m", "62bmkn0oxq", "o93crtl3m2", "unht0g3tqz", "50aynd1ltp", "d43uivb409", "aqujh09yo5", "b1anqkem7c", "h618adqt7b", "y7leai85xi", "a8an6sa0hi"},
// 	},

// 	"state": map[string]interface{}{
// 		"valid":           []interface{}{"ACTIVE", "INACTIVE"},
// 		"invalid":         []interface{}{"gv2xnyi3xn"},
// 		"multiple_valids": []interface{}{"ACTIVE", "INACTIVE"},
// 	},

// 	"connection_details": map[string]interface{}{
// 		"access_type": map[string]interface{}{
// 			"valid":           []interface{}{"role_delegation"},
// 			"invalid":         []interface{}{"rbnigdc04i"},
// 			"multiple_valids": []interface{}{"role_delegation"},
// 		},

// 		"account_id": map[string]interface{}{
// 			"valid":           []interface{}{"5o4my7oeu3", "1ae8g4wuhu", "zdkmjc328t", "2f1mo3hefr"},
// 			"invalid":         []interface{}{10, 12.43},
// 			"multiple_valids": []interface{}{"ekcpkqpi39", "15d0063eod", "gblz20y4a0", "hzx0snuy0z", "saohaxrk9o", "8a7sfzpa90", "x4bzk91fsb", "o8fcs0clww", "btbmye02a3", "ngnzexs7ql", "cmexbxct8r", "snkgxuxqx8", "ht5dekzu5i", "7dnjualg01", "u78tlrz8ll"},
// 		},
// 	},

// 	"configuration_details": map[string]interface{}{
// 		"regions": map[string]interface{}{
// 			"valid":           []interface{}{"liqfcpezaq", "b3ia0kcsyr", "r38l533rzt", "8a93v3ze7x"},
// 			"invalid":         []interface{}{10, 12.43},
// 			"multiple_valids": []interface{}{"9wg9t8uq0o", "gi5vj0m3cz", "3pe42f8c8g", "mqfkk21ode", "9lmswxuwct", "mcv1yhd0t2", "4g3e3bge6p", "3r936iows5", "7mpjx5e21d", "v6qqub585v", "lmfjx7lvc7", "uui9bs2ayg", "rvf6pr3eok", "w8qfy7roct", "pnzgtpu99j"},
// 		},

// 		"polling": map[string]interface{}{
// 			"interval": map[string]interface{}{
// 				"valid":           []interface{}{590, 721, 655, 751},
// 				"invalid":         []interface{}{"random", 10.023},
// 				"multiple_valids": []interface{}{-658, -499, -205, 112, 100, -987, -209, 130, 61, -70, -869, 232, -355, 936, -357},
// 			},

// 			"unit": map[string]interface{}{
// 				"valid":           []interface{}{"minute"},
// 				"invalid":         []interface{}{"fqwty4vjff"},
// 				"multiple_valids": []interface{}{"minute"},
// 			},
// 		},

// 		"import_tags": map[string]interface{}{
// 			"enabled": map[string]interface{}{
// 				"valid":           []interface{}{true, false},
// 				"invalid":         []interface{}{"random", 10},
// 				"multiple_valids": []interface{}{true, false},
// 			},

// 			"excluded_keys": map[string]interface{}{
// 				"valid":           []interface{}{"ncorsk8p01", "7daq85gffj", "59rn08r773", "6453ve0vu1"},
// 				"invalid":         []interface{}{10, 12.43},
// 				"multiple_valids": []interface{}{"r4eon0dtud", "uyn62a1tir", "zkw8u5yvl6", "zzbkcg10c6", "gu0p6j8o6q", "mldlkbyyl6", "5q95podnkv", "luct6mg7b3", "mfo2inb1jj", "j09bf9rtim", "9v5ojvo0n8", "e4oocutwk5", "so6plyamsg", "uavoj658m3", "82wq4e0go3"},
// 			},
// 		},

// 		"tag_filter": map[string]interface{}{
// 			"valid":           []interface{}{"aqyl7w0his", "n15yb5r3t1", "azxuvi2epj", "luk4e2v5c2"},
// 			"invalid":         []interface{}{10, 12.43},
// 			"multiple_valids": []interface{}{"579xckns8c", "llg9jv9oo9", "mli0kisgtu", "5fet6i655y", "w7gsufiiyh", "2ndurqr8in", "yrrlu62l4v", "wq931thsyw", "odu369k7tl", "nxd92w20tn", "b5vrodg12c", "oc1ljwwo9e", "lv499yjr9i", "hl7uwdkoja", "c4xixhord9"},
// 		},

// 		"services": map[string]interface{}{
// 			"name": map[string]interface{}{
// 				"valid":           []interface{}{"2uj20k35rw", "wi0426q1w0", "lhisrm7sqr", "k4r4c1xi9c"},
// 				"invalid":         []interface{}{10, 12.43},
// 				"multiple_valids": []interface{}{"g30am4bkc9", "hi6vv9iuu4", "tqdmtkdljb", "fdawy9qa65", "o4uenl79c8", "f4dsvwug86", "6c8q9gcswi", "of7ljkn8yf", "mcgbm2j2y4", "yh7pw6hq7o", "zpo5pfe7xc", "rpkh6xcwed", "c3cs1swbew", "vgrjrk3ntx", "v0a803kp0q"},
// 			},

// 			"import_tags": map[string]interface{}{
// 				"enabled": map[string]interface{}{
// 					"valid":           []interface{}{true, false},
// 					"invalid":         []interface{}{"random", 10},
// 					"multiple_valids": []interface{}{true, false},
// 				},

// 				"excluded_keys": map[string]interface{}{
// 					"valid":           []interface{}{"p9dqcp0hpe", "b78valfi6i", "fg82v78o15", "tnjtvwd9hh"},
// 					"invalid":         []interface{}{10, 12.43},
// 					"multiple_valids": []interface{}{"68qm4fj7sw", "lurbuydcoc", "nnno0crxrh", "eyz4nrgvrt", "ty5uo7ytaq", "fvqul7edjm", "98vkv4abl2", "iz1at0gda8", "car0ak3tai", "ly2x6zngvg", "e7d6bmmwd2", "ajvmyeuemc", "6qo40w0ptx", "2xj5zhj1qw", "pil2hvhwdg"},
// 				},
// 			},

// 			"tag_filter": map[string]interface{}{
// 				"valid":           []interface{}{"dz0zeld1nf", "vakx9lht89", "89q0bg440g", "6a8pt31z7f"},
// 				"invalid":         []interface{}{10, 12.43},
// 				"multiple_valids": []interface{}{"xgs8ksrqcn", "7ioeng8ww3", "0rexy8l8ps", "p5jxbonxv8", "0ywrnmhnti", "34yuh97aj2", "6gbfqh6uko", "pnj7kczzuc", "ow8bdjxt4s", "vk34mx9sp3", "6f89rqcwn1", "nseg8b4v19", "nlvgjbulus", "eqf6w3rva3", "u47kpp082m"},
// 			},

// 			"polling": map[string]interface{}{
// 				"interval": map[string]interface{}{
// 					"valid":           []interface{}{501, -144, 164, 741},
// 					"invalid":         []interface{}{"random", 10.023},
// 					"multiple_valids": []interface{}{-2, -484, 253, -618, 317, -488, -1000, -93, 626, -28, 958, -342, -583, 444, 203},
// 				},

// 				"unit": map[string]interface{}{
// 					"valid":           []interface{}{"minute"},
// 					"invalid":         []interface{}{"yx227qv1la"},
// 					"multiple_valids": []interface{}{"minute"},
// 				},
// 			},
// 		},
// 	},
// }

// func TestAccAppdynamicscloudConnectionAwsRole_Basic(t *testing.T) {
// 	var connectionAwsRole_default models.ConnectionAwsRole
// 	var connectionAwsRole_updated models.ConnectionAwsRole
// 	resourceName := "appdynamicscloud_connection_aws.test"

// 	rName := makeTestVariable(acctest.RandString(5))
// 	// rOther := makeTestVariable(acctest.RandString(5))

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsRoleDestroy,
// 		Steps: append([]resource.TestStep{
// 			{
// 				Config:      CreateAccConnectionAwsRoleWithoutDisplayName(rName),
// 				ExpectError: regexp.MustCompile(`Missing required argument`),
// 			},
// 			{
// 				Config:      CreateAccConnectionAwsRoleWithoutConnectionDetails(rName),
// 				ExpectError: regexp.MustCompile(`Missing required argument`),
// 			},
// 			{
// 				Config: CreateAccConnectionAwsRoleConfig(rName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, &connectionAwsRole_default),
// 					resource.TestCheckResourceAttr(resourceName, "display_name", rName),

// 					resource.TestCheckResourceAttr(resourceName, "description", ""),

// 					resource.TestCheckResourceAttr(resourceName, "state", ""),

// 					resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_type", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.0.account_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"))),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "0"),
// 				),
// 			},
// 			{
// 				Config: CreateAccConnectionAwsRoleConfigWithOptional(rName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, &connectionAwsRole_updated),
// 					resource.TestCheckResourceAttr(resourceName, "display_name", rName),
// 					resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "state", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_type", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.0.account_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"))),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.#", "2"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid.1"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"))),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.#", "2"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"))),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.name", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.#", "2"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"))),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))),

// 					testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(&connectionAwsRole_default, &connectionAwsRole_updated),
// 				),
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 			{
// 				Config: CreateAccConnectionAwsRoleConfig(rName),
// 			},
// 		}, generateStepForUpdatedRequiredAttrConnectionAwsRole(rName, resourceName, &connectionAwsRole_default, &connectionAwsRole_updated)...),
// 	})
// }

// func TestAccAppdynamicscloudConnectionAwsRole_Update(t *testing.T) {
// 	var connectionAwsRole_default models.ConnectionAwsRole
// 	var connectionAwsRole_updated models.ConnectionAwsRole
// 	resourceName := "appdynamicscloud_connection_aws.test"
// 	rName := makeTestVariable(acctest.RandString(5))

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsRoleDestroy,
// 		Steps: append([]resource.TestStep{
// 			{
// 				Config: CreateAccConnectionAwsRoleConfig(rName),
// 				Check:  testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, &connectionAwsRole_default),
// 			},
// 		}, generateStepForUpdatedAttrConnectionAwsRole(rName, resourceName, &connectionAwsRole_default, &connectionAwsRole_updated)...),
// 	})
// }

// func TestAccAppdynamicscloudConnectionAwsRole_NegativeCases(t *testing.T) {
// 	resourceName := "appdynamicscloud_connection_aws.test"

// 	// [TODO]: Add makeTestVariable() to utils.go file
// 	rName := makeTestVariable(acctest.RandString(5))

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAwsRoleDestroy,
// 		Steps: append([]resource.TestStep{
// 			{
// 				Config: CreateAccConnectionAwsRoleConfig(rName),
// 			},
// 		}, generateNegativeStepsConnectionAwsRole(rName, resourceName)...),
// 	})
// }

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

// func CreateAccConnectionAwsRoleWithoutDisplayName(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsRoleConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 				resource  "appdynamicscloud_connection_aws" "test" {

// 									description = "%v"

// 									state = "%v"

// 									connection_details {
    
									                        
//                                         access_type = "%v"
                        
//                                         account_id = "%v"

// 									}

// 									configuration_details {
    
									                        
//                                         regions = ["%v","%v"]

//                                         polling {
                                                    
//                                             interval = %v
                        
//                                             unit = "%v"

//                                           }

//                                             import_tags {
                                                    
//                                                 enabled = "%v"
                        
//                                                 excluded_keys = ["%v","%v"]

//                                               }
                        
//                                                 tag_filter = "%v"

//                                                 services {
                                                    
//                                                     name = "%v"

//                                                     import_tags {
                                                    
//                                                         enabled = "%v"
                        
//                                                         excluded_keys = ["%v","%v"]

//                                                       }
                        
//                                                         tag_filter = "%v"

//                                                         polling {
                                                    
//                                                             interval = %v
                        
//                                                             unit = "%v"

//                                                           }

//                                                           }

// 									}
// 				}
// 			`, searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
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
// func CreateAccConnectionAwsRoleWithoutConnectionDetails(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsRoleConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 				resource  "appdynamicscloud_connection_aws" "test" {

// 									display_name = "%v"

// 									description = "%v"

// 									state = "%v"

// 									configuration_details {
    
									                        
//                                         regions = ["%v","%v"]

//                                         polling {
                                                    
//                                             interval = %v
                        
//                                             unit = "%v"

//                                           }

//                                             import_tags {
                                                    
//                                                 enabled = "%v"
                        
//                                                 excluded_keys = ["%v","%v"]

//                                               }
                        
//                                                 tag_filter = "%v"

//                                                 services {
                                                    
//                                                     name = "%v"

//                                                     import_tags {
                                                    
//                                                         enabled = "%v"
                        
//                                                         excluded_keys = ["%v","%v"]

//                                                       }
                        
//                                                         tag_filter = "%v"

//                                                         polling {
                                                    
//                                                             interval = %v
                        
//                                                             unit = "%v"

//                                                           }

//                                                           }

// 									}
// 				}
// 			`, searchInObject(resourceConnectionAwsRoleTest, "display_name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
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

// func CreateAccConnectionAwsRoleConfig(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsRoleConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 		resource  "appdynamicscloud_connection_aws" "test" {


// 							display_name = "%v"

// 							connection_details {
    
							 

// 						          access_type = "%v"
 

// 						          account_id = "%v"

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"))
// 	return resource
// }

// func CreateAccConnectionAwsRoleConfigWithOptional(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]

// 	resource += createConnectionAwsRoleConfig(parentResources)

// 	resource += fmt.Sprintf(`
// 		resource  "appdynamicscloud_connection_aws" "test" {

// 						display_name = "%v"

// 						description = "%v"

// 						state = "%v"

//                         connection_details {
    
                                                
//                             access_type = "%v"
                        
//                             account_id = "%v"

//                         }

//                         configuration_details {
    
                                                
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

//                         }
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

// func generateStepForUpdatedRequiredAttrConnectionAwsRole(rName string, resourceName string, connectionAwsRole_default, connectionAwsRole_updated *models.ConnectionAwsRole) []resource.TestStep {
// 	testSteps := make([]resource.TestStep, 0, 1)
// 	var value interface{}
// 	value = searchInObject(resourceConnectionAwsRoleTest, "display_name.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAwsRoleUpdateRequiredDisplayName(rName),
// 		Check: resource.ComposeTestCheckFunc(
// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 			resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", value)),
// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdNotEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 		),
// 	})
// 	value = searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAwsRoleUpdateRequiredConnectionDetailsAccessType(rName),
// 		Check: resource.ComposeTestCheckFunc(
// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 			resource.TestCheckResourceAttr(resourceName, "connection_details.0.access_type", fmt.Sprintf("%v", value)),
// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdNotEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 		),
// 	})
// 	value = searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAwsRoleUpdateRequiredConnectionDetailsAccountId(rName),
// 		Check: resource.ComposeTestCheckFunc(
// 			testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 			resource.TestCheckResourceAttr(resourceName, "connection_details.0.account_id", fmt.Sprintf("%v", value)),
// 			testAccCheckAppdynamicscloudConnectionAwsRoleIdNotEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 		),
// 	})

// 	return testSteps
// }
// func CreateAccConnectionAwsRoleUpdateRequiredDisplayName(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAwsRole(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAwsRoleConfig(parentResources)
// 	value := searchInObject(resourceConnectionAwsRoleTest, "display_name.valid.1")
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
// 		`, value,
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

// func CreateAccConnectionAwsRoleUpdatedAttrDescription(rName string, value interface{}) string {
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
// 		}
// 	`, rName,
// 		value,
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
// func CreateAccConnectionAwsRoleUpdatedAttrState(rName string, value interface{}) string {
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
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
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
// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsRegions(rName string, value interface{}) string {
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
    
							
// 						    regions = ["%v"]
  

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
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		value,
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

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "polling.unit.valid.0"),
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
// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsPollingUnit(rName string, value interface{}) string {
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
					
// 						        unit = "%v"

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

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "polling.interval.valid.0"),
// 		value,
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
// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsImportTagsEnabled(rName string, value interface{}) string {
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

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "excluded_keys.valid.1"),
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
// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsImportTagsExcludedKeys(rName string, value interface{}) string {
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

// 						            excluded_keys = ["%v"]
  

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

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "import_tags.enabled.valid.0"),
// 		value,
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
// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsTagFilter(rName string, value interface{}) string {
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
					
// 						            tag_filter = "%v"

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

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesImportTagsEnabled(rName string, value interface{}) string {
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
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.name.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesImportTagsExcludedKeys(rName string, value interface{}) string {
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

// 						                    excluded_keys = ["%v"]
  

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
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "import_tags.enabled.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesTagFilter(rName string, value interface{}) string {
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
					
// 						                    tag_filter = "%v"

// 						                    polling {
                                                
//                                                 interval = %v
                        
//                                                 unit = "%v"

//                                               }

// 						                      }

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.excluded_keys.valid.1"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.polling.unit.valid.0"))
// 	return resource
// }
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

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.tag_filter.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAwsRoleTest, "polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesPollingUnit(rName string, value interface{}) string {
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
						                        
//                                                 interval = %v
					
// 						                        unit = "%v"

// 						                      }

// 						                      }

// 							}
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAwsRoleTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.access_type.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "connection_details.account_id.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "polling.interval.valid.0"),
// 		value)
// 	return resource
// }

// func generateStepForUpdatedAttrConnectionAwsRole(rName string, resourceName string, connectionAwsRole_default, connectionAwsRole_updated *models.ConnectionAwsRole) []resource.TestStep {
// 	testSteps := make([]resource.TestStep, 0, 1)
// 	var valid []interface{}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "description.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrDescription(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "description", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "state.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrState(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "state", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.regions.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsRegions(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.0", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.interval.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsPollingInterval(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.interval", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsPollingUnit(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.unit", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}

// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.enabled.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsImportTagsEnabled(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.enabled", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.import_tags.excluded_keys.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsImportTagsExcludedKeys(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.0", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}

// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsTagFilter(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.tag_filter", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesImportTagsEnabled(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.enabled", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesImportTagsExcludedKeys(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.0", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}

// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesTagFilter(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.tag_filter", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesPollingInterval(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.interval", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesPollingUnit(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAwsRoleExists(resourceName, connectionAwsRole_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.unit", v),
// 				testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole_default, connectionAwsRole_updated),
// 			),
// 		})
// 	}

// 	return testSteps
// }

// func generateNegativeStepsConnectionAwsRole(rName string, resourceName string) []resource.TestStep {
// 	//Use Update Config Function with false value
// 	testSteps := make([]resource.TestStep, 0, 1)
// 	var invalid []interface{}
// 	invalid = searchInObject(resourceConnectionAwsRoleTest, "state.invalid").([]interface{})
// 	for _, value := range invalid {
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config:      CreateAccConnectionAwsRoleUpdatedAttrState(rName, value),
// 			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
// 		})
// 	}
// 	invalid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.polling.unit.invalid").([]interface{})
// 	for _, value := range invalid {
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsPollingUnit(rName, value),
// 			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
// 		})
// 	}

// 	invalid = searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.invalid").([]interface{})
// 	for _, value := range invalid {
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config:      CreateAccConnectionAwsRoleUpdatedAttrConfigurationDetailsServicesPollingUnit(rName, value),
// 			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
// 		})
// 	}

// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAwsRoleConfig(rName),
// 	})
// 	return testSteps
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

// func testAccCheckAppdynamicscloudConnectionAwsRoleExists(name string, connectionAwsRole *models.ConnectionAwsRole) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		// [TODO]: Write your code here
// 	}
// }

// func testAccCheckAppdynamicscloudConnectionAwsRoleDestroy(s *terraform.State) error {
// 	client := testAccProvider.Meta().(*client.Client)

// 	for _, rs := range s.RootModule().Resources {

// 		if rs.Type == "appdynamicscloud_connection_aws" {
// 			// [TODO]: Write your code here
// 		}
// 	}
// 	return nil
// }

// func testAccCheckAppdynamicscloudConnectionAwsRoleIdEqual(connectionAwsRole1, connectionAwsRole2 *models.ConnectionAwsRole) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		Id1, err := getIdFromConnectionAwsRoleModel(connectionAwsRole1)
// 		if err != nil {
// 			return err
// 		}
// 		Id2, err := getIdFromConnectionAwsRoleModel(connectionAwsRole2)
// 		if err != nil {
// 			return err
// 		}
// 		if Id1 != Id2 {
// 			return fmt.Errorf("ConnectionAwsRole IDs are not equal")
// 		}
// 		return nil
// 	}
// }

// func testAccCheckAppdynamicscloudConnectionAwsRoleIdNotEqual(connectionAwsRole1, connectionAwsRole2 *models.ConnectionAwsRole) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		Id1, err := getIdFromConnectionAwsRoleModel(connectionAwsRole1)
// 		if err != nil {
// 			return err
// 		}
// 		Id2, err := getIdFromConnectionAwsRoleModel(connectionAwsRole2)
// 		if err != nil {
// 			return err
// 		}
// 		if Id1 == Id2 {
// 			return fmt.Errorf("ConnectionAwsRole IDs are equal")
// 		}
// 		return nil
// 	}
// }

// func getParentConnectionAwsRole(rName string) []string {
// 	t := []string{}
// 	t = append(t, connectionAwsRoleBlock(rName))
// 	return t
// }

// func connectionAwsRoleBlock(rName string) string {
// 	return fmt.Sprintf(`
// 		resource  "appdynamicscloud_connection_aws" "test" {

// 						display_name = "%v"


// 						description = "%v"


// 						state = "%v"


//                         connection_details {
    
                                                
//                             access_type = "%v"
                        
//                             account_id = "%v"

//                         }

//                         configuration_details {
    
                                                
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

//                         }
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
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAwsRoleTest, "configuration_details.services.polling.unit.valid.0"))
// }

// // To eliminate duplicate resource block from slice of resource blocks
// func createConnectionAwsRoleConfig(configSlice []string) string {
// 	keys := make(map[string]bool)
// 	str := ""

// 	for _, entry := range configSlice {
// 		if _, value := keys[entry]; !value {
// 			keys[entry] = true
// 			str += entry
// 		}
// 	}

// 	return str
// }
