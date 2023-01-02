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
  client_id     = var.tenant_client_id
  client_secret = var.tenant_client_secret

  # # Hashicorp Vault
  # client_id     = data.terraform_remote_state.vault.outputs.tenant["client_id"]
  # client_secret = data.terraform_remote_state.vault.outputs.tenant["client_secret"]

  # # AWS Secret Manager
  # client_id = data.terraform_remote_state.aws_secretmanager.outputs.credentials.client_id
  # client_secret = data.terraform_remote_state.aws_secretmanager.outputs.credentials.client_secret

  # omit everything to use values from environment variables
}

resource "appdynamicscloud_connection_aws" "example" {
  display_name = "AWS Dev"
  description  = "Description for this AWS connection"
  state        = "ACTIVE"

  connection_details {
    access_key_id     = var.aws_access_key
    secret_access_key = var.aws_secret_key
    access_type       = "access_key"
  }

  configuration_details {

    services {
      name = "elb"

      polling {
        interval = 5
        unit     = "minute"
      }
    }

    services {
      name       = "ec2"
      tag_filter = "tags(project) = 'cloudcollectors'"
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
