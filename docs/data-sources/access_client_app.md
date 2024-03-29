---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "appdynamicscloud_access_client_app Data Source - terraform-provider-appdynamicscloud"
subcategory: ""
description: |-
  
---

# appdynamicscloud_access_client_app (Data Source)



## Example Usage

```terraform
data "appdynamicscloud_access_client_app" "example" {
  client_id = "xxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `client_id` (String) The Client ID of the Service Principal

### Read-Only

- `auth_type` (String) Supported authentication methods used to request OAuth tokens: `client_secret_basic` - The client credentials will be sent in the authorization header `client_secret_post` - The client credentials will be sent in the request body.
- `created_at` (String) The RFC3339 timestamp when the client was created
- `description` (String) A user provided description of the client.
- `display_name` (String) The display name for the client.
- `has_rotated_secrets` (Boolean) Indicates if the client has rotated secrets. Rotated client secrets can be revoked.
- `id` (String) The ID of this resource.
- `rotated_secret_expires_at` (String) The RFC3339 timestamp when the rotated client secret will expire.
- `updated_at` (String) The RFC3339 timestamp when the client was last updated.


