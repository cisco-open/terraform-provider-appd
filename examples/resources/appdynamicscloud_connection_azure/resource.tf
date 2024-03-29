resource "appdynamicscloud_connection_azure" "example" {
  display_name = "Azure Dev"
  description  = "Description for this Azure connection"
  state        = "ACTIVE"

  connection_details {
    client_id       = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"
    client_secret   = ""
    tenant_id       = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"
    subscription_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"
  }

  configuration_details {
    import_tags {
      enabled       = true
      excluded_keys = ["key1", "key2"]
    }
    tag_filter      = "(tags(env) = 'prod' || tags(env) = 'production')) && tags(project) = 'cloudcollectors'"
    regions         = ["eastus", "westus"]
    resource_groups = ["resourceGroup1", "resourceGroup2", "resourceGroup3"]
    polling {
      interval = 5
      unit     = "minute"
    }

    services {
      name = "vm"
      import_tags {
        enabled       = false
        excluded_keys = []
      }
      tag_filter = "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)"
      polling {
        interval = 5
        unit     = "minute"
      }
    }
    services {
      name = "disk"
      import_tags {
        enabled       = true
        excluded_keys = ["key1", "key2"]
      }
      polling {
        interval = 5
        unit     = "minute"
      }
      tag_filter = "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse)"
    }
  }
}
