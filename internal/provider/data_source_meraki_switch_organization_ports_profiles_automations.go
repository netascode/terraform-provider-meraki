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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SwitchOrganizationPortsProfilesAutomationsDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchOrganizationPortsProfilesAutomationsDataSource{}
)

func NewSwitchOrganizationPortsProfilesAutomationsDataSource() datasource.DataSource {
	return &SwitchOrganizationPortsProfilesAutomationsDataSource{}
}

type SwitchOrganizationPortsProfilesAutomationsDataSource struct {
	client *meraki.Client
}

func (d *SwitchOrganizationPortsProfilesAutomationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_organization_ports_profiles_automations"
}

func (d *SwitchOrganizationPortsProfilesAutomationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Switch Organization Ports Profiles Automation` configuration in bulk.").AddEarlyAccessDescription().String,

		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Text describing the port profile automation.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the port profile automation.",
							Computed:            true,
						},
						"fallback_profile_id": schema.StringAttribute{
							MarkdownDescription: "Default port profile Id",
							Computed:            true,
						},
						"fallback_profile_name": schema.StringAttribute{
							MarkdownDescription: "Default port profile name",
							Computed:            true,
						},
						"assigned_switch_ports": schema.ListNestedAttribute{
							MarkdownDescription: "assigned switch ports",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"switch_serial": schema.StringAttribute{
										MarkdownDescription: "Serial number of the switch",
										Computed:            true,
									},
									"port_ids": schema.SetAttribute{
										MarkdownDescription: "List of port ids",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
						"rules": schema.ListNestedAttribute{
							MarkdownDescription: "Configuration settings for port profile automation rules",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"priority": schema.Int64Attribute{
										MarkdownDescription: "Priority of automation rule in sequence",
										Computed:            true,
									},
									"profile_id": schema.StringAttribute{
										MarkdownDescription: "ID of port profile",
										Computed:            true,
									},
									"profile_name": schema.StringAttribute{
										MarkdownDescription: "Name of port profile",
										Computed:            true,
									},
									"conditions": schema.ListNestedAttribute{
										MarkdownDescription: "Configuration settings for port profile automation conditions",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attribute": schema.StringAttribute{
													MarkdownDescription: "Type of the condition",
													Computed:            true,
												},
												"values": schema.ListAttribute{
													MarkdownDescription: "Value of the condition",
													ElementType:         types.StringType,
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *SwitchOrganizationPortsProfilesAutomationsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchOrganizationPortsProfilesAutomationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceSwitchOrganizationPortsProfilesAutomations

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "SwitchOrganizationPortsProfilesAutomationsDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "SwitchOrganizationPortsProfilesAutomationsDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
