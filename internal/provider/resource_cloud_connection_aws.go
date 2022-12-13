package provider

import (
	"context"
	"reflect"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudConnectionAWS() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceCloudConnectionAWSCreate,
		ReadContext:   resourceCloudConnectionAWSRead,
		UpdateContext: resourceCloudConnectionAWSUpdate,
		DeleteContext: resourceCloudConnectionDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,

		CustomizeDiff: customdiff.All(
			awsRequiredAttributesCustomizeDiff,
		),

		Schema: appendSchema(getCommonCloudConnectionSchema(), map[string]*schema.Schema{
			"details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_type": {
							Type:         schema.TypeString,
							Description:  "Connection type discriminator",
							ValidateFunc: validation.StringInSlice([]string{"role_delegation", "access_key"}, false),
							Required:     true,
							ForceNew:     true,
						},
						"access_key_id": {
							Type:          schema.TypeString,
							Description:   "AWS Access keys are long-term credentials for an AWS IAM user, or account root user. The access key ID is one of two access keys needed to authenticate to AWS. The other is a secret access key. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell.",
							Optional:      true,
							ConflictsWith: []string{"details.0.account_id"},
						},
						"secret_access_key": {
							Type:          schema.TypeString,
							Description:   "The secret access key is one of two access keys needed to authenticate to AWS. The other is an access key ID. The secret access key is only available once, when you create it. Download the generated secret access key and save in a secure location. If the secret access key is lost or deleted, you must create a new one. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell.",
							Optional:      true,
							Sensitive:     true,
							ConflictsWith: []string{"details.0.account_id"},
						},

						// computed for aws access_key
						"account_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// computed for access_type role_delegation
						"appdynamics_aws_account_id": {
							Type: schema.TypeString,
							Description: `AppDynamics AWS Account ID. Delegates a user to an Identity Access Management (IAM) role in AWS. The AWS IAM role provides AppDynamics access to resources.
							https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html`,
							Computed: true,
						},
						"external_id": {
							Type: schema.TypeString,
							Description: `Returns an external ID for AWS role delegation connections 
							https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html`,
							Computed: true,
						},
					},
				},
			},
		}),
	}
}

func resourceCloudConnectionAWSCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	err := checkStateConfiguration(ctx, d, m)
	if err != nil {
		return diag.FromErr(err)
	}

	connectionRequest := cloudconnectionapi.ConnectionRequest{}
	connectionRequest.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AWS))

	if v, ok := d.GetOk("display_name"); ok {
		connectionRequest.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		connectionRequest.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("state"); ok {
		connectionRequest.SetState(v.(string))
	}

	if v, ok := d.GetOk("configuration_id"); ok {
		connectionRequest.SetConfigurationId(v.(string))
	}

	if v, ok := d.GetOk("details"); ok {
		connectionRequestDetails := expandCloudConnectionAWSDetails(v, d)
		connectionRequest.SetDetails(connectionRequestDetails)
	}

	resp, httpResp, err := apiClient.ConnectionsApi.CreateConnection(myctx).ConnectionRequest(connectionRequest).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionAWSRead(ctx, d, m)
}

func resourceCloudConnectionAWSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()

	resp, httpResp, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	detailsSet := flattenCloudConnectionAWSDetails(resp, d)
	d.Set("details", detailsSet)

	flattenCloudConnectionCommons(resp, d)

	return nil
}

func resourceCloudConnectionAWSUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionUpdate := cloudconnectionapi.ConnectionUpdate{}

	err := checkStateConfiguration(ctx, d, m)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("display_name") {
		connectionUpdate.SetDisplayName(d.Get("display_name").(string))
	}

	if d.HasChange("description") {
		connectionUpdate.SetDescription(d.Get("description").(string))
	}

	if d.HasChange("configuration_id") {
		connectionUpdate.SetConfigurationId(d.Get("configuration_id").(string))
	}

	if d.HasChange("state") {
		connectionUpdate.SetState(d.Get("state").(string))
	}

	if d.HasChange("details") {
		connectionUpdateDetails, err := expandCloudConnectionAwsUpdateDetails(d.Get("details"), d)
		if err != nil {
			return diag.FromErr(err)
		}

		connectionUpdate.SetDetails(connectionUpdateDetails)
	}

	resp, httpResp, err := apiClient.ConnectionsApi.UpdateConnection(myctx, d.Id()).ConnectionUpdate(connectionUpdate).Execute()
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
	}
}

func flattenCloudConnectionAWSDetailsAccessKey(awsConnectionDetails *cloudconnectionapi.AWSConnectionResponseDetails, d *schema.ResourceData) interface{} {
	accessKeyDetails := awsConnectionDetails.AccessKeyConnectionResponseDetails

	secretAccessKey := ""
	if v, ok := d.GetOk("details"); ok {
		details, _ := singleListToMap(v)
		secretAccessKey = details["secret_access_key"].(string)
	} else {
		secretAccessKey = accessKeyDetails.SecretAccessKey
	}

	return map[string]interface{}{
		"access_type":       accessKeyDetails.AccessType,
		"access_key_id":     accessKeyDetails.AccessKeyId,
		"secret_access_key": secretAccessKey,
	}
}

func awsRequiredAttributesCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	details, _ := singleListToMap(d.Get("details"))

	accessType := details["access_type"].(string)

	err := checkRequiredNotRequired(d, accessType)
	if err != nil {
		return err
	}

	return nil
}
