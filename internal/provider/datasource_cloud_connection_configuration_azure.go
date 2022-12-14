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

func dataSourceCloudConnectionConfigurationAzureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("configuration_id").(string)

	d.SetId(id)

	return resourceCloudConnectionConfigurationAzureRead(ctx, d, meta)
}
