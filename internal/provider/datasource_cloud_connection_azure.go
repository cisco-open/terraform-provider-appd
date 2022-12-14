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
	id := d.Get("connection_id").(string)

	d.SetId(id)

	return resourceCloudConnectionAzureRead(ctx, d, meta)
}
