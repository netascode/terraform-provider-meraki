---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_settings Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Settings configuration.
---

# meraki_appliance_settings (Data Source)

This data source can read the `Appliance Settings` configuration.

## Example Usage

```terraform
data "meraki_appliance_settings" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `client_tracking_method` (String) Client tracking method of a network
- `deployment_mode` (String) Deployment mode of a network
- `dynamic_dns_enabled` (Boolean) Dynamic DNS enabled
- `dynamic_dns_prefix` (String) Dynamic DNS url prefix. DDNS must be enabled to update
- `id` (String) The id of the object
