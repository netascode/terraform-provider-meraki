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

//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/gen/yamlconfig"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"

	"golang.org/x/exp/maps"
)

const specPath = "./gen/models/spec3.json"
const definitionsPath = "./gen/definitions/"

const usage = `
Usage: go run gen/definition.go <endpoint> <resource_name>

Arguments:
  endpoint      The specific endpoint that is to be converted to generator specification
  resource_name The name that will be given to the resource

Example:
  go run gen/definition.go "/networks/{networkId}/groupPolicies/{groupPolicyId}" "Network Group Policy"`

const orgPrerequisites = `
data "meraki_organization" "test" {
  name = var.test_org
}
`

const networkPrerequisites = `
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance"]
}
`

const devicePrerequisites = `
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_switch_1_serial]
}
`

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Insufficient number of arguments")
		fmt.Println(usage)
		os.Exit(1)
	}

	endpointPath := os.Args[1]
	resourceName := os.Args[2]

	generateDefinition(endpointPath, resourceName)
}

func generateDefinition(endpointPath, resourceName string) {
	specData, err := os.ReadFile(specPath)
	if err != nil {
		fmt.Printf("Error reading OpenAPI spec file: %v\n", err)
		os.Exit(1)
	}

	// Load the OpenAPI spec
	var spec interface{}
	if strings.HasSuffix(specPath, ".json") {
		err = json.Unmarshal(specData, &spec)
	} else {
		err = yaml.Unmarshal(specData, &spec)
	}
	if err != nil {
		fmt.Printf("Error parsing OpenAPI spec: %v\n", err)
		os.Exit(1)
	}

	config := yamlconfig.YamlConfig{}
	urlResult := parseUrl(endpointPath, spec)

	example := urlResult.schema["schema"].(map[string]interface{})["example"].(map[string]interface{})
	exampleStr, err := json.Marshal(&example)
	if err != nil {
		panic(err)
	}

	if urlResult.resultPath[len(urlResult.resultPath)-1] == '/' {
		urlResult.resultPath = urlResult.resultPath[:len(urlResult.resultPath)-1]
	}
	config.RestEndpoint = urlResult.resultPath
	config.DocCategory = urlResult.category
	config.Name = resourceName
	config.PutCreate = urlResult.putCreate
	config.NoDelete = urlResult.noDelete
	config.GetFromAll = urlResult.getFromAll
	config.NoImport = urlResult.noImport
	config.NoResource = urlResult.noResource
	config.NoDataSource = urlResult.noDataSource
	config.NoUpdate = urlResult.noUpdate
	config.TestVariables = urlResult.testVariables

	attributes := []yamlconfig.YamlConfigAttribute{}
	for i, r := range urlResult.references {
		attr := yamlconfig.YamlConfigAttribute{}
		attr.TfName = yamlconfig.CamelToSnake(r[1 : len(r)-1])
		attr.Type = "String"
		attr.Reference = true
		if !urlResult.hasShortUrl && i == len(urlResult.references)-1 {
			attr.Id = true
		}
		if attr.TfName == "organization_id" {
			attr.Description = "Organization ID"
			attr.TestValue = "data.meraki_organization.test.id"
			attr.Example = "123456"
		} else if attr.TfName == "network_id" {
			attr.Description = "Network ID"
			attr.TestValue = "meraki_network.test.id"
			attr.Example = "L_123456"
		} else if attr.TfName == "serial" {
			attr.Description = "Device serial"
			attr.TestValue = "tolist(meraki_network_device_claim.test.serials)[0]"
			attr.Example = "1234-ABCD-1234"
		} else {
			attr.Description = "<<Description>>"
			attr.TestValue = "<<TestValue>>"
			attr.Example = "<<Example>>"
		}
		attributes = append(attributes, attr)
	}
	required := []string{}
	if r, ok := urlResult.schema["schema"].(map[string]interface{})["required"]; ok {
		required = toStringSlice(r.([]interface{}))
	}
	attributes = append(attributes, traverseProperties(urlResult.schema["schema"].(map[string]interface{})["properties"].(map[string]interface{}), []string{}, "", string(exampleStr), required)...)
	config.Attributes = attributes

	dataSourceNameQuery := false
	for _, a := range config.Attributes {
		if a.ModelName == "name" && len(a.DataPath) == 0 && !config.PutCreate {
			dataSourceNameQuery = true
			break
		}
	}
	config.DataSourceNameQuery = dataSourceNameQuery

	if slices.Contains(config.TestVariables, "test_switch_1_serial") {
		config.TestPrerequisites = devicePrerequisites
	} else if slices.Contains(config.TestVariables, "test_network") {
		config.TestPrerequisites = networkPrerequisites
	} else if slices.Contains(config.TestVariables, "test_org") {
		config.TestPrerequisites = orgPrerequisites
	}

	outputPath := definitionsPath + yamlconfig.SnakeCase(resourceName) + ".yaml"

	existingConfig := yamlconfig.YamlConfigP{}
	comments := ""
	commentsEndpoint := ""
	if yamlFile, err := os.ReadFile(outputPath); err == nil {
		// retain comments at the beginning of the definition file
		scanner := bufio.NewScanner(bytes.NewReader(yamlFile))
		for scanner.Scan() {
			line := scanner.Text()
			if line[0] == '#' {
				if strings.Contains(line, yamlconfig.EndpointToken) {
					commentsEndpoint = strings.Trim(strings.Split(line, yamlconfig.EndpointToken)[1], " ")
				} else {
					comments += line + "\n"
				}
			} else {
				break
			}
		}
		err = yaml.Unmarshal(yamlFile, &existingConfig)
		if err != nil {
			panic(err)
		}
	}
	if commentsEndpoint == "" {
		commentsEndpoint = endpointPath
	}

	newConfig := yamlconfig.MergeYamlConfig(existingConfig, config)

	var yamlBytes bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&yamlBytes)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(&newConfig)
	if err != nil {
		panic(err)
	}
	writeString := fmt.Sprintf("# %s %s\n%s%s", yamlconfig.EndpointToken, commentsEndpoint, comments, yamlBytes.String())
	os.WriteFile(outputPath, []byte(writeString), 0644)
}

func toStringSlice(i []interface{}) []string {
	ret := []string{}
	for _, v := range i {
		ret = append(ret, v.(string))
	}
	return ret
}

var jsonTypes = map[string]string{
	"integer": "Int64",
	"number":  "Float64",
	"boolean": "Bool",
	"string":  "String",
}

type parseUrlResult struct {
	resultPath    string
	references    []string
	category      string
	schema        map[string]interface{}
	putCreate     bool
	noDelete      bool
	getFromAll    bool
	hasShortUrl   bool
	noResource    bool
	noDataSource  bool
	noImport      bool
	noUpdate      bool
	testVariables []string
}

func parseUrl(url string, spec interface{}) parseUrlResult {
	ret := parseUrlResult{}

	shortUrl := ""
	if url[len(url)-1] == '}' {
		parts := strings.Split(url, "/")
		if len(parts) > 0 {
			parts = parts[:len(parts)-1]
		}
		shortUrl = strings.Join(parts, "/")
	}
	if shortUrl != "" {
		ret.hasShortUrl = true
	}

	hasPost := false
	hasShortPost := false
	hasPut := false
	hasShortPut := false
	hasGet := false
	hasShortGet := false
	hasDelete := false
	hasShortDelete := false

	paths := spec.(map[string]interface{})["paths"].(map[string]interface{})
	if p, ok := paths[shortUrl]; ok && shortUrl != "" {
		if _, ok := p.(map[string]interface{})["post"]; ok {
			hasShortPost = true
		}
		if _, ok := p.(map[string]interface{})["put"]; ok {
			hasShortPut = true
		}
		if _, ok := p.(map[string]interface{})["get"]; ok {
			hasShortGet = true
		}
		if _, ok := p.(map[string]interface{})["delete"]; ok {
			hasShortDelete = true
		}
	}
	if p, ok := paths[url]; ok {
		if _, ok := p.(map[string]interface{})["post"]; ok {
			hasPost = true
		}
		if _, ok := p.(map[string]interface{})["put"]; ok {
			hasPut = true
		}
		if _, ok := p.(map[string]interface{})["get"]; ok {
			hasGet = true
		}
		if _, ok := p.(map[string]interface{})["delete"]; ok {
			hasDelete = true
		}
	}

	if !hasPost && !hasShortPost {
		if hasPut || hasShortPut {
			ret.putCreate = true
		} else {
			ret.noResource = true
		}
	}
	if !ret.noResource {
		if !hasGet && hasShortGet {
			ret.getFromAll = true
		}
		if !hasDelete && !hasShortDelete {
			ret.noDelete = true
		}
		if !hasPut && !hasShortPut {
			ret.noUpdate = true
		}
		if !hasGet && !hasShortGet {
			ret.noDataSource = true
			ret.noImport = true
		}
	}

	var schema map[string]interface{}
	if hasShortPost {
		schema = paths[shortUrl].(map[string]interface{})["post"].(map[string]interface{})
	} else if hasPost {
		schema = paths[url].(map[string]interface{})["post"].(map[string]interface{})
	} else if hasShortPut {
		schema = paths[shortUrl].(map[string]interface{})["put"].(map[string]interface{})
	} else if hasPut {
		schema = paths[url].(map[string]interface{})["put"].(map[string]interface{})
	} else if hasShortGet {
		schema = paths[shortUrl].(map[string]interface{})["get"].(map[string]interface{})
	} else if hasGet {
		schema = paths[url].(map[string]interface{})["get"].(map[string]interface{})
	}
	ret.schema = schema["requestBody"].(map[string]interface{})["content"].(map[string]interface{})["application/json"].(map[string]interface{})

	r := regexp.MustCompile(`{[a-zA-Z]+}`)
	parts := r.Split(url, -1)
	ids := r.FindAllString(url, -1)
	if ret.hasShortUrl {
		ret.resultPath = strings.Join(parts[:len(parts)-1], "%v")
		ret.references = ids[:len(ids)-1]
	} else {
		ret.resultPath = strings.Join(parts, "%v")
		ret.references = ids
	}
	// Derive category from URL
	if strings.Contains(parts[0], "/organizations") {
		ret.category = "Organizations"
	} else if strings.Contains(parts[0], "/networks") {
		ret.category = "Networks"
	} else if strings.Contains(parts[0], "/devices") {
		ret.category = "Devices"
	}
	if len(parts) > 1 {
		if strings.Contains(parts[1], "/switch") {
			ret.category = "Switches"
		} else if strings.Contains(parts[1], "/wireless") {
			ret.category = "Wireless"
		} else if strings.Contains(parts[1], "/appliance") {
			ret.category = "Appliances"
		} else if strings.Contains(parts[1], "/camera") {
			ret.category = "Cameras"
		} else if strings.Contains(parts[1], "/cellularGateway") {
			ret.category = "Cellular Gateways"
		} else if strings.Contains(parts[1], "/insight") {
			ret.category = "Insight"
		} else if strings.Contains(parts[1], "/licensing") {
			ret.category = "Licensing"
		} else if strings.Contains(parts[1], "/sensor") {
			ret.category = "Sensors"
		} else if strings.Contains(parts[1], "/sm") {
			ret.category = "Systems Manager"
		}
	}

	if strings.Contains(parts[0], "/organizations") {
		ret.testVariables = append(ret.testVariables, "test_org")
	} else if strings.Contains(parts[0], "/networks") {
		ret.testVariables = append(ret.testVariables, "test_org")
		ret.testVariables = append(ret.testVariables, "test_network")
	} else if strings.Contains(parts[0], "/devices") {
		ret.testVariables = append(ret.testVariables, "test_org")
		ret.testVariables = append(ret.testVariables, "test_network")
		ret.testVariables = append(ret.testVariables, "test_switch_1_serial")
	}
	return ret
}

func traverseProperties(m map[string]interface{}, path []string, gjsonPath string, exampleStr string, requiredProperties []string) []yamlconfig.YamlConfigAttribute {
	ret := []yamlconfig.YamlConfigAttribute{}

	keys := maps.Keys(m)
	sort.Strings(keys)

	for _, propName := range keys {
		propMap := m[propName].(map[string]interface{})
		if propMap["type"] != "object" && propMap["type"] != "array" {
			// primitive value
			attr := yamlconfig.YamlConfigAttribute{}
			attr.DataPath = path
			attr.Type = jsonTypes[propMap["type"].(string)]
			attr.ModelName = propName
			childGjsonPath := (gjsonPath + "." + propName)[1:]
			res := gjson.Get(exampleStr, childGjsonPath)
			attr.Example = res.String()
			if desc, ok := propMap["description"]; ok {
				attr.Description = sanitizeDescription(desc.(string))
			}
			if enums, ok := propMap["enum"]; ok && attr.Type == "String" {
				for _, r := range enums.([]interface{}) {
					attr.EnumValues = append(attr.EnumValues, r.(string))
				}
			}
			if min, ok := propMap["minimum"]; ok {
				if attr.Type == "Int64" {
					attr.MinInt = min.(int64)
				} else if attr.Type == "Float64" {
					attr.MinFloat = min.(float64)
				}
			}
			if max, ok := propMap["maximum"]; ok {
				if attr.Type == "Int64" {
					attr.MaxInt = max.(int64)
				} else if attr.Type == "Float64" {
					attr.MaxFloat = max.(float64)
				}
			}
			if slices.Contains(requiredProperties, propName) && len(path) == 0 {
				attr.Mandatory = true
			}
			ret = append(ret, attr)
		}
	}
	for _, propName := range keys {
		propMap := m[propName].(map[string]interface{})
		if propMap["type"] == "object" {
			childPath := append(path, propName)
			childGjsonPath := gjsonPath + "." + propName
			childRequired := []string{}
			if prop, ok := propMap["properties"]; ok {
				if rp, ok := propMap["required"]; ok {
					childRequired = toStringSlice(rp.([]interface{}))
				}
				children := traverseProperties(prop.(map[string]interface{}), childPath, childGjsonPath, exampleStr, childRequired)
				ret = append(ret, children...)
			} else {
				attr := yamlconfig.YamlConfigAttribute{}
				attr.DataPath = path
				attr.Type = "Map"
				attr.ModelName = propName
				if desc, ok := propMap["description"]; ok {
					attr.Description = sanitizeDescription(desc.(string))
				}
				if slices.Contains(requiredProperties, propName) && len(path) == 0 {
					attr.Mandatory = true
				}
				ret = append(ret, attr)
			}
		}
	}
	for _, propName := range keys {
		propMap := m[propName].(map[string]interface{})
		if propMap["type"] == "array" {
			attr := yamlconfig.YamlConfigAttribute{}
			attr.DataPath = path
			attr.Type = "List"
			attr.ModelName = propName
			if slices.Contains(requiredProperties, propName) && len(path) == 0 {
				attr.Mandatory = true
			}
			items := propMap["items"].(map[string]interface{})
			if desc, ok := propMap["description"]; ok {
				attr.Description = sanitizeDescription(desc.(string))
			}
			if strings.Contains(attr.Description, "ordered array") {
				attr.OrderedList = true
			}
			if t, ok := jsonTypes[items["type"].(string)]; ok {
				attr.ElementType = t
				childGjsonPath := (gjsonPath + "." + propName + ".0")[1:]
				res := gjson.Get(exampleStr, childGjsonPath)
				attr.Example = res.String()
			} else if items["type"].(string) == "object" {
				childGjsonPath := gjsonPath + "." + propName + ".0"
				childRequired := []string{}
				if rp, ok := items["required"]; ok {
					childRequired = toStringSlice(rp.([]interface{}))
				}
				children := traverseProperties(items["properties"].(map[string]interface{}), []string{}, childGjsonPath, exampleStr, childRequired)
				attr.Attributes = children
			}
			ret = append(ret, attr)
		}
	}
	return ret
}

func sanitizeDescription(desc string) string {
	desc = strings.ReplaceAll(desc, "\n", " ")
	r := regexp.MustCompile("<.*?>")
	desc = r.ReplaceAllString(desc, "")
	desc = strings.ReplaceAll(desc, "'", "`")
	desc = strings.ReplaceAll(desc, "\"", "'")
	desc = strings.Join(strings.Fields(desc), " ") // Remove extra spaces
	return desc
}
