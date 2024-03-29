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
	"strings"

	cloudconnectionapi "github.com/cisco-open/appd-cloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()
	httpResp, err := apiClient.ConnectionsApi.DeleteConnection(myctx, connectionId).Execute()

	if err != nil && httpResp.StatusCode != 404 {
		return errRespToDiag(err, httpResp)
	}
	d.SetId("")
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

func flattenCloudConnectionConfigurationCommonsDetails(resp *cloudconnectionapi.ConfigurationDetail, d *schema.ResourceData, connectionType string) {
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
	if strings.Contains(connectionType, "azure") {
		detailsMap["resource_groups"] = toSliceInterface(resp.GetDetails().ResourceGroups)
	}

	detailsList := make([]interface{}, 0, 1)
	detailsList = append(detailsList, detailsMap)
	d.Set("configuration_details", detailsList)
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
		return cloudconnectionapi.ScheduleInterval{}, false
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

func checkRequiredNotRequired(d *schema.ResourceDiff, type_ string) error {
	requiredAttributes := map[string][]string{
		string(cloudconnectionapi.ACCESS_KEY):      {"access_key_id", "secret_access_key"},
		string(cloudconnectionapi.ROLE_DELEGATION): {"account_id"},
	}

	notRequiredAttributes := map[string][]string{
		string(cloudconnectionapi.ACCESS_KEY):      {"account_id"},
		string(cloudconnectionapi.ROLE_DELEGATION): {"access_key_id", "secret_access_key"},
	}

	details, _ := singleListToMap(d.Get("connection_details"))

	for _, k := range requiredAttributes[type_] {
		if details[k] == "" {
			return fmt.Errorf("%s is required with %s", k, type_)
		}
	}

	for _, k := range notRequiredAttributes[type_] {
		if v := details[k]; v != "" {
			return fmt.Errorf("%s should not be used with %s", k, type_)
		}
	}

	return nil
}
