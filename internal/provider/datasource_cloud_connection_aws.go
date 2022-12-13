package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudConnectionAWS() *schema.Resource {
	dsSchema := resourceSchemaToDataSourceSchema(resourceCloudConnectionAWS().Schema)

	dsSchema["connection_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}

	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionAWSRead,
		Schema:      dsSchema,
	}
}

func dataSourceCloudConnectionAWSRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("connection_id").(string)

	d.SetId(id)

	return resourceCloudConnectionAWSRead(ctx, d, meta)
}
