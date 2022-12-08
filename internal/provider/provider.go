package provider

import (
	"context"
	"fmt"

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
				Description: "ClientID of the AppDynamics API Client, in the format of [ClientName]@[TenantName] This can also be set as the APPDYNAMICS_CLIENT_ID environment variable.",
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
			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Tenant ID of the AppDynamics Platform.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"appdynamicscloud_connection_aws":                 resourceCloudConnectionAWS(),
			"appdynamicscloud_connection_azure":               resourceCloudConnectionAzure(),
			"appdynamicscloud_connection_configuration_aws":   resourceCloudConnectionConfigurationAWS(),
			"appdynamicscloud_connection_configuration_azure": resourceCloudConnectionConfigurationAzure(),
			// TODO: remaining resources
			// "appdynamicscloud_application_client_access": resourceCloudApplicationClientAccess(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			// "appdynamicscloud_connection_aws": dataSourceCloudConnectionAWS(),
			"appdynamicscloud_connection_azure": dataSourceCloudConnectionAzure(),
			// "appdynamicscloud_query":           dataSourceCloudQuery(),
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
	// TODO: Use custom configuration
	configuration := client.NewConfiguration()
	configuration.Debug = true

	tenantName := d.Get("tenant_name").(string)
	myctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
		"tenant-name": tenantName,
	})
	myctx = context.WithValue(myctx, client.ContextAccessToken, token.AccessToken)

	// TERRAFORM CONFIG
	config := config{configuration: configuration, ctx: myctx}

	return config, diags
}

func getAccessToken(d *schema.ResourceData) (*oauth2.Token, diag.Diagnostics) {
	var diags diag.Diagnostics

	conf := clientcredentials.Config{}

	if v, ok := d.GetOk("client_id"); ok {
		conf.ClientID = v.(string)
	}

	if v, ok := d.GetOk("client_secret"); ok {
		conf.ClientSecret = v.(string)
	}

	var tenantId, tenantName string

	if v, ok := d.GetOk("tenant_id"); ok {
		tenantId = v.(string)
	}

	if v, ok := d.GetOk("tenant_name"); ok {
		tenantName = v.(string)
		conf.TokenURL = fmt.Sprintf("https://%s.observe.appdynamics.com/auth/%s/default/oauth2/token", tenantName, tenantId)
	}

	// TODO: when get infra access
	// replace vars
	// if tenantId is not defined
	// GET https://observe-tenant-lookup-api.saas.appdynamics.com/tenants/lookup/{tenantName}.observe.appdynamics.com

	if conf.ClientID == "" || conf.ClientSecret == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create AppDynamics client",
			Detail:   "ClientID and ClientSecret are required",
		})
	}

	if tenantName == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create AppDynamics client",
			Detail:   "Tenant name is required",
		})
	}

	token, err := conf.Token(context.Background())
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return token, diags
}

type config struct {
	ctx           context.Context
	configuration *client.Configuration
}
