---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_device_cellular_sims Data Source - terraform-provider-meraki"
subcategory: "Devices"
description: |-
  This data source can read the Device Cellular SIMs configuration.
---

# meraki_device_cellular_sims (Data Source)

This data source can read the `Device Cellular SIMs` configuration.

## Example Usage

```terraform
data "meraki_device_cellular_sims" "example" {
  serial = "1234-1234-1234"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) Device serial

### Read-Only

- `id` (String) The id of the object
- `sim_failover_enabled` (Boolean) Failover to secondary SIM (optional)
- `sim_failover_timeout` (Number) Failover timeout in seconds (optional)
- `sim_ordering` (List of String) Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It`s required for devices with 3 or more SIMs and can be used in place of `isPrimary` for dual-SIM devices. To indicate eSIM, use `sim3`. Sim failover will occur only between primary and secondary sim slots.
- `sims` (Attributes List) List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged. (see [below for nested schema](#nestedatt--sims))

<a id="nestedatt--sims"></a>
### Nested Schema for `sims`

Read-Only:

- `apns` (Attributes List) APN configurations. If empty, the default APN will be used. (see [below for nested schema](#nestedatt--sims--apns))
- `is_primary` (Boolean) If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using `simOrdering`.
- `sim_order` (Number) Priority of SIM slot being configured. Use a value between 1 and total number of SIMs available. The value must be unique for each SIM.
- `slot` (String) SIM slot being configured. Must be `sim1` on single-sim devices. Use `sim3` for eSIM.

<a id="nestedatt--sims--apns"></a>
### Nested Schema for `sims.apns`

Read-Only:

- `allowed_ip_types` (Set of String) IP versions to support (permitted values include `ipv4`, `ipv6`).
- `authentication_password` (String) APN password, if type is set (if APN password is not supplied, the password is left unchanged).
- `authentication_type` (String) APN auth type.
- `authentication_username` (String) APN username, if type is set.
- `name` (String) APN name.
