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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &OrganizationAdaptivePolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationAdaptivePolicyDataSource{}
)

func NewOrganizationAdaptivePolicyDataSource() datasource.DataSource {
	return &OrganizationAdaptivePolicyDataSource{}
}

type OrganizationAdaptivePolicyDataSource struct {
	client *meraki.Client
}

func (d *OrganizationAdaptivePolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_adaptive_policy"
}

func (d *OrganizationAdaptivePolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Organization Adaptive Policy` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"last_entry_rule": schema.StringAttribute{
				MarkdownDescription: "The rule to apply if there is no matching ACL (default: 'default')",
				Computed:            true,
			},
			"destination_group_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the destination adaptive policy group",
				Computed:            true,
			},
			"destination_group_name": schema.StringAttribute{
				MarkdownDescription: "The name of the destination adaptive policy group",
				Computed:            true,
			},
			"destination_group_sgt": schema.Int64Attribute{
				MarkdownDescription: "The SGT of the destination adaptive policy group",
				Computed:            true,
			},
			"source_group_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the source adaptive policy group",
				Computed:            true,
			},
			"source_group_name": schema.StringAttribute{
				MarkdownDescription: "The name of the source adaptive policy group",
				Computed:            true,
			},
			"source_group_sgt": schema.Int64Attribute{
				MarkdownDescription: "The SGT of the source adaptive policy group",
				Computed:            true,
			},
			"acls": schema.ListNestedAttribute{
				MarkdownDescription: "An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy (default: [])",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The ID of the adaptive policy ACL",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the adaptive policy ACL",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationAdaptivePolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationAdaptivePolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OrganizationAdaptivePolicy

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res gjson.Result
	var err error

	if !res.Exists() {
		res, err = d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
