package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudConnectionConfigurationAzure() *schema.Resource {
	dsSchema := resourceSchemaToDataSourceSchema(resourceCloudConnectionConfigurationAzure().Schema)

	dsSchema["configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
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
