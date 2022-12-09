package provider

// import (
// 	"context"

// 	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
// )

// func resourceCloudConnectionConnection() *schema.Resource {
// 	return &schema.Resource{
// 		CreateContext: resourceCloudConnectionConnectionCreate,
// 		ReadContext:   resourceCloudConnectionConnectionRead,
// 		UpdateContext: resourceCloudConnectionConnectionUpdate,
// 		DeleteContext: resourceCloudConnectionConnectionDelete,

// 		Importer: &schema.ResourceImporter{
// 			StateContext: schema.ImportStatePassthroughContext,
// 		},

// 		SchemaVersion: 1,

// 		CustomizeDiff: customdiff.All(
// 			awsRequiredAttributesCustomizeDiff,
// 			azureRequiredAttributesCustomizeDiff,
// 		),

// 		Schema: map[string]*schema.Schema{
// 			"type": {
// 				Type:         schema.TypeString,
// 				Description:  "Provider type (also known as Connection type)",
// 				ValidateFunc: validation.StringInSlice([]string{"aws", "azure"}, false),
// 				Required:     true,
// 			},
// 			"display_name": {
// 				Type:        schema.TypeString,
// 				Description: "Name of the connection or configuration",
// 				Required:    true,
// 			},
// 			"description": {
// 				Type:        schema.TypeString,
// 				Description: "Description for this connection or configuration",
// 				Optional:    true,
// 			},
// 			"state": {
// 				Type:        schema.TypeString,
// 				Description: "Connection state",
// 				Optional:    true,
// 				Computed:    true,
// 			},
// 			"state_message": {
// 				Type:        schema.TypeString,
// 				Description: "Connection state message",
// 				Computed:    true,
// 			},
// 			"configuration_id": {
// 				Type:         schema.TypeString,
// 				ValidateFunc: validation.IsUUID,
// 				Optional:     true,
// 				Computed:     true,
// 			},
// 			"details": {
// 				Type:     schema.TypeSet,
// 				Required: true,
// 				MaxItems: 1,
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						// aws
// 						"access_type": {
// 							Type:         schema.TypeString,
// 							Description:  "Connection type discriminator",
// 							ValidateFunc: validation.StringInSlice([]string{"role_delegation", "access_key"}, false),
// 							Optional:     true,
// 						},
// 						"access_key_id": {
// 							Type:        schema.TypeString,
// 							Description: "AWS Access keys are long-term credentials for an AWS IAM user, or account root user. The access key ID is one of two access keys needed to authenticate to AWS. The other is a secret access key. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell.",
// 							Optional:    true,
// 						},
// 						"secret_access_key": {
// 							Type:        schema.TypeString,
// 							Description: "The secret access key is one of two access keys needed to authenticate to AWS. The other is an access key ID. The secret access key is only available once, when you create it. Download the generated secret access key and save in a secure location. If the secret access key is lost or deleted, you must create a new one. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell.",
// 							Optional:    true,
// 						},

// 						// The following attributes will be computed conditionally.
// 						// Keeping them as computed always as DiffSuppress won't work on Computed,
// 						// and ComputedIf will not work either as it is inside a Set
// 						// so these attributes will always show as (known after apply) in plan.

// 						// computed for aws access_key
// 						"account_id": {
// 							Type:     schema.TypeString,
// 							Optional: true,
// 							Computed: true,
// 						},

// 						// computed for access_type role_delegation
// 						"role_name": {
// 							Type:     schema.TypeString,
// 							Computed: true,
// 						},
// 						"appdynamics_aws_account_id": {
// 							Type: schema.TypeString,
// 							Description: `AppDynamics AWS Account ID. Delegates a user to an Identity Access Management (IAM) role in AWS. The AWS IAM role provides AppDynamics access to resources.
// 							https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html`,
// 							Computed: true,
// 						},
// 						"external_id": {
// 							Type: schema.TypeString,
// 							Description: `Returns an external ID for AWS role delegation connections
// 							https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html`,
// 							Computed: true,
// 						},

// 						// azure
// 						"client_id": {
// 							Type:        schema.TypeString,
// 							Description: "Client IDs, also known as Application IDs, are long-term credentials for an Azure user, or account root user. The Client ID is one of three properties needed to authenticate to Azure, the other two being Client Secret and Ten.ant (Directory) ID.",
// 							Optional:    true,
// 						},
// 						"client_secret": {
// 							Type:        schema.TypeString,
// 							Description: "A Client Secret allows an Azure application to provide its identity when requesting an access token. The Client Secret is one of three properties needed to authenticate to Azure, the other two being Client ID (Application ID) and Tenant (Directory) ID.",
// 							Optional:    true,
// 						},
// 						"tenant_id": {
// 							Type:        schema.TypeString,
// 							Description: "The Azure AD Tenant (Directory) IDis one of three properties needed to authenticate to Azure. The other two are Client Secret and Client ID (Application ID).",
// 							Optional:    true,
// 						},
// 						"subscription_id": {
// 							Type:        schema.TypeString,
// 							Description: "Specify a GUID Subscription ID to monitor. If monitoring all subscriptions, do not specify a Subscription ID.",
// 							Optional:    true,
// 						},
// 					},
// 				},
// 			},
// 			"created_at": {
// 				Type:     schema.TypeString,
// 				Computed: true,
// 			},
// 			"updated_at": {
// 				Type:     schema.TypeString,
// 				Computed: true,
// 			},
// 		},
// 	}
// }

// func resourceCloudConnectionConnectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	myctx, _, apiClient := initializeCloudConnectionClient(m)

// 	connectionRequest := cloudconnectionapi.ConnectionRequest{}

// 	if v, ok := d.GetOk("type"); ok {
// 		connectionRequest.SetType(cloudconnectionapi.ProviderType(v.(string)))
// 	}

// 	if v, ok := d.GetOk("display_name"); ok {
// 		connectionRequest.SetDisplayName(v.(string))
// 	}

// 	if v, ok := d.GetOk("description"); ok {
// 		connectionRequest.SetDescription(v.(string))
// 	}

// 	if v, ok := d.GetOk("state"); ok {
// 		connectionRequest.SetState(v.(string))
// 	}

// 	if v, ok := d.GetOk("configuration_id"); ok {
// 		connectionRequest.SetConfigurationId(v.(string))
// 	}

// 	if v, ok := d.GetOk("details"); ok {
// 		connectionRequestDetails := expandCloudConnectionConnectionDetails(v, d)
// 		connectionRequest.SetDetails(connectionRequestDetails)
// 	}

// 	resp, _, err := apiClient.ConnectionsApi.CreateConnection(myctx).ConnectionRequest(connectionRequest).Execute()
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	d.SetId(resp.Id)

// 	return resourceCloudConnectionConnectionRead(ctx, d, m)
// }

// func resourceCloudConnectionConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	myctx, _, apiClient := initializeCloudConnectionClient(m)

// 	connectionId := d.Id()

// 	resp, _, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	d.SetId(resp.GetId())
// 	d.Set("type", resp.GetType())
// 	d.Set("display_name", resp.GetDisplayName())
// 	d.Set("description", resp.GetDescription())
// 	d.Set("state", resp.GetState())
// 	// TODO: Later
// 	// d.Set("state_message", resp.StateMessage)
// 	// d.Set("details", resp.GetDetails())
// 	d.Set("configuration_id", resp.GetConfigurationId())
// 	d.Set("created_at", resp.GetCreatedAt())
// 	d.Set("updated_at", resp.GetUpdatedAt())

// 	return nil
// }
// func resourceCloudConnectionConnectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	// TODO: Later
// 	return nil
// }
// func resourceCloudConnectionConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	myctx, _, apiClient := initializeCloudConnectionClient(m)

// 	connectionId := d.Id()
// 	_, err := apiClient.ConnectionsApi.DeleteConnection(myctx, connectionId).Execute()

// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	return nil
// }

// func expandCloudConnectionConnectionDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConnectionRequestDetails {
// 	connectionRequestDetails := cloudconnectionapi.ConnectionRequestDetails{}

// 	type_ := d.Get("type").(string)
// 	details := v.(*schema.Set).List()[0].(map[string]interface{})

// 	if type_ == string(cloudconnectionapi.AWS) {
// 		connectionRequestDetails.AWSConnectionRequestDetails = expandCloudConnectionConnectionDetailsAWS(details, d)
// 	} else if type_ == string(cloudconnectionapi.AZURE) {
// 		connectionRequestDetails.AzureDetails = expandCloudConnectionConnectionDetailsAzure(details, d)
// 	}

// 	return connectionRequestDetails
// }

// func expandCloudConnectionConnectionDetailsAWS(v interface{}, d *schema.ResourceData) *cloudconnectionapi.AWSConnectionRequestDetails {
// 	awsConnectionRequestDetails := &cloudconnectionapi.AWSConnectionRequestDetails{}

// 	details := v.(map[string]interface{})

// 	accessType := details["access_type"].(string)

// 	if accessType == string(cloudconnectionapi.ROLE_DELEGATION) {
// 		awsConnectionRequestDetails.AWSRoleDelegationCreationDetails = expandCloudConnectionConnectionDetailsAWSRoleDelegation(details, d)
// 	} else if accessType == string(cloudconnectionapi.ACCESS_KEY) {
// 		awsConnectionRequestDetails.AWSAccessKeyDetails = expandCloudConnectionConnectionDetailsAWSAccessKey(details, d)
// 	}

// 	return awsConnectionRequestDetails
// }

// func expandCloudConnectionConnectionDetailsAzure(v interface{}, d *schema.ResourceData) *cloudconnectionapi.AzureDetails {
// 	details := v.(map[string]interface{})

// 	clientId := details["client_id"].(string)
// 	clientSecret := details["client_secret"].(string)
// 	tenantId := details["tenant_id"].(string)
// 	subscriptionId := details["subscription_id"].(string)

// 	return cloudconnectionapi.NewAzureDetails(clientId, clientSecret, tenantId, subscriptionId)
// }

// func expandCloudConnectionConnectionDetailsAWSRoleDelegation(v interface{}, d *schema.ResourceData) *cloudconnectionapi.AWSRoleDelegationCreationDetails {
// 	details := v.(map[string]interface{})

// 	accountId := details["account_id"].(string)

// 	return cloudconnectionapi.NewAWSRoleDelegationCreationDetails(cloudconnectionapi.ROLE_DELEGATION, accountId)
// }

// func expandCloudConnectionConnectionDetailsAWSAccessKey(v interface{}, d *schema.ResourceData) *cloudconnectionapi.AWSAccessKeyDetails {
// 	details := v.(map[string]interface{})

// 	accessKeyId := details["access_key_id"].(string)
// 	secretAccessKey := details["secret_access_key"].(string)

// 	return cloudconnectionapi.NewAWSAccessKeyDetails(accessKeyId, secretAccessKey, cloudconnectionapi.ACCESS_KEY)
// }

// func azureRequiredAttributesCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
// 	type_ := d.Get("type").(string)

// 	if type_ != string(cloudconnectionapi.AZURE) {
// 		return nil
// 	}

// 	err := checkRequiredNotRequired(d, type_)
// 	if err != nil {
// 		return err
// 	}

// 	err = checkRequiredNotRequired(d, type_)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
