package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccessClientApp() *schema.Resource {
	dsSchema := resourceSchemaToDataSourceSchema(resourceAccessClientApp().Schema)

	delete(dsSchema, "client_secret")
	delete(dsSchema, "rotate_secret")
	delete(dsSchema, "revoke_previous_secret_in")
	delete(dsSchema, "revoke_now")

	dsSchema["client_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The Client ID of the Service Principal",
		Required:    true,
	}

	return &schema.Resource{
		ReadContext: dataSourceAccessClientAppRead,
		Schema:      dsSchema,
	}
}

func dataSourceAccessClientAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeApplicationPrincipalManagementClient(m)

	clientId := d.Get("client_id").(string)

	resp, httpResp, err := apiClient.ServicesApi.GetServiceClientById(myctx, clientId).Execute()

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
