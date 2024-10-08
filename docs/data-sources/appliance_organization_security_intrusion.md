---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_organization_security_intrusion Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Organization Security Intrusion configuration.
---

# meraki_appliance_organization_security_intrusion (Data Source)

This data source can read the `Appliance Organization Security Intrusion` configuration.

## Example Usage

```terraform
data "meraki_appliance_organization_security_intrusion" "example" {
  organization_id = "123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) Organization ID

### Read-Only

- `allowed_rules` (Attributes List) Sets a list of specific SNORT signatures to allow (see [below for nested schema](#nestedatt--allowed_rules))
- `id` (String) The id of the object

<a id="nestedatt--allowed_rules"></a>
### Nested Schema for `allowed_rules`

Read-Only:

- `message` (String) Message is optional and is ignored on a PUT call. It is allowed in order for PUT to be compatible with GET
- `rule_id` (String) A rule identifier of the format meraki:intrusion/snort/GID//SID/. gid and sid can be obtained from either https://www.snort.org/rule-docs or as ruleIds from the security events in /organization/[orgId]/securityEvents
