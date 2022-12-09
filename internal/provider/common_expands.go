package provider

import (
	"context"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandCloudConnectionConfigurationDetailsPolling(v interface{}, d *schema.ResourceData) cloudconnectionapi.ScheduleInterval {
	details := v.(*schema.Set).List()[0].(map[string]interface{})

	scheduleData := details["polling"].(map[string]interface{})

	interval := scheduleData["interval"].(int32)
	unit := scheduleData["unit"].(string)

	return *cloudconnectionapi.NewScheduleInterval(interval, unit)
}

func expandCloudConnectionConfigurationDetailsImportTags(v interface{}, d *schema.ResourceData) cloudconnectionapi.ImportTagConfiguration {
	details := v.(*schema.Set).List()[0].(map[string]interface{})

	importTagsData := details["import_tags"].(map[string]interface{})

	enabled := importTagsData["enabled"].(bool)
	excludedKeys := importTagsData["excluded_keys"].([]string)

	tagConfiguration := cloudconnectionapi.NewImportTagConfiguration(enabled)

	tagConfiguration.SetExcludedKeys(excludedKeys)

	return *tagConfiguration
}

func resourceCloudConnectionConfigurationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	configurationId := d.Id()

	resp, _, err := apiClient.ConfigurationsApi.GetConfiguration(myctx, configurationId).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.GetId())
	d.Set("display_name", resp.GetDisplayName())
	d.Set("description", resp.GetDescription())
	// TODO: Later
	// d.Set("details", resp.GetDetails())
	d.Set("created_at", resp.GetCreatedAt())
	d.Set("updated_at", resp.GetUpdatedAt())

	return nil
}

func resourceCloudConnectionConfigurationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	configurationId := d.Id()
	_, err := apiClient.ConfigurationsApi.DeleteConfiguration(myctx, configurationId).Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceCloudConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()
	_, _, err := apiClient.ConnectionsApi.DeleteConnection(myctx, connectionId).Execute()

	if err != nil {
		return diag.FromErr(err)
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

	return
}
