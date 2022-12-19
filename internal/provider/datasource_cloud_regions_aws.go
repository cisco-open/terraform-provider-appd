package provider

import (
	"context"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudRegionsAWS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudRegionsAWSRead,
		Schema: map[string]*schema.Schema{
			"regions_aws": {
				Type:        schema.TypeList,
				Description: "AWS Regions",
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

func dataSourceCloudRegionsAWSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myCtx, _, apiClient := initializeCloudConnectionClient(m)
	resp, httpResp, err := apiClient.ResourcesApi.GetRegions(myCtx).Type_(cloudconnectionapi.AWS).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}
	regions := resp.Items
	regionsAWSList := make([]interface{}, 0)
	for _, region := range regions {
		regionMap := make(map[string]interface{})
		regionMap["id"] = region.Id
		regionMap["display_name"] = region.DisplayName
		regionMap["description"] = region.Description

		regionsAWSList = append(regionsAWSList, regionMap)
	}

	d.SetId("aws")
	d.Set("regions_aws", regionsAWSList)
	return nil
}
