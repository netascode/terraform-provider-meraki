---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_single_lan Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Single LAN configuration.
---

# meraki_appliance_single_lan (Data Source)

This data source can read the `Appliance Single LAN` configuration.

## Example Usage

```terraform
data "meraki_appliance_single_lan" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `appliance_ip` (String) The appliance IP address of the single LAN
- `id` (String) The id of the object
- `ipv6_enabled` (Boolean) Enable IPv6 on VLAN.
- `ipv6_prefix_assignments` (Attributes List) Prefix assignments on the VLAN (see [below for nested schema](#nestedatt--ipv6_prefix_assignments))
- `mandatory_dhcp_enabled` (Boolean) Enable Mandatory DHCP on LAN.
- `subnet` (String) The subnet of the single LAN configuration

<a id="nestedatt--ipv6_prefix_assignments"></a>
### Nested Schema for `ipv6_prefix_assignments`

Read-Only:

- `autonomous` (Boolean) Auto assign a /64 prefix from the origin to the VLAN
- `origin_interfaces` (List of String) Interfaces associated with the prefix
- `origin_type` (String) Type of the origin
- `static_appliance_ip6` (String) Manual configuration of the IPv6 Appliance IP
- `static_prefix` (String) Manual configuration of a /64 prefix on the VLAN
