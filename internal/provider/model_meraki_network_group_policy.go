// Copyright © 2024 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"slices"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-meraki/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkGroupPolicy struct {
	Id                                             types.String                                                     `tfsdk:"id"`
	SplashAuthSettings                             types.String                                                     `tfsdk:"splash_auth_settings"`
	VlanTaggingSettings                            types.String                                                     `tfsdk:"vlan_tagging_settings"`
	VlanTaggingVlanId                              types.String                                                     `tfsdk:"vlan_tagging_vlan_id"`
	BonjourForwardingSettings                      types.String                                                     `tfsdk:"bonjour_forwarding_settings"`
	BonjourForwardingRules                         []NetworkGroupPolicyBonjourForwardingRules                       `tfsdk:"bonjour_forwarding_rules"`
	Name                                           types.String                                                     `tfsdk:"name"`
	SchedulingMondayTo                             types.String                                                     `tfsdk:"scheduling_monday_to"`
	SchedulingMondayActive                         types.Bool                                                       `tfsdk:"scheduling_monday_active"`
	SchedulingMondayFrom                           types.String                                                     `tfsdk:"scheduling_monday_from"`
	SchedulingTuesdayActive                        types.Bool                                                       `tfsdk:"scheduling_tuesday_active"`
	SchedulingTuesdayFrom                          types.String                                                     `tfsdk:"scheduling_tuesday_from"`
	SchedulingTuesdayTo                            types.String                                                     `tfsdk:"scheduling_tuesday_to"`
	SchedulingWednesdayTo                          types.String                                                     `tfsdk:"scheduling_wednesday_to"`
	SchedulingWednesdayActive                      types.Bool                                                       `tfsdk:"scheduling_wednesday_active"`
	SchedulingWednesdayFrom                        types.String                                                     `tfsdk:"scheduling_wednesday_from"`
	SchedulingThursdayActive                       types.Bool                                                       `tfsdk:"scheduling_thursday_active"`
	SchedulingThursdayFrom                         types.String                                                     `tfsdk:"scheduling_thursday_from"`
	SchedulingThursdayTo                           types.String                                                     `tfsdk:"scheduling_thursday_to"`
	SchedulingFridayActive                         types.Bool                                                       `tfsdk:"scheduling_friday_active"`
	SchedulingFridayFrom                           types.String                                                     `tfsdk:"scheduling_friday_from"`
	SchedulingFridayTo                             types.String                                                     `tfsdk:"scheduling_friday_to"`
	SchedulingSaturdayTo                           types.String                                                     `tfsdk:"scheduling_saturday_to"`
	SchedulingSaturdayActive                       types.Bool                                                       `tfsdk:"scheduling_saturday_active"`
	SchedulingSaturdayFrom                         types.String                                                     `tfsdk:"scheduling_saturday_from"`
	SchedulingSundayFrom                           types.String                                                     `tfsdk:"scheduling_sunday_from"`
	SchedulingSundayTo                             types.String                                                     `tfsdk:"scheduling_sunday_to"`
	SchedulingSundayActive                         types.Bool                                                       `tfsdk:"scheduling_sunday_active"`
	SchedulingEnabled                              types.Bool                                                       `tfsdk:"scheduling_enabled"`
	BandwidthSettings                              types.String                                                     `tfsdk:"bandwidth_settings"`
	BandwidthBandwidthLimitsLimitUp                types.Int64                                                      `tfsdk:"bandwidth_bandwidth_limits_limit_up"`
	BandwidthBandwidthLimitsLimitDown              types.Int64                                                      `tfsdk:"bandwidth_bandwidth_limits_limit_down"`
	FirewallAndTrafficShapingL7FirewallRules       []NetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules     `tfsdk:"firewall_and_traffic_shaping_l7_firewall_rules"`
	FirewallAndTrafficShapingSettings              types.String                                                     `tfsdk:"firewall_and_traffic_shaping_settings"`
	FirewallAndTrafficShapingTrafficShapingRules   []NetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules `tfsdk:"firewall_and_traffic_shaping_traffic_shaping_rules"`
	FirewallAndTrafficShapingL3FirewallRules       []NetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules     `tfsdk:"firewall_and_traffic_shaping_l3_firewall_rules"`
	ContentFilteringAllowedUrlPatternsSettings     types.String                                                     `tfsdk:"content_filtering_allowed_url_patterns_settings"`
	ContentFilteringAllowedUrlPatternsPatterns     types.List                                                       `tfsdk:"content_filtering_allowed_url_patterns_patterns"`
	ContentFilteringBlockedUrlPatternsSettings     types.String                                                     `tfsdk:"content_filtering_blocked_url_patterns_settings"`
	ContentFilteringBlockedUrlPatternsPatterns     types.List                                                       `tfsdk:"content_filtering_blocked_url_patterns_patterns"`
	ContentFilteringBlockedUrlCategoriesSettings   types.String                                                     `tfsdk:"content_filtering_blocked_url_categories_settings"`
	ContentFilteringBlockedUrlCategoriesCategories types.List                                                       `tfsdk:"content_filtering_blocked_url_categories_categories"`
	NetworkId                                      types.String                                                     `tfsdk:"network_id"`
}

type NetworkGroupPolicyBonjourForwardingRules struct {
	Description types.String `tfsdk:"description"`
	VlanId      types.String `tfsdk:"vlan_id"`
	Services    types.List   `tfsdk:"services"`
}

type NetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules struct {
	Policy types.String `tfsdk:"policy"`
	Type   types.String `tfsdk:"type"`
	Value  types.String `tfsdk:"value"`
}

type NetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules struct {
	Definitions                                      []NetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions `tfsdk:"definitions"`
	PerClientBandwidthLimitsSettings                 types.String                                                                `tfsdk:"per_client_bandwidth_limits_settings"`
	PerClientBandwidthLimitsBandwidthLimitsLimitUp   types.Int64                                                                 `tfsdk:"per_client_bandwidth_limits_bandwidth_limits_limit_up"`
	PerClientBandwidthLimitsBandwidthLimitsLimitDown types.Int64                                                                 `tfsdk:"per_client_bandwidth_limits_bandwidth_limits_limit_down"`
	DscpTagValue                                     types.Int64                                                                 `tfsdk:"dscp_tag_value"`
	PcpTagValue                                      types.Int64                                                                 `tfsdk:"pcp_tag_value"`
	Priority                                         types.String                                                                `tfsdk:"priority"`
}

type NetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules struct {
	Protocol types.String `tfsdk:"protocol"`
	DestPort types.String `tfsdk:"dest_port"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	Comment  types.String `tfsdk:"comment"`
	Policy   types.String `tfsdk:"policy"`
}

type NetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkGroupPolicy) getPath() string {
	return fmt.Sprintf("/networks/%v/groupPolicies", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkGroupPolicy) toBody(ctx context.Context, state NetworkGroupPolicy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "groupPolicyId", data.Id.ValueString())
	}
	if !data.SplashAuthSettings.IsNull() {
		body, _ = sjson.Set(body, "splashAuthSettings", data.SplashAuthSettings.ValueString())
	}
	if !data.VlanTaggingSettings.IsNull() {
		body, _ = sjson.Set(body, "vlanTagging.settings", data.VlanTaggingSettings.ValueString())
	}
	if !data.VlanTaggingVlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanTagging.vlanId", data.VlanTaggingVlanId.ValueString())
	}
	if !data.BonjourForwardingSettings.IsNull() {
		body, _ = sjson.Set(body, "bonjourForwarding.settings", data.BonjourForwardingSettings.ValueString())
	}
	if len(data.BonjourForwardingRules) > 0 {
		body, _ = sjson.Set(body, "bonjourForwarding.rules", []interface{}{})
		for _, item := range data.BonjourForwardingRules {
			itemBody := ""
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueString())
			}
			if !item.Services.IsNull() {
				var values []string
				item.Services.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "services", values)
			}
			body, _ = sjson.SetRaw(body, "bonjourForwarding.rules.-1", itemBody)
		}
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.SchedulingMondayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.monday.to", data.SchedulingMondayTo.ValueString())
	}
	if !data.SchedulingMondayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.monday.active", data.SchedulingMondayActive.ValueBool())
	}
	if !data.SchedulingMondayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.monday.from", data.SchedulingMondayFrom.ValueString())
	}
	if !data.SchedulingTuesdayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.tuesday.active", data.SchedulingTuesdayActive.ValueBool())
	}
	if !data.SchedulingTuesdayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.tuesday.from", data.SchedulingTuesdayFrom.ValueString())
	}
	if !data.SchedulingTuesdayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.tuesday.to", data.SchedulingTuesdayTo.ValueString())
	}
	if !data.SchedulingWednesdayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.wednesday.to", data.SchedulingWednesdayTo.ValueString())
	}
	if !data.SchedulingWednesdayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.wednesday.active", data.SchedulingWednesdayActive.ValueBool())
	}
	if !data.SchedulingWednesdayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.wednesday.from", data.SchedulingWednesdayFrom.ValueString())
	}
	if !data.SchedulingThursdayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.thursday.active", data.SchedulingThursdayActive.ValueBool())
	}
	if !data.SchedulingThursdayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.thursday.from", data.SchedulingThursdayFrom.ValueString())
	}
	if !data.SchedulingThursdayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.thursday.to", data.SchedulingThursdayTo.ValueString())
	}
	if !data.SchedulingFridayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.friday.active", data.SchedulingFridayActive.ValueBool())
	}
	if !data.SchedulingFridayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.friday.from", data.SchedulingFridayFrom.ValueString())
	}
	if !data.SchedulingFridayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.friday.to", data.SchedulingFridayTo.ValueString())
	}
	if !data.SchedulingSaturdayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.saturday.to", data.SchedulingSaturdayTo.ValueString())
	}
	if !data.SchedulingSaturdayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.saturday.active", data.SchedulingSaturdayActive.ValueBool())
	}
	if !data.SchedulingSaturdayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.saturday.from", data.SchedulingSaturdayFrom.ValueString())
	}
	if !data.SchedulingSundayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.sunday.from", data.SchedulingSundayFrom.ValueString())
	}
	if !data.SchedulingSundayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.sunday.to", data.SchedulingSundayTo.ValueString())
	}
	if !data.SchedulingSundayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.sunday.active", data.SchedulingSundayActive.ValueBool())
	}
	if !data.SchedulingEnabled.IsNull() {
		body, _ = sjson.Set(body, "scheduling.enabled", data.SchedulingEnabled.ValueBool())
	}
	if !data.BandwidthSettings.IsNull() {
		body, _ = sjson.Set(body, "bandwidth.settings", data.BandwidthSettings.ValueString())
	}
	if !data.BandwidthBandwidthLimitsLimitUp.IsNull() {
		body, _ = sjson.Set(body, "bandwidth.bandwidthLimits.limitUp", data.BandwidthBandwidthLimitsLimitUp.ValueInt64())
	}
	if !data.BandwidthBandwidthLimitsLimitDown.IsNull() {
		body, _ = sjson.Set(body, "bandwidth.bandwidthLimits.limitDown", data.BandwidthBandwidthLimitsLimitDown.ValueInt64())
	}
	if len(data.FirewallAndTrafficShapingL7FirewallRules) > 0 {
		body, _ = sjson.Set(body, "firewallAndTrafficShaping.l7FirewallRules", []interface{}{})
		for _, item := range data.FirewallAndTrafficShapingL7FirewallRules {
			itemBody := ""
			if !item.Policy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "policy", item.Policy.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "firewallAndTrafficShaping.l7FirewallRules.-1", itemBody)
		}
	}
	if !data.FirewallAndTrafficShapingSettings.IsNull() {
		body, _ = sjson.Set(body, "firewallAndTrafficShaping.settings", data.FirewallAndTrafficShapingSettings.ValueString())
	}
	if len(data.FirewallAndTrafficShapingTrafficShapingRules) > 0 {
		body, _ = sjson.Set(body, "firewallAndTrafficShaping.trafficShapingRules", []interface{}{})
		for _, item := range data.FirewallAndTrafficShapingTrafficShapingRules {
			itemBody := ""
			if len(item.Definitions) > 0 {
				itemBody, _ = sjson.Set(itemBody, "definitions", []interface{}{})
				for _, childItem := range item.Definitions {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "definitions.-1", itemChildBody)
				}
			}
			if !item.PerClientBandwidthLimitsSettings.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.settings", item.PerClientBandwidthLimitsSettings.ValueString())
			}
			if !item.PerClientBandwidthLimitsBandwidthLimitsLimitUp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.bandwidthLimits.limitUp", item.PerClientBandwidthLimitsBandwidthLimitsLimitUp.ValueInt64())
			}
			if !item.PerClientBandwidthLimitsBandwidthLimitsLimitDown.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.bandwidthLimits.limitDown", item.PerClientBandwidthLimitsBandwidthLimitsLimitDown.ValueInt64())
			}
			if !item.DscpTagValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dscpTagValue", item.DscpTagValue.ValueInt64())
			}
			if !item.PcpTagValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "pcpTagValue", item.PcpTagValue.ValueInt64())
			}
			if !item.Priority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "priority", item.Priority.ValueString())
			}
			body, _ = sjson.SetRaw(body, "firewallAndTrafficShaping.trafficShapingRules.-1", itemBody)
		}
	}
	if len(data.FirewallAndTrafficShapingL3FirewallRules) > 0 {
		body, _ = sjson.Set(body, "firewallAndTrafficShaping.l3FirewallRules", []interface{}{})
		for _, item := range data.FirewallAndTrafficShapingL3FirewallRules {
			itemBody := ""
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.DestPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destPort", item.DestPort.ValueString())
			}
			if !item.DestCidr.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destCidr", item.DestCidr.ValueString())
			}
			if !item.Comment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "comment", item.Comment.ValueString())
			}
			if !item.Policy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "policy", item.Policy.ValueString())
			}
			body, _ = sjson.SetRaw(body, "firewallAndTrafficShaping.l3FirewallRules.-1", itemBody)
		}
	}
	if !data.ContentFilteringAllowedUrlPatternsSettings.IsNull() {
		body, _ = sjson.Set(body, "contentFiltering.allowedUrlPatterns.settings", data.ContentFilteringAllowedUrlPatternsSettings.ValueString())
	}
	if !data.ContentFilteringAllowedUrlPatternsPatterns.IsNull() {
		var values []string
		data.ContentFilteringAllowedUrlPatternsPatterns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "contentFiltering.allowedUrlPatterns.patterns", values)
	}
	if !data.ContentFilteringBlockedUrlPatternsSettings.IsNull() {
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlPatterns.settings", data.ContentFilteringBlockedUrlPatternsSettings.ValueString())
	}
	if !data.ContentFilteringBlockedUrlPatternsPatterns.IsNull() {
		var values []string
		data.ContentFilteringBlockedUrlPatternsPatterns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlPatterns.patterns", values)
	}
	if !data.ContentFilteringBlockedUrlCategoriesSettings.IsNull() {
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlCategories.settings", data.ContentFilteringBlockedUrlCategoriesSettings.ValueString())
	}
	if !data.ContentFilteringBlockedUrlCategoriesCategories.IsNull() {
		var values []string
		data.ContentFilteringBlockedUrlCategoriesCategories.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlCategories.categories", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkGroupPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("splashAuthSettings"); value.Exists() {
		data.SplashAuthSettings = types.StringValue(value.String())
	} else {
		data.SplashAuthSettings = types.StringNull()
	}
	if value := res.Get("vlanTagging.settings"); value.Exists() {
		data.VlanTaggingSettings = types.StringValue(value.String())
	} else {
		data.VlanTaggingSettings = types.StringNull()
	}
	if value := res.Get("vlanTagging.vlanId"); value.Exists() {
		data.VlanTaggingVlanId = types.StringValue(value.String())
	} else {
		data.VlanTaggingVlanId = types.StringNull()
	}
	if value := res.Get("bonjourForwarding.settings"); value.Exists() {
		data.BonjourForwardingSettings = types.StringValue(value.String())
	} else {
		data.BonjourForwardingSettings = types.StringNull()
	}
	if value := res.Get("bonjourForwarding.rules"); value.Exists() {
		data.BonjourForwardingRules = make([]NetworkGroupPolicyBonjourForwardingRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkGroupPolicyBonjourForwardingRules{}
			if value := res.Get("description"); value.Exists() {
				data.Description = types.StringValue(value.String())
			} else {
				data.Description = types.StringNull()
			}
			if value := res.Get("vlanId"); value.Exists() {
				data.VlanId = types.StringValue(value.String())
			} else {
				data.VlanId = types.StringNull()
			}
			if value := res.Get("services"); value.Exists() {
				data.Services = helpers.GetStringList(value.Array())
			} else {
				data.Services = types.ListNull(types.StringType)
			}
			(*parent).BonjourForwardingRules = append((*parent).BonjourForwardingRules, data)
			return true
		})
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("scheduling.monday.to"); value.Exists() {
		data.SchedulingMondayTo = types.StringValue(value.String())
	} else {
		data.SchedulingMondayTo = types.StringNull()
	}
	if value := res.Get("scheduling.monday.active"); value.Exists() {
		data.SchedulingMondayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingMondayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.monday.from"); value.Exists() {
		data.SchedulingMondayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingMondayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.tuesday.active"); value.Exists() {
		data.SchedulingTuesdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingTuesdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.tuesday.from"); value.Exists() {
		data.SchedulingTuesdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingTuesdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.tuesday.to"); value.Exists() {
		data.SchedulingTuesdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingTuesdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.wednesday.to"); value.Exists() {
		data.SchedulingWednesdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingWednesdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.wednesday.active"); value.Exists() {
		data.SchedulingWednesdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingWednesdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.wednesday.from"); value.Exists() {
		data.SchedulingWednesdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingWednesdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.thursday.active"); value.Exists() {
		data.SchedulingThursdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingThursdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.thursday.from"); value.Exists() {
		data.SchedulingThursdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingThursdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.thursday.to"); value.Exists() {
		data.SchedulingThursdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingThursdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.friday.active"); value.Exists() {
		data.SchedulingFridayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingFridayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.friday.from"); value.Exists() {
		data.SchedulingFridayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingFridayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.friday.to"); value.Exists() {
		data.SchedulingFridayTo = types.StringValue(value.String())
	} else {
		data.SchedulingFridayTo = types.StringNull()
	}
	if value := res.Get("scheduling.saturday.to"); value.Exists() {
		data.SchedulingSaturdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingSaturdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.saturday.active"); value.Exists() {
		data.SchedulingSaturdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingSaturdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.saturday.from"); value.Exists() {
		data.SchedulingSaturdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingSaturdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.sunday.from"); value.Exists() {
		data.SchedulingSundayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingSundayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.sunday.to"); value.Exists() {
		data.SchedulingSundayTo = types.StringValue(value.String())
	} else {
		data.SchedulingSundayTo = types.StringNull()
	}
	if value := res.Get("scheduling.sunday.active"); value.Exists() {
		data.SchedulingSundayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingSundayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.enabled"); value.Exists() {
		data.SchedulingEnabled = types.BoolValue(value.Bool())
	} else {
		data.SchedulingEnabled = types.BoolNull()
	}
	if value := res.Get("bandwidth.settings"); value.Exists() {
		data.BandwidthSettings = types.StringValue(value.String())
	} else {
		data.BandwidthSettings = types.StringNull()
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitUp"); value.Exists() {
		data.BandwidthBandwidthLimitsLimitUp = types.Int64Value(value.Int())
	} else {
		data.BandwidthBandwidthLimitsLimitUp = types.Int64Null()
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitDown"); value.Exists() {
		data.BandwidthBandwidthLimitsLimitDown = types.Int64Value(value.Int())
	} else {
		data.BandwidthBandwidthLimitsLimitDown = types.Int64Null()
	}
	if value := res.Get("firewallAndTrafficShaping.l7FirewallRules"); value.Exists() {
		data.FirewallAndTrafficShapingL7FirewallRules = make([]NetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkGroupPolicyFirewallAndTrafficShapingL7FirewallRules{}
			if value := res.Get("policy"); value.Exists() {
				data.Policy = types.StringValue(value.String())
			} else {
				data.Policy = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).FirewallAndTrafficShapingL7FirewallRules = append((*parent).FirewallAndTrafficShapingL7FirewallRules, data)
			return true
		})
	}
	if value := res.Get("firewallAndTrafficShaping.settings"); value.Exists() {
		data.FirewallAndTrafficShapingSettings = types.StringValue(value.String())
	} else {
		data.FirewallAndTrafficShapingSettings = types.StringNull()
	}
	if value := res.Get("firewallAndTrafficShaping.trafficShapingRules"); value.Exists() {
		data.FirewallAndTrafficShapingTrafficShapingRules = make([]NetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRules{}
			if value := res.Get("definitions"); value.Exists() {
				data.Definitions = make([]NetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := NetworkGroupPolicyFirewallAndTrafficShapingTrafficShapingRulesDefinitions{}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("value"); value.Exists() {
						data.Value = types.StringValue(value.String())
					} else {
						data.Value = types.StringNull()
					}
					(*parent).Definitions = append((*parent).Definitions, data)
					return true
				})
			}
			if value := res.Get("perClientBandwidthLimits.settings"); value.Exists() {
				data.PerClientBandwidthLimitsSettings = types.StringValue(value.String())
			} else {
				data.PerClientBandwidthLimitsSettings = types.StringNull()
			}
			if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitUp"); value.Exists() {
				data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Value(value.Int())
			} else {
				data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Null()
			}
			if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitDown"); value.Exists() {
				data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Value(value.Int())
			} else {
				data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Null()
			}
			if value := res.Get("dscpTagValue"); value.Exists() {
				data.DscpTagValue = types.Int64Value(value.Int())
			} else {
				data.DscpTagValue = types.Int64Null()
			}
			if value := res.Get("pcpTagValue"); value.Exists() {
				data.PcpTagValue = types.Int64Value(value.Int())
			} else {
				data.PcpTagValue = types.Int64Null()
			}
			if value := res.Get("priority"); value.Exists() {
				data.Priority = types.StringValue(value.String())
			} else {
				data.Priority = types.StringNull()
			}
			(*parent).FirewallAndTrafficShapingTrafficShapingRules = append((*parent).FirewallAndTrafficShapingTrafficShapingRules, data)
			return true
		})
	}
	if value := res.Get("firewallAndTrafficShaping.l3FirewallRules"); value.Exists() {
		data.FirewallAndTrafficShapingL3FirewallRules = make([]NetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkGroupPolicyFirewallAndTrafficShapingL3FirewallRules{}
			if value := res.Get("protocol"); value.Exists() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("destPort"); value.Exists() {
				data.DestPort = types.StringValue(value.String())
			} else {
				data.DestPort = types.StringNull()
			}
			if value := res.Get("destCidr"); value.Exists() {
				data.DestCidr = types.StringValue(value.String())
			} else {
				data.DestCidr = types.StringNull()
			}
			if value := res.Get("comment"); value.Exists() {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			if value := res.Get("policy"); value.Exists() {
				data.Policy = types.StringValue(value.String())
			} else {
				data.Policy = types.StringNull()
			}
			(*parent).FirewallAndTrafficShapingL3FirewallRules = append((*parent).FirewallAndTrafficShapingL3FirewallRules, data)
			return true
		})
	}
	if value := res.Get("contentFiltering.allowedUrlPatterns.settings"); value.Exists() {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.allowedUrlPatterns.patterns"); value.Exists() {
		data.ContentFilteringAllowedUrlPatternsPatterns = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringAllowedUrlPatternsPatterns = types.ListNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.settings"); value.Exists() {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.patterns"); value.Exists() {
		data.ContentFilteringBlockedUrlPatternsPatterns = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringBlockedUrlPatternsPatterns = types.ListNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.settings"); value.Exists() {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.categories"); value.Exists() {
		data.ContentFilteringBlockedUrlCategoriesCategories = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringBlockedUrlCategoriesCategories = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkGroupPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("splashAuthSettings"); value.Exists() && !data.SplashAuthSettings.IsNull() {
		data.SplashAuthSettings = types.StringValue(value.String())
	} else {
		data.SplashAuthSettings = types.StringNull()
	}
	if value := res.Get("vlanTagging.settings"); value.Exists() && !data.VlanTaggingSettings.IsNull() {
		data.VlanTaggingSettings = types.StringValue(value.String())
	} else {
		data.VlanTaggingSettings = types.StringNull()
	}
	if value := res.Get("vlanTagging.vlanId"); value.Exists() && !data.VlanTaggingVlanId.IsNull() {
		data.VlanTaggingVlanId = types.StringValue(value.String())
	} else {
		data.VlanTaggingVlanId = types.StringNull()
	}
	if value := res.Get("bonjourForwarding.settings"); value.Exists() && !data.BonjourForwardingSettings.IsNull() {
		data.BonjourForwardingSettings = types.StringValue(value.String())
	} else {
		data.BonjourForwardingSettings = types.StringNull()
	}
	for i := 0; i < len(data.BonjourForwardingRules); i++ {
		keys := [...]string{"description", "vlanId"}
		keyValues := [...]string{data.BonjourForwardingRules[i].Description.ValueString(), data.BonjourForwardingRules[i].VlanId.ValueString()}

		parent := &data
		data := (*parent).BonjourForwardingRules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("bonjourForwarding.rules").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing BonjourForwardingRules[%d] = %+v",
				i,
				(*parent).BonjourForwardingRules[i],
			))
			(*parent).BonjourForwardingRules = slices.Delete((*parent).BonjourForwardingRules, i, i+1)
			i--

			continue
		}
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
			data.VlanId = types.StringValue(value.String())
		} else {
			data.VlanId = types.StringNull()
		}
		if value := res.Get("services"); value.Exists() && !data.Services.IsNull() {
			data.Services = helpers.GetStringList(value.Array())
		} else {
			data.Services = types.ListNull(types.StringType)
		}
		(*parent).BonjourForwardingRules[i] = data
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("scheduling.monday.to"); value.Exists() && !data.SchedulingMondayTo.IsNull() {
		data.SchedulingMondayTo = types.StringValue(value.String())
	} else {
		data.SchedulingMondayTo = types.StringNull()
	}
	if value := res.Get("scheduling.monday.active"); value.Exists() && !data.SchedulingMondayActive.IsNull() {
		data.SchedulingMondayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingMondayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.monday.from"); value.Exists() && !data.SchedulingMondayFrom.IsNull() {
		data.SchedulingMondayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingMondayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.tuesday.active"); value.Exists() && !data.SchedulingTuesdayActive.IsNull() {
		data.SchedulingTuesdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingTuesdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.tuesday.from"); value.Exists() && !data.SchedulingTuesdayFrom.IsNull() {
		data.SchedulingTuesdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingTuesdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.tuesday.to"); value.Exists() && !data.SchedulingTuesdayTo.IsNull() {
		data.SchedulingTuesdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingTuesdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.wednesday.to"); value.Exists() && !data.SchedulingWednesdayTo.IsNull() {
		data.SchedulingWednesdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingWednesdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.wednesday.active"); value.Exists() && !data.SchedulingWednesdayActive.IsNull() {
		data.SchedulingWednesdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingWednesdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.wednesday.from"); value.Exists() && !data.SchedulingWednesdayFrom.IsNull() {
		data.SchedulingWednesdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingWednesdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.thursday.active"); value.Exists() && !data.SchedulingThursdayActive.IsNull() {
		data.SchedulingThursdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingThursdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.thursday.from"); value.Exists() && !data.SchedulingThursdayFrom.IsNull() {
		data.SchedulingThursdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingThursdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.thursday.to"); value.Exists() && !data.SchedulingThursdayTo.IsNull() {
		data.SchedulingThursdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingThursdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.friday.active"); value.Exists() && !data.SchedulingFridayActive.IsNull() {
		data.SchedulingFridayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingFridayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.friday.from"); value.Exists() && !data.SchedulingFridayFrom.IsNull() {
		data.SchedulingFridayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingFridayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.friday.to"); value.Exists() && !data.SchedulingFridayTo.IsNull() {
		data.SchedulingFridayTo = types.StringValue(value.String())
	} else {
		data.SchedulingFridayTo = types.StringNull()
	}
	if value := res.Get("scheduling.saturday.to"); value.Exists() && !data.SchedulingSaturdayTo.IsNull() {
		data.SchedulingSaturdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingSaturdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.saturday.active"); value.Exists() && !data.SchedulingSaturdayActive.IsNull() {
		data.SchedulingSaturdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingSaturdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.saturday.from"); value.Exists() && !data.SchedulingSaturdayFrom.IsNull() {
		data.SchedulingSaturdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingSaturdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.sunday.from"); value.Exists() && !data.SchedulingSundayFrom.IsNull() {
		data.SchedulingSundayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingSundayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.sunday.to"); value.Exists() && !data.SchedulingSundayTo.IsNull() {
		data.SchedulingSundayTo = types.StringValue(value.String())
	} else {
		data.SchedulingSundayTo = types.StringNull()
	}
	if value := res.Get("scheduling.sunday.active"); value.Exists() && !data.SchedulingSundayActive.IsNull() {
		data.SchedulingSundayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingSundayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.enabled"); value.Exists() && !data.SchedulingEnabled.IsNull() {
		data.SchedulingEnabled = types.BoolValue(value.Bool())
	} else {
		data.SchedulingEnabled = types.BoolNull()
	}
	if value := res.Get("bandwidth.settings"); value.Exists() && !data.BandwidthSettings.IsNull() {
		data.BandwidthSettings = types.StringValue(value.String())
	} else {
		data.BandwidthSettings = types.StringNull()
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitUp"); value.Exists() && !data.BandwidthBandwidthLimitsLimitUp.IsNull() {
		data.BandwidthBandwidthLimitsLimitUp = types.Int64Value(value.Int())
	} else {
		data.BandwidthBandwidthLimitsLimitUp = types.Int64Null()
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitDown"); value.Exists() && !data.BandwidthBandwidthLimitsLimitDown.IsNull() {
		data.BandwidthBandwidthLimitsLimitDown = types.Int64Value(value.Int())
	} else {
		data.BandwidthBandwidthLimitsLimitDown = types.Int64Null()
	}
	for i := 0; i < len(data.FirewallAndTrafficShapingL7FirewallRules); i++ {
		keys := [...]string{"policy", "type", "value"}
		keyValues := [...]string{data.FirewallAndTrafficShapingL7FirewallRules[i].Policy.ValueString(), data.FirewallAndTrafficShapingL7FirewallRules[i].Type.ValueString(), data.FirewallAndTrafficShapingL7FirewallRules[i].Value.ValueString()}

		parent := &data
		data := (*parent).FirewallAndTrafficShapingL7FirewallRules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("firewallAndTrafficShaping.l7FirewallRules").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing FirewallAndTrafficShapingL7FirewallRules[%d] = %+v",
				i,
				(*parent).FirewallAndTrafficShapingL7FirewallRules[i],
			))
			(*parent).FirewallAndTrafficShapingL7FirewallRules = slices.Delete((*parent).FirewallAndTrafficShapingL7FirewallRules, i, i+1)
			i--

			continue
		}
		if value := res.Get("policy"); value.Exists() && !data.Policy.IsNull() {
			data.Policy = types.StringValue(value.String())
		} else {
			data.Policy = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
			data.Value = types.StringValue(value.String())
		} else {
			data.Value = types.StringNull()
		}
		(*parent).FirewallAndTrafficShapingL7FirewallRules[i] = data
	}
	if value := res.Get("firewallAndTrafficShaping.settings"); value.Exists() && !data.FirewallAndTrafficShapingSettings.IsNull() {
		data.FirewallAndTrafficShapingSettings = types.StringValue(value.String())
	} else {
		data.FirewallAndTrafficShapingSettings = types.StringNull()
	}
	for i := 0; i < len(data.FirewallAndTrafficShapingTrafficShapingRules); i++ {
		keys := [...]string{"perClientBandwidthLimits.settings", "perClientBandwidthLimits.bandwidthLimits.limitUp", "perClientBandwidthLimits.bandwidthLimits.limitDown", "dscpTagValue", "pcpTagValue", "priority"}
		keyValues := [...]string{data.FirewallAndTrafficShapingTrafficShapingRules[i].PerClientBandwidthLimitsSettings.ValueString(), strconv.FormatInt(data.FirewallAndTrafficShapingTrafficShapingRules[i].PerClientBandwidthLimitsBandwidthLimitsLimitUp.ValueInt64(), 10), strconv.FormatInt(data.FirewallAndTrafficShapingTrafficShapingRules[i].PerClientBandwidthLimitsBandwidthLimitsLimitDown.ValueInt64(), 10), strconv.FormatInt(data.FirewallAndTrafficShapingTrafficShapingRules[i].DscpTagValue.ValueInt64(), 10), strconv.FormatInt(data.FirewallAndTrafficShapingTrafficShapingRules[i].PcpTagValue.ValueInt64(), 10), data.FirewallAndTrafficShapingTrafficShapingRules[i].Priority.ValueString()}

		parent := &data
		data := (*parent).FirewallAndTrafficShapingTrafficShapingRules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("firewallAndTrafficShaping.trafficShapingRules").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing FirewallAndTrafficShapingTrafficShapingRules[%d] = %+v",
				i,
				(*parent).FirewallAndTrafficShapingTrafficShapingRules[i],
			))
			(*parent).FirewallAndTrafficShapingTrafficShapingRules = slices.Delete((*parent).FirewallAndTrafficShapingTrafficShapingRules, i, i+1)
			i--

			continue
		}
		for i := 0; i < len(data.Definitions); i++ {
			keys := [...]string{"type", "value"}
			keyValues := [...]string{data.Definitions[i].Type.ValueString(), data.Definitions[i].Value.ValueString()}

			parent := &data
			data := (*parent).Definitions[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("definitions").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing Definitions[%d] = %+v",
					i,
					(*parent).Definitions[i],
				))
				(*parent).Definitions = slices.Delete((*parent).Definitions, i, i+1)
				i--

				continue
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).Definitions[i] = data
		}
		if value := res.Get("perClientBandwidthLimits.settings"); value.Exists() && !data.PerClientBandwidthLimitsSettings.IsNull() {
			data.PerClientBandwidthLimitsSettings = types.StringValue(value.String())
		} else {
			data.PerClientBandwidthLimitsSettings = types.StringNull()
		}
		if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitUp"); value.Exists() && !data.PerClientBandwidthLimitsBandwidthLimitsLimitUp.IsNull() {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Null()
		}
		if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitDown"); value.Exists() && !data.PerClientBandwidthLimitsBandwidthLimitsLimitDown.IsNull() {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Null()
		}
		if value := res.Get("dscpTagValue"); value.Exists() && !data.DscpTagValue.IsNull() {
			data.DscpTagValue = types.Int64Value(value.Int())
		} else {
			data.DscpTagValue = types.Int64Null()
		}
		if value := res.Get("pcpTagValue"); value.Exists() && !data.PcpTagValue.IsNull() {
			data.PcpTagValue = types.Int64Value(value.Int())
		} else {
			data.PcpTagValue = types.Int64Null()
		}
		if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
			data.Priority = types.StringValue(value.String())
		} else {
			data.Priority = types.StringNull()
		}
		(*parent).FirewallAndTrafficShapingTrafficShapingRules[i] = data
	}
	for i := 0; i < len(data.FirewallAndTrafficShapingL3FirewallRules); i++ {
		keys := [...]string{"protocol", "destPort", "destCidr", "comment", "policy"}
		keyValues := [...]string{data.FirewallAndTrafficShapingL3FirewallRules[i].Protocol.ValueString(), data.FirewallAndTrafficShapingL3FirewallRules[i].DestPort.ValueString(), data.FirewallAndTrafficShapingL3FirewallRules[i].DestCidr.ValueString(), data.FirewallAndTrafficShapingL3FirewallRules[i].Comment.ValueString(), data.FirewallAndTrafficShapingL3FirewallRules[i].Policy.ValueString()}

		parent := &data
		data := (*parent).FirewallAndTrafficShapingL3FirewallRules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("firewallAndTrafficShaping.l3FirewallRules").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing FirewallAndTrafficShapingL3FirewallRules[%d] = %+v",
				i,
				(*parent).FirewallAndTrafficShapingL3FirewallRules[i],
			))
			(*parent).FirewallAndTrafficShapingL3FirewallRules = slices.Delete((*parent).FirewallAndTrafficShapingL3FirewallRules, i, i+1)
			i--

			continue
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("destPort"); value.Exists() && !data.DestPort.IsNull() {
			data.DestPort = types.StringValue(value.String())
		} else {
			data.DestPort = types.StringNull()
		}
		if value := res.Get("destCidr"); value.Exists() && !data.DestCidr.IsNull() {
			data.DestCidr = types.StringValue(value.String())
		} else {
			data.DestCidr = types.StringNull()
		}
		if value := res.Get("comment"); value.Exists() && !data.Comment.IsNull() {
			data.Comment = types.StringValue(value.String())
		} else {
			data.Comment = types.StringNull()
		}
		if value := res.Get("policy"); value.Exists() && !data.Policy.IsNull() {
			data.Policy = types.StringValue(value.String())
		} else {
			data.Policy = types.StringNull()
		}
		(*parent).FirewallAndTrafficShapingL3FirewallRules[i] = data
	}
	if value := res.Get("contentFiltering.allowedUrlPatterns.settings"); value.Exists() && !data.ContentFilteringAllowedUrlPatternsSettings.IsNull() {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.allowedUrlPatterns.patterns"); value.Exists() && !data.ContentFilteringAllowedUrlPatternsPatterns.IsNull() {
		data.ContentFilteringAllowedUrlPatternsPatterns = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringAllowedUrlPatternsPatterns = types.ListNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.settings"); value.Exists() && !data.ContentFilteringBlockedUrlPatternsSettings.IsNull() {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.patterns"); value.Exists() && !data.ContentFilteringBlockedUrlPatternsPatterns.IsNull() {
		data.ContentFilteringBlockedUrlPatternsPatterns = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringBlockedUrlPatternsPatterns = types.ListNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.settings"); value.Exists() && !data.ContentFilteringBlockedUrlCategoriesSettings.IsNull() {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.categories"); value.Exists() && !data.ContentFilteringBlockedUrlCategoriesCategories.IsNull() {
		data.ContentFilteringBlockedUrlCategoriesCategories = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringBlockedUrlCategoriesCategories = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial
