resource "appdynamicscloud_connection_configuration_azure" "example" {
  display_name = "Example Azure Configuration"
  description  = "Azure Example Configuration"

  details {
    tag_filter = "(tags(env) = 'prod' || tags(env) = 'production')) && tags(project) = 'cloudcollectors'"
    regions    = ["eastus", "westus"]

    import_tags {
      enabled       = true
      excluded_keys = ["key1", "key2"]
    }

    polling {
      interval = 5
      unit     = "minute"
    }

    services {
      name       = "vm"
      tag_filter = "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)"
      import_tags {
        enabled      = true
        exclude_keys = ["dev"]
      }
      polling {
        interval = 5
        unit     = "minute"
      }
    }

    services {
      name = "emysql"
      polling {
        interval = 5
        unit     = "minute"
      }
    }

    resource_groups = ["resourceGroup1", "resourceGroup2"]
  }
}
