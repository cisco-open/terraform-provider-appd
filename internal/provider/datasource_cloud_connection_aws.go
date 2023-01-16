package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCloudConnectionAWS() *schema.Resource {
	dsSchema := resourceSchemaToDataSourceSchema(resourceCloudConnectionAWS().Schema)

	dsSchema["connection_id"] = &schema.Schema{
		Type:             schema.TypeString,
		Description:      "The Connection ID of the AWS Connection",
		Required:         true,
		ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
	}

	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionAWSRead,
		Schema:      dsSchema,
	}
}

func dataSourceCloudConnectionAWSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connectionId := d.Get("connection_id").(string)

	myctx, _, apiClient := initializeCloudConnectionClient(m)

	respConnection, httpRespConnection, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		return errRespToDiag(err, httpRespConnection)
	}
	d.SetId(respConnection.Id)

	d.Set("connection_details", flattenCloudConnectionAWSDetails(respConnection, d))

	flattenCloudConnectionCommons(respConnection, d)

	configurationId := respConnection.ConfigurationId
	d.Set("configuration_id", configurationId)
	respConfiguration, httpRespConfiguration, err := apiClient.ConfigurationsApi.GetConfiguration(myctx, *configurationId).Execute()
	if err != nil {
		return errRespToDiag(err, httpRespConfiguration)
	}
	flattenCloudConnectionConfigurationCommonsDetails(respConfiguration, d, "AWS")

	return nil
}
