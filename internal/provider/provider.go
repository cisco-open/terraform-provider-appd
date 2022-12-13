package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	client "github.com/aniketk-crest/appdynamicscloud-go-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New() func() *schema.Provider {
	return func() *schema.Provider {
		return Provider()
	}
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPDYNAMICS_CLIENT_ID", nil),
				Description: "ClientID of the AppDynamics API Client, this can also be set as the APPDYNAMICS_CLIENT_ID environment variable.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("APPDYNAMICS_CLIENT_SECRET", nil),
				Description: "ClientSecret of the AppDynamics API Client. This can also be set as the APPDYNAMICS_CLIENT_SECRET environment variable.",
			},
			"tenant_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Tenant name of the AppDynamics Platform.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"appdynamicscloud_connection_aws":                 resourceCloudConnectionAWS(),
			"appdynamicscloud_connection_azure":               resourceCloudConnectionAzure(),
			"appdynamicscloud_connection_aws_role_attachment": resourceCloudConnectionAWSRoleAttachment(),
			"appdynamicscloud_connection_configuration_aws":   resourceCloudConnectionConfigurationAWS(),
			"appdynamicscloud_connection_configuration_azure": resourceCloudConnectionConfigurationAzure(),
			// TODO: remaining resources
			// "appdynamicscloud_application_client_access": resourceCloudApplicationClientAccess(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			// "appdynamics_cloud_connection_aws": dataSourceCloudConnectionAWS(),
			"appdynamicscloud_query": dataSourceCloudQuery(),
		},
		ConfigureContextFunc: configureClient,
	}
}

func configureClient(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// GET ACCESS TOKEN
	token, diags := getAccessToken(d)
	if diags.HasError() {
		return nil, diags
	}

	// CONFIGURE API CLIENT
	configuration := client.NewConfiguration()
	configuration.Debug = true

	tenantName := d.Get("tenant_name").(string)
	myctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
		"tenant-name": tenantName,
	})
	myctx = context.WithValue(myctx, client.ContextAccessToken, token.AccessToken)

	// TERRAFORM CONFIG
	config := config{configuration: configuration, ctx: myctx}

	return config, nil
}

func getAccessToken(d *schema.ResourceData) (*oauth2.Token, diag.Diagnostics) {
	conf := clientcredentials.Config{}

	conf.ClientID = d.Get("client_id").(string)
	conf.ClientSecret = d.Get("client_secret").(string)

	tenantName := d.Get("tenant_name").(string)
	tenantId, err := lookupTenantId(tenantName)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	conf.TokenURL = fmt.Sprintf("https://%s.observe.appdynamics.com/auth/%s/default/oauth2/token", tenantName, tenantId)
	log.Printf("calling %s", conf.TokenURL)

	token, err := conf.Token(context.Background())

	// The error string contains response body as json.
	// parse it to display a more concise error message
	if err != nil {
		var resp map[string]interface{}

		errRespJson := strings.Split(err.Error(), "Response: ")[1]
		json.Unmarshal([]byte(errRespJson), &resp)

		d := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  resp["cause"].(string),
			Detail:   resp["error_description"].(string),
		}

		return nil, diag.Diagnostics{d}
	}

	return token, nil
}

func lookupTenantId(tenantName string) (string, error) {
	var tenantIdLookup map[string]interface{}

	tenantLookupUrl := fmt.Sprintf("https://observe-tenant-lookup-api.saas.appdynamics.com/tenants/lookup/%s.observe.appdynamics.com", tenantName)

	resp, err := http.Get(tenantLookupUrl)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	json.Unmarshal(body, &tenantIdLookup)

	if v, ok := tenantIdLookup["tenantId"]; ok {
		return v.(string), nil
	}

	return "", fmt.Errorf("%s, is tenant_name valid?", tenantIdLookup["message"].(string))
}

type config struct {
	ctx           context.Context
	configuration *client.Configuration
}
