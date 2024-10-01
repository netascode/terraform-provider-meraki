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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceFirewalledService struct {
	Id         types.String `tfsdk:"id"`
	NetworkId  types.String `tfsdk:"network_id"`
	Service    types.String `tfsdk:"service"`
	Access     types.String `tfsdk:"access"`
	AllowedIps types.List   `tfsdk:"allowed_ips"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceFirewalledService) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/firewall/firewalledServices/%v", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Service.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceFirewalledService) toBody(ctx context.Context, state ApplianceFirewalledService) string {
	body := ""
	if !data.Access.IsNull() {
		body, _ = sjson.Set(body, "access", data.Access.ValueString())
	}
	if !data.AllowedIps.IsNull() {
		var values []string
		data.AllowedIps.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "allowedIps", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceFirewalledService) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("access"); value.Exists() && value.Value() != nil {
		data.Access = types.StringValue(value.String())
	} else {
		data.Access = types.StringNull()
	}
	if value := res.Get("allowedIps"); value.Exists() && value.Value() != nil {
		data.AllowedIps = helpers.GetStringList(value.Array())
	} else {
		data.AllowedIps = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceFirewalledService) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("access"); value.Exists() && !data.Access.IsNull() {
		data.Access = types.StringValue(value.String())
	} else {
		data.Access = types.StringNull()
	}
	if value := res.Get("allowedIps"); value.Exists() && !data.AllowedIps.IsNull() {
		data.AllowedIps = helpers.GetStringList(value.Array())
	} else {
		data.AllowedIps = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial
