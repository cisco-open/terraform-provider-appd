---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "appdynamicscloud_regions_azure Data Source - terraform-provider-appdynamicscloud"
subcategory: ""
description: |-
  
---

# appdynamicscloud_regions_azure (Data Source)



## Example Usage

```terraform
data "appdynamicscloud_regions_azure" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `regions_azure` (List of Object) All supported hosting regions for the Azure cloud provider (see [below for nested schema](#nestedatt--regions_azure))

<a id="nestedatt--regions_azure"></a>
### Nested Schema for `regions_azure`

Read-Only:

- `description` (String)
- `display_name` (String)
- `id` (String)


