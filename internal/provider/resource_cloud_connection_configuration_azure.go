package provider

import (
	"context"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudConnectionConfigurationAzure() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudConnectionConfigurationAzureCreate,
		ReadContext:   resourceCloudConnectionConfigurationRead,
		UpdateContext: resourceCloudConnectionConfigurationAzureUpdate,
		DeleteContext: resourceCloudConnectionConfigurationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,

		Schema: getCloudConnectionConfigurationAzureSchema(),
	}
}

func resourceCloudConnectionConfigurationAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	configuration := cloudconnectionapi.Configuration{}
	configuration.AzureConfiguration.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AZURE))

	if v, ok := d.GetOk("display_name"); ok {
		configuration.AzureConfiguration.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		configuration.AzureConfiguration.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("details"); ok {
		configurationDetails := expandCloudConnectionConfigurationAzureDetails(v, d)
		configuration.AzureConfiguration.SetDetails(configurationDetails)
	}

	resp, _, err := apiClient.ConfigurationsApi.CreateConfiguration(myctx).Configuration(configuration).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionConfigurationRead(ctx, d, m)
}

func resourceCloudConnectionConfigurationAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO: Later
	return nil
}

func expandCloudConnectionConfigurationAzureDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.AzureConfigurationDetails {
	azureConfigurationDetails := cloudconnectionapi.AzureConfigurationDetails{}

	details := v.(*schema.Set).List()[0].(map[string]interface{})

	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)
	// TODO: Make sure this works 👇
	services := details["services"].([]interface{})
	resourceGroups := details["resource_groups"].([]string)

	azureConfigurationDetails.SetRegions(regions)
	azureConfigurationDetails.SetPolling(expandCloudConnectionConfigurationDetailsPolling(details, d))
	azureConfigurationDetails.SetImportTags(expandCloudConnectionConfigurationDetailsImportTags(details, d))
	azureConfigurationDetails.SetTagFilter(tagFilter)
	azureConfigurationDetails.SetServices(services)
	azureConfigurationDetails.SetResourceGroups(resourceGroups)

	return azureConfigurationDetails
}
