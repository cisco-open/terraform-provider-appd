terraform {
  required_providers {
    appdynamicscloud = {
      # The provider source here will be ciscodevnet/appdynamicscloud
      source = "appdynamicscloud"
    }
  }

  # backend "s3" {
  #   bucket = "terraform-appd"
  #   key = "terraform.tfstate"
  #   region = "ap-south-1"
  # }
}

data "terraform_remote_state" "vault" {
  backend = "local"

  config = {
    path = var.vault_workspace_path
  }
}

data "terraform_remote_state" "aws_secretmanager" {
  backend = "local"

  config = {
    path = var.aws_secretmanager_workspace_path
  }
}

provider "appdynamicscloud" {
  tenant_name = "partner-demo"
  login_mode  = "service_principal"

  # # hard-coded
  # client_id = "srv_6i9OwmrLVqaycbz6iflFmQ"
  # client_secret = "7JqM3auWPz_dmv4_hkbe_JyfauGt9A21ZI5HfWpPMks"

  # Terraform Variables
  client_id = var.tenant_client_id
  client_secret = var.tenant_client_secret

  # # Hashicorp Vault
  # client_id     = data.terraform_remote_state.vault.outputs.tenant["client_id"]
  # client_secret = data.terraform_remote_state.vault.outputs.tenant["client_secret"]

  # # AWS Secret Manager
  # client_id = data.terraform_remote_state.aws_secretmanager.outputs.credentials.client_id
  # client_secret = data.terraform_remote_state.aws_secretmanager.outputs.credentials.client_secret

  # omit everything to use values from environment variables
}

resource "appdynamicscloud_connection_azure" "example" {
  display_name = "Azure Dev"
  description  = "Description for this Azure connection"
  state        = "ACTIVE"

  # we can use the same method used above in provider block to use
  # credentials from tfvars, env vars, vautl or aws secret manager.
  connection_details {
    client_id       = var.azure_client_id
    client_secret   = var.azure_client_secret
    tenant_id       = var.azure_tenant_id
    subscription_id = var.azure_subscription_id
  }

  configuration_details {
    tag_filter      = "tags(env) = 'prod'"
    regions         = ["eastus", "westus"]
    resource_groups = ["resourceGroup1", "resourceGroup2", "resourceGroup3"]
    
    import_tags {
      enabled       = true
      excluded_keys = ["key1", "key2"]
    }

    polling {
      interval = 5
      unit     = "minute"
    }

    services {
      name = "vm"
      tag_filter = "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE']"
      
      import_tags {
        enabled       = false
        excluded_keys = []
      }
      
      polling {
        interval = 5
        unit     = "minute"
      }
    }

    services {
      name = "disk"
      tag_filter = "tags(region) IN ['US','IN'] && HAS tags(monitorEnabled)"
      
      import_tags {
        enabled       = true
        excluded_keys = ["key1", "key2"]
      }
      
      polling {
        interval = 5
        unit     = "minute"
      }
    }
  }
}
