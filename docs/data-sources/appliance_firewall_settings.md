---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_firewall_settings Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Firewall Settings configuration.
---

# meraki_appliance_firewall_settings (Data Source)

This data source can read the `Appliance Firewall Settings` configuration.

## Example Usage

```terraform
data "meraki_appliance_firewall_settings" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `id` (String) The id of the object
- `spoofing_protection_ip_source_guard_mode` (String) Mode of protection
