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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchOrganizationPortsProfilesAutomation struct {
	Id                  types.String                                                   `tfsdk:"id"`
	OrganizationId      types.String                                                   `tfsdk:"organization_id"`
	Description         types.String                                                   `tfsdk:"description"`
	Name                types.String                                                   `tfsdk:"name"`
	FallbackProfileId   types.String                                                   `tfsdk:"fallback_profile_id"`
	FallbackProfileName types.String                                                   `tfsdk:"fallback_profile_name"`
	AssignedSwitchPorts []SwitchOrganizationPortsProfilesAutomationAssignedSwitchPorts `tfsdk:"assigned_switch_ports"`
	Rules               []SwitchOrganizationPortsProfilesAutomationRules               `tfsdk:"rules"`
}

type SwitchOrganizationPortsProfilesAutomationAssignedSwitchPorts struct {
	SwitchSerial types.String `tfsdk:"switch_serial"`
	PortIds      types.Set    `tfsdk:"port_ids"`
}

type SwitchOrganizationPortsProfilesAutomationRules struct {
	Priority    types.Int64                                                `tfsdk:"priority"`
	ProfileId   types.String                                               `tfsdk:"profile_id"`
	ProfileName types.String                                               `tfsdk:"profile_name"`
	Conditions  []SwitchOrganizationPortsProfilesAutomationRulesConditions `tfsdk:"conditions"`
}

type SwitchOrganizationPortsProfilesAutomationRulesConditions struct {
	Attribute types.String `tfsdk:"attribute"`
	Values    types.List   `tfsdk:"values"`
}

type SwitchOrganizationPortsProfilesAutomationIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchOrganizationPortsProfilesAutomation) getPath() string {
	return fmt.Sprintf("/organizations/%v/switch/ports/profiles/automations", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchOrganizationPortsProfilesAutomation) toBody(ctx context.Context, state SwitchOrganizationPortsProfilesAutomation) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.FallbackProfileId.IsNull() {
		body, _ = sjson.Set(body, "fallbackProfile.id", data.FallbackProfileId.ValueString())
	} else if !state.FallbackProfileId.IsNull() {
		body, _ = sjson.Set(body, "fallbackProfile.id", nil)
	}
	if !data.FallbackProfileName.IsNull() {
		body, _ = sjson.Set(body, "fallbackProfile.name", data.FallbackProfileName.ValueString())
	} else if !state.FallbackProfileName.IsNull() {
		body, _ = sjson.Set(body, "fallbackProfile.name", nil)
	}
	if len(data.AssignedSwitchPorts) > 0 {
		body, _ = sjson.Set(body, "assignedSwitchPorts", []interface{}{})
		for _, item := range data.AssignedSwitchPorts {
			itemBody := ""
			if !item.SwitchSerial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "switch.serial", item.SwitchSerial.ValueString())
			}
			if !item.PortIds.IsNull() {
				var values []string
				item.PortIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "portIds", values)
			}
			body, _ = sjson.SetRaw(body, "assignedSwitchPorts.-1", itemBody)
		}
	}
	{
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Priority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "priority", item.Priority.ValueInt64())
			}
			if !item.ProfileId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "profile.id", item.ProfileId.ValueString())
			}
			if !item.ProfileName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "profile.name", item.ProfileName.ValueString())
			}
			{
				itemBody, _ = sjson.Set(itemBody, "conditions", []interface{}{})
				for _, childItem := range item.Conditions {
					itemChildBody := ""
					if !childItem.Attribute.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "attribute", childItem.Attribute.ValueString())
					}
					if !childItem.Values.IsNull() {
						var values []string
						childItem.Values.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "values", values)
					}
					itemBody, _ = sjson.SetRaw(itemBody, "conditions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchOrganizationPortsProfilesAutomation) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && value.Value() != nil {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("fallbackProfile.id"); value.Exists() && value.Value() != nil {
		data.FallbackProfileId = types.StringValue(value.String())
	} else {
		data.FallbackProfileId = types.StringNull()
	}
	if value := res.Get("fallbackProfile.name"); value.Exists() && value.Value() != nil {
		data.FallbackProfileName = types.StringValue(value.String())
	} else {
		data.FallbackProfileName = types.StringNull()
	}
	if value := res.Get("assignedSwitchPorts"); value.Exists() && value.Value() != nil {
		data.AssignedSwitchPorts = make([]SwitchOrganizationPortsProfilesAutomationAssignedSwitchPorts, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchOrganizationPortsProfilesAutomationAssignedSwitchPorts{}
			if value := res.Get("switch.serial"); value.Exists() && value.Value() != nil {
				data.SwitchSerial = types.StringValue(value.String())
			} else {
				data.SwitchSerial = types.StringNull()
			}
			if value := res.Get("portIds"); value.Exists() && value.Value() != nil {
				data.PortIds = helpers.GetStringSet(value.Array())
			} else {
				data.PortIds = types.SetNull(types.StringType)
			}
			(*parent).AssignedSwitchPorts = append((*parent).AssignedSwitchPorts, data)
			return true
		})
	}
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]SwitchOrganizationPortsProfilesAutomationRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchOrganizationPortsProfilesAutomationRules{}
			if value := res.Get("priority"); value.Exists() && value.Value() != nil {
				data.Priority = types.Int64Value(value.Int())
			} else {
				data.Priority = types.Int64Null()
			}
			if value := res.Get("profile.id"); value.Exists() && value.Value() != nil {
				data.ProfileId = types.StringValue(value.String())
			} else {
				data.ProfileId = types.StringNull()
			}
			if value := res.Get("profile.name"); value.Exists() && value.Value() != nil {
				data.ProfileName = types.StringValue(value.String())
			} else {
				data.ProfileName = types.StringNull()
			}
			if value := res.Get("conditions"); value.Exists() && value.Value() != nil {
				data.Conditions = make([]SwitchOrganizationPortsProfilesAutomationRulesConditions, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := SwitchOrganizationPortsProfilesAutomationRulesConditions{}
					if value := res.Get("attribute"); value.Exists() && value.Value() != nil {
						data.Attribute = types.StringValue(value.String())
					} else {
						data.Attribute = types.StringNull()
					}
					if value := res.Get("values"); value.Exists() && value.Value() != nil {
						data.Values = helpers.GetStringList(value.Array())
					} else {
						data.Values = types.ListNull(types.StringType)
					}
					(*parent).Conditions = append((*parent).Conditions, data)
					return true
				})
			}
			(*parent).Rules = append((*parent).Rules, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchOrganizationPortsProfilesAutomation) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("fallbackProfile.id"); value.Exists() && !data.FallbackProfileId.IsNull() {
		data.FallbackProfileId = types.StringValue(value.String())
	} else {
		data.FallbackProfileId = types.StringNull()
	}
	if value := res.Get("fallbackProfile.name"); value.Exists() && !data.FallbackProfileName.IsNull() {
		data.FallbackProfileName = types.StringValue(value.String())
	} else {
		data.FallbackProfileName = types.StringNull()
	}
	for i := 0; i < len(data.AssignedSwitchPorts); i++ {
		keys := [...]string{"switch.serial"}
		keyValues := [...]string{data.AssignedSwitchPorts[i].SwitchSerial.ValueString()}

		parent := &data
		data := (*parent).AssignedSwitchPorts[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("assignedSwitchPorts").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AssignedSwitchPorts[%d] = %+v",
				i,
				(*parent).AssignedSwitchPorts[i],
			))
			(*parent).AssignedSwitchPorts = slices.Delete((*parent).AssignedSwitchPorts, i, i+1)
			i--

			continue
		}
		if value := res.Get("switch.serial"); value.Exists() && !data.SwitchSerial.IsNull() {
			data.SwitchSerial = types.StringValue(value.String())
		} else {
			data.SwitchSerial = types.StringNull()
		}
		if value := res.Get("portIds"); value.Exists() && !data.PortIds.IsNull() {
			data.PortIds = helpers.GetStringSet(value.Array())
		} else {
			data.PortIds = types.SetNull(types.StringType)
		}
		(*parent).AssignedSwitchPorts[i] = data
	}
	for i := 0; i < len(data.Rules); i++ {
		keys := [...]string{"priority"}
		keyValues := [...]string{strconv.FormatInt(data.Rules[i].Priority.ValueInt64(), 10)}

		parent := &data
		data := (*parent).Rules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("rules").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Rules[%d] = %+v",
				i,
				(*parent).Rules[i],
			))
			(*parent).Rules = slices.Delete((*parent).Rules, i, i+1)
			i--

			continue
		}
		if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
			data.Priority = types.Int64Value(value.Int())
		} else {
			data.Priority = types.Int64Null()
		}
		if value := res.Get("profile.id"); value.Exists() && !data.ProfileId.IsNull() {
			data.ProfileId = types.StringValue(value.String())
		} else {
			data.ProfileId = types.StringNull()
		}
		if value := res.Get("profile.name"); value.Exists() && !data.ProfileName.IsNull() {
			data.ProfileName = types.StringValue(value.String())
		} else {
			data.ProfileName = types.StringNull()
		}
		for i := 0; i < len(data.Conditions); i++ {
			keys := [...]string{"attribute"}
			keyValues := [...]string{data.Conditions[i].Attribute.ValueString()}

			parent := &data
			data := (*parent).Conditions[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("conditions").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Conditions[%d] = %+v",
					i,
					(*parent).Conditions[i],
				))
				(*parent).Conditions = slices.Delete((*parent).Conditions, i, i+1)
				i--

				continue
			}
			if value := res.Get("attribute"); value.Exists() && !data.Attribute.IsNull() {
				data.Attribute = types.StringValue(value.String())
			} else {
				data.Attribute = types.StringNull()
			}
			if value := res.Get("values"); value.Exists() && !data.Values.IsNull() {
				data.Values = helpers.GetStringList(value.Array())
			} else {
				data.Values = types.ListNull(types.StringType)
			}
			(*parent).Conditions[i] = data
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchOrganizationPortsProfilesAutomation) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *SwitchOrganizationPortsProfilesAutomationIdentity) toIdentity(ctx context.Context, plan *SwitchOrganizationPortsProfilesAutomation) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *SwitchOrganizationPortsProfilesAutomation) fromIdentity(ctx context.Context, identity *SwitchOrganizationPortsProfilesAutomationIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data SwitchOrganizationPortsProfilesAutomation) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
