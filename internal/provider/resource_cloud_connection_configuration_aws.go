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
		ReadContext:   resourceCloudConnectionConfigurationAWSRead,
		UpdateContext: resourceCloudConnectionConfigurationAWSUpdate,
		DeleteContext: resourceCloudConnectionConfigurationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,

		Schema: getCloudConnectionConfigurationAWSSchema(),
	}
}
func resourceCloudConnectionConfigurationAWSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	configurationId := d.Id()

	resp, _, err := apiClient.ConfigurationsApi.GetConfiguration(myctx, configurationId).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	flattenCloudConnectionConfigurationCommons(resp, d)
	flattenCloudConnectionConfigurationCommonsDetails(resp, d, "aws")

	return nil
}
func resourceCloudConnectionConfigurationAWSCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	awsConfiguration := cloudconnectionapi.AWSConfiguration{}
	awsConfiguration.BaseEntity.BaseEntityAllOf.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AWS))

	if v, ok := d.GetOk("display_name"); ok {
		awsConfiguration.BaseEntity.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		awsConfiguration.BaseEntity.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("details"); ok {
		awsConfigurationDetails := expandCloudConnectionConfigurationAWSCreateDetails(v, d)
		awsConfiguration.SetDetails(awsConfigurationDetails)
	} else {
		awsConfiguration.SetDetails(cloudconnectionapi.AWSConfigurationDetails{})
	}

	configuration := cloudconnectionapi.AWSConfigurationAsConfiguration(&awsConfiguration)

	resp, _, err := apiClient.ConfigurationsApi.CreateConfiguration(myctx).Configuration(configuration).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionConfigurationAWSRead(ctx, d, m)
}

func expandCloudConnectionConfigurationAWSCreateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.AWSConfigurationDetails {
	awsConfigurationDetails := cloudconnectionapi.AWSConfigurationDetails{}

	details, _ := singleListToMap(v)
	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)

	services := details["services"].([]interface{})
	servicesList := make([]map[string]interface{}, 0, len(services))
	for _, v := range services {
		service := v.(map[string]interface{})
		serviceMap := make(map[string]interface{})
		serviceMap["name"] = service["name"]
		if service["tag_filter"].(string) != "" {
			serviceMap["tagFilter"] = service["tag_filter"]
		}
		if polling, ok := singleListToMap(service["polling"]); ok {
			serviceMap["polling"] = polling
		}

		if importTags, ok := singleListToMap(service["import_tags"]); ok {
			tags := importTags
			tags["excludedKeys"] = tags["excluded_keys"]
			delete(tags, "excluded_keys")
			serviceMap["importTags"] = tags
		}

		servicesList = append(servicesList, serviceMap)
	}

	if len(regions) > 0 {
		awsConfigurationDetails.SetRegions(toSliceString(regions))
	}
	if polling, ok := expandCloudConnectionConfigurationDetailsPolling(details, d); ok {
		awsConfigurationDetails.SetPolling(polling)
	}
	if importTags, ok := expandCloudConnectionConfigurationDetailsImportTags(details, d); ok {
		awsConfigurationDetails.SetImportTags(importTags)
	}
	if tagFilter != "" {
		awsConfigurationDetails.SetTagFilter(tagFilter)
	}
	if len(servicesList) > 0 {
		awsConfigurationDetails.SetServices(servicesList)
	}

	return awsConfigurationDetails
}

func resourceCloudConnectionConfigurationAWSUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	myctx, _, apiClient := initializeCloudConnectionClient(m)

	awsConfigurationUpdate := cloudconnectionapi.ConfigurationUpdate{}
	if v, ok := d.GetOk("display_name"); ok {
		awsConfigurationUpdate.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		awsConfigurationUpdate.SetDescription(v.(string))
	}
	if v, ok := d.GetOk("details"); ok {
		awsConfigurationUpdateDetails := expandCloudConnectionConfigurationAWSUpdateDetails(v, d)
		awsConfigurationUpdate.SetDetails(awsConfigurationUpdateDetails)
	}
	resp, _, err := apiClient.ConfigurationsApi.UpdateConfiguration(myctx, d.Id()).ConfigurationUpdate(awsConfigurationUpdate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)
	return resourceCloudConnectionConfigurationAWSRead(ctx, d, m)
}

func expandCloudConnectionConfigurationAWSUpdateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConfigurationUpdateDetails {
	awsConfigurationDetails := cloudconnectionapi.AWSConfigurationDetails{}

	details, _ := singleListToMap(v)

	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)

	services := details["services"].([]interface{})
	servicesList := make([]map[string]interface{}, 0, len(services))
	for _, v := range services {
		service := v.(map[string]interface{})
		serviceMap := make(map[string]interface{})
		serviceMap["name"] = service["name"]
		if service["tag_filter"].(string) != "" {
			serviceMap["tagFilter"] = service["tag_filter"]
		}
		if polling, ok := singleListToMap(service["polling"]); ok {
			serviceMap["polling"] = polling
		}

		if importTags, ok := singleListToMap(service["import_tags"]); ok {
			tags := importTags
			tags["excludedKeys"] = tags["excluded_keys"]
			delete(tags, "excluded_keys")
			serviceMap["importTags"] = tags
		}

		servicesList = append(servicesList, serviceMap)
	}

	if len(regions) > 0 {
		awsConfigurationDetails.SetRegions(toSliceString(regions))
	} else {
		awsConfigurationDetails.SetRegions(make([]string, 0))
	}
	if polling, ok := expandCloudConnectionConfigurationDetailsPolling(details, d); ok {
		awsConfigurationDetails.SetPolling(polling)
	}
	if importTags, ok := expandCloudConnectionConfigurationDetailsImportTags(details, d); ok {
		awsConfigurationDetails.SetImportTags(importTags)
	} else {
		importTags := cloudconnectionapi.ImportTagConfiguration{}
		importTags.Enabled = true
		importTags.ExcludedKeys = make([]string, 0)
		awsConfigurationDetails.SetImportTags(importTags)
	}

	awsConfigurationDetails.SetTagFilter(tagFilter)

	if len(servicesList) > 0 {
		awsConfigurationDetails.SetServices(servicesList)
	}

	awsUpdateDetails := AWSConfigurationDetailsAsConfigurationUpdateDetails(&awsConfigurationDetails)
	return awsUpdateDetails
}

func AWSConfigurationDetailsAsConfigurationUpdateDetails(v *cloudconnectionapi.AWSConfigurationDetails) cloudconnectionapi.ConfigurationUpdateDetails {
	return cloudconnectionapi.ConfigurationUpdateDetails{
		AWSConfigurationDetails: v,
	}
}
