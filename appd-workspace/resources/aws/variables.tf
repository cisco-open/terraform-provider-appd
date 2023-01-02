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

variable "aws_access_key" {
  type = string
}

variable "aws_secret_key" {
  type = string
}

variable "aws_account_id" {
  type = string
}

# ===========

variable "vault_workspace_path" {
  type = string
}

variable "aws_secretmanager_workspace_path" {
  type = string
}
