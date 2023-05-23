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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCloudConnectionAWS() *schema.Resource {
	dsSchema := resourceSchemaToDataSourceSchema(resourceCloudConnectionAWS().Schema)

	dsSchema["connection_id"] = &schema.Schema{
		Type:             schema.TypeString,
		Description:      "The Connection ID of the AWS Connection",
		Required:         true,
		ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
	}

	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionAWSRead,
		Schema:      dsSchema,
	}
}

func dataSourceCloudConnectionAWSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connectionId := d.Get("connection_id").(string)

	myctx, _, apiClient := initializeCloudConnectionClient(m)

	respConnection, httpRespConnection, err := apiClient.ConnectionsApi.GetConnection(myctx, connectionId).Execute()
	if err != nil {
		return errRespToDiag(err, httpRespConnection)
	}
	d.SetId(respConnection.Id)

	d.Set("connection_details", flattenCloudConnectionAWSDetails(respConnection, d))

	flattenCloudConnectionCommons(respConnection, d)

	configurationId := respConnection.ConfigurationId
	d.Set("configuration_id", configurationId)
	respConfiguration, httpRespConfiguration, err := apiClient.ConfigurationsApi.GetConfiguration(myctx, *configurationId).Execute()
	if err != nil {
		return errRespToDiag(err, httpRespConfiguration)
	}
	flattenCloudConnectionConfigurationCommonsDetails(respConfiguration, d, "AWS")

	return nil
}
