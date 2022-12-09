resource "appdynamicscloud_connection_configuration_aws" "example" {
  display_name = "Example AWS Configuration"
  description  = "AWS Example Configuration"

  details {
    tag_filter = "(tags(env) = 'prod' || tags(env) = 'production'))"
    regions    = ["us-east-1", "us-west-2"]

    import_tags {
      enabled       = true
      excluded_keys = ["local", "sandbox"]
    }

    polling {
      interval = 5
      unit     = "minute"
    }

    services {
      name = "elb"
      polling {
        interval = 5
        unit     = "minute"
      }
    }

    services {
      name = "ec2"
      polling {
        interval = 5
        unit     = "minute"
      }
    }
  }
}
