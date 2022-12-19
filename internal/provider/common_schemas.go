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
			Type:             schema.TypeString,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
			Optional:         true,
		},
	}
}
