package provider

import (
	"context"
	"io"
	"log"
	"strings"

	client "github.com/aniketk-crest/appdynamicscloud-go-client"
	cloudQueryApi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudquery"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudQuery() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudQueryRead,
		Schema: map[string]*schema.Schema{
			"query": {
				Type:        schema.TypeString,
				Description: "Query String",
				Required:    true,
			},
			"response": {
				Type:        schema.TypeString,
				Description: "Query Response",
				Computed:    true,
			},
		},
	}
}

func dataSourceCloudQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	myCtx, _, apiClient := initializeCloudQueryClient(m)
	query := cloudQueryApi.QueryRequestBody{}
	query.SetQuery(d.Get("query").(string))

	_, httpResp, err := apiClient.ExecuteQueryApi.ExecuteQuery(myCtx).QueryRequestBody(query).Execute()
	if err != nil {
		return errRespToDiag(err, httpResp)
	}

	bytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		log.Fatal(err)
	}
	response := string(bytes)
	response = strings.Replace(response, "\n", "", -1)
	response = strings.Replace(response, "\"", "'", -1)
	d.SetId(*query.Query)
	d.Set("response", response)
	return nil
}

func initializeCloudQueryClient(m interface{}) (context.Context, *client.Configuration, *cloudQueryApi.APIClient) {
	config := m.(config)

	configuration := config.configuration
	apiClient := cloudQueryApi.NewAPIClient(configuration)

	myCtx := context.WithValue(config.ctx, client.ContextServerIndex, client.SERVER_INDEX_CLOUD_QUERY)

	return myCtx, configuration, apiClient
}
