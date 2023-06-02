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
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/cisco-open/appd-cloud-go-client/apis/v1/applicationprincipalmanagement"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var resourceAccessClientAppTest = map[string]interface{}{
	"display_name": map[string]interface{}{
		"valid":           []interface{}{"k8tq9w4lbt", "iyc3m20zsm", "grv6or925e", "gld9q2vkg6"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"ct836hydhd", "aa451d5fyf", "588v1wg0dy"},
	},

	"description": map[string]interface{}{
		"valid":           []interface{}{"TestAcc_u2fbqws7od", "TestAcc_ivnrov2ug5", "TestAcc_64lxo2tqlu", "TestAcc_j2i5ocro03"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"TestAcc_vhlr5zo9b0", "TestAcc_92mq72kzg9", "TestAcc_wox9azec5l", "TestAcc_clplwuhu3q", "TestAcc_xxkabfo4y9", "TestAcc_oyyaiqjkpm"},
	},

	"auth_type": map[string]interface{}{
		"valid":           []interface{}{"client_secret_basic", "client_secret_post"},
		"invalid":         []interface{}{"4yskmyfv0m"},
		"multiple_valids": []interface{}{"client_secret_basic", "client_secret_post"},
	},

	"rotate_secret": map[string]interface{}{
		"valid":           []interface{}{true, false},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{true, false},
	},

	"revoke_previous_secret_in": map[string]interface{}{
		"valid":           []interface{}{"NOW", "1D", "3D", "7D", "30D"},
		"invalid":         []interface{}{"k3674u7g41"},
		"multiple_valids": []interface{}{"NOW", "1D", "3D", "7D", "30D"},
	},

	"revoke_now": map[string]interface{}{
		"valid":           []interface{}{true, false},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{true, false},
	},
}

// BASIC - checks that creation is successful with valid
// 		   combination required and optional values
// REQUIRED ATTRIBUTES: display_name, description, auth_type
// OPTIONAL ATTRIBUTES: rotate_secret, revoke_previous, revoke_now
//
// === CASES:
// 		1. without required attributes
// 		2. with only required attributes

// revoke_previous depends on value of rotate_secret
// revoke_now conflicts with revoke_previous
//		3. with rotate_secret and revoke_previous optional attributes
//		4. with revoke_now
// 		5. import and compare state
// 		6. update required attributes with valid values
//
// === EXPECT ERRORS:
//		1. invalid value of revoke_previous (enum validation) [case 3]
//		2. revoke_previous without its dependency revoke_now  [case 3]
// 		3. revoke_previous without revoke_now: false		  [case 3]
// 		4. rotate_now: true without revoke_previous			  [case 3]
//		5. revoke_now with its conflicting attribute          [case 4]
//
// this case will test creating resource with all combinations of valid values.
// expected errors for negative combinations and values during creation will be
// tested in _NegativeCases method.

func TestAccAppdynamicscloudAccessClientApp_Basic(t *testing.T) {
	var accessClientApp_default applicationprincipalmanagement.ServiceClientResponse
	var accessClientApp_updated applicationprincipalmanagement.ServiceClientResponse

	resourceName := "appdynamicscloud_access_client_app.test"
	rName := makeTestVariable(acctest.RandString(5))

	testSteps := make([]resource.TestStep, 0, 1)

	basicTests := []resource.TestStep{
		// [case 1]
		// without required
		{
			Config:      CreateAccAccessClientAppWithoutDisplayName(rName),
			ExpectError: regexp.MustCompile(`Missing required argument`),
		},
		{
			Config:      CreateAccAccessClientAppWithoutDescription(rName),
			ExpectError: regexp.MustCompile(`Missing required argument`),
		},
		{
			Config:      CreateAccAccessClientAppWithoutAuthType(rName),
			ExpectError: regexp.MustCompile(`Missing required argument`),
		},

		// [case 2]
		// with only required
		{
			Config: CreateAccAccessClientAppConfig(rName),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, &accessClientApp_default),

				resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "display_name.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "description.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "auth_type", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))),
			),
		},
	}

	testSteps = append(testSteps, basicTests...)

	// [case 3]
	// with optional rotate_secret
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccAccessClientAppConfigWithOptionalRotate(rName, false, nil),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, &accessClientApp_updated),

			resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "display_name.valid.0"))),
			resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "description.valid.0"))),
			resource.TestCheckResourceAttr(resourceName, "auth_type", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))),
			resource.TestCheckResourceAttr(resourceName, "rotate_secret", "false"),
			resource.TestCheckResourceAttr(resourceName, "revoke_now", "false"),

			testAccCheckAppdynamicscloudAccessClientAppIdEqual(&accessClientApp_default, &accessClientApp_updated),
		),
	})

	// [case 3]
	// with optional rotate_secret and revoke_previous dependency
	validRevokePrevious := searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid").([]interface{})
	for _, value := range validRevokePrevious {
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccAccessClientAppConfigWithOptionalRotate(rName, true, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, &accessClientApp_updated),

				resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "display_name.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "description.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "auth_type", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "rotate_secret", "false"),
				resource.TestCheckResourceAttr(resourceName, "revoke_previous_secret_in", ""),
				resource.TestCheckResourceAttr(resourceName, "revoke_now", "false"),

				testAccCheckAppdynamicscloudAccessClientAppIdEqual(&accessClientApp_default, &accessClientApp_updated),
			),
			ExpectNonEmptyPlan: true,
		})
	}

	// [case 4]
	// with revoke_now
	testSteps = append(testSteps, []resource.TestStep{
		{
			Config: CreateAccAccessClientAppConfigWithOptionalRevoke(rName, true),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, &accessClientApp_updated),

				resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "display_name.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "description.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "auth_type", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "rotate_secret", "false"),
				resource.TestCheckResourceAttr(resourceName, "revoke_previous_secret_in", ""),
				resource.TestCheckResourceAttr(resourceName, "revoke_now", "false"),

				testAccCheckAppdynamicscloudAccessClientAppIdEqual(&accessClientApp_default, &accessClientApp_updated),
			),
			ExpectNonEmptyPlan: true,
		},
		{
			Config: CreateAccAccessClientAppConfigWithOptionalRevoke(rName, false),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, &accessClientApp_updated),

				resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "display_name.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "description.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "auth_type", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))),
				resource.TestCheckResourceAttr(resourceName, "rotate_secret", "false"),
				resource.TestCheckResourceAttr(resourceName, "revoke_previous_secret_in", ""),
				resource.TestCheckResourceAttr(resourceName, "revoke_now", "false"),

				testAccCheckAppdynamicscloudAccessClientAppIdEqual(&accessClientApp_default, &accessClientApp_updated),
			),
		}}...,
	)

	// [case 5]
	// import and compare state
	testSteps = append(testSteps, []resource.TestStep{
		{
			ResourceName:      resourceName,
			ImportState:       true,
			ImportStateVerify: true,
			// client_secret is not sent in the read call. revoke_previous_secret_in, revoke_now, and rotate_secret are
			// meta arguments introduced by us in terraform and not returned by the API. rotated_secret_expires_at is only
			// fetched during rotate call, it will not be part of the GET call.
			ImportStateVerifyIgnore: []string{"client_secret", "revoke_previous_secret_in", "revoke_now", "rotate_secret", "rotated_secret_expires_at"},
		},
		{
			Config: CreateAccAccessClientAppConfig(rName),
		},
	}...)

	// [case 6]
	// update required attributes with valid values
	testSteps = append(
		testSteps,
		generateStepForUpdatedRequiredAttrAccessClientApp(rName, resourceName, &accessClientApp_default, &accessClientApp_updated)...,
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudAccessClientAppDestroy,
		Steps:             testSteps,
	})
}

// UPDATE - check that update is successful with all possible values
// 			for optional attributes, separately and combined.
// OPTIONAL ATTRIBUTES: rotate_secret, revoke_previous, revoke_now
//
// === CASES:
// rotate_secret and revoke_previous depend on each other
//		1. rotate_secret: present, revoke_previous: omit
//		2. rotate_secret: omit,    revoke_previous: present
//		3. rotate_secret: present, revoke_previous: present
//
// revoke_now conflicts with revoke_previous
// above test cases will confirm that revoke_previous can
// only be set with rotate_secret: true
//		4. revoke_now: present, revoke_previous: absent
//
// === EXPECT ERRORS:
//		1. rotate_secret: true,  revoke_previous: omit     [case 1]
//		2. rotate_secret: omit,  revoke_previous: present  [case 2]
//		3. rotate_secret: false, revoke_previous: present  [case 3]
//		4. revoke_now: true,     revoke_previous: present  [case 5]
//
// when the attribute is present, it shall be checked with all combinations of valid values.
// except when checking against conflictsWith attributes [case 5]
// once the presence of both attribute throw error, it dees not matter what the values are.

func TestAccAppdynamicscloudAccessClientApp_Update(t *testing.T) {
	var accessClientApp_default applicationprincipalmanagement.ServiceClientResponse
	var accessClientApp_updated applicationprincipalmanagement.ServiceClientResponse

	resourceName := "appdynamicscloud_access_client_app.test"

	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudAccessClientAppDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccAccessClientAppConfig(rName),
				Check:  testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, &accessClientApp_default),
			},
		}, generateStepForUpdatedAttrAccessClientApp(rName, resourceName, &accessClientApp_default, &accessClientApp_updated)...),
	})
}

// NEGATIVE - checks that resource cannot be created with invalid
// 			  combination of attribute or values
//
// === CASES:
// 	[BASIC]
//		1. invalid value of revoke_previous (enum validation)
//		2. revoke_previous without its dependency revoke_now
// 		3. revoke_previous without revoke_now: false
// 		4. rotate_now: true without revoke_previous
//		5. revoke_now with its conflicting attribute
//
// 	[UPDATE]
//		1. rotate_secret: omit,  revoke_previous: present [same as basic 2, ignored]
//		2. rotate_secret: false, revoke_previous: present [same as basic 3, ignored]
//		3. rotate_secret: true,  revoke_previous: omit 	  [same as basic 4, ignored]
//		4. revoke_now: true,     revoke_previous: present [same as basic 5, ignored]
//
// The attributes here will be checked for all possible combination of invalid values
// except when checking against conflictsWith attributes.
// once the presence of both attribute throw error, it dees not matter what the values are.

func TestAccAppdynamicscloudAccessClientApp_NegativeCases(t *testing.T) {
	resourceName := "appdynamicscloud_access_client_app.test"

	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudAccessClientAppDestroy,
		Steps: append([]resource.TestStep{
			// create with invalid display name
			{
				Config:      CreateAccAccessClientAppConfigDisplayNameInvalid(rName),
				ExpectError: regexp.MustCompile("Bad Request - The provided description is invalid."),
			},
			{
				Config: CreateAccAccessClientAppConfig(rName),
			},
			// update with invalid display name
			{
				Config:      CreateAccAccessClientAppConfigDisplayNameInvalid(rName),
				ExpectError: regexp.MustCompile("Bad Request - The provided description is invalid."),
			},
			// revoke_previous without rotate_secret
			{
				Config:      CreateAccAccessClientAppConfigRotateNegative(rName, nil, false, "NOW"),
				ExpectError: regexp.MustCompile(ERROR_ROTATE_NOT_PRESENT_OR_FALSE),
			},
			// revoke_previous with rotate_secret: false
			{
				Config:      CreateAccAccessClientAppConfigRotateNegative(rName, false, true, "NOW"),
				ExpectError: regexp.MustCompile(ERROR_ROTATE_NOT_PRESENT_OR_FALSE),
			},
			// rotate_secret: true without revoke_previous
			{
				Config:      CreateAccAccessClientAppUpdatedAttrRotateSecretOnly(rName, true),
				ExpectError: regexp.MustCompile(ERROR_REVOKE_TIMEOUT_NOT_PRESENT),
			},
			{
				Config: CreateAccAccessClientAppConfig(rName),
			},
			{
				Config:      CreateAccAccessClientAppUpdatedAttrRevokeNowConflictsWith(rName),
				ExpectError: regexp.MustCompile("Conflicting configuration arguments"),
			},
			// enum validation
		}, generateNegativeStepsAccessClientApp(rName, resourceName)...),
	})
}

func CreateAccAccessClientAppConfigDisplayNameInvalid(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "appdynamicscloud_access_client_app" "test" {

									display_name = "%v"

									description = "%v"

									auth_type = "%v"
				}
			`, acctest.RandStringFromCharSet(256, "abcdef0123456789"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))
	return resource
}

func CreateAccAccessClientAppConfigRotateNegative(rName string, rotate, includeRotate, revokeIn interface{}) string {
	var resource string

	resource = fmt.Sprintf(`
		resource  "appdynamicscloud_access_client_app" "test" {

						display_name = "%v"

						description = "%v"

						auth_type = "%v"

						revoke_previous_secret_in = "%v"

	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		revokeIn)

	if includeRotate.(bool) {
		resource += fmt.Sprintf(`
						rotate_secret = %v
		`, rotate)
	}

	resource += `
	}`

	return resource
}

func TestAccAppdynamicscloudAccessClientApp_MultipleCreateDelete(t *testing.T) {

	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudAccessClientAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: CreateAccAccessClientAppMultipleConfig(rName),
			},
		},
	})
}

func CreateAccAccessClientAppWithoutDisplayName(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "appdynamicscloud_access_client_app" "test" {

									description = "%v"

									auth_type = "%v"
				}
			`, searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))
	return resource
}
func CreateAccAccessClientAppWithoutDescription(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "appdynamicscloud_access_client_app" "test" {

									display_name = "%v"

									auth_type = "%v"
				}
			`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))
	return resource
}
func CreateAccAccessClientAppWithoutAuthType(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
				resource  "appdynamicscloud_access_client_app" "test" {

									display_name = "%v"

									description = "%v"
				}
			`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"))
	return resource
}

func CreateAccAccessClientAppConfig(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"

							auth_type = "%v"
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))
	return resource
}

func CreateAccAccessClientAppConfigWithOptionalRotate(rName string, rotate bool, revokeIn interface{}) string {
	var resource string

	resource = fmt.Sprintf(`
		resource  "appdynamicscloud_access_client_app" "test" {

						display_name = "%v"

						description = "%v"

						auth_type = "%v"

						rotate_secret = %v

	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		rotate)

	if rotate {
		resource += fmt.Sprintf(`
						revoke_previous_secret_in = "%v"
		`, revokeIn)
	}

	resource += `
	}`

	return resource
}

func CreateAccAccessClientAppConfigWithOptionalRevoke(rName string, revoke interface{}) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]

	resource += createAccessClientAppConfig(parentResources)

	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_access_client_app" "test" {

						display_name = "%v"

						description = "%v"

						auth_type = "%v"

						revoke_now = %v
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		revoke)
	return resource
}

func generateStepForUpdatedRequiredAttrAccessClientApp(rName string, resourceName string, accessClientApp_default, accessClientApp_updated *applicationprincipalmanagement.ServiceClientResponse) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var value interface{}
	value = searchInObject(resourceAccessClientAppTest, "display_name.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccAccessClientAppUpdateRequiredDisplayName(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
			resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
		),
	})
	value = searchInObject(resourceAccessClientAppTest, "description.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccAccessClientAppUpdateRequiredDescription(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
			resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
		),
	})
	value = searchInObject(resourceAccessClientAppTest, "auth_type.valid.1")
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccAccessClientAppUpdateRequiredAuthType(rName),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
			resource.TestCheckResourceAttr(resourceName, "auth_type", fmt.Sprintf("%v", value)),
			testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
		),
	})
	return testSteps
}
func CreateAccAccessClientAppUpdateRequiredDisplayName(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	value := searchInObject(resourceAccessClientAppTest, "display_name.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {
							
							display_name = "%v"

							description = "%v"

							auth_type = "%v"
			}
		`, value,
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))
	return resource
}
func CreateAccAccessClientAppUpdateRequiredDescription(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	value := searchInObject(resourceAccessClientAppTest, "description.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"
							
							description = "%v"

							auth_type = "%v"
			}
		`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		value,
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))

	return resource
}
func CreateAccAccessClientAppUpdateRequiredAuthType(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	value := searchInObject(resourceAccessClientAppTest, "auth_type.valid.1")
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"
							
							auth_type = "%v"
			}
		`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		value)
	return resource
}

func CreateAccAccessClientAppUpdatedAttrRotateSecretOnly(rName string, value interface{}) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"

							auth_type = "%v"
							
							rotate_secret = %v
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		value)
	return resource
}
func CreateAccAccessClientAppUpdatedAttrRevokePreviousSecretInOnly(rName string, value interface{}) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"

							auth_type = "%v"
							
							revoke_previous_secret_in = "%v"
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		value)
	return resource
}

func CreateAccAccessClientAppUpdatedAttrRotateWithRevokePreviousSecretIn(rName string, rotate, revokeIn interface{}) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"

							auth_type = "%v"
							
							rotate_secret = %v

							revoke_previous_secret_in = "%v"
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		rotate,
		revokeIn)
	return resource
}

func CreateAccAccessClientAppUpdatedAttrRevokeNow(rName string, value interface{}) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"

							auth_type = "%v"
							
							revoke_now = %v
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		value)
	return resource
}

func CreateAccAccessClientAppUpdatedAttrRevokeNowConflictsWith(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"

							auth_type = "%v"
							
							rotate_secret = true

							revoke_previous_secret_in = "NOW"

							revoke_now = true
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))
	return resource
}

func generateStepForUpdatedAttrAccessClientApp(rName string, resourceName string, accessClientApp_default, accessClientApp_updated *applicationprincipalmanagement.ServiceClientResponse) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)

	// [case 1]
	// rotate_secret: present, revoke_previous: absent,

	// rotate_secret: false, revoke_previous is not required
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccAccessClientAppUpdatedAttrRotateSecretOnly(rName, false),
		Check: resource.ComposeTestCheckFunc(
			testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
			resource.TestCheckResourceAttr(resourceName, "rotate_secret", "false"),
			testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
		),
	})

	// [case 1, error 1]
	// rotate_secret: true, revoke_previous is required but absent
	// handled in _NegativeCases

	// [case 2, error 2]
	// rotate_secret: absent, revoke_previous: *,
	// handled in _NegativeCases
	validRevokePrevious := searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid").([]interface{})

	// [case 3]
	// rotate_secret: present, revoke_previous: present
	for _, revokePreviousValue := range validRevokePrevious {
		// [error 3]
		// rotate: false, revoke_previous: *
		// revoke_previous can only be used when rotate_secret is present and is set to true
		// handled in _NegativeCases

		// rotate_secret: true, revoke_previous: *
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccAccessClientAppUpdatedAttrRotateWithRevokePreviousSecretIn(rName, true, revokePreviousValue),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
				resource.TestCheckResourceAttr(resourceName, "rotate_secret", "false"),
				resource.TestCheckResourceAttr(resourceName, "revoke_previous_secret_in", ""),
				testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
			),
			ExpectNonEmptyPlan: true,
		})
	}

	// [case 4]
	// revoke_now: present, revoke_previous: absent
	testSteps = append(testSteps, []resource.TestStep{
		{
			Config: CreateAccAccessClientAppUpdatedAttrRevokeNow(rName, true),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
				resource.TestCheckResourceAttr(resourceName, "revoke_now", "false"),
				testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
			),
			ExpectNonEmptyPlan: true,
		},
		{
			Config: CreateAccAccessClientAppUpdatedAttrRevokeNow(rName, false),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
				resource.TestCheckResourceAttr(resourceName, "revoke_now", "false"),
				testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
			),
		},
	}...)

	return testSteps
}

func generateNegativeStepsAccessClientApp(rName string, resourceName string) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)

	invalid := searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccAccessClientAppUpdatedAttrRevokePreviousSecretInOnly(rName, value),
			ExpectError: regexp.MustCompile(expectErrorMap["StringInSlice"]),
		})
	}
	testSteps = append(testSteps, resource.TestStep{
		Config: CreateAccAccessClientAppConfig(rName),
	})
	return testSteps
}

func CreateAccAccessClientAppMultipleConfig(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	multipleValues := searchInObject(resourceAccessClientAppTest, "display_name.multiple_valids").([]interface{})
	for i, val := range multipleValues {
		resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test%d" {
							
							display_name = "%v"

							description = "%v"

							auth_type = "%v"
			}
		`, i, val,
			searchInObject(resourceAccessClientAppTest, "description.valid.0"),
			searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))
	}
	return resource
}

func testAccCheckAppdynamicscloudAccessClientAppExists(name string, accessClientApp *applicationprincipalmanagement.ServiceClientResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Access Client App - Service Principal %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Client ID was set")
		}

		config := testAccProvider.Meta().(config)
		myCtx, _, apiClient := initializeApplicationPrincipalManagementClient(config)

		resp, _, err := apiClient.ServicesApi.GetServiceClientById(myCtx, rs.Primary.ID).Execute()
		if err != nil {
			return err
		}

		if resp.GetId() != rs.Primary.ID {
			return fmt.Errorf("Access Client App - Service Principal %s not found during read", rs.Primary.ID)
		}

		*accessClientApp = *resp

		return nil
	}
}

func testAccCheckAppdynamicscloudAccessClientAppDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(config)

	myCtx, _, apiClient := initializeApplicationPrincipalManagementClient(config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "appdynamicscloud_access_client_app" {
			_, _, err := apiClient.ServicesApi.GetServiceClientById(myCtx, rs.Primary.ID).Execute()

			if err == nil {
				return fmt.Errorf("Access Client App - Service Principal %s Still exists", rs.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp1, accessClientApp2 *applicationprincipalmanagement.ServiceClientResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if accessClientApp1.GetId() != accessClientApp2.GetId() {
			return fmt.Errorf("AccessClientApp IDs are not equal")
		}

		return nil
	}
}

func getParentAccessClientApp(rName string) []string {
	t := []string{}
	t = append(t, accessClientAppBlock(rName))
	return t
}

func accessClientAppBlock(rName string) string {
	return fmt.Sprintf(`
		resource  "appdynamicscloud_access_client_app" "test" {

						display_name = "%v"

						description = "%v"

						auth_type = "%v"

		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))
}

// To eliminate duplicate resource block from slice of resource blocks
func createAccessClientAppConfig(configSlice []string) string {
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

func TestAccessClientApp_CustomDiff_RotateSecretValidCombinations(t *testing.T) {
	cases := []struct {
		rotateSecretIsPresent  bool
		rotateSecretVal        bool
		revokeTimeoutIsPresent bool
		expectErr              bool
		expectErrMsg           string
	}{
		{
			rotateSecretIsPresent:  false,
			rotateSecretVal:        false,
			revokeTimeoutIsPresent: false,
			expectErr:              false,
		},
		{
			rotateSecretIsPresent:  false,
			rotateSecretVal:        false,
			revokeTimeoutIsPresent: true,
			expectErr:              true,
			expectErrMsg:           "revoke_previous_secret_in can only be used when rotate_secret is set to true",
		},
		{
			rotateSecretIsPresent:  true,
			rotateSecretVal:        false,
			revokeTimeoutIsPresent: false,
			expectErr:              false,
		},
		{
			rotateSecretIsPresent:  true,
			rotateSecretVal:        false,
			revokeTimeoutIsPresent: true,
			expectErr:              true,
			expectErrMsg:           "revoke_previous_secret_in can only be used when rotate_secret is set to true",
		},
		{
			rotateSecretIsPresent:  true,
			rotateSecretVal:        true,
			revokeTimeoutIsPresent: false,
			expectErr:              true,
			expectErrMsg:           "revoke_previous_secret_in must be set with rotate_secret",
		},
		{
			rotateSecretIsPresent:  true,
			rotateSecretVal:        true,
			revokeTimeoutIsPresent: true,
			expectErr:              false,
		},
	}

	for _, c := range cases {
		err := customDiffRotateSecretCheck(c.rotateSecretIsPresent, c.rotateSecretVal, c.revokeTimeoutIsPresent)

		if c.expectErr && err == nil {
			// if error is expected but got nil
			t.Fatalf("Expected combination of %v, %v and %v to throw err, but got nil", c.rotateSecretIsPresent, c.rotateSecretVal, c.revokeTimeoutIsPresent)
		} else if c.expectErr && err.Error() != c.expectErrMsg {
			// if error is expected but got different error
			t.Fatalf("Expected combination of %v, %v and %v to throw err %v, but got %v", c.rotateSecretIsPresent, c.rotateSecretVal, c.revokeTimeoutIsPresent, c.expectErrMsg, err.Error())
		}

	}
}

func TestAccessClientApp_InvalidCRUD(t *testing.T) {
	m := sharedClient()

	d := resourceAccessClientApp().TestResourceData()
	d.SetId("foo")

	diag := resourceAccessClientApp().ReadContext(context.Background(), d, m)

	if diag == nil {
		t.Fatalf("expected read to fail, but it did not")
	}

	d.SetId("srv_3ZSmQDb7bsfxaI_invalid")
	diag = resourceAccessClientApp().ReadContext(context.Background(), d, m)

	if diag != nil {
		t.Fatalf("expected read to succeed, but it did not")
	}

	diag = resourceAccessClientApp().DeleteContext(context.Background(), d, m)
	if diag == nil {
		t.Fatalf("expected delete to fail, but it did not")

	}

	d.Set("revoke_now", true)
	diag = revokeSecret(d, m)
	if diag == nil {
		t.Fatalf("expected revoke to fail, but it did not")
	}

	d.Set("rotate_secret", true)
	d.Set("revoke_previous_secret_in", "invalidValue")
	diag = rotateSecret(d, m)

	if diag == nil {
		t.Fatalf("expected rotate to fail, but it did not")
	}
}

func TestAccessClientApp_GetRotationRequest(t *testing.T) {
	cases := []struct {
		revokePreviousSecretIn string
		expectedValue          string
	}{
		{
			revokePreviousSecretIn: "NOW",
			expectedValue:          "P0D",
		},
		{
			revokePreviousSecretIn: "1D",
			expectedValue:          "P1D",
		},
		{
			revokePreviousSecretIn: "3D",
			expectedValue:          "P3D",
		},
		{
			revokePreviousSecretIn: "7D",
			expectedValue:          "P7D",
		},
		{
			revokePreviousSecretIn: "30D",
			expectedValue:          "P31D",
		},
	}

	for _, c := range cases {
		out := getRotationRequest(c.revokePreviousSecretIn)

		if out.GetRevokeRotatedAfter().(string) != c.expectedValue {
			testErrorMessage(c.revokePreviousSecretIn, c.expectedValue, out)
		}
	}
}
