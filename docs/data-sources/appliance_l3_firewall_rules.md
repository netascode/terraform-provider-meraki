---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_l3_firewall_rules Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance L3 Firewall Rules configuration.
---

# meraki_appliance_l3_firewall_rules (Data Source)

This data source can read the `Appliance L3 Firewall Rules` configuration.

## Example Usage

```terraform
data "meraki_appliance_l3_firewall_rules" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `id` (String) The id of the object
- `rules` (Attributes List) An ordered array of the firewall rules (not including the default rule) (see [below for nested schema](#nestedatt--rules))
- `syslog_default_rule` (Boolean) Log the special default rule (boolean value - enable only if you`ve configured a syslog server) (optional)

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Read-Only:

- `comment` (String) Description of the rule (optional)
- `dest_cidr` (String) Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or `Any`
- `dest_port` (String) Comma-separated list of destination port(s) (integer in the range 1-65535), or `Any`
- `policy` (String) `allow` or `deny` traffic specified by this rule
- `protocol` (String) The type of protocol (must be `tcp`, `udp`, `icmp`, `icmp6` or `any`)
- `src_cidr` (String) Comma-separated list of source IP address(es) (in IP or CIDR notation), or `Any` (note: FQDN not supported for source addresses)
- `src_port` (String) Comma-separated list of source port(s) (integer in the range 1-65535), or `Any`
- `syslog_enabled` (Boolean) Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
