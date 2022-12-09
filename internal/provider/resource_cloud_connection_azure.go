package provider

import (
	"context"
	"fmt"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudConnectionAzure() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceCloudConnectionAzureCreate,
		ReadContext:   resourceCloudConnectionAzureRead,
		UpdateContext: resourceCloudConnectionAzureUpdate,
		DeleteContext: resourceCloudConnectionDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,

		Schema: appendSchema(getCommonCloudConnectionSchema(), map[string]*schema.Schema{
			"details": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem:     detailsSchemaAzure(),
			},
		}),
	}
}

func detailsSchemaAzure() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Description: "Client IDs, also known as Application IDs, are long-term credentials for an Azure user, or account root user. The Client ID is one of three properties needed to authenticate to Azure, the other two being Client Secret and Tenant (Directory) ID",
				Required:    true,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Description: "A Client Secret allows an Azure application to provide its identity when requesting an access token. The Client Secret is one of three properties needed to authenticate to Azure, the other two being Client ID (Application ID) and Tenant (Directory) ID",
				Sensitive:   true,
				Required:    true,
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Description: "The Azure AD Tenant (Directory) IDis one of three properties needed to authenticate to Azure. The other two are Client Secret and Client ID (Application ID).",
				Required:    true,
			},
			"subscription_id": {
				Type:        schema.TypeString,
				Description: "Specify a GUID Subscription ID to monitor. If monitoring all subscriptions, do not specify a Subscription ID.",
				Required:    true,
			},
		},
	}
}

func resourceCloudConnectionAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionRequest := cloudconnectionapi.ConnectionRequest{}
	connectionRequest.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AZURE))

	err := conditionalStateCheck(ctx, d, m)
	if err != nil {
		return diag.FromErr(err)
	}

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
		connectionRequestDetails := expandCloudConnectionAzureDetails(v, d)
		connectionRequest.SetDetails(connectionRequestDetails)
	}

	resp, httpResp, err := apiClient.ConnectionsApi.CreateConnection(myctx).ConnectionRequest(connectionRequest).Execute()

	// TODO: this must be common
	if err != nil {
		m, ok := httpRespToMap(httpResp)
		if !ok {
			return diag.FromErr(err)
		}

		title, isPresentTitle := m["title"]
		detail, isPresentDetail := m["detail"]

		if !isPresentTitle {
			return diag.FromErr(err)
		}

		d := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  title.(string),
		}

		if isPresentDetail {
			d.Detail = detail.(string)
		}

		return diag.Diagnostics{d}
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionAzureRead(ctx, d, m)
}

func resourceCloudConnectionAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()

	resp, _, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		// TODO: httpRespToDiag
		return diag.FromErr(err)
	}

	var clientSecret string = ""
	if v, ok := d.GetOk("details"); ok {
		details := v.(*schema.Set).List()[0].(map[string]interface{})
		clientSecret = details["client_secret"].(string)
	} else {
		clientSecret = resp.GetDetails().AzureDetails.ClientSecret
	}

	detailsSet := schema.NewSet(schema.HashResource(detailsSchemaAzure()), []interface{}{})
	detailsSet.Add(map[string]interface{}{
		"client_id":       resp.GetDetails().AzureDetails.ClientId,
		"client_secret":   clientSecret,
		"tenant_id":       resp.GetDetails().AzureDetails.TenantId,
		"subscription_id": resp.GetDetails().AzureDetails.SubscriptionId,
	})

	d.Set("details", detailsSet)

	flattenCloudConnectionCommons(resp, d)

	return nil
}

func resourceCloudConnectionAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionUpdate := cloudconnectionapi.ConnectionUpdate{}

	err := conditionalStateCheck(ctx, d, m)
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
		connectionUpdateDetails := expandCloudConnectionAzureUpdateDetails(d.Get("details"), d)
		connectionUpdate.SetDetails(connectionUpdateDetails)
	}

	resp, _, err := apiClient.ConnectionsApi.UpdateConnection(myctx, d.Id()).ConnectionUpdate(connectionUpdate).Execute()
	if err != nil {
		// TODO: httpResp to Diag
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionAzureRead(ctx, d, m)
}

func expandCloudConnectionAzureDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConnectionRequestDetails {
	connectionRequestDetails := cloudconnectionapi.ConnectionRequestDetails{}

	details := v.(*schema.Set).List()[0].(map[string]interface{})

	clientId := details["client_id"].(string)
	clientSecret := details["client_secret"].(string)
	tenantId := details["tenant_id"].(string)
	subscriptionId := details["subscription_id"].(string)

	connectionRequestDetails.AzureDetails = cloudconnectionapi.NewAzureDetails(clientId, clientSecret, tenantId, subscriptionId)

	return connectionRequestDetails
}

func expandCloudConnectionAzureUpdateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConnectionUpdateDetails {
	details := v.(*schema.Set).List()[0].(map[string]interface{})

	clientId := details["client_id"].(string)
	clientSecret := details["client_secret"].(string)

	connectionUpdateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf2{}
	connectionUpdateDetails.SetClientId(clientId)
	connectionUpdateDetails.SetClientSecret(clientSecret)

	updateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf2AsConnectionUpdateDetails(&connectionUpdateDetails)
	return updateDetails
}

func conditionalStateCheck(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	_, isPresentState := d.GetOk("state")
	_, isPresentConfigurationId := d.GetOk("configuration_id")

	if !isPresentConfigurationId && isPresentState {
		return fmt.Errorf("the configuration ID must be provided to assign a connection state")
	}

	return nil
}
