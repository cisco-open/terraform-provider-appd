package provider

import (
	"context"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudConnectionConfigurationAWS() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudConnectionConfigurationAWSCreate,
		ReadContext:   resourceCloudConnectionConfigurationRead,
		UpdateContext: resourceCloudConnectionConfigurationAWSUpdate,
		DeleteContext: resourceCloudConnectionConfigurationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,

		Schema: getCloudConnectionConfigurationAWSSchema(),
	}
}

func resourceCloudConnectionConfigurationAWSCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(meta)

	configuration := cloudconnectionapi.Configuration{}
	configuration.AWSConfiguration.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AWS))

	if v, ok := d.GetOk("display_name"); ok {
		configuration.AWSConfiguration.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		configuration.AWSConfiguration.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("details"); ok {
		configurationDetails := expandCloudConnectionConfigurationAWSDetails(v, d)
		configuration.AWSConfiguration.SetDetails(configurationDetails)
	}

	resp, _, err := apiClient.ConfigurationsApi.CreateConfiguration(myctx).Configuration(configuration).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionConfigurationRead(ctx, d, meta)
}

func resourceCloudConnectionConfigurationAWSUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func expandCloudConnectionConfigurationAWSDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.AWSConfigurationDetails {
	AWSConfigurationDetails := cloudconnectionapi.AWSConfigurationDetails{}

	details := v.(*schema.Set).List()[0].(map[string]interface{})

	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)
	// TODO: Make sure this works 👇
	services := details["services"].([]interface{})

	AWSConfigurationDetails.SetRegions(regions)
	AWSConfigurationDetails.SetPolling(expandCloudConnectionConfigurationDetailsPolling(details, d))
	AWSConfigurationDetails.SetImportTags(expandCloudConnectionConfigurationDetailsImportTags(details, d))
	AWSConfigurationDetails.SetTagFilter(tagFilter)
	AWSConfigurationDetails.SetServices(services)

	return AWSConfigurationDetails
}
