---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_firewalled_service Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Firewalled Service configuration.
---

# meraki_appliance_firewalled_service (Data Source)

This data source can read the `Appliance Firewalled Service` configuration.

## Example Usage

```terraform
data "meraki_appliance_firewalled_service" "example" {
  network_id = "L_123456"
  service    = "ICMP"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `service` (String) Service

### Read-Only

- `access` (String) A string indicating the rule for which IPs are allowed to use the specified service. Acceptable values are 'blocked' (no remote IPs can access the service), 'restricted' (only allowed IPs can access the service), and 'unrestriced' (any remote IP can access the service). This field is required
- `allowed_ips` (List of String) An array of allowed IPs that can access the service. This field is required if 'access' is set to 'restricted'. Otherwise this field is ignored
- `id` (String) The id of the object
