terraform {
  required_providers {
    appdynamicscloud = {
      source = "appdynamicscloud"
    }
  }
}

provider "appdynamicscloud" {
  tenant_name   = "partner-demo"
  login_mode = "headless"
  
  username = var.appd_username
  password = var.appd_password
  
  save_token = true
}

resource "appdynamicscloud_access_client_app" "example" {
  display_name = "basic service principal"
  description = "orchestrated by terraform"
  auth_type = "client_secret_basic" 

#   rotate_secret = true
#   revoke_previous_secret_in = "1D"
#   revoke_now = true
}