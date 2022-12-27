package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	client "github.com/aniketk-crest/appdynamicscloud-go-client"
	cloudappprincipalmgmtapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/applicationprincipalmanagement"
	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	cloudqueryapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudquery"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func initializeCloudConnectionClient(m interface{}) (context.Context, *client.Configuration, *cloudconnectionapi.APIClient) {
	config := m.(config)

	configuration := config.configuration
	apiClient := cloudconnectionapi.NewAPIClient(configuration)

	myctx := context.WithValue(config.ctx, client.ContextServerIndex, client.SERVER_INDEX_CLOUD_CONNECTION)

	return myctx, configuration, apiClient
}

func initializeCloudQueryClient(m interface{}) (context.Context, *client.Configuration, *cloudqueryapi.APIClient) {
	config := m.(config)

	configuration := config.configuration
	apiClient := cloudqueryapi.NewAPIClient(configuration)

	myCtx := context.WithValue(config.ctx, client.ContextServerIndex, client.SERVER_INDEX_CLOUD_QUERY)

	return myCtx, configuration, apiClient
}

func initializeApplicationPrincipalManagementClient(m interface{}) (context.Context, *client.Configuration, *cloudappprincipalmgmtapi.APIClient) {
	config := m.(config)

	configuration := config.configuration
	apiClient := cloudappprincipalmgmtapi.NewAPIClient(configuration)

	myCtx := context.WithValue(config.ctx, client.ContextServerIndex, client.SERVER_INDEX_APPLICATION_PRINCIPAL_MANAGEMENT)

	return myCtx, configuration, apiClient
}

func resourceSchemaToDataSourceSchema(resourceSchema map[string]*schema.Schema) map[string]*schema.Schema {
	dataSourceSchema := make(map[string]*schema.Schema, len(resourceSchema))

	for k, attributeSchema := range resourceSchema {
		dataSourceAttributeSchema := &schema.Schema{
			Type:        attributeSchema.Type,
			Description: attributeSchema.Description,
			Computed:    true,
		}

		switch attributeSchema.Type {
		case schema.TypeSet:
			fallthrough
		case schema.TypeList:
			if elem, ok := attributeSchema.Elem.(*schema.Resource); ok {
				dataSourceAttributeSchema.Elem = &schema.Resource{
					Schema: resourceSchemaToDataSourceSchema(elem.Schema),
				}
			} else {
				dataSourceAttributeSchema.Elem = attributeSchema.Elem
			}
		}

		dataSourceSchema[k] = dataSourceAttributeSchema
	}

	return dataSourceSchema
}

// func addRequiredFieldsToSchema(schema map[string]*schema.Schema, key string) {
// 	schema[key].Computed = false
// 	schema[key].Required = true
// }

// func addOptionalFieldsToSchema(schema map[string]*schema.Schema, key string) {
// 	schema[key].Computed = false
// 	schema[key].Optional = true
// }

func appendSchema(a, b map[string]*schema.Schema) map[string]*schema.Schema {
	c := map[string]*schema.Schema{}

	for k, v := range a {
		c[k] = v
	}

	for k, v := range b {
		c[k] = v
	}

	return c
}

func appendSchemas(maps ...map[string]*schema.Schema) (result map[string]*schema.Schema) {
	result = make(map[string]*schema.Schema)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func httpRespToMap(resp *http.Response) (map[string]interface{}, bool) {
	var m map[string]interface{}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false
	}

	json.Unmarshal(body, &m)

	return m, true
}

func utcTimeToString(t time.Time) string {
	v, _ := t.UTC().MarshalText()
	return string(v)
}

func errRespToDiag(err error, errResp *http.Response) diag.Diagnostics {
	m, ok := httpRespToMap(errResp)
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

// func singleSetToMap(v interface{}) (map[string]interface{}, bool) {
// 	schemaSet := v.(*schema.Set).List()

// 	if len(schemaSet) > 0 {
// 		return schemaSet[0].(map[string]interface{}), true
// 	}

// 	return nil, false
// }

func singleListToMap(v interface{}) (map[string]interface{}, bool) {
	if len(v.([]interface{})) == 0 {
		return nil, false
	}

	return v.([]interface{})[0].(map[string]interface{}), true
}

func toSliceString(data []interface{}) []string {
	s := make([]string, 0)
	for _, v := range data {
		s = append(s, fmt.Sprint(v))
	}
	return s
}

func toSliceInterface(data []string) []interface{} {
	s := make([]interface{}, 0)
	for _, v := range data {
		var v1 interface{}
		v1 = v
		s = append(s, v1)
	}
	return s
}
