package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/applicationprincipalmanagement"
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
		"valid":           []interface{}{"u2fbqws7od", "ivnrov2ug5", "64lxo2tqlu", "j2i5ocro03"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"vhlr5zo9b0", "92mq72kzg9", "wox9azec5l", "clplwuhu3q", "xxkabfo4y9", "oyyaiqjkpm"},
	},

	"auth_type": map[string]interface{}{
		"valid":           []interface{}{"client_secret_basic", "client_secret_post"},
		"invalid":         []interface{}{"4yskmyfv0m"},
		"multiple_valids": []interface{}{"client_secret_basic", "client_secret_post"},
	},

	"rotate_secret": map[string]interface{}{
		"valid":           []interface{}{"12/20/2019", "05/16/2013", "03/12/2019", "08/01/2022"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"12/20/2019", "05/16/2013", "03/12/2019", "08/01/2022", "11/05/2019", "02/12/2015", "07/18/2020"},
	},

	"revoke_previous_secret_in": map[string]interface{}{
		"valid":           []interface{}{"NOW", "1D", "3D", "7D", "30D"},
		"invalid":         []interface{}{"k3674u7g41"},
		"multiple_valids": []interface{}{"NOW", "1D", "3D", "7D", "30D"},
	},

	"revoked_all_previous_at": map[string]interface{}{
		"valid":           []interface{}{"12/20/2019", "05/16/2013", "03/12/2019", "08/01/2022"},
		"invalid":         []interface{}{10, 12.43},
		"multiple_valids": []interface{}{"12/20/2019", "05/16/2013", "03/12/2019", "08/01/2022", "11/05/2019", "02/12/2015", "07/18/2020"},
	},
}

func TestAccAppdynamicscloudAccessClientApp_Basic(t *testing.T) {
	var accessClientApp_default applicationprincipalmanagement.ServiceClientResponse
	var accessClientApp_updated applicationprincipalmanagement.ServiceClientResponse

	resourceName := "appdynamicscloud_access_client_app.test"

	rName := makeTestVariable(acctest.RandString(5))
	// rOther := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudAccessClientAppDestroy,
		Steps: append([]resource.TestStep{
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
			{
				Config: CreateAccAccessClientAppConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, &accessClientApp_default),

					resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "display_name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "description.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "auth_type", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))),
				),
			},
			{
				Config: CreateAccAccessClientAppConfigWithOptional(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, &accessClientApp_updated),

					resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "display_name.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "description", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "description.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "auth_type", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "rotate_secret", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "revoke_previous_secret_in", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"))),
					resource.TestCheckResourceAttr(resourceName, "revoked_all_previous_at", fmt.Sprintf("%v", searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))),

					testAccCheckAppdynamicscloudAccessClientAppIdEqual(&accessClientApp_default, &accessClientApp_updated),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_secret", "revoke_previous_secret_in", "revoked_all_previous_at", "rotate_secret", "rotated_secret_expires_at"},
			},
			{
				Config: CreateAccAccessClientAppConfig(rName),
			},
		}, generateStepForUpdatedRequiredAttrAccessClientApp(rName, resourceName, &accessClientApp_default, &accessClientApp_updated)...),
	})
}

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

func TestAccAppdynamicscloudAccessClientApp_NegativeCases(t *testing.T) {
	resourceName := "appdynamicscloud_access_client_app.test"

	// [TODO]: Add makeTestVariable() to utils.go file
	rName := makeTestVariable(acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckAppdynamicscloudAccessClientAppDestroy,
		Steps: append([]resource.TestStep{
			{
				Config: CreateAccAccessClientAppConfig(rName),
			},
		}, generateNegativeStepsAccessClientApp(rName, resourceName)...),
	})
}

func TestAccAppdynamicscloudAccessClientApp_MultipleCreateDelete(t *testing.T) {

	// [TODO]: Add makeTestVariable() to utils.go file
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

									rotate_secret = "%v"

									revoke_previous_secret_in = "%v"

									revoked_all_previous_at = "%v"
				}
			`, searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
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

									rotate_secret = "%v"

									revoke_previous_secret_in = "%v"

									revoked_all_previous_at = "%v"
				}
			`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
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

									rotate_secret = "%v"

									revoke_previous_secret_in = "%v"

									revoked_all_previous_at = "%v"
				}
			`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
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

func CreateAccAccessClientAppConfigWithOptional(rName string) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]

	resource += createAccessClientAppConfig(parentResources)

	resource += fmt.Sprintf(`
		resource  "appdynamicscloud_access_client_app" "test" {

						display_name = "%v"

						description = "%v"

						auth_type = "%v"

						rotate_secret = "%v"

						revoke_previous_secret_in = "%v"

						revoked_all_previous_at = "%v"
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
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

							rotate_secret = "%v"

							revoke_previous_secret_in = "%v"

							revoked_all_previous_at = "%v"
			}
		`, value,
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
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

							rotate_secret = "%v"

							revoke_previous_secret_in = "%v"

							revoked_all_previous_at = "%v"
			}
		`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		value,
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
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

							rotate_secret = "%v"

							revoke_previous_secret_in = "%v"

							revoked_all_previous_at = "%v"
			}
		`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		value,
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
	return resource
}

func CreateAccAccessClientAppUpdatedAttrRotateSecret(rName string, value interface{}) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"

							auth_type = "%v"
							
							rotate_secret = "%v"

							revoke_previous_secret_in = "%v"

							revoked_all_previous_at = "%v"
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		value,
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
	return resource
}
func CreateAccAccessClientAppUpdatedAttrRevokePreviousSecretIn(rName string, value interface{}) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"

							auth_type = "%v"

							rotate_secret = "%v"
							
							revoke_previous_secret_in = "%v"

							revoked_all_previous_at = "%v"
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		value,
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
	return resource
}
func CreateAccAccessClientAppUpdatedAttrRevokedAllPreviousAt(rName string, value interface{}) string {
	var resource string
	parentResources := getParentAccessClientApp(rName)
	parentResources = parentResources[:len(parentResources)-1]
	resource += createAccessClientAppConfig(parentResources)
	resource += fmt.Sprintf(`
			resource "appdynamicscloud_access_client_app" "test" {

							display_name = "%v"

							description = "%v"

							auth_type = "%v"

							rotate_secret = "%v"

							revoke_previous_secret_in = "%v"
							
							revoked_all_previous_at = "%v"
		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		value)
	return resource
}

func generateStepForUpdatedAttrAccessClientApp(rName string, resourceName string, accessClientApp_default, accessClientApp_updated *applicationprincipalmanagement.ServiceClientResponse) []resource.TestStep {
	testSteps := make([]resource.TestStep, 0, 1)
	var valid []interface{}
	valid = searchInObject(resourceAccessClientAppTest, "rotate_secret.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccAccessClientAppUpdatedAttrRotateSecret(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
				resource.TestCheckResourceAttr(resourceName, "rotate_secret", v),
				testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
			),
		})
	}
	valid = searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccAccessClientAppUpdatedAttrRevokePreviousSecretIn(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
				resource.TestCheckResourceAttr(resourceName, "revoke_previous_secret_in", v),
				testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
			),
		})
	}
	valid = searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid").([]interface{})
	for _, value := range valid {
		v := fmt.Sprintf("%v", value)
		testSteps = append(testSteps, resource.TestStep{
			Config: CreateAccAccessClientAppUpdatedAttrRevokedAllPreviousAt(rName, value),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckAppdynamicscloudAccessClientAppExists(resourceName, accessClientApp_updated),
				resource.TestCheckResourceAttr(resourceName, "revoked_all_previous_at", v),
				testAccCheckAppdynamicscloudAccessClientAppIdEqual(accessClientApp_default, accessClientApp_updated),
			),
		})
	}
	return testSteps
}

func generateNegativeStepsAccessClientApp(rName string, resourceName string) []resource.TestStep {
	//Use Update Config Function with false value
	testSteps := make([]resource.TestStep, 0, 1)
	//lint:ignore S1021 searchInObject returns interface, we need slice of interface
	var invalid []interface{}
	invalid = searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.invalid").([]interface{})
	for _, value := range invalid {
		testSteps = append(testSteps, resource.TestStep{
			Config:      CreateAccAccessClientAppUpdatedAttrRevokePreviousSecretIn(rName, value),
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

							rotate_secret = "%v"

							revoke_previous_secret_in = "%v"

							revoked_all_previous_at = "%v"
			}
		`, i, val,
			searchInObject(resourceAccessClientAppTest, "description.valid.0"),
			searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
			searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
			searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
			searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
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


						rotate_secret = "%v"


						revoke_previous_secret_in = "%v"


						revoked_all_previous_at = "%v"

		}
	`, searchInObject(resourceAccessClientAppTest, "display_name.valid.0"),
		searchInObject(resourceAccessClientAppTest, "description.valid.0"),
		searchInObject(resourceAccessClientAppTest, "auth_type.valid.0"),
		searchInObject(resourceAccessClientAppTest, "rotate_secret.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoke_previous_secret_in.valid.0"),
		searchInObject(resourceAccessClientAppTest, "revoked_all_previous_at.valid.0"))
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
