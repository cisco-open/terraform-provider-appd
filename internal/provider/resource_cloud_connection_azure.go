package provider

// import (
// 	"context"
// 	"fmt"

// 	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
// )

// func resourceCloudConnectionAzure() *schema.Resource {

// 	return &schema.Resource{
// 		CreateContext: resourceCloudConnectionAzureCreate,
// 		ReadContext:   resourceCloudConnectionAzureRead,
// 		UpdateContext: resourceCloudConnectionAzureUpdate,
// 		DeleteContext: resourceCloudConnectionDelete,

// 		Importer: &schema.ResourceImporter{
// 			StateContext: resourceCloudConnectionAzureImport,
// 		},

// 		SchemaVersion: 1,

// 		Schema: appendSchema(getCommonCloudConnectionSchema(), map[string]*schema.Schema{
// 			"details": {
// 				Type:     schema.TypeList,
// 				Required: true,
// 				MaxItems: 1,
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						"client_id": {
// 							Type:             schema.TypeString,
// 							ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
// 							Description:      "Client IDs, also known as Application IDs, are long-term credentials for an Azure user, or account root user. The Client ID is one of three properties needed to authenticate to Azure, the other two being Client Secret and Tenant (Directory) ID",
// 							Required:         true,
// 						},
// 						"client_secret": {
// 							Type:        schema.TypeString,
// 							Description: "A Client Secret allows an Azure application to provide its identity when requesting an access token. The Client Secret is one of three properties needed to authenticate to Azure, the other two being Client ID (Application ID) and Tenant (Directory) ID",
// 							Sensitive:   true,
// 							Required:    true,
// 						},
// 						"tenant_id": {
// 							Type:             schema.TypeString,
// 							ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
// 							Description:      "The Azure AD Tenant (Directory) IDis one of three properties needed to authenticate to Azure. The other two are Client Secret and Client ID (Application ID).",
// 							Required:         true,
// 							ForceNew:         true,
// 						},
// 						"subscription_id": {
// 							Type:             schema.TypeString,
// 							ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
// 							Description:      "Specify a GUID Subscription ID to monitor. If monitoring all subscriptions, do not specify a Subscription ID.",
// 							Required:         true,
// 							ForceNew:         true,
// 						},
// 					},
// 				},
// 			},
// 		}),
// 	}
// }

// func resourceCloudConnectionAzureImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
// 	myctx, _, apiClient := initializeCloudConnectionClient(m)

// 	connectionId := d.Id()

// 	resp, _, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
// 	if err != nil {
// 		return nil, err
// 	}

// 	d.Set("details", flattenCloudConnectionAzureDetails(resp, d))

// 	flattenCloudConnectionCommons(resp, d)

// 	return []*schema.ResourceData{d}, nil
// }

// func resourceCloudConnectionAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	myctx, _, apiClient := initializeCloudConnectionClient(m)

// 	connectionRequest := cloudconnectionapi.ConnectionRequest{}
// 	connectionRequest.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AZURE))

// 	err := checkStateConfiguration(ctx, d, m)
// 	if err != nil {
// 		return diag.FromErr(err)
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
// 		connectionRequestDetails := expandCloudConnectionAzureDetails(v, d)
// 		connectionRequest.SetDetails(connectionRequestDetails)
// 	}

// 	resp, httpResp, err := apiClient.ConnectionsApi.CreateConnection(myctx).ConnectionRequest(connectionRequest).Execute()
// 	if err != nil {
// 		return errRespToDiag(err, httpResp)
// 	}

// 	d.SetId(resp.Id)

// 	return resourceCloudConnectionAzureRead(ctx, d, m)
// }

// func resourceCloudConnectionAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	myctx, _, apiClient := initializeCloudConnectionClient(m)

// 	connectionId := d.Id()

// 	resp, httpResp, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
// 	if err != nil {
// 		if httpResp.StatusCode == 404 {
// 			d.SetId("")
// 			return nil
// 		}
// 		return errRespToDiag(err, httpResp)
// 	}

// 	d.Set("details", flattenCloudConnectionAzureDetails(resp, d))

// 	flattenCloudConnectionCommons(resp, d)

// 	return nil
// }

// func resourceCloudConnectionAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	myctx, _, apiClient := initializeCloudConnectionClient(m)

// 	connectionUpdate := cloudconnectionapi.ConnectionUpdate{}

// 	err := checkStateConfiguration(ctx, d, m)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	if d.HasChange("display_name") {
// 		connectionUpdate.SetDisplayName(d.Get("display_name").(string))
// 	}

// 	if d.HasChange("description") {
// 		connectionUpdate.SetDescription(d.Get("description").(string))
// 	}

// 	if d.HasChange("configuration_id") {
// 		connectionUpdate.SetConfigurationId(d.Get("configuration_id").(string))
// 	}

// 	if d.HasChange("state") {
// 		connectionUpdate.SetState(d.Get("state").(string))
// 	}

// 	if d.HasChange("details") {
// 		connectionUpdateDetails, err := expandCloudConnectionAzureUpdateDetails(d.Get("details"), d)
// 		if err != nil {
// 			return diag.FromErr(err)
// 		}

// 		connectionUpdate.SetDetails(connectionUpdateDetails)
// 	}

// 	resp, httpResp, err := apiClient.ConnectionsApi.UpdateConnection(myctx, d.Id()).ConnectionUpdate(connectionUpdate).Execute()
// 	if err != nil {
// 		return errRespToDiag(err, httpResp)
// 	}

// 	d.SetId(resp.Id)

// 	return resourceCloudConnectionAzureRead(ctx, d, m)
// }

// func expandCloudConnectionAzureDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConnectionRequestDetails {
// 	connectionRequestDetails := cloudconnectionapi.ConnectionRequestDetails{}

// 	details, _ := singleListToMap(v)

// 	clientId := details["client_id"].(string)
// 	clientSecret := details["client_secret"].(string)
// 	tenantId := details["tenant_id"].(string)
// 	subscriptionId := details["subscription_id"].(string)

// 	connectionRequestDetails.AzureDetails = cloudconnectionapi.NewAzureDetails(clientId, clientSecret, tenantId, subscriptionId)

// 	return connectionRequestDetails
// }

// func expandCloudConnectionAzureUpdateDetails(v interface{}, d *schema.ResourceData) (cloudconnectionapi.ConnectionUpdateDetails, error) {
// 	connectionDetails := expandCloudConnectionAzureDetails(v, d).AzureDetails

// 	connectionUpdateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf2{}
// 	connectionUpdateDetails.SetClientId(connectionDetails.ClientId)
// 	connectionUpdateDetails.SetClientSecret(connectionDetails.ClientSecret)

// 	updateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf2AsConnectionUpdateDetails(&connectionUpdateDetails)
// 	return updateDetails, nil
// }

// func flattenCloudConnectionAzureDetails(resp *cloudconnectionapi.ConnectionResponse, d *schema.ResourceData) interface{} {
// 	var clientSecret string = ""

// 	// For datasource, if details block does not exists
// 	// set client secret as recieved from response.
// 	if v, ok := d.GetOk("details"); ok {
// 		details, _ := singleListToMap(v)
// 		clientSecret = details["client_secret"].(string)
// 	} else {
// 		clientSecret = resp.GetDetails().AzureDetails.ClientSecret
// 	}

// 	detailsList := []interface{}{}
// 	detailsList = append(detailsList, map[string]interface{}{
// 		"client_id":       resp.GetDetails().AzureDetails.ClientId,
// 		"client_secret":   clientSecret,
// 		"tenant_id":       resp.GetDetails().AzureDetails.TenantId,
// 		"subscription_id": resp.GetDetails().AzureDetails.SubscriptionId,
// 	})

// 	return detailsList
// }

// func checkStateConfiguration(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
// 	valueState, isPresentState := d.GetOk("state")
// 	_, isPresentConfigurationId := d.GetOk("configuration_id")

// 	// State cannot be specified without configuration id
// 	if !isPresentConfigurationId && isPresentState && (valueState == "ACTIVE" || valueState == "INACTIVE") {
// 		return fmt.Errorf("the configuration ID must be provided to assign a connection state")
// 	}

// 	return nil
// }
