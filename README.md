# terraform-provider-appd

The AppDynamics Cloud Terraform provider is a plugin that allows Terraform to manage resources on AppDynamics Cloud Platform.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.18

## Building the Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
make build
```

## Using the Provider

If you are building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory, run `terraform init` to initialize it:

```terraform
# Configure provider with your Cisco AppDynamics credentials.

terraform {
  required_providers {
    appdynamicscloud = {
      source = "CiscoDevNet/appdynamicscloud"
    }
  }
}

provider "appdynamicscloud" {
  tenant_name = "tenant-name"
  login_mode  = "service_principal"

  client_id     = "xxxxxxxxxxxxxxxxx"
  client_secret = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run:

```shell
make testacc
```

To remove dangling resources created during acceptance tests, run `make sweep`:

```shell
make sweep
```
