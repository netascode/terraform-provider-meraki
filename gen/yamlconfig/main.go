package yamlconfig

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const EndpointToken = "@endpoint:"

type YamlConfig struct {
	Name                   string                `yaml:"name,omitempty"`
	TfName                 string                `yaml:"tf_name,omitempty"`
	BulkName               string                `yaml:"bulk_name,omitempty"`
	SpecEndpoint           string                `yaml:"spec_endpoint,omitempty"`
	RestEndpoint           string                `yaml:"rest_endpoint,omitempty"`
	NoDataSource           bool                  `yaml:"no_data_source,omitempty"`
	NoResource             bool                  `yaml:"no_resource,omitempty"`
	Action                 bool                  `yaml:"action,omitempty"`
	BulkDataSource         bool                  `yaml:"bulk_data_source,omitempty"`
	BulkResource           bool                  `yaml:"bulk_resource,omitempty"`
	PutCreate              bool                  `yaml:"put_create,omitempty"`
	GetFromAll             bool                  `yaml:"get_from_all,omitempty"`
	NoUpdate               bool                  `yaml:"no_update,omitempty"`
	NoDelete               bool                  `yaml:"no_delete,omitempty"`
	NoImport               bool                  `yaml:"no_import,omitempty"`
	NoRead                 bool                  `yaml:"no_read,omitempty"`
	PostAndPut             bool                  `yaml:"post_and_put,omitempty"`
	IdName                 string                `yaml:"id_name,omitempty"`
	EarlyAccess            bool                  `yaml:"early_access,omitempty"`
	DataSourceNameQuery    bool                  `yaml:"data_source_name_query,omitempty"`
	MinimumVersion         string                `yaml:"minimum_version,omitempty"`
	DsDescription          string                `yaml:"ds_description,omitempty"`
	ResDescription         string                `yaml:"res_description,omitempty"`
	DsBulkDescription      string                `yaml:"ds_bulk_description,omitempty"`
	ResBulkDescription     string                `yaml:"res_bulk_description,omitempty"`
	ActionDescription      string                `yaml:"action_description,omitempty"`
	DocCategory            string                `yaml:"doc_category,omitempty"`
	ExcludeTest            bool                  `yaml:"exclude_test,omitempty"`
	SkipMinimumTest        bool                  `yaml:"skip_minimum_test,omitempty"`
	SkipBulkResourceTest   bool                  `yaml:"skip_bulk_resource_test,omitempty"`
	SkipBulkDataSourceTest bool                  `yaml:"skip_bulk_data_source_test,omitempty"`
	TestTags               []string              `yaml:"test_tags,omitempty,flow"`
	TestVariables          []string              `yaml:"test_variables,omitempty,flow"`
	IgnoreAttributes       []string              `yaml:"ignore_attributes,omitempty,flow"`
	Attributes             []YamlConfigAttribute `yaml:"attributes,omitempty"`
	TestPrerequisites      string                `yaml:"test_prerequisites,omitempty"`
	AdditionalTests        []string              `yaml:"additional_tests,omitempty"`
}

type YamlConfigP struct {
	Name                   *string                 `yaml:"name,omitempty"`
	TfName                 *string                 `yaml:"tf_name,omitempty"`
	BulkName               *string                 `yaml:"bulk_name,omitempty"`
	SpecEndpoint           *string                 `yaml:"spec_endpoint,omitempty"`
	RestEndpoint           *string                 `yaml:"rest_endpoint,omitempty"`
	NoDataSource           *bool                   `yaml:"no_data_source,omitempty"`
	NoResource             *bool                   `yaml:"no_resource,omitempty"`
	Action                 *bool                   `yaml:"action,omitempty"`
	BulkDataSource         *bool                   `yaml:"bulk_data_source,omitempty"`
	BulkResource           *bool                   `yaml:"bulk_resource,omitempty"`
	PutCreate              *bool                   `yaml:"put_create,omitempty"`
	GetFromAll             *bool                   `yaml:"get_from_all,omitempty"`
	NoUpdate               *bool                   `yaml:"no_update,omitempty"`
	NoDelete               *bool                   `yaml:"no_delete,omitempty"`
	NoImport               *bool                   `yaml:"no_import,omitempty"`
	NoRead                 *bool                   `yaml:"no_read,omitempty"`
	PostAndPut             *bool                   `yaml:"post_and_put,omitempty"`
	IdName                 *string                 `yaml:"id_name,omitempty"`
	EarlyAccess            *bool                   `yaml:"early_access,omitempty"`
	DataSourceNameQuery    *bool                   `yaml:"data_source_name_query,omitempty"`
	MinimumVersion         *string                 `yaml:"minimum_version,omitempty"`
	DsDescription          *string                 `yaml:"ds_description,omitempty"`
	ResDescription         *string                 `yaml:"res_description,omitempty"`
	DsBulkDescription      *string                 `yaml:"ds_bulk_description,omitempty"`
	ResBulkDescription     *string                 `yaml:"res_bulk_description,omitempty"`
	ActionDescription      *string                 `yaml:"action_description,omitempty"`
	DocCategory            *string                 `yaml:"doc_category,omitempty"`
	ExcludeTest            *bool                   `yaml:"exclude_test,omitempty"`
	SkipMinimumTest        *bool                   `yaml:"skip_minimum_test,omitempty"`
	SkipBulkResourceTest   *bool                   `yaml:"skip_bulk_resource_test,omitempty"`
	SkipBulkDataSourceTest *bool                   `yaml:"skip_bulk_data_source_test,omitempty"`
	TestTags               *[]string               `yaml:"test_tags,omitempty,flow"`
	TestVariables          *[]string               `yaml:"test_variables,omitempty,flow"`
	IgnoreAttributes       *[]string               `yaml:"ignore_attributes,omitempty,flow"`
	Attributes             *[]YamlConfigAttributeP `yaml:"attributes,omitempty"`
	TestPrerequisites      *string                 `yaml:"test_prerequisites,omitempty"`
	AdditionalTests        *[]string               `yaml:"additional_tests,omitempty"`
}

type YamlConfigAttribute struct {
	ModelName          string                `yaml:"model_name,omitempty"`
	TfName             string                `yaml:"tf_name,omitempty"`
	Type               string                `yaml:"type,omitempty"`
	ElementType        string                `yaml:"element_type,omitempty"`
	DataPath           []string              `yaml:"data_path,omitempty,flow"`
	Id                 bool                  `yaml:"id,omitempty"`
	BulkId             bool                  `yaml:"bulk_id,omitempty"`
	Reference          bool                  `yaml:"reference,omitempty"`
	RequiresReplace    bool                  `yaml:"requires_replace,omitempty"`
	Mandatory          bool                  `yaml:"mandatory,omitempty"`
	WriteOnly          bool                  `yaml:"write_only,omitempty"`
	DataSourceOnly     bool                  `yaml:"data_source_only,omitempty"`
	WriteChangesOnly   bool                  `yaml:"write_changes_only,omitempty"`
	SendNullOnClear    bool                  `yaml:"send_null_on_clear,omitempty"`
	Sensitive          bool                  `yaml:"sensitive,omitempty"`
	WriteEmptyList     bool                  `yaml:"write_empty_list,omitempty"`
	ExcludeTest        bool                  `yaml:"exclude_test,omitempty"`
	ExcludeExample     bool                  `yaml:"exclude_example,omitempty"`
	AllowImportChanges bool                  `yaml:"allow_import_changes,omitempty"`
	Description        string                `yaml:"description,omitempty"`
	Example            string                `yaml:"example,omitempty"`
	MapKeyExample      string                `yaml:"map_key_example,omitempty"`
	EnumValues         []string              `yaml:"enum_values,omitempty,flow"`
	MinList            int64                 `yaml:"min_list,omitempty"`
	MaxList            int64                 `yaml:"max_list,omitempty"`
	MinInt             int64                 `yaml:"min_int,omitempty"`
	MaxInt             int64                 `yaml:"max_int,omitempty"`
	MinFloat           float64               `yaml:"min_float,omitempty"`
	MaxFloat           float64               `yaml:"max_float,omitempty"`
	OrderedList        bool                  `yaml:"ordered_list,omitempty"`
	StringPatterns     []string              `yaml:"string_patterns,omitempty,flow"`
	StringMinLength    int64                 `yaml:"string_min_length,omitempty"`
	StringMaxLength    int64                 `yaml:"string_max_length,omitempty"`
	Computed           bool                  `yaml:"computed,omitempty"`
	DefaultValue       string                `yaml:"default_value,omitempty"`
	Value              string                `yaml:"value,omitempty"`
	TestValue          string                `yaml:"test_value,omitempty"`
	MinimumTestValue   string                `yaml:"minimum_test_value,omitempty"`
	DestroyValue       string                `yaml:"destroy_value,omitempty"`
	IgnoreImportValues []string              `yaml:"ignore_import_values,omitempty,flow"`
	TestTags           []string              `yaml:"test_tags,omitempty,flow"`
	Attributes         []YamlConfigAttribute `yaml:"attributes,omitempty"`
	GoTypeName         string                `yaml:"gotypename,omitempty"`
	GoTypeBulkName     string                `yaml:"gobulktypename,omitempty"`
}

type YamlConfigAttributeP struct {
	ModelName          *string                 `yaml:"model_name,omitempty"`
	TfName             *string                 `yaml:"tf_name,omitempty"`
	Type               *string                 `yaml:"type,omitempty"`
	ElementType        *string                 `yaml:"element_type,omitempty"`
	DataPath           *[]string               `yaml:"data_path,omitempty,flow"`
	Id                 *bool                   `yaml:"id,omitempty"`
	BulkId             *bool                   `yaml:"bulk_id,omitempty"`
	Reference          *bool                   `yaml:"reference,omitempty"`
	RequiresReplace    *bool                   `yaml:"requires_replace,omitempty"`
	Mandatory          *bool                   `yaml:"mandatory,omitempty"`
	WriteOnly          *bool                   `yaml:"write_only,omitempty"`
	DataSourceOnly     *bool                   `yaml:"data_source_only,omitempty"`
	WriteChangesOnly   *bool                   `yaml:"write_changes_only,omitempty"`
	SendNullOnClear    *bool                   `yaml:"send_null_on_clear,omitempty"`
	Sensitive          *bool                   `yaml:"sensitive,omitempty"`
	WriteEmptyList     *bool                   `yaml:"write_empty_list,omitempty"`
	ExcludeTest        *bool                   `yaml:"exclude_test,omitempty"`
	ExcludeExample     *bool                   `yaml:"exclude_example,omitempty"`
	AllowImportChanges *bool                   `yaml:"allow_import_changes,omitempty"`
	Description        *string                 `yaml:"description,omitempty"`
	Example            *string                 `yaml:"example,omitempty"`
	MapKeyExample      *string                 `yaml:"map_key_example,omitempty"`
	EnumValues         *[]string               `yaml:"enum_values,omitempty,flow"`
	MinList            *int64                  `yaml:"min_list,omitempty"`
	MaxList            *int64                  `yaml:"max_list,omitempty"`
	MinInt             *int64                  `yaml:"min_int,omitempty"`
	MaxInt             *int64                  `yaml:"max_int,omitempty"`
	MinFloat           *float64                `yaml:"min_float,omitempty"`
	MaxFloat           *float64                `yaml:"max_float,omitempty"`
	OrderedList        *bool                   `yaml:"ordered_list,omitempty"`
	StringPatterns     *[]string               `yaml:"string_patterns,omitempty,flow"`
	StringMinLength    *int64                  `yaml:"string_min_length,omitempty"`
	StringMaxLength    *int64                  `yaml:"string_max_length,omitempty"`
	Computed           *bool                   `yaml:"computed,omitempty"`
	DefaultValue       *string                 `yaml:"default_value,omitempty"`
	Value              *string                 `yaml:"value,omitempty"`
	TestValue          *string                 `yaml:"test_value,omitempty"`
	MinimumTestValue   *string                 `yaml:"minimum_test_value,omitempty"`
	DestroyValue       *string                 `yaml:"destroy_value,omitempty"`
	IgnoreImportValues *[]string               `yaml:"ignore_import_values,omitempty,flow"`
	TestTags           *[]string               `yaml:"test_tags,omitempty,flow"`
	Attributes         *[]YamlConfigAttributeP `yaml:"attributes,omitempty"`
	GoTypeName         *string                 `yaml:"gotypename,omitempty"`
	GoTypeBulkName     *string                 `yaml:"gobulktypename,omitempty"`
}

func P[T any](v T) *T {
	return &v
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "/", "")
	p := strings.Split(s, "_")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "/", "")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "/", "")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// Templating helper function to fail a template mid-way
func Errorf(s string, args ...any) (struct{}, error) {
	return struct{}{}, fmt.Errorf(s, args...)
}

// Templating helper function to build a SJSON path
func BuildPath(s []string) string {
	return strings.Join(s, ".")
}

// Templating helper function to return true if id included in attributes
func HasId(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Id {
			return true
		}
	}
	return false
}

// Templating helper function to return the id
func GetId(attributes []YamlConfigAttribute) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.Id {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to return true if reference included in attributes
func HasReference(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Reference {
			return true
		}
	}
	return false
}

// Templating helper function to return true if type is a list or set without nested elements
func IsListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list without nested elements
func IsList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set without nested elements
func IsSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of strings without nested elements
func IsStringListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "String" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of integers without nested elements
func IsInt64ListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "Int64" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set with nested elements
func IsNestedListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list, set or map with nested elements
func IsNestedListSetMap(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set" || attribute.Type == "Map") && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list with nested elements
func IsNestedList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set with nested elements
func IsNestedSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a map with nested elements
func IsNestedMap(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Map" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return all import attributes
func ImportAttributes(config YamlConfig) []YamlConfigAttribute {
	attributes := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if attr.Reference {
			attributes = append(attributes, attr)
		} else if attr.Id {
			attributes = append(attributes, attr)
		} else if attr.ModelName == "" {
			attributes = append(attributes, attr)
		}
	}
	if !HasId(config.Attributes) {
		attributes = append(attributes, YamlConfigAttribute{TfName: "id", Type: "String"})
	}

	return attributes
}

// Templating helper function to subtract one number from another
func Subtract(a, b int) int {
	return a - b
}

// Templating helper function to iterate over a count
func Iterate(count int) []int {
	var i int
	var Items []int
	for i = 0; i < count; i++ {
		Items = append(Items, i)
	}
	return Items
}

// GetImportExcludes returns a list of attributes to exclude from import testing
func GetImportExcludes(attributes []YamlConfigAttribute) []string {
	var excludes []string
	for _, attr := range attributes {
		if attr.WriteOnly || attr.AllowImportChanges {
			excludes = append(excludes, attr.TfName)
		}
		if len(attr.Attributes) > 0 {
			ca := GetImportExcludes(attr.Attributes)
			for _, c := range ca {
				excludes = append(excludes, attr.TfName+".0."+c)
			}
		}
	}
	return excludes
}

// GetFullModelName returns the full model name for an attribute, including data path
func GetFullModelName(attr YamlConfigAttribute, write bool) string {
	r := ""
	for i := range attr.DataPath {
		if _, err := strconv.Atoi(attr.DataPath[i]); err == nil && write {
			r += fmt.Sprintf(":%s.", attr.DataPath[i])
		} else {
			r += attr.DataPath[i] + "."
		}
	}
	return r + attr.ModelName
}

// HasComputedAttributes returns true if any attributes are computed
func HasComputedAttributes(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Computed {
			return true
		}
		if len(attr.Attributes) > 0 && HasComputedAttributes(attr.Attributes) {
			return true
		}
	}
	return false
}

// Templating helper function to build a test path
func BuildTestPath(attr ...YamlConfigAttribute) string {
	var path []string
	for _, a := range attr {
		path = append(path, a.TfName)
		if a.Type == "Map" && a.MapKeyExample != "" {
			path = append(path, a.MapKeyExample)
		} else {
			path = append(path, "0")

		}
	}
	return strings.Join(path, ".") + "."
}

// HasDestroyValues returns true if any attributes have explicit destroy values
func HasDestroyValues(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.DestroyValue != "" {
			return true
		}
		if len(attr.Attributes) > 0 && HasDestroyValues(attr.Attributes) {
			return true
		}
	}
	return false
}

// GetBulkPath returns the bulk path for a given path
func GetBulkPath(path string) string {
	// Remove the last element from the slash separated path, if the last element starts with "%"
	parts := strings.Split(path, "/")
	if len(parts) > 1 && strings.HasPrefix(parts[len(parts)-1], "%") {
		return strings.Join(parts[:len(parts)-1], "/")
	}
	return path
}

// GetBulkParentAttributes returns a list of parent attributes that are used in bulk operations
func GetBulkParentAttributes(config YamlConfig) []YamlConfigAttribute {
	var parentAttributes []YamlConfigAttribute
	for _, attr := range config.Attributes {
		if !attr.Reference {
			continue
		}
		if attr.Id && config.PutCreate {
			continue
		}
		parentAttributes = append(parentAttributes, attr)
	}
	return parentAttributes
}

// GetBulkItemAttributes returns a list of item attributes that are visible to the resource in bulk operations
// (i.e. excluding data-source-only attributes, which are only ever emitted into the data source's model/schema).
func GetBulkItemAttributes(config YamlConfig) []YamlConfigAttribute {
	var itemAttributes []YamlConfigAttribute
	for _, attr := range config.Attributes {
		if attr.Reference && attr.Id && config.PutCreate {
			itemAttributes = append(itemAttributes, attr)
			continue
		} else if attr.Reference || attr.DataSourceOnly {
			continue
		}
		itemAttributes = append(itemAttributes, attr)
	}
	return itemAttributes
}

// GetBulkImportAttributes returns a list of import attributes that are used in bulk operations
func GetBulkImportAttributes(config YamlConfig) []YamlConfigAttribute {
	var importAttributes []YamlConfigAttribute
	if !HasOrganizationId(config) {
		importAttributes = append(importAttributes, YamlConfigAttribute{ModelName: "organizationId", TfName: "organization_id", Type: "String"})
	}
	for _, attr := range config.Attributes {
		if !attr.Reference && attr.ModelName != "" {
			continue
		}
		if attr.Id && config.PutCreate {
			continue
		}
		importAttributes = append(importAttributes, attr)
	}
	return importAttributes
}

// HasOrganizationId checks if the list of attributes contains an organization_id attribute
func HasOrganizationId(config YamlConfig) bool {
	for _, attr := range config.Attributes {
		if attr.TfName == "organization_id" {
			return true
		}
	}
	return false
}

// GetBulkItemId returns the Go attribute name of the bulk item id
func GetBulkItemId(config YamlConfig) string {
	// Get the id attribute from the bulk item attributes
	for _, attr := range config.Attributes {
		if attr.Id {
			return ToGoName(attr.TfName)
		}
	}
	// If no id attribute is found return default "id"
	return ToGoName("id")
}

// GetBulkItemIdTfName returns the Terraform attribute name of the bulk item id
func GetBulkItemIdTfName(config YamlConfig) string {
	// Get the id attribute from the bulk item attributes
	for _, attr := range config.Attributes {
		if attr.Id {
			return attr.TfName
		}
	}
	// If no id attribute is found return default "id"
	return "id"
}

// IsSingleton returns true if the resource is a singleton
func IsSingleton(config YamlConfig) bool {
	return config.PutCreate && config.NoDelete && !config.NoRead
}

func HasSensitiveAttr(attrs []YamlConfigAttribute) bool {
	for _, attr := range attrs {
		if attr.Sensitive && attr.Type == "String" {
			return true
		}
	}
	return false
}

func HasSensitiveAttrRecursive(attrs []YamlConfigAttribute) bool {
	for _, attr := range attrs {
		if attr.Sensitive && attr.Type == "String" {
			return true
		}
		if HasSensitiveAttrRecursive(attr.Attributes) {
			return true
		}
	}
	return false
}

func HasWriteOnly(attrs []YamlConfigAttribute) bool {
	for _, attr := range attrs {
		if attr.WriteOnly && !attr.ExcludeTest {
			return true
		}
	}
	return false
}

// Map of templating functions
var Functions = template.FuncMap{
	"toGoName":                  ToGoName,
	"camelCase":                 CamelCase,
	"snakeCase":                 SnakeCase,
	"sprintf":                   fmt.Sprintf,
	"errorf":                    Errorf,
	"toLower":                   strings.ToLower,
	"path":                      BuildPath,
	"hasId":                     HasId,
	"getId":                     GetId,
	"hasReference":              HasReference,
	"isListSet":                 IsListSet,
	"isList":                    IsList,
	"isSet":                     IsSet,
	"isStringListSet":           IsStringListSet,
	"isInt64ListSet":            IsInt64ListSet,
	"isNestedListSet":           IsNestedListSet,
	"isNestedListSetMap":        IsNestedListSetMap,
	"isNestedList":              IsNestedList,
	"isNestedSet":               IsNestedSet,
	"isNestedMap":               IsNestedMap,
	"importAttributes":          ImportAttributes,
	"subtract":                  Subtract,
	"iterate":                   Iterate,
	"getImportExcludes":         GetImportExcludes,
	"getFullModelName":          GetFullModelName,
	"hasComputedAttributes":     HasComputedAttributes,
	"buildTestPath":             BuildTestPath,
	"hasDestroyValues":          HasDestroyValues,
	"getBulkPath":               GetBulkPath,
	"getBulkParentAttributes":   GetBulkParentAttributes,
	"getBulkItemAttributes":     GetBulkItemAttributes,
	"getBulkImportAttributes":   GetBulkImportAttributes,
	"hasOrganizationId":         HasOrganizationId,
	"getBulkItemId":             GetBulkItemId,
	"getBulkItemIdTfName":       GetBulkItemIdTfName,
	"isSingleton":               IsSingleton,
	"hasSensitiveAttr":          HasSensitiveAttr,
	"hasSensitiveAttrRecursive": HasSensitiveAttrRecursive,
	"hasWriteOnly":              HasWriteOnly,
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func CamelToSnake(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (attr *YamlConfigAttribute) Init(parentGoTypeName, parentGoTypeBulkName string, parentDataSourceOnly bool) error {
	// Augument
	if attr.TfName == "" {
		fullString := ""
		for _, s := range attr.DataPath {
			fullString += CamelToSnake(strings.ToUpper(string(s[0]))+s[1:]) + "_"
		}
		fullString += CamelToSnake(strings.ToUpper(string(attr.ModelName[0])) + attr.ModelName[1:])

		fullString = strings.ReplaceAll(fullString, "/", "_")
		attr.TfName = fullString
	}

	attr.GoTypeName = parentGoTypeName + ToGoName(attr.TfName)
	attr.GoTypeBulkName = parentGoTypeBulkName + ToGoName(attr.TfName)

	// A data-source-only container cascades the flag onto every descendant attribute.
	attr.DataSourceOnly = attr.DataSourceOnly || parentDataSourceOnly

	// Validate
	if len(attr.Attributes) > 0 && attr.Type != "List" && attr.Type != "Map" && attr.Type != "Set" {
		return fmt.Errorf("%q has type %q which cannot have `attributes`: instead use type List, Map, Set",
			attr.TfName, attr.Type)
	}

	if len(attr.Attributes) > 0 && attr.ElementType != "" {
		return fmt.Errorf("%q: either `attributes` or `element_type` can be specified, but not both", attr.TfName)
	}

	if attr.Type == "Map" && attr.ElementType != "" {
		return fmt.Errorf("%q: the `element_type` is not yet implemented for type Map", attr.TfName)
	}

	if attr.OrderedList {
		if attr.Type != "List" {
			return fmt.Errorf("%q has type %q which cannot use `ordered_list`: instead use type List",
				attr.TfName, attr.Type)
		}
		if HasId(attr.Attributes) {
			return fmt.Errorf("%q: the `ordered_list: true` conflicts with sub-attributes having `id: true`, as it treats list index ([i]) as the only unique id",
				attr.TfName)
		}
	}

	if attr.Type == "Map" && HasId(attr.Attributes) {
		return fmt.Errorf("Map %q cannot contain sub-attributes with `id: true`, as it treats map key ([k]) as the only unique id",
			attr.TfName)
	}

	if attr.DataSourceOnly {
		if attr.Mandatory {
			return fmt.Errorf("%q: `data_source_only` cannot be combined with `mandatory`, since data-source-only attributes are never part of the resource's writable input", attr.TfName)
		}
		if attr.RequiresReplace {
			return fmt.Errorf("%q: `data_source_only` cannot be combined with `requires_replace`, since that only applies to resource plan diffs", attr.TfName)
		}
		if attr.WriteChangesOnly {
			return fmt.Errorf("%q: `data_source_only` cannot be combined with `write_changes_only`, since that only applies to the resource's PUT body diffing", attr.TfName)
		}
		if attr.SendNullOnClear {
			return fmt.Errorf("%q: `data_source_only` cannot be combined with `send_null_on_clear`, since that only applies to the resource's PUT body clearing", attr.TfName)
		}
		if attr.WriteOnly {
			return fmt.Errorf("%q: `data_source_only` cannot be combined with `write_only`, they are opposites (write-only means unreadable, data-source-only means unwritable)", attr.TfName)
		}
		if attr.Reference || attr.Id {
			return fmt.Errorf("%q: `data_source_only` cannot be combined with `reference`/`id`, since those drive the resource's REST path/import key", attr.TfName)
		}
		if attr.BulkId {
			return fmt.Errorf("%q: `data_source_only` cannot be combined with `bulk_id`, since that drives the resource's per-item composite key", attr.TfName)
		}
	}

	// Recurse
	for i := range attr.Attributes {
		if err := attr.Attributes[i].Init(attr.GoTypeName, attr.GoTypeBulkName, attr.DataSourceOnly); err != nil {
			return err
		}
	}

	return nil
}

func NewYamlConfig(bytes []byte) (YamlConfig, error) {
	var config YamlConfig

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return config, err
	}

	if config.BulkName == "" {
		if string(config.Name[len(config.Name)-1]) == "y" {
			config.BulkName = config.Name[:len(config.Name)-1] + "ies"
		} else {
			config.BulkName = config.Name + "s"
		}
	}

	for i := range config.Attributes {
		if err := config.Attributes[i].Init(CamelCase(config.Name), CamelCase(config.BulkName), false); err != nil {
			return YamlConfig{}, err
		}
	}
	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the `%s` configuration.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage the `%s` configuration.", config.Name)
	}
	if config.DsBulkDescription == "" {
		config.DsBulkDescription = fmt.Sprintf("This data source can read the `%s` configuration in bulk.", config.Name)
	}
	if config.ResBulkDescription == "" {
		config.ResBulkDescription = fmt.Sprintf("This resource can manage the `%s` configuration in bulk.", config.Name)
	}
	if config.ActionDescription == "" {
		config.ActionDescription = fmt.Sprintf("This action can invoke the `%s` operation.", config.Name)
	}
	if config.TfName == "" {
		config.TfName = strings.Replace(config.Name, " ", "_", -1)
	}
	if config.IdName == "" {
		config.IdName = "id"
	}

	return config, nil
}

func MergeYamlConfig(existing *YamlConfigP, new *YamlConfigP) *YamlConfigP {
	// Merge existing into new
	if existing.Name != nil {
		new.Name = existing.Name
	}
	if existing.TfName != nil {
		new.TfName = existing.TfName
	}
	if existing.BulkName != nil {
		new.BulkName = existing.BulkName
	}
	if existing.SpecEndpoint != nil {
		new.SpecEndpoint = existing.SpecEndpoint
	}
	if existing.RestEndpoint != nil {
		new.RestEndpoint = existing.RestEndpoint
	}
	if existing.NoDataSource != nil {
		new.NoDataSource = existing.NoDataSource
	}
	if existing.NoResource != nil {
		new.NoResource = existing.NoResource
	}
	if existing.Action != nil {
		new.Action = existing.Action
	}
	if existing.BulkDataSource != nil {
		new.BulkDataSource = existing.BulkDataSource
	}
	if existing.BulkResource != nil {
		new.BulkResource = existing.BulkResource
	}
	if existing.PutCreate != nil {
		new.PutCreate = existing.PutCreate
	}
	if existing.GetFromAll != nil {
		new.GetFromAll = existing.GetFromAll
	}
	if existing.NoUpdate != nil {
		new.NoUpdate = existing.NoUpdate
	}
	if existing.NoDelete != nil {
		new.NoDelete = existing.NoDelete
	}
	if existing.NoImport != nil {
		new.NoImport = existing.NoImport
	}
	if existing.NoRead != nil {
		new.NoRead = existing.NoRead
	}
	if existing.PostAndPut != nil {
		new.PostAndPut = existing.PostAndPut
	}
	if existing.IdName != nil {
		new.IdName = existing.IdName
	}
	if existing.DataSourceNameQuery != nil {
		new.DataSourceNameQuery = existing.DataSourceNameQuery
	}
	if existing.MinimumVersion != nil {
		new.MinimumVersion = existing.MinimumVersion
	}
	if existing.DsDescription != nil {
		new.DsDescription = existing.DsDescription
	}
	if existing.ResDescription != nil {
		new.ResDescription = existing.ResDescription
	}
	if existing.DsBulkDescription != nil {
		new.DsBulkDescription = existing.DsBulkDescription
	}
	if existing.ResBulkDescription != nil {
		new.ResBulkDescription = existing.ResBulkDescription
	}
	if existing.ActionDescription != nil {
		new.ActionDescription = existing.ActionDescription
	}
	if existing.DocCategory != nil {
		new.DocCategory = existing.DocCategory
	}
	if existing.ExcludeTest != nil {
		new.ExcludeTest = existing.ExcludeTest
	}
	if existing.SkipMinimumTest != nil {
		new.SkipMinimumTest = existing.SkipMinimumTest
	}
	if existing.SkipBulkResourceTest != nil {
		new.SkipBulkResourceTest = existing.SkipBulkResourceTest
	}
	if existing.SkipBulkDataSourceTest != nil {
		new.SkipBulkDataSourceTest = existing.SkipBulkDataSourceTest
	}
	if existing.TestTags != nil {
		new.TestTags = existing.TestTags
	}
	if existing.TestVariables != nil {
		new.TestVariables = existing.TestVariables
	}
	if existing.IgnoreAttributes != nil {
		new.IgnoreAttributes = existing.IgnoreAttributes
	}
	if existing.Attributes != nil {
		new.Attributes = MergeYamlConfigAttributes(existing.Attributes, new.Attributes)
	}
	if existing.TestPrerequisites != nil {
		new.TestPrerequisites = existing.TestPrerequisites
	}
	if existing.AdditionalTests != nil {
		new.AdditionalTests = existing.AdditionalTests
	}
	return new
}

func cutYamlConfigAttribute(i int, attributes []YamlConfigAttributeP) []YamlConfigAttributeP {
	cut := append(attributes[:i], attributes[i+1:]...)
	return cut
}

func MergeYamlConfigAttributes(existing *[]YamlConfigAttributeP, new *[]YamlConfigAttributeP) *[]YamlConfigAttributeP {
	if new == nil {
		return existing
	}
	// Merge existing into new
	var c []YamlConfigAttributeP
	for _, existingAttr := range *existing {
		found := -1
		for i, attr := range *new {
			if existingAttr.ModelName != nil && *existingAttr.ModelName != "" && attr.ModelName != nil && *attr.ModelName != "" && *existingAttr.ModelName == *attr.ModelName {
				if (existingAttr.DataPath == nil && attr.DataPath == nil) || (existingAttr.DataPath != nil && attr.DataPath != nil && slices.Equal(*existingAttr.DataPath, *attr.DataPath)) {
					c = append(c, *MergeYamlConfigAttribute(&existingAttr, &attr))
					found = i
					break
				}
			} else if existingAttr.TfName != nil && *existingAttr.TfName != "" && attr.TfName != nil && *attr.TfName != "" && *existingAttr.TfName == *attr.TfName {
				c = append(c, *MergeYamlConfigAttribute(&existingAttr, &attr))
				found = i
				break
			}
		}
		if found > -1 {
			// we have found a matching new attribute, so we remove it from the list
			new = P(cutYamlConfigAttribute(found, *new))
			// check if there are more elements from the new list and add them if there is no matching attribute in existing list
			foundCount := 0
			for _, attr := range *new {
				foundExisting := false
				for _, eAttr := range *existing {
					if eAttr.ModelName != nil && *eAttr.ModelName != "" && attr.ModelName != nil && *attr.ModelName != "" && *eAttr.ModelName == *attr.ModelName {
						if (eAttr.DataPath == nil && attr.DataPath == nil) || (eAttr.DataPath != nil && attr.DataPath != nil && slices.Equal(*eAttr.DataPath, *attr.DataPath)) {
							foundExisting = true
							break
						}
					} else if eAttr.TfName != nil && *eAttr.TfName != "" && attr.TfName != nil && *attr.TfName != "" && *eAttr.TfName == *attr.TfName {
						foundExisting = true
						break
					}
				}
				if foundExisting {
					// if there is a matching attribute in existing list, we stop as it will be added later
					break
				} else {
					// if there is no matching attribute in existing list, it is a new element and we add it to the list
					foundCount += 1
					c = append(c, attr)
				}
			}
			for i := 0; i < foundCount; i++ {
				// if we have found a new elements that were not in the existing list, we remove them from the new list
				new = P(cutYamlConfigAttribute(0, *new))
			}
		} else {
			// we haven't found a matching new attribute, so we add the existing to the list
			c = append(c, *MergeYamlConfigAttribute(&YamlConfigAttributeP{}, &existingAttr))
		}
	}
	// add remaining elements from existing list
	c = append(c, *new...)
	return &c
}

func MergeYamlConfigAttribute(existing *YamlConfigAttributeP, new *YamlConfigAttributeP) *YamlConfigAttributeP {
	// Merge existing into new
	if existing.ModelName != nil {
		new.ModelName = existing.ModelName
	}
	if existing.TfName != nil {
		new.TfName = existing.TfName
	}
	if existing.Type != nil {
		new.Type = existing.Type
	}
	if existing.ElementType != nil {
		new.ElementType = existing.ElementType
	}
	if existing.DataPath != nil {
		new.DataPath = existing.DataPath
	}
	if existing.Id != nil {
		new.Id = existing.Id
	}
	if existing.Reference != nil {
		new.Reference = existing.Reference
	}
	if existing.BulkId != nil {
		new.BulkId = existing.BulkId
	}
	if existing.RequiresReplace != nil {
		new.RequiresReplace = existing.RequiresReplace
	}
	if existing.Mandatory != nil {
		new.Mandatory = existing.Mandatory
	}
	if existing.WriteOnly != nil {
		new.WriteOnly = existing.WriteOnly
	}
	if existing.DataSourceOnly != nil {
		new.DataSourceOnly = existing.DataSourceOnly
	}
	if existing.WriteChangesOnly != nil {
		new.WriteChangesOnly = existing.WriteChangesOnly
	}
	if existing.SendNullOnClear != nil {
		new.SendNullOnClear = existing.SendNullOnClear
	}
	if existing.Sensitive != nil {
		new.Sensitive = existing.Sensitive
	}
	if existing.WriteEmptyList != nil {
		new.WriteEmptyList = existing.WriteEmptyList
	}
	if existing.ExcludeTest != nil {
		new.ExcludeTest = existing.ExcludeTest
	}
	if existing.ExcludeExample != nil {
		new.ExcludeExample = existing.ExcludeExample
	}
	if existing.AllowImportChanges != nil {
		new.AllowImportChanges = existing.AllowImportChanges
	}
	if existing.Description != nil {
		new.Description = existing.Description
	}
	if existing.Example != nil {
		new.Example = existing.Example
	}
	if existing.MapKeyExample != nil {
		new.MapKeyExample = existing.MapKeyExample
	}
	if existing.EnumValues != nil {
		new.EnumValues = existing.EnumValues
	}
	if existing.MinList != nil {
		new.MinList = existing.MinList
	}
	if existing.MaxList != nil {
		new.MaxList = existing.MaxList
	}
	if existing.MinInt != nil {
		new.MinInt = existing.MinInt
	}
	if existing.MaxInt != nil {
		new.MaxInt = existing.MaxInt
	}
	if existing.MinFloat != nil {
		new.MinFloat = existing.MinFloat
	}
	if existing.MaxFloat != nil {
		new.MaxFloat = existing.MaxFloat
	}
	if existing.OrderedList != nil {
		new.OrderedList = existing.OrderedList
	}
	if existing.StringPatterns != nil {
		new.StringPatterns = existing.StringPatterns
	}
	if existing.StringMinLength != nil {
		new.StringMinLength = existing.StringMinLength
	}
	if existing.StringMaxLength != nil {
		new.StringMaxLength = existing.StringMaxLength
	}
	if existing.Computed != nil {
		new.Computed = existing.Computed
	}
	if existing.DefaultValue != nil {
		new.DefaultValue = existing.DefaultValue
	}
	if existing.Value != nil {
		new.Value = existing.Value
	}
	if existing.TestValue != nil {
		new.TestValue = existing.TestValue
	}
	if existing.MinimumTestValue != nil {
		new.MinimumTestValue = existing.MinimumTestValue
	}
	if existing.DestroyValue != nil {
		new.DestroyValue = existing.DestroyValue
	}
	if existing.IgnoreImportValues != nil {
		new.IgnoreImportValues = existing.IgnoreImportValues
	}
	if existing.TestTags != nil {
		new.TestTags = existing.TestTags
	}
	if existing.Attributes != nil {
		new.Attributes = MergeYamlConfigAttributes(existing.Attributes, new.Attributes)
	}
	return new
}
