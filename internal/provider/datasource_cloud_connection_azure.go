package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCloudConnectionAzure() *schema.Resource {
	dsSchema := resourceSchemaToDataSourceSchema(resourceCloudConnectionAzure().Schema)

	dsSchema["connection_id"] = &schema.Schema{
		Type:             schema.TypeString,
		ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
		Required:         true,
	}

	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionAzureRead,
		Schema:      dsSchema,
	}
}

func dataSourceCloudConnectionAzureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	connectionId := d.Get("connection_id").(string)

	myctx, _, apiClient := initializeCloudConnectionClient(meta)

	resp, httpResp, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	d.Set("details", flattenCloudConnectionAzureDetails(resp, d))

	flattenCloudConnectionCommons(resp, d)

	return nil
}
