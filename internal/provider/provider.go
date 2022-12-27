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
	"github.com/aniketk-crest/terraform-provider-appdynamics/internal/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"golang.org/x/oauth2/clientcredentials"
)

func init() {
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
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPDYNAMICS_CLIENT_ID", nil),
				Description: "ClientID of the AppDynamics API Client, this can also be set as the APPDYNAMICS_CLIENT_ID environment variable. To be used with login mode service_principal.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("APPDYNAMICS_CLIENT_SECRET", nil),
				Description: "ClientSecret of the AppDynamics API Client. This can also be set as the APPDYNAMICS_CLIENT_SECRET environment variable. To be used with login mode service_principal.",
			},
			"tenant_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPDYNAMICS_TENANT_NAME", nil),
				Description: "Tenant name of the AppDynamics Platform. This can also be set as the APPDYNAMICS_TENANT_NAME environment variable.",
			},

			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPDYNAMICS_USERNAME", nil),
				Description: "Username to login to the AppDynamics Platform. This can also be set as the APPDYNAMICS_USERNAME environment variable. To be used with login mode headless.",
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPDYNAMICS_PASSWORD", nil),
				Description: "Password to login to the AppDynamics Platform. This can also be set as the APPDYNAMICS_PASSWORD environment variable. To be used with login mode headless.",
			},

			"login_mode": {
				Type:             schema.TypeString,
				Required:         true,
				DefaultFunc:      schema.EnvDefaultFunc("APPDYNAMICS_LOGIN_MODE", nil),
				Description:      "Mode of login. Possible values are: service_principal, browser and headless. This can also be set as the APPDYNAMICS_LOGIN_MODE environment variable.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"browser", "headless", "service_principal"}, true)),
			},

			"save_token": {
				Type:        schema.TypeBool,
				DefaultFunc: schema.EnvDefaultFunc("APPDYNAMICS_SAVE_TOKEN", nil),
				Description: "Whether or not to store the access token acquired by login mode browser and headless. This is for convenience and if you store the token, it would not prompt you to login again until it expires. The value is ignored with login mode service_principal. This can also be set as the APPDYNAMICS_SAVE_TOKEN environment variable.",
				Required:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"appdynamicscloud_query":          dataSourceCloudQuery(),
			"appdynamicscloud_services_azure": dataSourceCloudServicesAzure(),
			"appdynamicscloud_services_aws":   dataSourceCloudServicesAWS(),
			"appdynamicscloud_regions_aws":    dataSourceCloudRegionsAWS(),
			"appdynamicscloud_regions_azure":  dataSourceCloudRegionsAzure(),
		},
		ConfigureContextFunc: configureClient,
	}
}

var required = map[string][]string{
	"service_principal": {"client_id", "client_secret"},
	"headless":          {"username", "password"},
	"browser":           {},
}

func configureClient(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// GET ACCESS TOKEN
	var token string
	var diags diag.Diagnostics

	loginMode := d.Get("login_mode").(string)

	for _, attribute := range required[loginMode] {
		v, ok := d.GetOk(attribute)
		if !ok || v == "" {
			return nil, diag.Errorf("%v is required with to login with %v", attribute, loginMode)
		}
	}

	tenantName := d.Get("tenant_name").(string)
	tenantId, err := lookupTenantId(tenantName)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	if loginMode == "service_principal" {
		token, diags = getAccessToken(d)
		if diags.HasError() {
			return nil, diags
		}
	} else {
		ctx := context.WithValue(context.Background(), auth.MODE, loginMode)
		if loginMode == "headless" {
			ctx = context.WithValue(ctx, auth.USERNAME, d.Get("username"))
			ctx = context.WithValue(ctx, auth.PASSWORD, d.Get("password"))
		}

		token, err = auth.Login(tenantName, tenantId, d.Get("save_token").(bool), ctx)
		if err != nil {
			return nil, diag.FromErr(err)
		}
	}

	// CONFIGURE API CLIENT
	configuration := client.NewConfiguration()
	configuration.Debug = true

	myctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
		"tenant-name": tenantName,
	})
	myctx = context.WithValue(myctx, client.ContextAccessToken, token)

	// TERRAFORM CONFIG
	config := config{configuration: configuration, ctx: myctx}

	return config, nil
}

func getAccessToken(d *schema.ResourceData) (string, diag.Diagnostics) {
	conf := clientcredentials.Config{}

	conf.ClientID = d.Get("client_id").(string)
	conf.ClientSecret = d.Get("client_secret").(string)

	tenantName := d.Get("tenant_name").(string)
	tenantId, err := lookupTenantId(tenantName)
	if err != nil {
		return "", diag.FromErr(err)
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

		return "", diag.Diagnostics{d}
	}

	return token.AccessToken, nil
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
