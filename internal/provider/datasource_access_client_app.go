// Copyright 2023 Cisco Systems, Inc.
//
// Licensed under the MPL License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.mozilla.org/en-US/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
