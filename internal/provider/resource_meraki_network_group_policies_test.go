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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiNetworkGroupPolicies(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policies.test", "name", "test_group_policy"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policies.test", "bonjour_forwarding_settings", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policies.test", "bonjour_forwarding_rules.0.description", "a simple bonjour rule"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policies.test", "bonjour_forwarding_rules.0.services.0", "All Services"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policies.test", "bonjour_forwarding_rules.0.vlan_id", "2"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkGroupPoliciesPrerequisitesConfig + testAccMerakiNetworkGroupPoliciesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkGroupPoliciesPrerequisitesConfig + testAccMerakiNetworkGroupPoliciesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiNetworkGroupPoliciesPrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Network1"
  product_types   = ["switch", "wireless"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiNetworkGroupPoliciesConfig_minimum() string {
	config := `resource "meraki_network_group_policies" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	name = "test_group_policy"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiNetworkGroupPoliciesConfig_all() string {
	config := `resource "meraki_network_group_policies" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	name = "test_group_policy"` + "\n"
	config += `	bonjour_forwarding_settings = "custom"` + "\n"
	config += `	bonjour_forwarding_rules = [{` + "\n"
	config += `		description = "a simple bonjour rule"` + "\n"
	config += `		services = ["All Services"]` + "\n"
	config += `		vlan_id = "2"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
