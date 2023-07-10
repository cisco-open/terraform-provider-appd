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
	"os"
	"strconv"
	"strings"
	"testing"

	client "github.com/cisco-open/appd-cloud-go-client"
	cloudconnectionapi "github.com/cisco-open/appd-cloud-go-client/apis/v1/cloudconnections"
	"github.com/cisco-open/terraform-provider-appd/internal/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testPrefix = "TestAcc"

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("appdynamicscloud_connection_azure",
		&resource.Sweeper{
			Name: "appdynamicscloud_connection_azure",
			F:    connectionAzureSweeper,
		})
	resource.AddTestSweepers("appdynamicscloud_access_client_app",
		&resource.Sweeper{
			Name: "appdynamicscloud_access_client_app",
			F:    accessClientAppSweeper,
		})
	resource.AddTestSweepers("appdynamicscloud_connection_aws",
		&resource.Sweeper{
			Name: "appdynamicscloud_connection_aws",
			F:    connectionAWSSweeper,
		})
}

func connectionAzureSweeper(_ string) error {
	myctx, _, apiClient := initializeCloudConnectionClient(sharedClient())
	connections, _, err := apiClient.ConnectionsApi.GetConnections(myctx).Execute()
	if err != nil {
		return err
	}
	for _, connection := range connections.GetItems() {
		if (connection.Type == cloudconnectionapi.AZURE) && (strings.Contains(*connection.Description, testPrefix)) {
			httpResp, err := apiClient.ConnectionsApi.DeleteConnection(myctx, connection.Id).Execute()
			if err != nil && httpResp.StatusCode != 404 {
				return err
			}
		}
	}
	return nil
}
func accessClientAppSweeper(_ string) error {
	myctx, _, apiClient := initializeApplicationPrincipalManagementClient(sharedClient())
	servicePrincipals, _, err := apiClient.ServicesApi.ListAllServiceClients(myctx).Execute()
	if err != nil {
		return err
	}
	for _, servicePrincipal := range servicePrincipals.GetItems() {
		if strings.Contains(*servicePrincipal.Description, testPrefix) {
			httpResp, err := apiClient.ServicesApi.DeleteServiceClient(myctx, *servicePrincipal.Id).Execute()
			if err != nil && httpResp.StatusCode != 404 {
				return err
			}
		}
	}
	return nil
}
func connectionAWSSweeper(_ string) error {
	myctx, _, apiClient := initializeCloudConnectionClient(sharedClient())
	connections, _, err := apiClient.ConnectionsApi.GetConnections(myctx).Execute()
	if err != nil {
		return err
	}
	for _, connection := range connections.GetItems() {
		if (connection.Type == cloudconnectionapi.AWS) && (strings.Contains(*connection.Description, testPrefix)) {
			httpResp, err := apiClient.ConnectionsApi.DeleteConnection(myctx, connection.Id).Execute()
			if err != nil && httpResp.StatusCode != 404 {
				return err
			}
		}
	}
	return nil
}
func sharedClient() interface{} {
	client_id := os.Getenv("APPDYNAMICS_CLIENT_ID")
	client_secret := os.Getenv("APPDYNAMICS_CLIENT_SECRET")
	tenant_name := os.Getenv("APPDYNAMICS_TENANT_NAME")
	username := os.Getenv("APPDYNAMICS_USERNAME")
	password := os.Getenv("APPDYNAMICS_PASSWORD")
	login_mode := os.Getenv("APPDYNAMICS_LOGIN_MODE")
	save_token, _ := strconv.ParseBool(os.Getenv("APPDYNAMICS_SAVE_TOKEN"))
	tenantId, _ := lookupTenantId(tenant_name)

	loginCtx := context.WithValue(context.Background(), auth.MODE, login_mode)
	if login_mode == "headless" {
		loginCtx = context.WithValue(loginCtx, auth.USERNAME, username)
		loginCtx = context.WithValue(loginCtx, auth.PASSWORD, password)
	} else if login_mode == "service_principal" {
		loginCtx = context.WithValue(loginCtx, auth.CLIENT_ID, client_id)
		loginCtx = context.WithValue(loginCtx, auth.CLIENT_SECRET, client_secret)
	}

	token, _ := auth.Login(tenant_name, tenantId, save_token, loginCtx)

	// CONFIGURE API CLIENT
	configuration := client.NewConfiguration()
	configuration.Debug = true

	myctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
		"tenant-name": tenant_name,
	})
	myctx = context.WithValue(myctx, client.ContextAccessToken, token)

	// TERRAFORM CONFIG
	config := config{configuration: configuration, ctx: myctx}

	return config
}
