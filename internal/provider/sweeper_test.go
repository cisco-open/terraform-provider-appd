package provider

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	d := schema.ResourceData{}
	d.Set("client_id", os.Getenv("APPDYNAMICS_CLIENT_ID"))
	d.Set("client_secret", os.Getenv("APPDYNAMICS_CLIENT_SECRET"))
	d.Set("tenant_name", os.Getenv("APPDYNAMICS_TENANT_NAME"))
	d.Set("username", os.Getenv("APPDYNAMICS_USERNAME"))
	d.Set("password", os.Getenv("APPDYNAMICS_PASSWORD"))
	d.Set("login_mode", os.Getenv("APPDYNAMICS_LOGIN_MODE"))
	d.Set("save_token", os.Getenv("APPDYNAMICS_SAVE_TOKEN"))

	config, diags := configureClient(context.Background(), &d)
	for _, diag := range diags {
		fmt.Errorf("%v %v %v", diag.Detail, diag.Severity, diag.Summary)
		return nil
	}
	return config
}
