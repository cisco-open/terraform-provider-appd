resource "appdynamicscloud_connection_azure" "example" {
  display_name     = "Azure Dev"
  description      = "Description for this Azure connection"
  configuration_id = appdynamicscloud_connection_configuration_azure.example.id
  state            = "ACTIVE"

  details {
    client_id       = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"
    client_secret   = ""
    tenant_id       = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"
    subscription_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"
  }
}
