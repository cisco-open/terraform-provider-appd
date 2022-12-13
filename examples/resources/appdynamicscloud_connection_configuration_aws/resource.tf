resource "appdynamicscloud_connection_configuration_aws" "example" {
  display_name = "AWS Dev"
  description  = "Description for this AWS connection configuration"

  details {
    import_tags {
      enabled       = true
      excluded_keys = ["key1", "key2"]
    }
    tag_filter = "(tags(env) = 'prod' || tags(env) = 'production'))"
    regions    = ["us-east-1", "us-west-2"]
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
      import_tags {
        enabled       = false
        excluded_keys = []
      }
      tag_filter = "(tags(env) = 'prod' || tags(env) = 'production'))"
    }
    services {
      name = "ec2"
      polling {
        interval = 5
        unit     = "minute"
      }
      import_tags {
        enabled       = true
        excluded_keys = ["key1", "key2"]
      }
      tag_filter = "(tags(env) = 'prod' || tags(env) = 'production'))"
    }
  }
}