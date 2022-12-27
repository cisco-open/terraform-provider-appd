package provider

import (
	"context"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudRegionsAzure() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudRegionsAzureRead,
		Schema: map[string]*schema.Schema{
			"regions_azure": {
				Type:        schema.TypeList,
				Description: "All supported hosting regions for the Azure cloud provider",
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

func dataSourceCloudRegionsAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myCtx, _, apiClient := initializeCloudConnectionClient(m)
	resp, httpResp, err := apiClient.ResourcesApi.GetRegions(myCtx).Type_(cloudconnectionapi.AZURE).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}
	regions := resp.Items
	regionsAzureList := make([]interface{}, 0)
	for _, region := range regions {
		regionMap := make(map[string]interface{})
		regionMap["id"] = region.Id
		regionMap["display_name"] = region.DisplayName
		regionMap["description"] = region.Description

		regionsAzureList = append(regionsAzureList, regionMap)
	}

	d.SetId("azure")
	d.Set("regions_azure", regionsAzureList)
	return nil
}
