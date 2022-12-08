package provider

import (
	"context"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudConnectionAWS() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceCloudConnectionAWSCreate,
		ReadContext:   resourceCloudConnectionRead,
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
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_type": {
							Type:         schema.TypeString,
							Description:  "Connection type discriminator",
							ValidateFunc: validation.StringInSlice([]string{"role_delegation", "access_key"}, false),
							Required:     true,
						},
						"access_key_id": {
							Type:        schema.TypeString,
							Description: "AWS Access keys are long-term credentials for an AWS IAM user, or account root user. The access key ID is one of two access keys needed to authenticate to AWS. The other is a secret access key. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell.",
							Optional:    true,
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Description: "The secret access key is one of two access keys needed to authenticate to AWS. The other is an access key ID. The secret access key is only available once, when you create it. Download the generated secret access key and save in a secure location. If the secret access key is lost or deleted, you must create a new one. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell.",
							Optional:    true,
						},

						// The following attributes will be computed conditionally.
						// Keeping them as computed always as DiffSuppress won't work on Computed,
						// and ComputedIf will not work either as it is inside a Set
						// so these attributes will always show as (known after apply) in plan.

						// computed for aws access_key
						"account_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// computed for access_type role_delegation
						"role_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
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

	resp, _, err := apiClient.ConnectionsApi.CreateConnection(myctx).ConnectionRequest(connectionRequest).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionRead(ctx, d, m)
}

func resourceCloudConnectionAWSUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO: Later
	return nil
}

func expandCloudConnectionAWSDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConnectionRequestDetails {
	connectionRequestDetails := cloudconnectionapi.ConnectionRequestDetails{}

	details := v.(*schema.Set).List()[0].(map[string]interface{})
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

func awsRequiredAttributesCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	details := d.Get("details").(*schema.Set).List()[0].(map[string]interface{})

	accessType := details["access_type"].(string)

	err := checkRequiredNotRequired(d, accessType)
	if err != nil {
		return err
	}

	err = checkRequiredNotRequired(d, accessType)
	if err != nil {
		return err
	}

	return nil
}
