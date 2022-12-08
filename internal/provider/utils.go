package provider

import (
	"context"
	"fmt"

	client "github.com/aniketk-crest/appdynamicscloud-go-client"
	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
	cloudqueryapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudquery"
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

	myctx := context.WithValue(config.ctx, client.ContextServerIndex, client.SERVER_INDEX_CLOUD_QUERY)

	return myctx, configuration, apiClient
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

func checkRequiredNotRequired(d *schema.ResourceDiff, type_ string) error {
	requiredAttributes := map[string][]string{
		string(cloudconnectionapi.ACCESS_KEY):      {"access_key_id", "secret_access_key"},
		string(cloudconnectionapi.ROLE_DELEGATION): {"account_id"},
	}

	// TODO: Use ConflictsWith
	notRequiredAttributes := map[string][]string{
		string(cloudconnectionapi.ACCESS_KEY):      {"account_id"},
		string(cloudconnectionapi.ROLE_DELEGATION): {"access_key_id", "secret_access_key"},
	}

	details := d.Get("details").(*schema.Set).List()[0].(map[string]interface{})

	for _, k := range requiredAttributes[type_] {
		if details[k] == "" {
			return fmt.Errorf("%s is required with %s", k, type_)
		}
	}

	for _, k := range notRequiredAttributes[type_] {
		if v, ok := details[k]; ok || v != "" {
			return fmt.Errorf("%s should not be used with %s", k, type_)
		}
	}

	return nil
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
