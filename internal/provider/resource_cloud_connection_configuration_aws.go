package provider

import (
	"context"

	cloudConnectionApi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
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

		CustomizeDiff: customdiff.All(serviceAtLeastOne),
		SchemaVersion: 1,

		Schema: getCloudConnectionConfigurationAWSSchema(),
	}
}

func resourceCloudConnectionConfigurationAWSCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myCtx, _, apiClient := initializeCloudConnectionClient(m)

	awsConfiguration := cloudConnectionApi.AWSConfiguration{}
	awsConfiguration.BaseEntity.BaseEntityAllOf.SetType(cloudConnectionApi.ProviderType(cloudConnectionApi.AWS))

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
		//Create Default Details with default services.
		awsConfiguration.SetDetails(cloudConnectionApi.AWSConfigurationDetails{})
	}

	configuration := cloudConnectionApi.AWSConfigurationAsConfiguration(&awsConfiguration)

	resp, httpResp, err := apiClient.ConfigurationsApi.CreateConfiguration(myCtx).Configuration(configuration).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionConfigurationAWSRead(ctx, d, m)
}

func resourceCloudConnectionConfigurationAWSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myCtx, _, apiClient := initializeCloudConnectionClient(m)

	configurationId := d.Id()

	resp, httpResp, err := apiClient.ConfigurationsApi.GetConfiguration(myCtx, configurationId).Execute()
	if err != nil {
		if httpResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return errRespToDiag(err, httpResp)
	}

	flattenCloudConnectionConfigurationCommons(resp, d)
	flattenCloudConnectionConfigurationCommonsDetails(resp, d, "aws")

	return nil
}

func resourceCloudConnectionConfigurationAWSUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myCtx, _, apiClient := initializeCloudConnectionClient(m)

	AWSConfigurationUpdate := cloudConnectionApi.ConfigurationUpdate{}

	if v, ok := d.GetOk("display_name"); ok {
		AWSConfigurationUpdate.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		AWSConfigurationUpdate.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("details"); ok {
		AWSConfigurationUpdateDetails:= expandCloudConnectionConfigurationAWSUpdateDetails(v, d)
		AWSConfigurationUpdate.SetDetails(AWSConfigurationUpdateDetails)
	}

	resp, httpResp, err := apiClient.ConfigurationsApi.UpdateConfiguration(myCtx, d.Id()).ConfigurationUpdate(AWSConfigurationUpdate).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionConfigurationAWSRead(ctx, d, m)
}

func expandCloudConnectionConfigurationAWSCreateDetails(v interface{}, d *schema.ResourceData) cloudConnectionApi.AWSConfigurationDetails {
	awsConfigurationDetails := cloudConnectionApi.AWSConfigurationDetails{}

	details, _ := singleListToMap(v)
	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)

	services := details["services"].(*schema.Set).List()
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

	if len(servicesList) > 0 {
		awsConfigurationDetails.SetServices(servicesList)
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

	return awsConfigurationDetails
}

func expandCloudConnectionConfigurationAWSUpdateDetails(v interface{}, d *schema.ResourceData) cloudConnectionApi.ConfigurationUpdateDetails {
	awsConfigurationDetails := cloudConnectionApi.AWSConfigurationDetails{}

	details, _ := singleListToMap(v)
	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)

	services := details["services"].(*schema.Set).List()
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

	if len(servicesList) > 0 {
		awsConfigurationDetails.SetServices(servicesList)
	}

	if polling, ok := expandCloudConnectionConfigurationDetailsPolling(details, d); ok {
		awsConfigurationDetails.SetPolling(polling)
	}

	if importTags, ok := expandCloudConnectionConfigurationDetailsImportTags(details, d); ok {
		awsConfigurationDetails.SetImportTags(importTags)
	} else {
		importTags := cloudConnectionApi.ImportTagConfiguration{}
		importTags.Enabled = true
		importTags.ExcludedKeys = make([]string, 0)
		awsConfigurationDetails.SetImportTags(importTags)
	}

	awsConfigurationDetails.SetTagFilter(tagFilter)

	awsUpdateDetails := awsConfigurationDetailsAsConfigurationUpdateDetails(&awsConfigurationDetails)
	return awsUpdateDetails
}

func awsConfigurationDetailsAsConfigurationUpdateDetails(v *cloudConnectionApi.AWSConfigurationDetails) cloudConnectionApi.ConfigurationUpdateDetails {
	return cloudConnectionApi.ConfigurationUpdateDetails{
		AWSConfigurationDetails: v,
	}
}
