package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCloudConnectionConfigurationAWS() *schema.Resource {
	dsSchema := resourceSchemaToDataSourceSchema(resourceCloudConnectionConfigurationAWS().Schema)

	dsSchema["configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
		Required: true,
	}

	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionConfigurationAWSRead,
		Schema:      dsSchema,
	}
}

func dataSourceCloudConnectionConfigurationAWSRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("configuration_id").(string)

	d.SetId(id)

	return resourceCloudConnectionConfigurationAWSRead(ctx, d, meta)
}
