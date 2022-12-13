package provider

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	client "github.com/aniketk-crest/appdynamicscloud-go-client"
	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
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
