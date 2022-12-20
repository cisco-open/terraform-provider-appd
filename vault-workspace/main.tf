terraform {
  required_providers {
    vault = {
      source  = "hashicorp/vault"
      version = "3.11.0"
    }
  }
}

provider "vault" {
    address = "http://127.0.0.1:8200/"
}


data "vault_kv_secret" "tenant_credentials" {
  path = "appd-partner-demo/data/tenant_credentails_test_service_principal"
}

data "vault_kv_secret" "azure_credentials" {
  path = "appd-partner-demo/data/azure_credentials"
}

output "tenant" {
  value = jsondecode(data.vault_kv_secret.tenant_credentials.data.data)
  sensitive = true
}

output "azure" {
  value = jsondecode(data.vault_kv_secret.azure_credentials.data.data)
  sensitive = true
}