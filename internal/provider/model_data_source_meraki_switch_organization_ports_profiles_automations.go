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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceSwitchOrganizationPortsProfilesAutomations struct {
	OrganizationId types.String                                                `tfsdk:"organization_id"`
	Items          []DataSourceSwitchOrganizationPortsProfilesAutomationsItems `tfsdk:"items"`
}

type DataSourceSwitchOrganizationPortsProfilesAutomationsItems struct {
	Id                  types.String                                                              `tfsdk:"id"`
	Description         types.String                                                              `tfsdk:"description"`
	Name                types.String                                                              `tfsdk:"name"`
	FallbackProfileId   types.String                                                              `tfsdk:"fallback_profile_id"`
	FallbackProfileName types.String                                                              `tfsdk:"fallback_profile_name"`
	AssignedSwitchPorts []DataSourceSwitchOrganizationPortsProfilesAutomationsAssignedSwitchPorts `tfsdk:"assigned_switch_ports"`
	Rules               []DataSourceSwitchOrganizationPortsProfilesAutomationsRules               `tfsdk:"rules"`
}

type DataSourceSwitchOrganizationPortsProfilesAutomationsAssignedSwitchPorts struct {
	SwitchSerial types.String `tfsdk:"switch_serial"`
	PortIds      types.Set    `tfsdk:"port_ids"`
}

type DataSourceSwitchOrganizationPortsProfilesAutomationsRules struct {
	Priority    types.Int64                                                           `tfsdk:"priority"`
	ProfileId   types.String                                                          `tfsdk:"profile_id"`
	ProfileName types.String                                                          `tfsdk:"profile_name"`
	Conditions  []DataSourceSwitchOrganizationPortsProfilesAutomationsRulesConditions `tfsdk:"conditions"`
}

type DataSourceSwitchOrganizationPortsProfilesAutomationsRulesConditions struct {
	Attribute types.String `tfsdk:"attribute"`
	Values    types.List   `tfsdk:"values"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSwitchOrganizationPortsProfilesAutomations) getPath() string {
	return fmt.Sprintf("/organizations/%v/switch/ports/profiles/automations", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchOrganizationPortsProfilesAutomations) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceSwitchOrganizationPortsProfilesAutomationsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceSwitchOrganizationPortsProfilesAutomationsItems{}
		data.Id = types.StringValue(res.Get("id").String())
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
			data.AssignedSwitchPorts = make([]DataSourceSwitchOrganizationPortsProfilesAutomationsAssignedSwitchPorts, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceSwitchOrganizationPortsProfilesAutomationsAssignedSwitchPorts{}
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
			data.Rules = make([]DataSourceSwitchOrganizationPortsProfilesAutomationsRules, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceSwitchOrganizationPortsProfilesAutomationsRules{}
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
					data.Conditions = make([]DataSourceSwitchOrganizationPortsProfilesAutomationsRulesConditions, 0)
					value.ForEach(func(k, res gjson.Result) bool {
						parent := &data
						data := DataSourceSwitchOrganizationPortsProfilesAutomationsRulesConditions{}
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
