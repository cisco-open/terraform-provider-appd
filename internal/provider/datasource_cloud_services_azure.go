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

	// TODO: Update url
	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudServicesAzure() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudServicesAzureRead,
		Schema: map[string]*schema.Schema{
			"services_azure": {
				Type:        schema.TypeList,
				Description: "All supported services for Azure provider.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "Unique Identifier",
							Computed:    true,
						},
						"display_name": {
							Type:        schema.TypeString,
							Description: "Display Name",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudServicesAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myCtx, _, apiClient := initializeCloudConnectionClient(m)
	resp, httpResp, err := apiClient.ResourcesApi.GetServices(myCtx).Type_(cloudconnectionapi.AZURE).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}
	services := resp.Items
	servicesAzureList := make([]interface{}, 0)
	for _, service := range services {
		serviceMap := make(map[string]interface{})
		serviceMap["id"] = service.Id
		serviceMap["display_name"] = service.DisplayName
		serviceMap["description"] = service.Description

		servicesAzureList = append(servicesAzureList, serviceMap)
	}

	d.SetId("azure")
	d.Set("services_azure", servicesAzureList)
	return nil
}
