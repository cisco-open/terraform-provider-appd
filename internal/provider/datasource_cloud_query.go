package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	cloudqueryapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudquery"
)

func dataSourceCloudQuery() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudQueryRead,
		Schema: map[string]*schema.Schema{
			"query": {
				Type:     schema.TypeString,
				Required: true,
			},

			"query_results": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCloudQueryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	myctx, _, apiClient := initializeCloudQueryClient(meta)

	query := d.Get("query").(string)

	queryRequestBody := cloudqueryapi.NewQueryRequestBody()
	queryRequestBody.SetQuery(query)

	// TODO: convert to json string and store into query_result
	// resp, r, err := apiClient.ExecuteQueryApi.ExecuteQuery(myctx).QueryRequestBody(*queryRequestBody).Execute()
	_, _, _ = apiClient.ExecuteQueryApi.ExecuteQuery(myctx).QueryRequestBody(*queryRequestBody).Execute()

	d.SetId(query)

	return nil
}
