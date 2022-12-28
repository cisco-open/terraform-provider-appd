package provider

import (
	"context"
	"time"

	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudConnectionAWSRoleAttachment() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceCloudConnectionAWSRoleAttachmentCreate,
		ReadContext:   resourceCloudConnectionAWSRoleAttachmentRead,
		UpdateContext: resourceCloudConnectionAWSRoleAttachmentUpdate,
		DeleteContext: resourceCloudConnectionAWSRoleAttachmentDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"connection_id": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
			},

			"role_name": {
				Type:        schema.TypeString,
				Description: "Role name for AWS iam role",
				Required:    true,
			},
		},
	}
}

func resourceCloudConnectionAWSRoleAttachmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	time.Sleep(10 * time.Second)
	return resourceCloudConnectionAWSRoleAttachmentUpdate(ctx, d, m)
}

func resourceCloudConnectionAWSRoleAttachmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionId := d.Id()

	resp, httpResp, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		if httpResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return errRespToDiag(err, httpResp)
	}

	d.Set("connection_id", resp.GetId())
	d.Set("role_name", resp.GetDetails().AWSConnectionResponseDetails.RoleDelegationConnectionResponseDetails.RoleName)

	return nil
}

func resourceCloudConnectionAWSRoleAttachmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudConnectionClient(m)

	connectionUpdate := cloudconnectionapi.ConnectionUpdate{}

	connectionId := d.Get("connection_id").(string)

	if d.HasChange("role_name") {
		connectionUpdateDetailsOneOf := cloudconnectionapi.NewConnectionUpdateDetailsOneOf(d.Get("role_name").(string))
		connectionUpdateDetails := cloudconnectionapi.ConnectionUpdateDetailsOneOfAsConnectionUpdateDetails(connectionUpdateDetailsOneOf)
		connectionUpdate.SetDetails(connectionUpdateDetails)
	}

	resp, httpResp, err := apiClient.ConnectionsApi.UpdateConnection(myctx, connectionId).ConnectionUpdate(connectionUpdate).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	d.SetId(resp.Id)

	return resourceCloudConnectionAWSRoleAttachmentRead(ctx, d, m)
}

func resourceCloudConnectionAWSRoleAttachmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	diags := diag.Diagnostics{}
	diagWarn := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Warning",
		Detail:   `The Resource has been deleted but the role name cannot be detached from connection. Unless the connection itself is deleted.`,
	}
	diags = append(diags, diagWarn)
	d.SetId("")
	return diags
}
