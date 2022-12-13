package provider

import (
	"context"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()
	_, httpResp, err := apiClient.ConnectionsApi.DeleteConnection(myctx, connectionId).Execute()

	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	return nil
}

func flattenCloudConnectionCommons(resp *cloudconnectionapi.ConnectionResponse, d *schema.ResourceData) {
	d.SetId(resp.GetId())
	d.Set("display_name", resp.GetDisplayName())
	d.Set("description", resp.GetDescription())
	d.Set("state", resp.GetState())
	d.Set("state_message", resp.StateMessage)
	d.Set("configuration_id", resp.GetConfigurationId())
	d.Set("created_at", utcTimeToString(resp.GetCreatedAt()))
	d.Set("updated_at", utcTimeToString(resp.GetUpdatedAt()))
}
