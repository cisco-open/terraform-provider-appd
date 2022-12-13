package provider

import (
	"context"
	"strings"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()
	httpResp, err := apiClient.ConnectionsApi.DeleteConnection(myctx, connectionId).Execute()

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

func resourceCloudConnectionConfigurationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	configurationId := d.Id()
	_, err := apiClient.ConfigurationsApi.DeleteConfiguration(myctx, configurationId).Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
func flattenCloudConnectionConfigurationCommons(resp *cloudconnectionapi.ConfigurationDetail, d *schema.ResourceData) {
	d.SetId(resp.GetId())
	d.Set("display_name", resp.GetDisplayName())
	d.Set("description", resp.GetDescription())
	d.Set("created_at", utcTimeToString(resp.GetCreatedAt()))
	d.Set("updated_at", utcTimeToString(resp.GetUpdatedAt()))
}

func flattenCloudConnectionConfigurationCommonsDetails(resp *cloudconnectionapi.ConfigurationDetail, d *schema.ResourceData,connectionType string) {
	detailsMap := make(map[string]interface{})

	var servicesList []interface{}

	if resp.Details.HasPolling() {

		polling := map[string]interface{}{
			"interval": int(resp.GetDetails().Polling.Interval),
			"unit":     resp.GetDetails().Polling.Unit,
		}
		pollingList := []interface{}{}
		pollingList = append(pollingList, polling)

		detailsMap["polling"] = pollingList
	}

	if resp.Details.HasImportTags() {

		importTags := map[string]interface{}{
			"enabled": resp.GetDetails().ImportTags.Enabled,
		}
		excludedKeys := toSliceInterface(resp.GetDetails().ImportTags.ExcludedKeys)
		if len(excludedKeys) > 0 {
			importTags["excluded_keys"] = excludedKeys
		}

		importTagsList := []interface{}{}
		importTagsList = append(importTagsList, importTags)

		detailsMap["import_tags"] = importTagsList
	}

	services := resp.GetDetails().Services.([]interface{})
	servicesList = make([]interface{}, 0, len(services))
	for _, v := range services {
		service := v.(map[string]interface{})
		serviceMap := make(map[string]interface{})
		serviceMap["name"] = service["name"].(string)

		if _, ok := service["polling"]; ok {

			polling := map[string]interface{}{
				"interval": int((service["polling"].(map[string]interface{})["interval"]).(float64)),
				"unit":     service["polling"].(map[string]interface{})["unit"],
			}
			pollingList := []interface{}{}
			pollingList = append(pollingList, polling)
			serviceMap["polling"] = pollingList
		}

		if _, ok := service["importTags"]; ok {

			importTags := map[string]interface{}{
				"enabled": service["importTags"].(map[string]interface{})["enabled"],
			}

			excludedKeys := service["importTags"].(map[string]interface{})["excludedKeys"].([]interface{})
			if len(excludedKeys) > 0 {
				importTags["excluded_keys"] = excludedKeys
			}

			importTagsList := []interface{}{}
			importTagsList = append(importTagsList, importTags)
			serviceMap["import_tags"] = importTagsList
		}

		if _, ok := service["tagFilter"]; ok {
			serviceMap["tag_filter"] = service["tagFilter"].(string)
		}

		servicesList = append(servicesList, serviceMap)
	}
	if len(servicesList) > 0 {
		detailsMap["services"] = servicesList
	}

	detailsMap["regions"] = resp.GetDetails().Regions
	detailsMap["tag_filter"] = *resp.GetDetails().TagFilter
	if strings.Contains(connectionType,"azure"){
		detailsMap["resource_groups"] = toSliceInterface(resp.GetDetails().ResourceGroups)
	}

	detailsList := make([]interface{}, 0, 1)
	detailsList = append(detailsList, detailsMap)
	d.Set("details", detailsList)
}

func expandCloudConnectionConfigurationDetailsPolling(v interface{}, d *schema.ResourceData) (cloudconnectionapi.ScheduleInterval, bool) {
	details := v.(map[string]interface{})
	var (
		interval int32
		unit     string
	)

	polling := details["polling"].([]interface{})
	if len(polling) > 0 {
		scheduleData := polling[0].(map[string]interface{})
		interval = int32(scheduleData["interval"].(int))
		unit = scheduleData["unit"].(string)
		return *cloudconnectionapi.NewScheduleInterval(interval, unit), true
	} else {
		return *&cloudconnectionapi.ScheduleInterval{}, false
	}

}

func expandCloudConnectionConfigurationDetailsImportTags(v interface{}, d *schema.ResourceData) (cloudconnectionapi.ImportTagConfiguration, bool) {
	details := v.(map[string]interface{})
	var (
		enabled      bool
		excludedKeys []string
	)

	importTags := details["import_tags"].([]interface{})
	if len(importTags) > 0 {

		tags := importTags[0].(map[string]interface{})
		enabled = tags["enabled"].(bool)
		excludedKeys = toSliceString(tags["excluded_keys"].([]interface{}))
		tagConfiguration := cloudconnectionapi.ImportTagConfiguration{}
		tagConfiguration.SetEnabled(enabled)
		tagConfiguration.SetExcludedKeys(excludedKeys)

		return tagConfiguration, true
	} else {
		return cloudconnectionapi.ImportTagConfiguration{}, false
	}

}

