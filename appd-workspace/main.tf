variable "vault_workspace_path" {
  default = "../vault-workspace/terraform.tfstate"
}

terraform {
  required_providers {
    appdynamicscloud = {
      source = "appdynamics/appdynamicscloud"
    }
  }

#   backend "s3" {
#     bucket = "terraform-appd"
#     key = "terraform.tfstate"
#     region = "ap-south-1"
#   }
}

data "terraform_remote_state" "credentials" {
  backend = "local"

  config = {
    path = var.vault_workspace_path
  }
}


provider "appdynamicscloud" {
  client_id     = data.terraform_remote_state.credentials.outputs.tenant["client_id"]
  client_secret = data.terraform_remote_state.credentials.outputs.tenant["client_secret"]
  tenant_name   = "partner-demo"
}

resource "appdynamicscloud_connection_configuration_azure" "example" {
  display_name = "Azure Dev"
  description  = "Description for this Azure connection configuration"
  
  details {
    tag_filter      = "tags(project) = 'cloudcollectors'"
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
      tag_filter = "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'])"
      
      polling {
        interval = 5
        unit     = "minute"
      }
    }

    services {
      name = "disk"
      
      polling {
        interval = 5
        unit     = "minute"
      }
    }
  }
}

resource "appdynamicscloud_connection_azure" "example2" {
  display_name = "Chagned Azure Dev"
  description  = "Changed Description for this Azure connection"
  configuration_id = appdynamicscloud_connection_configuration_azure.example.id
  state            = "INACTIVE"

  details {
    client_id       = data.terraform_remote_state.credentials.outputs.azure["client_id"]
    client_secret   = data.terraform_remote_state.credentials.outputs.azure["client_secret"]
    tenant_id       = data.terraform_remote_state.credentials.outputs.azure["tenant_id"]
    subscription_id = data.terraform_remote_state.credentials.outputs.azure["subscription_id"]
  }
}
