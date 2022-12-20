package provider

import (
	"context"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudConnectionAzure() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudConnectionAzureCreate,
		ReadContext:   resourceCloudConnectionAzureRead,
		UpdateContext: resourceCloudConnectionAzureUpdate,
		DeleteContext: resourceCloudConnectionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		CustomizeDiff: customdiff.All(serviceAtLeastOne),
		SchemaVersion: 1,

		Schema: getCloudConnectionAzureSchema(),
	}
}

// ====================================CREATE======================================
func resourceCloudConnectionAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionRequest := cloudconnectionapi.ConnectionRequest{}
	connectionRequest.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AZURE))

	azureConfiguration := cloudconnectionapi.AzureConfiguration{}
	azureConfiguration.BaseEntity.BaseEntityAllOf.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AZURE))

	if v, ok := d.GetOk("display_name"); ok {
		connectionRequest.SetDisplayName(v.(string))
		azureConfiguration.BaseEntity.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		connectionRequest.SetDescription(v.(string))
		azureConfiguration.BaseEntity.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("state"); ok {
		connectionRequest.SetState(v.(string))
	}

	if v, ok := d.GetOk("connection_details"); ok {
		azureConnectionDetails := expandCloudConnectionAzureDetails(v, d)
		connectionRequest.SetDetails(azureConnectionDetails)
	}

	if v, ok := d.GetOk("configuration_details"); ok {
		azureConnectionConfigurationDetails := expandCloudConnectionConfigurationAzureCreateDetails(v, d)
		azureConfiguration.SetDetails(azureConnectionConfigurationDetails)
	} else {
		//Create Default Details with default services.
		azureConfiguration.SetDetails(cloudconnectionapi.AzureConfigurationDetails{})
	}

	// Create a configuration

	configuration := cloudconnectionapi.AzureConfigurationAsConfiguration(&azureConfiguration)

	respConfiguration, httpRespConfiguration, err := apiClient.ConfigurationsApi.CreateConfiguration(myctx).Configuration(configuration).Execute()
	if err != nil {
		return errRespToDiag(err, httpRespConfiguration)
	}

	d.Set("configuration_id", respConfiguration.Id)

	// Create a connection
	if v, ok := d.GetOk("configuration_id"); ok {
		connectionRequest.SetConfigurationId(v.(string))
	}

	respConnection, httpRespConnection, err := apiClient.ConnectionsApi.CreateConnection(myctx).ConnectionRequest(connectionRequest).Execute()
	if err != nil {
		return errRespToDiag(err, httpRespConnection)
	}

	d.SetId(respConnection.Id)

	return resourceCloudConnectionAzureRead(ctx, d, m)
	// return nil
}

func expandCloudConnectionAzureDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConnectionRequestDetails {
	connectionRequestDetails := cloudconnectionapi.ConnectionRequestDetails{}

	details, _ := singleListToMap(v)

	clientId := details["client_id"].(string)
	clientSecret := details["client_secret"].(string)
	tenantId := details["tenant_id"].(string)
	subscriptionId := details["subscription_id"].(string)

	connectionRequestDetails.AzureDetails = cloudconnectionapi.NewAzureDetails(clientId, clientSecret, tenantId, subscriptionId)

	return connectionRequestDetails
}

func expandCloudConnectionConfigurationAzureCreateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.AzureConfigurationDetails {
	azureConfigurationDetails := cloudconnectionapi.AzureConfigurationDetails{}

	details, _ := singleListToMap(v)
	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)
	resourceGroups := details["resource_groups"].([]interface{})

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
		azureConfigurationDetails.SetRegions(toSliceString(regions))
	}

	if len(resourceGroups) > 0 {
		azureConfigurationDetails.SetResourceGroups(toSliceString(resourceGroups))
	}

	if len(servicesList) > 0 {
		azureConfigurationDetails.SetServices(servicesList)
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

	return azureConfigurationDetails
}

// ====================================READ======================================

func resourceCloudConnectionAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()

	respConnection, httpRespConnection, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		if httpRespConnection.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return errRespToDiag(err, httpRespConnection)
	}

	d.Set("connection_details", flattenCloudConnectionAzureDetails(respConnection, d))

	flattenCloudConnectionCommons(respConnection, d)

	configurationId := respConnection.ConfigurationId

	respConfiguration, httpRespConfiguration, err := apiClient.ConfigurationsApi.GetConfiguration(myctx, *configurationId).Execute()
	if err != nil {
		if httpRespConfiguration.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return errRespToDiag(err, httpRespConfiguration)
	}
	flattenCloudConnectionConfigurationCommonsDetails(respConfiguration, d, "azure")

	return nil
}

func flattenCloudConnectionAzureDetails(resp *cloudconnectionapi.ConnectionResponse, d *schema.ResourceData) interface{} {
	var clientSecret string = ""

	// For datasource, if details block does not exists
	// set client secret as recieved from response.
	if v, ok := d.GetOk("connection_details"); ok {
		details, _ := singleListToMap(v)
		clientSecret = details["client_secret"].(string)
	} else {
		clientSecret = resp.GetDetails().AzureDetails.ClientSecret
	}

	detailsList := []interface{}{}
	detailsList = append(detailsList, map[string]interface{}{
		"client_id":       resp.GetDetails().AzureDetails.ClientId,
		"client_secret":   clientSecret,
		"tenant_id":       resp.GetDetails().AzureDetails.TenantId,
		"subscription_id": resp.GetDetails().AzureDetails.SubscriptionId,
	})

	return detailsList
}

// ====================================UPDATE======================================

func resourceCloudConnectionAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	azureConnectionUpdate := cloudconnectionapi.ConnectionUpdate{}
	azureConfigurationUpdate := cloudconnectionapi.ConfigurationUpdate{}

	updateConfigurationFlag := false

	if d.HasChange("display_name") {
		updateConfigurationFlag = true
		azureConnectionUpdate.SetDisplayName(d.Get("display_name").(string))
		azureConfigurationUpdate.SetDisplayName(d.Get("display_name").(string))
	}

	if d.HasChange("description") {
		updateConfigurationFlag = true
		azureConnectionUpdate.SetDescription(d.Get("description").(string))
		azureConfigurationUpdate.SetDescription(d.Get("description").(string))
	}

	if d.HasChange("configuration_details") {
		updateConfigurationFlag = true
		connectionConfigurationUpdateDetails := expandCloudConnectionConfigurationAzureUpdateDetails(d.Get("configuration_details"), d)

		azureConfigurationUpdate.SetDetails(connectionConfigurationUpdateDetails)
	}

	if updateConfigurationFlag {
		resp, httpResp, err := apiClient.ConfigurationsApi.UpdateConfiguration(myctx, d.Id()).ConfigurationUpdate(azureConfigurationUpdate).Execute()
		if err != nil {
			return errRespToDiag(err, httpResp)
		}

		d.Set("configuration_id", resp.Id)
	}

	azureConnectionUpdate.SetConfigurationId(d.Get("configuration_id").(string))

	if d.HasChange("state") {
		azureConnectionUpdate.SetState(d.Get("state").(string))
	}

	if d.HasChange("connection_details") {
		connectionUpdateDetails, err := expandCloudConnectionAzureUpdateDetails(d.Get("connection_details"), d)
		if err != nil {
			return diag.FromErr(err)
		}

		azureConnectionUpdate.SetDetails(connectionUpdateDetails)
	}

	resp, httpResp, err := apiClient.ConnectionsApi.UpdateConnection(myctx, d.Id()).ConnectionUpdate(azureConnectionUpdate).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionAzureRead(ctx, d, m)
}

func expandCloudConnectionAzureUpdateDetails(v interface{}, d *schema.ResourceData) (cloudconnectionapi.ConnectionUpdateDetails, error) {
	connectionDetails := expandCloudConnectionAzureDetails(v, d).AzureDetails

	connectionUpdateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf2{}
	connectionUpdateDetails.SetClientId(connectionDetails.ClientId)
	connectionUpdateDetails.SetClientSecret(connectionDetails.ClientSecret)

	updateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf2AsConnectionUpdateDetails(&connectionUpdateDetails)
	return updateDetails, nil
}

func expandCloudConnectionConfigurationAzureUpdateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConfigurationUpdateDetails {
	azureConfigurationDetails := cloudconnectionapi.AzureConfigurationDetails{}

	details, _ := singleListToMap(v)
	regions := details["regions"].([]interface{})
	tagFilter := details["tag_filter"].(string)
	resourceGroups := details["resource_groups"].([]interface{})

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
		azureConfigurationDetails.SetRegions(toSliceString(regions))
	} else {
		azureConfigurationDetails.SetRegions(make([]string, 0))
	}

	if len(resourceGroups) > 0 {
		azureConfigurationDetails.SetResourceGroups(toSliceString(resourceGroups))
	} else {
		azureConfigurationDetails.SetResourceGroups(make([]string, 0))
	}

	if len(servicesList) > 0 {
		azureConfigurationDetails.SetServices(servicesList)
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

	azureUpdateDetails := azureConfigurationDetailsAsConfigurationUpdateDetails(&azureConfigurationDetails)
	return azureUpdateDetails
}

func azureConfigurationDetailsAsConfigurationUpdateDetails(v *cloudconnectionapi.AzureConfigurationDetails) cloudconnectionapi.ConfigurationUpdateDetails {
	return cloudconnectionapi.ConfigurationUpdateDetails{
		AzureConfigurationDetails: v,
	}
}
