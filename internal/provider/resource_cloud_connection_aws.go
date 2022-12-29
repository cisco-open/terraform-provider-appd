package provider

import (
	"context"
	"fmt"
	"reflect"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudConnectionAWS() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceCloudConnectionAWSCreate,
		ReadContext:   resourceCloudConnectionAWSRead,
		UpdateContext: resourceCloudConnectionAWSUpdate,
		DeleteContext: resourceCloudConnectionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudConnectionAWSImport,
		},

		SchemaVersion: 1,

		CustomizeDiff: customdiff.All(
			serviceAtLeastOne,
			awsRequiredAttributesCustomizeDiff,
		),

		Schema: getCloudConnectionAWSSchema(),
	}
}

func resourceCloudConnectionAWSImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()

	respConnection, _, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		return nil, err
	}

	d.Set("connection_details", flattenCloudConnectionAWSDetails(respConnection, d))

	flattenCloudConnectionCommons(respConnection, d)

	configurationId := respConnection.ConfigurationId

	d.Set("configuration_id", configurationId)

	respConfiguration, _, err := apiClient.ConfigurationsApi.GetConfiguration(myctx, *configurationId).Execute()
	if err != nil {
		return nil, err
	}
	flattenCloudConnectionConfigurationCommonsDetails(respConfiguration, d, "aws")

	return []*schema.ResourceData{d}, nil
}

func resourceCloudConnectionAWSCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionRequest := cloudconnectionapi.ConnectionRequest{}
	connectionRequest.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AWS))

	awsConfiguration := cloudconnectionapi.AWSConfiguration{}
	awsConfiguration.BaseEntity.BaseEntityAllOf.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AWS))

	if v, ok := d.GetOk("display_name"); ok {
		connectionRequest.SetDisplayName(v.(string))
		awsConfiguration.BaseEntity.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		connectionRequest.SetDescription(v.(string))
		awsConfiguration.BaseEntity.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("state"); ok {
		connectionRequest.SetState(v.(string))
	}

	if v, ok := d.GetOk("connection_details"); ok {
		connectionRequestDetails := expandCloudConnectionAWSDetails(v, d)
		connectionRequest.SetDetails(connectionRequestDetails)
	}

	if v, ok := d.GetOk("configuration_details"); ok {
		awsConnectionConfigurationDetails := expandCloudConnectionConfigurationAWSCreateDetails(v, d)
		awsConfiguration.SetDetails(awsConnectionConfigurationDetails)
	} else {
		//Create Default Details with default services.
		awsConfiguration.SetDetails(cloudconnectionapi.AWSConfigurationDetails{})
	}

	// Create a configuration

	configuration := cloudconnectionapi.AWSConfigurationAsConfiguration(&awsConfiguration)

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
		// Delete configuration created when error obtained while creating connection.
		httpResp, err := apiClient.ConfigurationsApi.DeleteConfiguration(myctx, respConfiguration.Id).Execute()
		if err != nil {
			return errRespToDiag(err, httpResp)
		}
		return errRespToDiag(err, httpRespConnection)
	}

	d.SetId(respConnection.Id)

	return resourceCloudConnectionAWSRead(ctx, d, m)
}

func resourceCloudConnectionAWSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	diags := diag.Diagnostics{}
	connectionId := d.Id()

	respConnection, httpRespConnection, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		if httpRespConnection.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return errRespToDiag(err, httpRespConnection)
	}

	d.Set("connection_details", flattenCloudConnectionAWSDetails(respConnection, d))

	flattenCloudConnectionCommons(respConnection, d)
	if !((d.Get("state") == "ACTIVE") || (d.Get("state") == "INACTIVE") || (d.Get("state") == "CONFIGURED")) {
		d := diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "State Warning",
			Detail:   fmt.Sprintf("Current State:%v\nState Message: %v", d.Get("state"), d.Get("state_message")),
		}
		diags = append(diags, d)
	}

	configurationId := respConnection.ConfigurationId

	respConfiguration, httpRespConfiguration, err := apiClient.ConfigurationsApi.GetConfiguration(myctx, *configurationId).Execute()
	if err != nil {
		if httpRespConfiguration.StatusCode == 404 {
			d.Set("configuration_id", "")
			return nil
		}
		return errRespToDiag(err, httpRespConfiguration)
	}
	d.Set("configuration_id", configurationId)
	flattenCloudConnectionConfigurationCommonsDetails(respConfiguration, d, "aws")

	return diags
}

func resourceCloudConnectionAWSUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	awsConnectionUpdate := cloudconnectionapi.ConnectionUpdate{}
	awsConfigurationUpdate := cloudconnectionapi.ConfigurationUpdate{}

	updateConfigurationFlag := false

	if d.HasChange("display_name") {
		updateConfigurationFlag = true
		awsConnectionUpdate.SetDisplayName(d.Get("display_name").(string))
		awsConfigurationUpdate.SetDisplayName(d.Get("display_name").(string))

	}

	if d.HasChange("description") {
		updateConfigurationFlag = true
		awsConnectionUpdate.SetDescription(d.Get("description").(string))
		awsConfigurationUpdate.SetDescription(d.Get("description").(string))
	}

	if d.HasChange("configuration_details") {
		updateConfigurationFlag = true
		connectionConfigurationUpdateDetails := expandCloudConnectionConfigurationAWSUpdateDetails(d.Get("configuration_details"), d)

		awsConfigurationUpdate.SetDetails(connectionConfigurationUpdateDetails)
	}

	if updateConfigurationFlag {
		resp, httpResp, err := apiClient.ConfigurationsApi.UpdateConfiguration(myctx, d.Get("configuration_id").(string)).ConfigurationUpdate(awsConfigurationUpdate).Execute()
		if err != nil {
			return errRespToDiag(err, httpResp)
		}

		d.Set("configuration_id", resp.Id)
	}

	awsConnectionUpdate.SetConfigurationId(d.Get("configuration_id").(string))

	if d.HasChange("state") {
		awsConnectionUpdate.SetState(d.Get("state").(string))
	}

	if d.HasChange("connection_details") {
		connectionUpdateDetails, err := expandCloudConnectionAwsUpdateDetails(d.Get("connection_details"), d)
		if err != nil {
			return diag.FromErr(err)
		}

		awsConnectionUpdate.SetDetails(connectionUpdateDetails)
	}

	resp, httpResp, err := apiClient.ConnectionsApi.UpdateConnection(myctx, d.Id()).ConnectionUpdate(awsConnectionUpdate).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionAWSRead(ctx, d, m)

}

func expandCloudConnectionAWSDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConnectionRequestDetails {
	connectionRequestDetails := cloudconnectionapi.ConnectionRequestDetails{}

	details, _ := singleListToMap(v)
	connectionRequestDetails.AWSConnectionRequestDetails = expandCloudConnectionAWSDetailsAWSCredentials(details, d)

	return connectionRequestDetails
}

func expandCloudConnectionAWSDetailsAWSCredentials(v interface{}, d *schema.ResourceData) *cloudconnectionapi.AWSConnectionRequestDetails {
	awsConnectionRequestDetails := &cloudconnectionapi.AWSConnectionRequestDetails{}

	details := v.(map[string]interface{})

	accessType := details["access_type"].(string)

	if accessType == string(cloudconnectionapi.ROLE_DELEGATION) {
		awsConnectionRequestDetails.AWSRoleDelegationCreationDetails = expandCloudConnectionAWSDetailsAWSRoleDelegation(details, d)
	} else if accessType == string(cloudconnectionapi.ACCESS_KEY) {
		awsConnectionRequestDetails.AWSAccessKeyDetails = expandCloudConnectionAWSDetailsAWSAccessKey(details, d)
	}

	return awsConnectionRequestDetails
}

func expandCloudConnectionAWSDetailsAWSRoleDelegation(v interface{}, d *schema.ResourceData) *cloudconnectionapi.AWSRoleDelegationCreationDetails {
	details := v.(map[string]interface{})

	accountId := details["account_id"].(string)

	return cloudconnectionapi.NewAWSRoleDelegationCreationDetails(cloudconnectionapi.ROLE_DELEGATION, accountId)
}

func expandCloudConnectionAWSDetailsAWSAccessKey(v interface{}, d *schema.ResourceData) *cloudconnectionapi.AWSAccessKeyDetails {
	details := v.(map[string]interface{})

	accessKeyId := details["access_key_id"].(string)
	secretAccessKey := details["secret_access_key"].(string)

	return cloudconnectionapi.NewAWSAccessKeyDetails(accessKeyId, secretAccessKey, cloudconnectionapi.ACCESS_KEY)
}

func expandCloudConnectionAwsUpdateDetails(v interface{}, d *schema.ResourceData) (cloudconnectionapi.ConnectionUpdateDetails, error) {
	connectionDetails := expandCloudConnectionAWSDetails(v, d).AWSConnectionRequestDetails

	connectionUpdateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf1{}
	connectionUpdateDetails.SetAccessKeyId(connectionDetails.AWSAccessKeyDetails.AccessKeyId)
	connectionUpdateDetails.SetSecretAccessKey(connectionDetails.AWSAccessKeyDetails.SecretAccessKey)

	updateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf1AsConnectionUpdateDetails(&connectionUpdateDetails)
	return updateDetails, nil
}

func flattenCloudConnectionAWSDetails(resp *cloudconnectionapi.ConnectionResponse, d *schema.ResourceData) interface{} {
	awsConnectionDetails := resp.GetDetails().AWSConnectionResponseDetails
	isAccessKey := reflect.ValueOf(awsConnectionDetails.RoleDelegationConnectionResponseDetails).IsNil()

	detailsSet := []interface{}{}

	if isAccessKey {
		accessKeyDetails := flattenCloudConnectionAWSDetailsAccessKey(awsConnectionDetails, d)
		detailsSet = append(detailsSet, accessKeyDetails)
	} else {
		roleDelegationDetails := flattenCloudConnectionAWSDetailsRoleDelegation(awsConnectionDetails, d)
		detailsSet = append(detailsSet, roleDelegationDetails)
	}

	return detailsSet
}

func flattenCloudConnectionAWSDetailsRoleDelegation(awsConnectionDetails *cloudconnectionapi.AWSConnectionResponseDetails, d *schema.ResourceData) interface{} {
	roleDelegationDetails := awsConnectionDetails.RoleDelegationConnectionResponseDetails

	return map[string]interface{}{
		"access_type":                roleDelegationDetails.AccessType,
		"account_id":                 roleDelegationDetails.AccountId,
		"appdynamics_aws_account_id": roleDelegationDetails.GetAppDynamicsAwsAccountId(),
		"external_id":                roleDelegationDetails.GetExternalId(),
		"role_name":                  roleDelegationDetails.GetRoleName(),
	}
}

func flattenCloudConnectionAWSDetailsAccessKey(awsConnectionDetails *cloudconnectionapi.AWSConnectionResponseDetails, d *schema.ResourceData) interface{} {
	accessKeyDetails := awsConnectionDetails.AccessKeyConnectionResponseDetails

	secretAccessKey := ""
	if v, ok := d.GetOk("connection_details"); ok {
		details, _ := singleListToMap(v)
		secretAccessKey = details["secret_access_key"].(string)
	} else {
		secretAccessKey = accessKeyDetails.SecretAccessKey
	}

	return map[string]interface{}{
		"access_type":       cloudconnectionapi.ACCESS_KEY,
		"access_key_id":     accessKeyDetails.AccessKeyId,
		"secret_access_key": secretAccessKey,
		"aws_account_id":    accessKeyDetails.AccountId,
	}
}

// =========================================CONFIGURATION=========================
func expandCloudConnectionConfigurationAWSCreateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.AWSConfigurationDetails {
	awsConfigurationDetails := cloudconnectionapi.AWSConfigurationDetails{}

	details, _ := singleListToMap(v)
	regions := details["regions"].(*schema.Set).List()
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

func expandCloudConnectionConfigurationAWSUpdateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConfigurationUpdateDetails {
	awsConfigurationDetails := cloudconnectionapi.AWSConfigurationDetails{}

	details, _ := singleListToMap(v)
	regions := details["regions"].(*schema.Set).List()
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
		importTags := cloudconnectionapi.ImportTagConfiguration{}
		importTags.Enabled = true
		importTags.ExcludedKeys = make([]string, 0)
		awsConfigurationDetails.SetImportTags(importTags)
	}

	awsConfigurationDetails.SetTagFilter(tagFilter)

	awsUpdateDetails := awsConfigurationDetailsAsConfigurationUpdateDetails(&awsConfigurationDetails)
	return awsUpdateDetails
}

func awsConfigurationDetailsAsConfigurationUpdateDetails(v *cloudconnectionapi.AWSConfigurationDetails) cloudconnectionapi.ConfigurationUpdateDetails {
	return cloudconnectionapi.ConfigurationUpdateDetails{
		AWSConfigurationDetails: v,
	}
}

func awsRequiredAttributesCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	details, _ := singleListToMap(d.Get("connection_details"))

	accessType := details["access_type"].(string)

	err := checkRequiredNotRequired(d, accessType)
	if err != nil {
		return err
	}

	return nil
}
