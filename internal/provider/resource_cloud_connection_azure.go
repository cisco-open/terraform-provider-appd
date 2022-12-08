package provider

import (
	"context"
	"log"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudConnectionAzure() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceCloudConnectionAzureCreate,
		ReadContext:   resourceCloudConnectionAzureRead,
		UpdateContext: resourceCloudConnectionAzureUpdate,
		DeleteContext: resourceCloudConnectionAzureDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,

		Schema: appendSchema(getCommonCloudConnectionSchema(), map[string]*schema.Schema{
			"details": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_id": {
							Type:        schema.TypeString,
							Description: "Client IDs, also known as Application IDs, are long-term credentials for an Azure user, or account root user. The Client ID is one of three properties needed to authenticate to Azure, the other two being Client Secret and Tenant (Directory) ID",
							Required:    true,
						},
						"client_secret": {
							Type:        schema.TypeString,
							Description: "A Client Secret allows an Azure application to provide its identity when requesting an access token. The Client Secret is one of three properties needed to authenticate to Azure, the other two being Client ID (Application ID) and Tenant (Directory) ID",
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
				},
			},
		}),
	}
}

func resourceCloudConnectionAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionRequest := cloudconnectionapi.ConnectionRequest{}
	connectionRequest.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AZURE))

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
		connectionRequestDetails := expandCloudConnectionAzureCreateDetails(v, d)
		connectionRequest.SetDetails(connectionRequestDetails)
	}

	resp, _, err := apiClient.ConnectionsApi.CreateConnection(myctx).ConnectionRequest(connectionRequest).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionAzureRead(ctx, d, m)
}

func resourceCloudConnectionAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()

	resp, _, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.GetId())
	d.Set("display_name", resp.GetDisplayName())
	d.Set("description", resp.GetDescription())
	d.Set("state", resp.GetState())
	d.Set("state_message", resp.StateMessage)

	var clientSecret string = ""
	if v, ok := d.GetOk("details"); ok {
		details := v.(*schema.Set).List()[0].(map[string]interface{})
		clientSecret = details["client_secret"].(string)
	} else {
		clientSecret = resp.GetDetails().AzureDetails.ClientSecret
	}

	detailsSet := schema.NewSet(schema.HashResource(resourceCloudConnectionAzure().Schema["details"].Elem.(*schema.Resource)), []interface{}{})
	detailsSet.Add(map[string]interface{}{
		"client_id":       resp.GetDetails().AzureDetails.ClientId,
		"client_secret":   clientSecret,
		"tenant_id":       resp.GetDetails().AzureDetails.TenantId,
		"subscription_id": resp.GetDetails().AzureDetails.SubscriptionId,
	})
	log.Println("Details______", detailsSet)
	d.Set("details", detailsSet)
	d.Set("configuration_id", resp.GetConfigurationId())
	d.Set("created_at", resp.GetCreatedAt().Format("2022-12-06T06:58:01.615Z"))
	d.Set("updated_at", resp.GetUpdatedAt().Format("2022-12-06T06:58:01.615Z"))

	return nil
}
func resourceCloudConnectionAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionUpdate := cloudconnectionapi.ConnectionUpdate{}
	// connectionUpdate.SetType(cloudconnectionapi.ProviderType(cloudconnectionapi.AZURE))

	// if ok := d.HasChanges(tenant_id,)
	if v, ok := d.GetOk("display_name"); ok {
		connectionUpdate.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		connectionUpdate.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("configuration_id"); ok {
		connectionUpdate.SetConfigurationId(v.(string))

		if v, ok := d.GetOk("state"); ok {
			connectionUpdate.SetState(v.(string))
		}
	}

	if v, ok := d.GetOk("details"); ok {
		connectionUpdateDetails := expandCloudConnectionAzureUpdateDetails(v, d)
		connectionUpdate.SetDetails(connectionUpdateDetails)
	}

	resp, _, err := apiClient.ConnectionsApi.UpdateConnection(myctx, d.Id()).ConnectionUpdate(connectionUpdate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionAzureRead(ctx, d, m)
}
func resourceCloudConnectionAzureDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()
	_,_, err := apiClient.ConnectionsApi.DeleteConnection(myctx, connectionId).Execute()

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func expandCloudConnectionAzureCreateDetails(v interface{}, d *schema.ResourceData) cloudconnectionapi.ConnectionRequestDetails {
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
	// tenantId := details["tenant_id"].(string)
	// subscriptionId := details["subscription_id"].(string)

	connectionUpdateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf2{}
	connectionUpdateDetails.SetClientId(clientId)
	connectionUpdateDetails.SetClientSecret(clientSecret)

	updateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOf2AsConnectionUpdateDetails(&connectionUpdateDetails)
	return updateDetails
}
