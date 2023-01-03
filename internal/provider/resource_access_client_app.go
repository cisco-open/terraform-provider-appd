package provider

import (
	"context"
	"fmt"

	applicationprincipalmanagement "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/applicationprincipalmanagement"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var maxRetries = 1

func resourceAccessClientApp() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceAccessClientAppCreate,
		ReadContext:   resourceAccessClientAppRead,
		UpdateContext: resourceAccessClientAppUpdate,
		DeleteContext: resourceAccessClientAppDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceAccessClientAppImport,
		},

		SchemaVersion: 1,

		CustomizeDiff: customdiff.All(
			func(ctx context.Context, d *schema.ResourceDiff, i interface{}) error {
				val, rotateSecretIsPreset := d.GetOk("rotate_secret")
				_, revokeTimeoutIsPresent := d.GetOk("revoke_previous_secret_in")

				errString := "revoke_previous_secret_in can only be used when rotate_secret is set to true"
				err := fmt.Errorf(errString)

				if !rotateSecretIsPreset && revokeTimeoutIsPresent {
					return err
				}

				if rotateSecretIsPreset && !val.(bool) && revokeTimeoutIsPresent {
					return err
				}

				return nil
			},

			customdiff.If(
				func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) bool {
					_, ok := d.GetOk("rotate_secret")
					return ok
				},
				func(ctx context.Context, d *schema.ResourceDiff, i interface{}) error {
					_, ok := d.GetOk("revoke_previous_secret_in")
					if !ok {
						return fmt.Errorf("revoke_previous_secret_in must be set with rotate_secret")
					}

					return nil
				},
			),
		),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Description: "The display name for the client.",
				Required:    true,
			},
			"description": {
				Type:             schema.TypeString,
				Description:      "A user provided description of the client.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Required:         true,
			},
			"auth_type": {
				Type:             schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"client_secret_basic", "client_secret_post"}, true)),
				Description:      "Supported authentication methods used to request oAuth tokens: `client_secret_basic` - The client credentials will be sent in the authorization header `client_secret_post` - The client credentials will be sent in the request body.",
				Required:         true,
			},
			"has_rotated_secrets": {
				Type:        schema.TypeBool,
				Description: "Indicates if the client has rotated secrets. Rotated client secrets can be revoked.",
				Computed:    true,
			},
			"created_at": {
				Type:        schema.TypeString,
				Description: "The RFC3339 timestamp when the client was created",
				Computed:    true,
			},
			"updated_at": {
				Type:        schema.TypeString,
				Description: "The RFC3339 timestamp when the client was last updated.",
				Computed:    true,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Description: "The client's secret, used to authenticate during an oAuth token request",
				Sensitive:   true,
				Computed:    true,
			},

			"rotate_secret": {
				Type:        schema.TypeBool,
				Description: "Rotates the client secret of the specified service client.",
				Optional:    true,
				Default:     false,
			},

			"revoke_previous_secret_in": {
				Type:             schema.TypeString,
				Description:      "Time duration of how long the previous secret should be active for. Acceptable values are `NOW`, `1D`, `3D`, `7D` and `30D`. Must be set when rotating a secret with rotate_secret.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"NOW", "1D", "3D", "7D", "30D"}, false)),
				Optional:         true,
			},

			"rotated_secret_expires_at": {
				Type:        schema.TypeString,
				Description: "The RFC3339 timestamp when the rotated client secret will expire.",
				Computed:    true,
			},

			"revoke_now": {
				Type:        schema.TypeBool,
				Description: "Revokes all the rotated client secrets of the specified client.",
				Optional:    true,
				Default:     false,
				ValidateDiagFunc: validation.ToDiagFunc(func(i interface{}, s string) ([]string, []error) {
					if i.(bool) {
						return []string{"All previous secrets will be revoked with this action"}, nil
					}

					return nil, nil
				}),
			},
		},
	}
}

func flattenAccessClientApp(d *schema.ResourceData, resp *applicationprincipalmanagement.ServiceClientResponse) {
	d.SetId(resp.GetId())
	d.Set("display_name", resp.GetDisplayName())
	d.Set("description", resp.GetDescription())
	d.Set("auth_type", resp.GetAuthType())
	d.Set("has_rotated_secrets", resp.GetHasRotatedSecrets())
	d.Set("created_at", resp.GetCreatedAt())
	d.Set("updated_at", resp.GetUpdatedAt())
}

func resourceAccessClientAppImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	myctx, _, apiClient := initializeApplicationPrincipalManagementClient(m)

	clientId := d.Id()

	resp, _, err := apiClient.ServicesApi.GetServiceClientById(myctx, clientId).Execute()
	if err != nil {
		return nil, err
	}

	flattenAccessClientApp(d, resp)

	return []*schema.ResourceData{d}, nil
}

func resourceAccessClientAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeApplicationPrincipalManagementClient(m)

	serviceClientRequest := applicationprincipalmanagement.ServiceClientRequest{}

	if v, ok := d.GetOk("display_name"); ok {
		serviceClientRequest.SetDisplayName(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		serviceClientRequest.SetDescription(v.(string))
	}

	if v, ok := d.GetOk("auth_type"); ok {
		serviceClientRequest.SetAuthType(v.(string))
	}

	resp, httpResp, err := apiClient.ServicesApi.CreateServiceClient(myctx).ServiceClientRequest(serviceClientRequest).Execute()
	if err != nil {
		if httpResp.StatusCode == 504 && maxRetries != 0 {
			maxRetries -= 1
			return resourceAccessClientAppCreate(ctx, d, m)
		}
		return errRespToDiag(err, httpResp)
	}

	d.SetId(resp.GetId())
	d.Set("client_secret", resp.GetClientSecret())

	rotateSecret(d, m)
	revokeSecret(d, m)

	return resourceAccessClientAppRead(ctx, d, m)
}

func resourceAccessClientAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeApplicationPrincipalManagementClient(m)

	connectionId := d.Id()

	resp, httpResp, err := apiClient.ServicesApi.GetServiceClientById(myctx, connectionId).Execute()
	if err != nil {
		if httpResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return errRespToDiag(err, httpResp)
	}

	flattenAccessClientApp(d, resp)

	return nil
}

func resourceAccessClientAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeApplicationPrincipalManagementClient(m)

	if d.HasChange("display_name") || d.HasChange("description") || d.HasChange("auth_type") {
		serviceClientRequest := applicationprincipalmanagement.NewServiceClientRequest()

		serviceClientRequest.SetDisplayName(d.Get("display_name").(string))
		serviceClientRequest.SetDescription(d.Get("description").(string))
		serviceClientRequest.SetAuthType(d.Get("auth_type").(string))

		r, httpResp, err := apiClient.ServicesApi.UpdateServiceClient(myctx, d.Id()).ServiceClientRequest(*serviceClientRequest).Execute()
		if err != nil {
			return errRespToDiag(err, httpResp)
		}

		flattenAccessClientApp(d, r)
	}

	rotateSecret(d, m)

	revokeSecret(d, m)

	// after rotate, get updatedAt, hasRotated etc..
	return resourceAccessClientAppRead(ctx, d, m)
}

func resourceAccessClientAppDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeApplicationPrincipalManagementClient(m)

	clientId := d.Id()

	httpResp, err := apiClient.ServicesApi.DeleteServiceClient(myctx, clientId).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	return nil
}

func rotateSecret(d *schema.ResourceData, m interface{}) {
	myctx, _, apiClient := initializeApplicationPrincipalManagementClient(m)

	val, ok := d.GetOk("rotate_secret")
	if val.(bool) && ok {
		rotationRequest := getRotationRequest(d.Get("revoke_previous_secret_in").(string))

		resp, httpResp, err := apiClient.ServicesApi.RotateServiceClientSecret(myctx, d.Id()).RotationRequest(rotationRequest).Execute()
		if err != nil {
			errRespToDiag(err, httpResp)
		}

		// d.SetId(resp.GetClientId())
		d.Set("client_secret", resp.GetClientSecret())
		d.Set("rotated_secret_expires_at", resp.GetRotatedSecretExpiresAt())
		d.Set("rotate_secret", false)
		d.Set("revoke_previous_secret_in", "")
	}
}

func revokeSecret(d *schema.ResourceData, m interface{}) {
	myctx, _, apiClient := initializeApplicationPrincipalManagementClient(m)

	val, ok := d.GetOk("revoke_now")
	if ok && val.(bool) {
		// resp only contains status and message acknowledging that secrets have been revoked
		// thus no need to flatten here
		_, httpResp, err := apiClient.ServicesApi.RevokeServiceClientSecret(myctx, d.Id()).Execute()
		if err != nil {
			errRespToDiag(err, httpResp)
		}

		d.Set("revoke_now", false)
	}
}

func getRotationRequest(revokePreviousIn string) applicationprincipalmanagement.RotationRequest {
	rotationRequest := *applicationprincipalmanagement.NewRotationRequest()

	revokePreviousSecretIn := revokePreviousIn
	if revokePreviousSecretIn == "NOW" {
		revokePreviousSecretIn = "P0D"
	} else {
		revokePreviousSecretIn = "P" + revokePreviousSecretIn
	}

	rotationRequest.SetRevokeRotatedAfter(revokePreviousSecretIn)

	return rotationRequest
}
