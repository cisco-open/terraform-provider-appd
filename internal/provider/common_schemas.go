package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const serviceEmptyErrorMsg = "At Least one services is required while updating, services cannot be updated as empty."


func getCloudConnectionAzureSchema() map[string]*schema.Schema {
	return appendSchemas(
		cloudConnectionCommonSchema(),
		cloudConnectionCommonSchemaExtras(),
		cloudConnectionDetailsAzureSchema(),
		cloudConnectionConfigurationAzureSchema(),
	)
}

func cloudConnectionConfigurationAWSSchema() map[string]*schema.Schema {
	return cloudConnectionConfigurationSchema("AWS")
}

func cloudConnectionConfigurationAzureSchema() map[string]*schema.Schema {
	return cloudConnectionConfigurationSchema("AZURE")
}

func cloudConnectionCommonSchema() map[string]*schema.Schema {
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

func cloudConnectionCommonSchemaExtras() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"state": {
			Type:             schema.TypeString,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"ACTIVE", "INACTIVE"}, true)),
			Description:      "Connection state. This can only be used if configuration_id is specified. Possible values: [\"ACTIVE\", \"INACTIVE\"]",
			Optional:         true,
			Computed:         true,
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				// Connection can go into inactive state only after it is
				// activated at least once.
				//
				// Thus, if it is created with inactive state at the
				// time of creation, it will go into configured state,
				// which is technically the same hence suppressing diff.
				return oldValue == "CONFIGURED" && newValue == "INACTIVE"
			},
		},
		"state_message": {
			Type:        schema.TypeString,
			Description: "Connection state message",
			Computed:    true,
		},
		"configuration_id": {
			Type: schema.TypeString,
			// ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
			Computed: true,
		},
	}
}

func cloudConnectionDetailsAzureSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connection_details": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"client_id": {
						Type:             schema.TypeString,
						ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
						Description:      "Client IDs, also known as Application IDs, are long-term credentials for an Azure user, or account root user. The Client ID is one of three properties needed to authenticate to Azure, the other two being Client Secret and Tenant (Directory) ID",
						Required:         true,
					},
					"client_secret": {
						Type:        schema.TypeString,
						Description: "A Client Secret allows an Azure application to provide its identity when requesting an access token. The Client Secret is one of three properties needed to authenticate to Azure, the other two being Client ID (Application ID) and Tenant (Directory) ID",
						Sensitive:   true,
						Required:    true,
					},
					"tenant_id": {
						Type:             schema.TypeString,
						ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
						Description:      "The Azure AD Tenant (Directory) IDis one of three properties needed to authenticate to Azure. The other two are Client Secret and Client ID (Application ID).",
						Required:         true,
						ForceNew:         true,
					},
					"subscription_id": {
						Type:             schema.TypeString,
						ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
						Description:      "Specify a GUID Subscription ID to monitor. If monitoring all subscriptions, do not specify a Subscription ID.",
						Required:         true,
						ForceNew:         true,
					},
				},
			},
		},
	}
}

func cloudConnectionConfigurationSchema(provider string) map[string]*schema.Schema {
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
		Type:        schema.TypeSet,
		Description: "services for which we will fetch metrics",
		Optional:    true,
		Computed:    true,
		Set:         calculateHashStringWithPolling,
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

	return map[string]*schema.Schema{
		"configuration_details": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: detailsSchema,
			},
		},
		"configuration_details_service_default": {
			Type:        schema.TypeBool,
			Description: "Whether default services are present in configuration details",
			Computed:    true,
		},
	}
}

func cloudConnectionConfigurationDetailsServices() map[string]*schema.Schema {
	return cloudConnectionConfigurationDetails()
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
			Type:             schema.TypeList,
			Description:      "Configuration for importing tags of resources that are being monitored",
			Optional:         true,
			DiffSuppressFunc: importTagsSuppressFunc,
			MaxItems:         1,
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

func calculateHashStringWithPolling(val interface{}) int {
	if val == nil {
		return 0
	}
	var hash int = 0
	v := val.(map[string]any)
	name := v["name"]
	hash += schema.HashString(name)
	if polling, ok := singleListToMap(v["polling"]); ok {
		interval := polling["interval"].(int)
		hash += schema.HashString(fmt.Sprint(interval))
	} else {
		hash += schema.HashString("5")
	}
	if importTags, ok := singleListToMap(v["import_tags"]); ok {
		enabled := importTags["enabled"]
		excludedKeys := importTags["excluded_keys"]
		hash += schema.HashString(fmt.Sprint(enabled)) + schema.HashString(fmt.Sprint(excludedKeys))
	}
	if v["tag_filter"].(string) != "" {
		hash += schema.HashString(v["tag_filter"])
	}

	return hash
}

func importTagsSuppressFunc(k, oldValue, newValue string, d *schema.ResourceData) bool {
	detailsOld, _ := d.GetChange("configuration_details")
	if len(detailsOld.([]interface{})) > 0 {
		detailsMap := detailsOld.([]interface{})[0].(map[string]interface{})
		importTags := detailsMap["import_tags"].([]interface{})

		isImportTagsPresent := len(importTags) > 0

		previousWasEmpty := false
		if isImportTagsPresent && len(importTags[0].(map[string]interface{})["excluded_keys"].([]interface{})) == 0 {
			previousWasEmpty = true
		}

		if k == "configuration_details.0.import_tags.#" && oldValue == "1" && newValue == "0" && previousWasEmpty {
			return true
		}
	}
	return false
}
func serviceAtLeastOne(ctx context.Context, rd *schema.ResourceDiff, i interface{}) error {
	var length int
	if len(rd.GetRawConfig().GetAttr("configuration_details").AsValueSlice()) > 0 {
		length = rd.GetRawConfig().GetAttr("configuration_details").AsValueSlice()[0].GetAttr("services").AsValueSet().Length()
	}

	val, exist := rd.GetOkExists("configuration_details_service_default")
	if !exist && length == 0 {
		rd.SetNew("configuration_details_service_default", true)
	} else if exist && length > 0 {
		rd.SetNew("configuration_details_service_default", false)
	} else if length == 0 && exist == true && val == false {
		return fmt.Errorf(serviceEmptyErrorMsg)
	}
	return nil
}
