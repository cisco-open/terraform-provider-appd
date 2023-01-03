package provider

import (
	"context"
	"encoding/json"
	"io"
	"reflect"
	"strings"

	container "github.com/Jeffail/gabs/v2"
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
				Description: "Query String.",
				Required:    true,
			},
			"response": {
				Type:        schema.TypeString,
				Description: "Query Response.",
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
		return diag.FromErr(err)
	}

	listOfQueryResponse := make([]interface{}, 0, 1)

	queryResponse := []interface{}{}

	err = json.Unmarshal(bytes, &queryResponse)
	if err != nil {
		return diag.FromErr(err)
	}
	listOfQueryResponse = append(listOfQueryResponse, queryResponse)

	cursor, flag := getCursor(bytes)

	for flag {
		_, httpResp, err = apiClient.ResultPaginationApi.ResultPagination(myCtx).Cursor(cursor).Execute()
		if err != nil {
			return errRespToDiag(err, httpResp)
		}

		bytes, err = io.ReadAll(httpResp.Body)
		if err != nil {
			return diag.FromErr(err)
		}
		queryResponse := []interface{}{}
		err = json.Unmarshal(bytes, &queryResponse)
		if err != nil {
			return diag.FromErr(err)
		}

		listOfQueryResponse = append(listOfQueryResponse, queryResponse)
		cursor, flag = getCursor(bytes)
	}

	bytes, err = json.Marshal(&listOfQueryResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*query.Query)
	d.Set("response", string(bytes))
	return nil
}

func getCursor(bytes []byte) (string, bool) {
	contBytes, _ := container.ParseJSON(bytes)
	link := contBytes.Index(1).Search("_links", "next", "href").Data()

	if link == nil || (reflect.ValueOf(link).Kind() == reflect.Ptr && reflect.ValueOf(link).IsNil()) {
		return "", false
	} else {
		curSlice := strings.Split(contBytes.Index(1).Search("_links", "next", "href").Data().(string), "=")[1]
		cur := strings.Split(curSlice, "%")[0]
		return cur, true
	}
}
