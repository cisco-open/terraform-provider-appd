# Example for Role Delegation AWS Connction.

resource "appdynamicscloud_connection_aws" "example1" {
  display_name = "AWS Dev"
  description  = "Description for this AWS connection"
  state        = "ACTIVE"

  connection_details {
    access_type = "role_delegation"
    account_id  = "xxxxxxxxxxxx"
  }

  configuration_details {
    import_tags {
      enabled       = true
      excluded_keys = ["key1", "key2"]
    }
    tag_filter = "(tags(env) = 'prod' || tags(env) = 'production')) && tags(project) = 'cloudcollectors'"
    regions    = ["us-east-1", "us-west-1"]
    polling {
      interval = 5
      unit     = "minute"
    }

    services {
      name = "ebs"
      import_tags {
        enabled       = false
        excluded_keys = []
      }
      tag_filter = "tags(project) = 'cloudcollectors' && tags(jira) IN ['XTNSBL','ACE'] && !(tags(region) IN ['US','IN']) && HAS tags(monitorEnabled) && !(HAS tags(restrictedUse))"
      polling {
        interval = 5
        unit     = "minute"
      }
    }
    services {
      name = "ec2"
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

# Example for Access Key AWS Connction.

resource "appdynamicscloud_connection_aws" "example2" {
  display_name = "AWS Dev"
  description  = "Description for this AWS connection"
  state        = "ACTIVE"

  connection_details {
    access_type       = "access_key"
    access_key_id     = "xxxxxxxxxxxxxxxxxxx"
    secret_access_key = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  }

  configuration_details {
    import_tags {
      enabled       = true
      excluded_keys = ["key1", "key2"]
    }
    tag_filter = "(tags(env) = 'prod' || tags(env) = 'production')) && tags(project) = 'cloudcollectors'"
    regions    = ["us-east-1", "us-west-1"]
    polling {
      interval = 5
      unit     = "minute"
    }

    services {
      name = "ebs"
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
      name = "ec2"
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