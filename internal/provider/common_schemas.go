package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// The following methods will be used outside of this file
func getCommonCloudConnectionSchema() map[string]*schema.Schema {
	return appendSchema(
		cloudConnectionSchema(),
		cloudConnectionSchemaExtras())
}

// The following methods are helper methods to the methods defined
// above and defines the actual schema
func cloudConnectionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Type:        schema.TypeString,
			Description: "Name of the connection or configuration",
			Required:    true,
		},
		"description": {
			Type:        schema.TypeString,
			Description: "Description for this connection or configuration",
			Optional:    true,
		},

		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func cloudConnectionSchemaExtras() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"state": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"ACTIVE", "INACTIVE"}, true),
			Description:  "Connection state. This can only be used if configuration_id is specified. Possible values: [\"ACTIVE\", \"INACTIVE\"]",
			Optional:     true,
			Computed:     true,
		},
		"state_message": {
			Type:        schema.TypeString,
			Description: "Connection state message",
			Computed:    true,
		},
		"configuration_id": {
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
			Optional:     true,
		},
	}
}

func getCloudConnectionConfigurationAWSSchema() map[string]*schema.Schema {
	return cloudConnectionConfigurationSchema("AWS")
}

func getCloudConnectionConfigurationAzureSchema() map[string]*schema.Schema {
	return cloudConnectionConfigurationSchema("AZURE")
}

func cloudConnectionConfigurationSchema(provider string) map[string]*schema.Schema {
	rootSchema := cloudConnectionSchema()
	detailsSchema := cloudConnectionConfigurationDetails()
	servicesSchema := cloudConnectionConfigurationDetailsServices()

	servicesSchema["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}

	detailsSchema["regions"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Geographic locations used to fetch metrics",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	detailsSchema["services"] = &schema.Schema{
		Type:        schema.TypeList,
		Description: "services for which we will fetch metrics",
		Optional:    true,
		Computed:    true,
		Elem: &schema.Resource{
			Schema: servicesSchema,
		},
	}

	if provider == "AZURE" {
		detailsSchema["resource_groups"] = &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Azure Resource groups used to fetch metrics",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		}
	}

	rootSchema["details"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: detailsSchema,
		},
	}

	return rootSchema
}

func cloudConnectionConfigurationDetails() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"polling": {
			Type:        schema.TypeList,
			Description: "How often the selected connection is polled for information",
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"interval": {
						Type:        schema.TypeInt,
						Description: "The default polling interval is five (5) minutes",
						Optional:    true,
						Default:     5,
					},

					"unit": {
						Type:             schema.TypeString,
						Description:      "The unit of polling interval, currently only support 'minute'. Defaults to the same",
						Optional:         true,
						Default:          "minute",
						ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"minute"}, true)),
					},
				},
			},
		},

		"import_tags": {
			Type:        schema.TypeList,
			Description: "Configuration for importing tags of resources that are being monitored",
			Optional:    true,
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				detailsOld, _ := d.GetChange("details")
				if len(detailsOld.([]interface{})) > 0 {
					detailsMap := detailsOld.([]interface{})[0].(map[string]interface{})
					importTags := detailsMap["import_tags"].([]interface{})

					isImportTagsPresent := len(importTags) > 0

					previousWasEmpty := false
					if isImportTagsPresent && len(importTags[0].(map[string]interface{})["excluded_keys"].([]interface{})) == 0 {
						previousWasEmpty = true
					}

					if k == "details.0.import_tags.#" && oldValue == "1" && newValue == "0" && previousWasEmpty {
						return true
					}
				}
				return false
			},
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:        schema.TypeBool,
						Description: "It is true by default. Tags will be imported for all the resources that are being monitored by default",
						Optional:    true,
						Default:     true,
					},

					"excluded_keys": {
						Type:        schema.TypeList,
						Description: "Array of tag keys that need to be excluded from being imported. It can be set only when enabled is true",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},

		"tag_filter": {
			Type:        schema.TypeString,
			Description: "Expression for filtering resources to be monitored, based on tags. Example: (tags(env) = 'prod' || tags(env) = 'production')) && tags(project) = 'cloudcollectors'",
			Optional:    true,
		},
	}
}

func cloudConnectionConfigurationDetailsServices() map[string]*schema.Schema {
	return cloudConnectionConfigurationDetails()
}
