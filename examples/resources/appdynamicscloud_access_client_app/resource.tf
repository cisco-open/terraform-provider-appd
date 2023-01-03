resource "appdynamicscloud_access_client_app" "example" {
  display_name = "AppDGrafanaPlugin"
  description  = "This Service principal can be used for authentication which help to connect AppD cloud to grafana"
  auth_type    = "client_secret_basic"

  rotate_secret             = true
  revoke_previous_secret_in = "3D"
  revoked_all_previous_at   = true
}