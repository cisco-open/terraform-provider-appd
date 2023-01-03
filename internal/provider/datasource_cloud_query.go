package provider

import (
	"context"
	"encoding/json"
	"fmt"
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

	listOfData := []interface{}{}

	err = json.Unmarshal(bytes, &listOfData)
	if err != nil {
		return diag.FromErr(err)
	}
	// response := string(bytes)
	// d.SetId(*query.Query)
	// d.Set("response", response)

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
		tempList := []interface{}{}
		err = json.Unmarshal(bytes, &tempList)
		if err != nil {
			return diag.FromErr(err)
		}

		listOfData = append(listOfData, tempList...)
		cursor, flag = getCursor(bytes)
	}

	bytes, err = json.Marshal(&listOfData)
	if err != nil {
		return diag.FromErr(err)
	}
	response := string(bytes)
	d.SetId(*query.Query)
	d.Set("response", response)
	return nil
}

func getCursor(bytes []byte) (string, bool) {
	contBytes, _ := container.ParseJSON(bytes)
	link := contBytes.Index(1).Search("_links", "next", "href").Data()
	fmt.Printf("CURSOR:%v", link)
	if link == nil || (reflect.ValueOf(link).Kind() == reflect.Ptr && reflect.ValueOf(link).IsNil()) {
		return "", false
	} else {
		curSlice := strings.Split(contBytes.Index(1).Search("_links", "next", "href").Data().(string), "=")
		if len(curSlice) > 1 {
			cur := curSlice[1]
			cur = strings.Split(cur, "%")[0]
			return cur, true
		}
	}

	return "", false
}
