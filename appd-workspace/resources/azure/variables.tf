variable "tenant_client_id" {
  type = string
}

variable "tenant_client_secret" {
  type = string
}

variable "appd_username" {
  type = string
}

variable "appd_password" {
  type = string
}

# ===== AZURE ===== #

variable "azure_client_id" {
  type = string
}

variable "azure_client_secret" {
  type = string
}

variable "azure_tenant_id" {
  type = string
}

variable "azure_subscription_id" {
  type = string
}

# ===========

variable "vault_workspace_path" {
  type = string
}

variable "aws_secretmanager_workspace_path" {
  type = string
}
