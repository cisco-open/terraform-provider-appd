package provider

// const connectionAzureSelfRequiredCount = 1

// var resourceConnectionAzureTest = map[string]interface{}{
// 	"display_name": map[string]interface{}{
// 		"valid":           []interface{}{"98x59w2sfs", "fgekehkf98", "nvugshto6o", "ev88tb70y5"},
// 		"invalid":         []interface{}{10, 12.43},
// 		"multiple_valids": []interface{}{"sgji2x93uz", "unqs91vkej", "3x1w9fp4uz", "91wihgws94", "k1mnmwntru", "abw4locy8l", "86j50my56r", "uv83ulysy8", "i6s30lxr6q", "dze1kfe3lh", "txiho8iolk", "dj1itm8oaj", "03n1hnwufd", "sf5f57m23a", "hrtt19bd96"},
// 	},

// 	"description": map[string]interface{}{
// 		"valid":           []interface{}{"p5bmblr2eg", "w0m0pobo1v", "h9mauw3xiw", "6ofktj75lf"},
// 		"invalid":         []interface{}{10, 12.43},
// 		"multiple_valids": []interface{}{"2n37hadx3z", "d4jvucj3hm", "b258wb7hrp", "xnwi5exfnb", "6qvq7w9tw7", "qytqvhv28h", "giv8hqi4yv", "aecbg3ws1v", "t2ujm0uoc2", "68p2wgp8m5", "mydjz5hn14", "6g4tz5t04o", "xx6212mbaw", "gf4zfyoj7g", "o1w9033049"},
// 	},

// 	"state": map[string]interface{}{
// 		"valid":           []interface{}{"ACTIVE", "INACTIVE"},
// 		"invalid":         []interface{}{"ai0vkwvyju"},
// 		"multiple_valids": []interface{}{"ACTIVE", "INACTIVE"},
// 	},

// 	"connection_details": map[string]interface{}{
// 		"client_id": map[string]interface{}{
// 			"valid":           searchInObject(Test, "client_id.valid"),
// 			"invalid":         searchInObject(Test, "client_id.invalid"),
// 			"multiple_valids": searchInObject(Test, "client_id.multiple_valids"),
// 		},

// 		"client_secret": map[string]interface{}{
// 			"valid":           []interface{}{"ug13uhrx0w", "1ieuuza11y", "up6orjjatx", "sw2kbyhcj8"},
// 			"invalid":         []interface{}{10, 12.43},
// 			"multiple_valids": []interface{}{"2b1ub8svz4", "ch6dkdwlpp", "1m9jfwvclr", "7rm3rvhumv", "vgjj2gzpxn", "tkqccpuayq", "ts1vcy9wbv", "729io0591i", "a8ijqshhbm", "4ymwsicxd9", "9npfqov06d", "tq7vcvsn7n", "vgk6f6046u", "aogk3q4z2b", "webhu85ij3"},
// 		},

// 		"tenant_id": map[string]interface{}{
// 			"valid":           []interface{}{"w7ikz73b4b", "69vsqk4phu", "1sdfj0ulm0", "ctq34a9ib9"},
// 			"invalid":         []interface{}{10, 12.43},
// 			"multiple_valids": []interface{}{"xknqow3aj2", "5l0p5xesoe", "by10e8pbps", "dudzlyxnfj", "7jhggif8xi", "rgdzqpqo5i", "xqrddeefpa", "thq7icyzmy", "s4layatjh8", "gel0ss0r8y", "v1ycr4kvlv", "tewxceassz", "fd5zwzm6fc", "juv8adlsp2", "k8bj1ht0ml"},
// 		},

// 		"subscription_id": map[string]interface{}{
// 			"valid":           searchInObject(Test, "subscription_id.valid"),
// 			"invalid":         searchInObject(Test, "subscription_id.invalid"),
// 			"multiple_valids": searchInObject(Test, "subscription_id.multiple_valids"),
// 		},
// 	},

// 	"configuration_details": map[string]interface{}{
// 		"regions": map[string]interface{}{
// 			"valid":           []interface{}{"vjpugsry3h", "lo0g9snlus", "t50xqbkk6m", "huqa2a5mlt"},
// 			"invalid":         []interface{}{10, 12.43},
// 			"multiple_valids": []interface{}{"6ukrh8098a", "dhgk6fidco", "w68ba78ktx", "kojweazi2o", "t0kamn445p", "9wup9idm78", "atoqxxzw8g", "2kzpdbpj2n", "f6rucrz9bs", "w6cp5rkycn", "tzp22gcj5b", "wuzactqcv8", "8lbdu54ja6", "a88hnv1428", "fffuwwym7e"},
// 		},

// 		"resource_groups": map[string]interface{}{
// 			"valid":           []interface{}{"wyp34max3g", "qjwmmn03al", "14tqrysy97", "dy7zydzl6a"},
// 			"invalid":         []interface{}{10, 12.43},
// 			"multiple_valids": []interface{}{"ff2ehhzz6w", "e5dqvg4534", "ivr6trnxib", "vci2r8cccv", "9pn209vwta", "bxtqtzz4hc", "g7vpime1it", "e89tp5xz05", "hri1n4ft8n", "3q130c28r8", "u95fxvfsf8", "j2xx06x3mc", "pqu952ow6v", "m81syxqi7h", "2l2f46tt52"},
// 		},

// 		"polling": map[string]interface{}{
// 			"interval": map[string]interface{}{
// 				"valid":           []interface{}{-404, 67, -753, 129},
// 				"invalid":         []interface{}{"random", 10.023},
// 				"multiple_valids": []interface{}{-147, -849, 433, 847, 550, -69, -424, -61, -501, 1, -699, -754, -129, 785, -450},
// 			},

// 			"unit": map[string]interface{}{
// 				"valid":           []interface{}{"minute"},
// 				"invalid":         []interface{}{"tphqj2i6cb"},
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
// 				"valid":           []interface{}{"j0c8yg2w9u", "08rw88zrgx", "ypg30vq7ya", "bbmn5sr18j"},
// 				"invalid":         []interface{}{10, 12.43},
// 				"multiple_valids": []interface{}{"q1l5yw0dwm", "qomyte9mgd", "nea1hrghlg", "7m1kwkgbpg", "i4x2trfu1y", "kw1px3pvr7", "wdc3wip8j5", "q0ms78lway", "s3oofcpdsv", "kud9r0c6nt", "bv74ohbrhe", "0llxadx6v2", "vfr0z4ltkt", "jgmnz221vm", "1hslpndgld"},
// 			},
// 		},

// 		"tag_filter": map[string]interface{}{
// 			"valid":           []interface{}{"klnxyhwiu0", "tjvc40zpvk", "xelo9pcd5r", "uhz9lnue11"},
// 			"invalid":         []interface{}{10, 12.43},
// 			"multiple_valids": []interface{}{"jkvl6qj5bx", "oze1z6l49b", "foh39vdcor", "8wycy5zjls", "0xqvi8cnzo", "z8n48xwtv2", "y9bi7kqzdc", "ukolbsxxg8", "39ivvgpr5h", "noqgiqiquj", "c6rokykhon", "5l0o6mgce8", "31pgvcqy27", "p76yj7h2uz", "xp5cb7uuoy"},
// 		},

// 		"services": map[string]interface{}{
// 			"name": map[string]interface{}{
// 				"valid":           []interface{}{"qjvmr2ksh1", "vk2pwe65hw", "hwzspufhqt", "z82pkga0e7"},
// 				"invalid":         []interface{}{10, 12.43},
// 				"multiple_valids": []interface{}{"ifz0lm0h58", "5dkkv1w90w", "t2e45wrbkg", "a658297cr0", "16apxv35rd", "zxu7j6gy74", "93bnyflbn2", "gciwea3goo", "gahlkqm4g4", "cc89eoeabx", "f427ph6w74", "0ika8lc3ju", "rek2ko0cr1", "vn7iix467u", "7yyfmyxrig"},
// 			},

// 			"import_tags": map[string]interface{}{
// 				"enabled": map[string]interface{}{
// 					"valid":           []interface{}{true, false},
// 					"invalid":         []interface{}{"random", 10},
// 					"multiple_valids": []interface{}{true, false},
// 				},

// 				"excluded_keys": map[string]interface{}{
// 					"valid":           []interface{}{"jpuxo9p7cz", "fyopn0g1of", "rhobn8kj5p", "lavhldvuyj"},
// 					"invalid":         []interface{}{10, 12.43},
// 					"multiple_valids": []interface{}{"pr04ucsrrz", "kk1aecqh3h", "qq0rkd8zu9", "ns64mf8ohf", "ewnsxyu7z5", "0l5xu506yk", "i8wovqc4y4", "im1fcs0xii", "cpo91zyga9", "1byzi4i4cz", "k5jkne2ud3", "xvzqhv2hh2", "1vb3h94ay9", "qaq19rmrbh", "w80fb0owxh"},
// 				},
// 			},

// 			"tag_filter": map[string]interface{}{
// 				"valid":           []interface{}{"vpatjoad4b", "sutccbv48z", "dn67r47z8d", "arlczgpoci"},
// 				"invalid":         []interface{}{10, 12.43},
// 				"multiple_valids": []interface{}{"fjnwgtnzwt", "63bmoixilt", "cp8ro3q72f", "5x5ub1hjy6", "6mq07t5cwu", "xl17g9qar6", "oug78ovacj", "0isaeyvnbf", "macckilefc", "v7st8s7kmt", "dibjhdhdcr", "57czl2k1gr", "lcxgth45bs", "v6u81b0fcz", "5kbn2ufqky"},
// 			},

// 			"polling": map[string]interface{}{
// 				"interval": map[string]interface{}{
// 					"valid":           []interface{}{-961, -991, 602, -165},
// 					"invalid":         []interface{}{"random", 10.023},
// 					"multiple_valids": []interface{}{152, 652, 45, -279, -920, 55, -842, 206, -118, -491, 412, -582, 826, -243, 819},
// 				},

// 				"unit": map[string]interface{}{
// 					"valid":           []interface{}{"minute"},
// 					"invalid":         []interface{}{"gbg097pvls"},
// 					"multiple_valids": []interface{}{"minute"},
// 				},
// 			},
// 		},
// 	},
// }

// func TestAccAppdynamicscloudConnectionAzure_Basic(t *testing.T) {
// 	var connectionAzure_default models.ConnectionAzure
// 	var connectionAzure_updated models.ConnectionAzure
// 	resourceName := "appdynamicscloud_connection_azure.test"

// 	rName := makeTestVariable(acctest.RandString(5))
// 	// rOther := makeTestVariable(acctest.RandString(5))

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAzureDestroy,
// 		Steps: append([]resource.TestStep{
// 			{
// 				Config:      CreateAccConnectionAzureWithoutDisplayName(rName),
// 				ExpectError: regexp.MustCompile(`Missing required argument`),
// 			},
// 			{
// 				Config: CreateAccConnectionAzureConfig(rName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, &connectionAzure_default),
// 					resource.TestCheckResourceAttr(resourceName, "display_name", rName),

// 					resource.TestCheckResourceAttr(resourceName, "description", ""),

// 					resource.TestCheckResourceAttr(resourceName, "state", ""),

// 					resource.TestCheckResourceAttr(resourceName, "connection_details.#", "0"),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "0"),
// 				),
// 			},
// 			{
// 				Config: CreateAccConnectionAzureConfigWithOptional(rName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, &connectionAzure_updated),
// 					resource.TestCheckResourceAttr(resourceName, "display_name", rName),
// 					resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "description.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "state", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "state.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.0.client_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.0.client_secret", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.0.tenant_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "connection_details.0.subscription_id", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"))),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.#", "2"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.1"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.resource_groups.#", "2"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.resource_groups.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.resource_groups.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.1"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"))),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.#", "2"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"))),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.name", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.services.name.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.enabled", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.enabled.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.#", "2"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.0", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.1", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.1"))),

// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.tag_filter", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.services.tag_filter.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.#", "1"),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.interval", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.interval.valid.0"))),
// 					resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.unit", fmt.Sprintf("%v", searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.valid.0"))),

// 					testAccCheckAppdynamicscloudConnectionAzureIdEqual(&connectionAzure_default, &connectionAzure_updated),
// 				),
// 			},
// 			{
// 				ResourceName:      resourceName,
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 			{
// 				Config: CreateAccConnectionAzureConfig(rName),
// 			},
// 		}, generateStepForUpdatedRequiredAttrConnectionAzure(rName, resourceName, &connectionAzure_default, &connectionAzure_updated)...),
// 	})
// }

// func TestAccAppdynamicscloudConnectionAzure_Update(t *testing.T) {
// 	var connectionAzure_default models.ConnectionAzure
// 	var connectionAzure_updated models.ConnectionAzure
// 	resourceName := "appdynamicscloud_connection_azure.test"
// 	rName := makeTestVariable(acctest.RandString(5))

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAzureDestroy,
// 		Steps: append([]resource.TestStep{
// 			{
// 				Config: CreateAccConnectionAzureConfig(rName),
// 				Check:  testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, &connectionAzure_default),
// 			},
// 		}, generateStepForUpdatedAttrConnectionAzure(rName, resourceName, &connectionAzure_default, &connectionAzure_updated)...),
// 	})
// }

// func TestAccAppdynamicscloudConnectionAzure_NegativeCases(t *testing.T) {
// 	resourceName := "appdynamicscloud_connection_azure.test"

// 	// [TODO]: Add makeTestVariable() to utils.go file
// 	rName := makeTestVariable(acctest.RandString(5))

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCheckAppdynamicscloudConnectionAzureDestroy,
// 		Steps: append([]resource.TestStep{
// 			{
// 				Config: CreateAccConnectionAzureConfig(rName),
// 			},
// 		}, generateNegativeStepsConnectionAzure(rName, resourceName)...),
// 	})
// }

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

// func CreateAccConnectionAzureWithoutDisplayName(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 				resource  "appdynamicscloud_connection_azure" "test" {

// 									description = "%v"

// 									state = "%v"

// 									connection_details {
    
									                        
//                                         client_id = "%v"
                        
//                                         client_secret = "%v"
                        
//                                         tenant_id = "%v"
                        
//                                         subscription_id = "%v"

// 									}

// 									configuration_details {
    
									                        
//                                         regions = ["%v","%v"]
                        
//                                         resource_groups = ["%v","%v"]

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
// 			`, searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
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

// func CreateAccConnectionAzureConfig(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 		resource  "appdynamicscloud_connection_azure" "test" {


// 							display_name = "%v"
// 		}
// 	`, rName)
// 	return resource
// }

// func CreateAccConnectionAzureConfigWithOptional(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]

// 	resource += createConnectionAzureConfig(parentResources)

// 	resource += fmt.Sprintf(`
// 		resource  "appdynamicscloud_connection_azure" "test" {

// 						display_name = "%v"

// 						description = "%v"

// 						state = "%v"

//                         connection_details {
    
                                                
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

//                         }

//                         configuration_details {
    
                                                
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

//                         }
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
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

// func generateStepForUpdatedRequiredAttrConnectionAzure(rName string, resourceName string, connectionAzure_default, connectionAzure_updated *models.ConnectionAzure) []resource.TestStep {
// 	testSteps := make([]resource.TestStep, 0, 1)
// 	var value interface{}
// 	value = searchInObject(resourceConnectionAzureTest, "display_name.valid.1")
// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAzureUpdateRequiredDisplayName(rName),
// 		Check: resource.ComposeTestCheckFunc(
// 			testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 			resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", value)),
// 			testAccCheckAppdynamicscloudConnectionAzureIdNotEqual(connectionAzure_default, connectionAzure_updated),
// 		),
// 	})
// 	return testSteps
// }
// func CreateAccConnectionAzureUpdateRequiredDisplayName(rName string) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	value := searchInObject(resourceConnectionAzureTest, "display_name.valid.1")
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {
							
// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

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
// 		`, value,
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
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

// func CreateAccConnectionAzureUpdatedAttrDescription(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"
							
// 							description = "%v"

// 							state = "%v"

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
// 		}
// 	`, rName,
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
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
// func CreateAccConnectionAzureUpdatedAttrState(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"
							
// 							state = "%v"

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
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
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
// 	return resource
// }
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsRegions(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
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
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsResourceGroups(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		value,
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
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsPollingInterval(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "polling.unit.valid.0"),
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
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsPollingUnit(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "polling.interval.valid.0"),
// 		value,
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
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsImportTagsEnabled(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "excluded_keys.valid.1"),
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
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsImportTagsExcludedKeys(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "import_tags.enabled.valid.0"),
// 		value,
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
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsTagFilter(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesImportTagsEnabled(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.name.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesImportTagsExcludedKeys(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "import_tags.enabled.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesTagFilter(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.import_tags.excluded_keys.valid.1"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "services.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesPollingInterval(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "services.tag_filter.valid.0"),
// 		value,
// 		searchInObject(resourceConnectionAzureTest, "polling.unit.valid.0"))
// 	return resource
// }
// func CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesPollingUnit(rName string, value interface{}) string {
// 	var resource string
// 	parentResources := getParentConnectionAzure(rName)
// 	parentResources = parentResources[:len(parentResources)-1]
// 	resource += createConnectionAzureConfig(parentResources)
// 	resource += fmt.Sprintf(`
// 			resource "appdynamicscloud_connection_azure" "test" {

// 							display_name = "%v"

// 							description = "%v"

// 							state = "%v"

// 							connection_details {
    
							                        
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

// 							}

// 							configuration_details {
    
							
// 						    regions = ["%v", "%v"]
  

// 						    resource_groups = ["%v", "%v"]
  

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
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "regions.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "resource_groups.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.name.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.import_tags.enabled.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.import_tags.excluded_keys.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "services.import_tags.excluded_keys.valid.1"),
// 		searchInObject(resourceConnectionAzureTest, "services.tag_filter.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "polling.interval.valid.0"),
// 		value)
// 	return resource
// }

// func generateStepForUpdatedAttrConnectionAzure(rName string, resourceName string, connectionAzure_default, connectionAzure_updated *models.ConnectionAzure) []resource.TestStep {
// 	testSteps := make([]resource.TestStep, 0, 1)
// 	var valid []interface{}
// 	valid = searchInObject(resourceConnectionAzureTest, "description.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrDescription(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "description", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAzureTest, "state.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrState(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "state", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}

// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsRegions(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.regions.0", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsResourceGroups(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.resource_groups.0", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsPollingInterval(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.interval", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsPollingUnit(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.polling.0.unit", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}

// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsImportTagsEnabled(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.enabled", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsImportTagsExcludedKeys(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.import_tags.0.excluded_keys.0", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}

// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsTagFilter(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.tag_filter", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.enabled.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesImportTagsEnabled(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.enabled", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesImportTagsExcludedKeys(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.import_tags.0.excluded_keys.0", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}

// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.services.tag_filter.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesTagFilter(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.tag_filter", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.interval.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesPollingInterval(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.interval", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}
// 	valid = searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.valid").([]interface{})
// 	for _, value := range valid {
// 		v := fmt.Sprintf("%v", value)
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config: CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesPollingUnit(rName, value),
// 			Check: resource.ComposeTestCheckFunc(
// 				testAccCheckAppdynamicscloudConnectionAzureExists(resourceName, connectionAzure_updated),
// 				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.services.0.polling.0.unit", v),
// 				testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure_default, connectionAzure_updated),
// 			),
// 		})
// 	}

// 	return testSteps
// }

// func generateNegativeStepsConnectionAzure(rName string, resourceName string) []resource.TestStep {
// 	//Use Update Config Function with false value
// 	testSteps := make([]resource.TestStep, 0, 1)
// 	var invalid []interface{}
// 	invalid = searchInObject(resourceConnectionAzureTest, "state.invalid").([]interface{})
// 	for _, value := range invalid {
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config:      CreateAccConnectionAzureUpdatedAttrState(rName, value),
// 			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
// 		})
// 	}

// 	invalid = searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.invalid").([]interface{})
// 	for _, value := range invalid {
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config:      CreateAccConnectionAzureUpdatedAttrConfigurationDetailsPollingUnit(rName, value),
// 			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
// 		})
// 	}

// 	invalid = searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.invalid").([]interface{})
// 	for _, value := range invalid {
// 		testSteps = append(testSteps, resource.TestStep{
// 			Config:      CreateAccConnectionAzureUpdatedAttrConfigurationDetailsServicesPollingUnit(rName, value),
// 			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
// 		})
// 	}

// 	testSteps = append(testSteps, resource.TestStep{
// 		Config: CreateAccConnectionAzureConfig(rName),
// 	})
// 	return testSteps
// }

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
// 		`, i, val,
// 			searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "state.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "connection_details.client_id.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "connection_details.client_secret.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "connection_details.tenant_id.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "connection_details.subscription_id.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.regions.valid.1"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.resource_groups.valid.1"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.polling.interval.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.polling.unit.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.enabled.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.import_tags.excluded_keys.valid.1"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.tag_filter.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.services.name.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.enabled.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.services.import_tags.excluded_keys.valid.1"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.services.tag_filter.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.interval.valid.0"),
// 			searchInObject(resourceConnectionAzureTest, "configuration_details.services.polling.unit.valid.0"))
// 	}
// 	return resource
// }

// func testAccCheckAppdynamicscloudConnectionAzureExists(name string, connectionAzure *models.ConnectionAzure) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		// [TODO]: Write your code here
// 	}
// }

// func testAccCheckAppdynamicscloudConnectionAzureDestroy(s *terraform.State) error {
// 	client := testAccProvider.Meta().(*client.Client)

// 	for _, rs := range s.RootModule().Resources {

// 		if rs.Type == "appdynamicscloud_connection_azure" {
// 			// [TODO]: Write your code here
// 		}
// 	}
// 	return nil
// }

// func testAccCheckAppdynamicscloudConnectionAzureIdEqual(connectionAzure1, connectionAzure2 *models.ConnectionAzure) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		Id1, err := getIdFromConnectionAzureModel(connectionAzure1)
// 		if err != nil {
// 			return err
// 		}
// 		Id2, err := getIdFromConnectionAzureModel(connectionAzure2)
// 		if err != nil {
// 			return err
// 		}
// 		if Id1 != Id2 {
// 			return fmt.Errorf("ConnectionAzure IDs are not equal")
// 		}
// 		return nil
// 	}
// }

// func testAccCheckAppdynamicscloudConnectionAzureIdNotEqual(connectionAzure1, connectionAzure2 *models.ConnectionAzure) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		Id1, err := getIdFromConnectionAzureModel(connectionAzure1)
// 		if err != nil {
// 			return err
// 		}
// 		Id2, err := getIdFromConnectionAzureModel(connectionAzure2)
// 		if err != nil {
// 			return err
// 		}
// 		if Id1 == Id2 {
// 			return fmt.Errorf("ConnectionAzure IDs are equal")
// 		}
// 		return nil
// 	}
// }

// func getParentConnectionAzure(rName string) []string {
// 	t := []string{}
// 	t = append(t, connectionAzureBlock(rName))
// 	return t
// }

// func connectionAzureBlock(rName string) string {
// 	return fmt.Sprintf(`
// 		resource  "appdynamicscloud_connection_azure" "test" {

// 						display_name = "%v"


// 						description = "%v"


// 						state = "%v"


//                         connection_details {
    
                                                
//                             client_id = "%v"
                        
//                             client_secret = "%v"
                        
//                             tenant_id = "%v"
                        
//                             subscription_id = "%v"

//                         }

//                         configuration_details {
    
                                                
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

//                         }
// 		}
// 	`, rName,
// 		searchInObject(resourceConnectionAzureTest, "description.valid.0"),
// 		searchInObject(resourceConnectionAzureTest, "state.valid.0"),
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
// }

// // To eliminate duplicate resource block from slice of resource blocks
// func createConnectionAzureConfig(configSlice []string) string {
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
