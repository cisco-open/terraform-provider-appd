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
		ReadContext:   resourceCloudConnectionConfigurationAzureRead,
		UpdateContext: resourceCloudConnectionConfigurationAzureUpdate,
		DeleteContext: resourceCloudConnectionConfigurationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,

		Schema: getCloudConnectionConfigurationAzureSchema(),
	}
}
func resourceCloudConnectionConfigurationAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	configurationId := d.Id()

	resp, _, err := apiClient.ConfigurationsApi.GetConfiguration(myctx, configurationId).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.GetId())
	d.Set("display_name", resp.GetDisplayName())
	d.Set("description", resp.GetDescription())

	pollingSet := schema.NewSet(schema.HashResource(cloudConnectionConfigurationDetails()["polling"].Elem.(*schema.Resource)), []interface{}{})
	pollingSet.Add(map[string]interface{}{
		"interval": int(resp.GetDetails().Polling.Interval),
		"unit":     resp.GetDetails().Polling.Unit,
	})
	
	importTagsSet := schema.NewSet(schema.HashResource(cloudConnectionConfigurationDetails()["import_tags"].Elem.(*schema.Resource)), []interface{}{})
	importTagsSet.Add(map[string]interface{}{
		"enabled":       resp.GetDetails().ImportTags.Enabled,
		"excluded_keys": toSliceInterface(resp.GetDetails().ImportTags.ExcludedKeys),
	})
	
	detailsSet := schema.NewSet(schema.HashResource(resourceCloudConnectionConfigurationAzure().Schema["details"].Elem.(*schema.Resource)), []interface{}{})
	detailsSet.Add(map[string]interface{}{
		"regions":     resp.GetDetails().Regions,
		"polling":     pollingSet,
		"import_tags": importTagsSet,
		"tag_filter":  *resp.GetDetails().TagFilter,
		// "services":        resp.GetDetails().Services,
		"resource_groups": toSliceInterface(resp.GetDetails().ResourceGroups),
	})

	d.Set("details", detailsSet)
	// d.Set("details", resp.GetDetails())
	d.Set("created_at", resp.GetCreatedAt())
	d.Set("updated_at", resp.GetUpdatedAt())

	return nil
}
func resourceCloudConnectionConfigurationAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	azureConfiguration := cloudconnectionapi.AzureConfiguration{}
	azureConfiguration.BaseEntity.BaseEntityAllOf.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AZURE))

	if v, ok := d.GetOk("display_name"); ok {
		azureConfiguration.BaseEntity.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		azureConfiguration.BaseEntity.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("details"); ok {
		azureConfigurationDetails := expandCloudConnectionConfigurationAzureDetails(v, d)
		azureConfiguration.SetDetails(azureConfigurationDetails)
	}

	configuration := cloudconnectionapi.AzureConfigurationAsConfiguration(&azureConfiguration)

	resp, _, err := apiClient.ConfigurationsApi.CreateConfiguration(myctx).Configuration(configuration).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionConfigurationAzureRead(ctx, d, m)
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
	// services := details["services"].(interface{})
	resourceGroups := details["resource_groups"].([]interface{})

	azureConfigurationDetails.SetRegions(toSliceString(regions))
	azureConfigurationDetails.SetPolling(expandCloudConnectionConfigurationDetailsPolling(details, d))
	azureConfigurationDetails.SetImportTags(expandCloudConnectionConfigurationDetailsImportTags(details, d))
	azureConfigurationDetails.SetTagFilter(tagFilter)
	// azureConfigurationDetails.SetServices(services)
	azureConfigurationDetails.SetResourceGroups(toSliceString(resourceGroups))

	return azureConfigurationDetails
}
