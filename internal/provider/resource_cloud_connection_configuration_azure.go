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

	flattenCloudConnectionConfigurationCommons(resp, d)
	flattenCloudConnectionConfigurationCommonsDetails(resp, d, "azure")

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
		azureConfigurationDetails := expandCloudConnectionConfigurationAzureCreateDetails(v, d)
		azureConfiguration.SetDetails(azureConfigurationDetails)
	} else {
		azureConfiguration.SetDetails(cloudconnectionapi.AzureConfigurationDetails{})
	}

	configuration := cloudconnectionapi.AzureConfigurationAsConfiguration(&azureConfiguration)

	resp, _, err := apiClient.ConfigurationsApi.CreateConfiguration(myctx).Configuration(configuration).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionConfigurationAzureRead(ctx, d, m)
}

func expandCloudConnectionConfigurationAzureCreateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.AzureConfigurationDetails {
	azureConfigurationDetails := cloudconnectionapi.AzureConfigurationDetails{}

	details,_:=singleListToMap(v)
	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)
	resourceGroups := details["resource_groups"].([]interface{})

	services := details["services"].([]interface{})
	servicesList := make([]map[string]interface{}, 0, len(services))
	for _, v := range services {
		service := v.(map[string]interface{})
		serviceMap := make(map[string]interface{})
		serviceMap["name"] = service["name"]
		if service["tag_filter"].(string) != "" {
			serviceMap["tagFilter"] = service["tag_filter"]
		}
		if polling,ok:=singleListToMap(service["polling"]);ok{
			serviceMap["polling"] = polling
		}
	
		if importTags,ok:=singleListToMap(service["import_tags"]);ok{
			tags := importTags
			tags["excludedKeys"] = tags["excluded_keys"]
			delete(tags, "excluded_keys")
			serviceMap["importTags"] = tags
		}
	
		servicesList = append(servicesList, serviceMap)
	}

	if len(regions) > 0 {
		azureConfigurationDetails.SetRegions(toSliceString(regions))
	}
	if polling, ok := expandCloudConnectionConfigurationDetailsPolling(details, d); ok {
		azureConfigurationDetails.SetPolling(polling)
	}
	if importTags, ok := expandCloudConnectionConfigurationDetailsImportTags(details, d); ok {
		azureConfigurationDetails.SetImportTags(importTags)
	}
	if tagFilter != "" {
		azureConfigurationDetails.SetTagFilter(tagFilter)
	}
	if len(servicesList) > 0 {
		azureConfigurationDetails.SetServices(servicesList)
	}
	if len(resourceGroups) > 0 {
		azureConfigurationDetails.SetResourceGroups(toSliceString(resourceGroups))
	}

	return azureConfigurationDetails
}

func resourceCloudConnectionConfigurationAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	myctx, _, apiClient := initializeCloudConnectionClient(m)

	azureConfigurationUpdate := cloudconnectionapi.ConfigurationUpdate{}
	if v, ok := d.GetOk("display_name"); ok {
		azureConfigurationUpdate.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		azureConfigurationUpdate.SetDescription(v.(string))
	}
	if v, ok := d.GetOk("details"); ok {
		azureConfigurationUpdateDetails := expandCloudConnectionConfigurationAzureUpdateDetails(v, d)
		azureConfigurationUpdate.SetDetails(azureConfigurationUpdateDetails)
	}
	resp, _, err := apiClient.ConfigurationsApi.UpdateConfiguration(myctx, d.Id()).ConfigurationUpdate(azureConfigurationUpdate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)
	return resourceCloudConnectionConfigurationAzureRead(ctx, d, m)
}

func expandCloudConnectionConfigurationAzureUpdateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConfigurationUpdateDetails {
	azureConfigurationDetails := cloudconnectionapi.AzureConfigurationDetails{}

	details,_:=singleListToMap(v)

	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)
	resourceGroups := details["resource_groups"].([]interface{})

	services := details["services"].([]interface{})
	servicesList := make([]map[string]interface{}, 0, len(services))
	for _, v := range services {
		service := v.(map[string]interface{})
		serviceMap := make(map[string]interface{})
		serviceMap["name"] = service["name"]
		if service["tag_filter"].(string) != "" {
			serviceMap["tagFilter"] = service["tag_filter"]
		}
		if polling,ok:=singleListToMap(service["polling"]);ok{
			serviceMap["polling"] = polling
		}
	
		if importTags,ok:=singleListToMap(service["import_tags"]);ok{
			tags := importTags
			tags["excludedKeys"] = tags["excluded_keys"]
			delete(tags, "excluded_keys")
			serviceMap["importTags"] = tags
		}

		servicesList = append(servicesList, serviceMap)
	}

	if len(regions) > 0 {
		azureConfigurationDetails.SetRegions(toSliceString(regions))
	} else {
		azureConfigurationDetails.SetRegions(make([]string, 0))
	}
	if polling, ok := expandCloudConnectionConfigurationDetailsPolling(details, d); ok {
		azureConfigurationDetails.SetPolling(polling)
	}
	if importTags, ok := expandCloudConnectionConfigurationDetailsImportTags(details, d); ok {
		azureConfigurationDetails.SetImportTags(importTags)
	} else {
		importTags := cloudconnectionapi.ImportTagConfiguration{}
		importTags.Enabled = true
		importTags.ExcludedKeys = make([]string, 0)
		azureConfigurationDetails.SetImportTags(importTags)
	}

	azureConfigurationDetails.SetTagFilter(tagFilter)

	if len(servicesList) > 0 {
		azureConfigurationDetails.SetServices(servicesList)
	}
	if len(resourceGroups) > 0 {
		azureConfigurationDetails.SetResourceGroups(toSliceString(resourceGroups))
	} else {
		azureConfigurationDetails.SetResourceGroups(make([]string, 0))
	}

	azureUpdateDetails := AzureConfigurationDetailsAsConfigurationUpdateDetails(&azureConfigurationDetails)
	return azureUpdateDetails
}

func AzureConfigurationDetailsAsConfigurationUpdateDetails(v *cloudconnectionapi.AzureConfigurationDetails) cloudconnectionapi.ConfigurationUpdateDetails {
	return cloudconnectionapi.ConfigurationUpdateDetails{
		AzureConfigurationDetails: v,
	}
}
