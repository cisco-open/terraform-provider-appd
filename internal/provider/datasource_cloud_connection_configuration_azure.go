package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCloudConnectionConfigurationAzure() *schema.Resource {
	dsSchema := resourceSchemaToDataSourceSchema(resourceCloudConnectionConfigurationAzure().Schema)

	dsSchema["configuration_id"] = &schema.Schema{
		Type:             schema.TypeString,
		ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
		Required:         true,
	}

	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionConfigurationAzureRead,
		Schema:      dsSchema,
	}
}

func dataSourceCloudConnectionConfigurationAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Get("configuration_id").(string)

	myctx, _, apiClient := initializeCloudConnectionClient(m)

	resp, httpResp, err := apiClient.ConfigurationsApi.GetConfiguration(myctx, id).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}
	d.SetId(resp.Id)
	flattenCloudConnectionConfigurationCommons(resp, d)
	flattenCloudConnectionConfigurationCommonsDetails(resp, d, "azure")

	return nil
}
