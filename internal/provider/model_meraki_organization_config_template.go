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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationConfigTemplate struct {
	Id                types.String `tfsdk:"id"`
	OrganizationId    types.String `tfsdk:"organization_id"`
	CopyFromNetworkId types.String `tfsdk:"copy_from_network_id"`
	Name              types.String `tfsdk:"name"`
	TimeZone          types.String `tfsdk:"time_zone"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationConfigTemplate) getPath() string {
	return fmt.Sprintf("/organizations/%v/configTemplates", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationConfigTemplate) toBody(ctx context.Context, state OrganizationConfigTemplate) string {
	body := ""
	if !data.CopyFromNetworkId.IsNull() {
		body, _ = sjson.Set(body, "copyFromNetworkId", data.CopyFromNetworkId.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.TimeZone.IsNull() {
		body, _ = sjson.Set(body, "timeZone", data.TimeZone.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationConfigTemplate) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("timeZone"); value.Exists() {
		data.TimeZone = types.StringValue(value.String())
	} else {
		data.TimeZone = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationConfigTemplate) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("timeZone"); value.Exists() && !data.TimeZone.IsNull() {
		data.TimeZone = types.StringValue(value.String())
	} else {
		data.TimeZone = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial
