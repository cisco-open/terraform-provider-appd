resource "appdynamicscloud_connection_aws" "example" {
  display_name     = "AWS Dev"
  description      = "Description for this AWS role delegation connection"
  configuration_id = appdynamicscloud_connection_configuration_aws.example
  state            = "ACTIVE"
  details {
    access_type = "role_delegation"
    account_id  = "258762700219"
  }

  # Role Delegation
  # details {
  #   access_type       = "role_delegation"
  #   access_key_id     = "something"
  #   secret_access_key = "something"
  # }
}
