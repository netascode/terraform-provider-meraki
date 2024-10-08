---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_dhcp_server_policy_arp_inspection_trusted_server Data Source - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This data source can read the Switch DHCP Server Policy ARP Inspection Trusted Server configuration.
---

# meraki_switch_dhcp_server_policy_arp_inspection_trusted_server (Data Source)

This data source can read the `Switch DHCP Server Policy ARP Inspection Trusted Server` configuration.

## Example Usage

```terraform
data "meraki_switch_dhcp_server_policy_arp_inspection_trusted_server" "example" {
  id         = "12345678"
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object
- `network_id` (String) Network ID

### Read-Only

- `ipv4_address` (String) The IPv4 address of the trusted server being added
- `mac` (String) The mac address of the trusted server being added
- `vlan` (Number) The VLAN of the trusted server being added. It must be between 1 and 4094
