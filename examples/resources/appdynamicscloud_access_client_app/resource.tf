resource "appdynamicscloud_access_client_app" "example" {
  display_name = "AppDGrafanaPlugin "
  description  = "This Service principal can be used for authentication which help to connect AppD cloud to grafana"
  auth_type    = "client_secret_basic"

  rotate_secret             = "12/20/2022"
  revoke_previous_secret_in = "3D"
  revoked_all_previous_at   = "12/29/2022"
}