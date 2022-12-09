resource "appdynamicscloud_connection_azure" "example" {
  display_name     = "Azure Dev"
  description      = "Description for this Azure connection"
  configuration_id = appdynamicscloud_connection_configuration_azure.example
  state            = "ACTIVE"

  details {
    client_id       = "17760330-4ecd-47d5-818b-40308d3e67d1"
    client_secret   = "k33C_.5b0VrXSiQIf1t-c--p18~Ea_L~2H"
    tenant_id       = "5ae1af62-9505-4097-a69a-c1553ef7840e"
    subscription_id = "00000000-0000-0000-0000-000000000000"
  }
}
