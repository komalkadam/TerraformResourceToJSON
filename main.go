package main

import (
	//"github.com/hashicorp/terraform/helper/schema"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/terraform-providers/terraform-provider-aws/aws"
)

type TerraformAttribute struct {
	Name           string   `json:"name"`
	AttributeType  string   `json:"type"`
	Required       string   `json:"required"`
	DefaultValue   string   `json:"defaultValue,omitempty"`
	PossibleValues []string `json:"possibleValues,omitempty"`
	SampleValue    string   `json:"sampleValue,omitempty"`
}

type TerraformResource struct {
	Name        string `json:"name"`
	ShortName   string `json:"shortName"`
	DisplayName string `json:"displayName"`

	TagsSupport bool                 `json:"tagsSupport"`
	IdAttribute string               `json:"idAttribute"`
	Style       string               `json:"style"`
	Image       string               `json:"image"`
	Provider    string               `json:"provider"`
	Attributes  []TerraformAttribute `json:attributes`
}

func main() {
	var v = aws.GetResourceSchema()
	var resource TerraformResource = TerraformResource{Name: "aws_instance"}
	var terraformAttributes []TerraformAttribute = []TerraformAttribute{}
	resource.Attributes = terraformAttributes

	for k, va := range v.Schema {
		var attribute TerraformAttribute = TerraformAttribute{}
		attribute.Name = k
		var dataTypeStr = va.Type.String()
		if va.Required {
			attribute.Required = "true"
		} else {
			attribute.Required = "false"
		}

		if va.Default != nil {
			attribute.DefaultValue = fmt.Sprintf("%v", va.Default)
		}

		switch dataTypeStr {
		case "TypeString":
			attribute.AttributeType = "String"

		case "TypeInt":
			attribute.AttributeType = "Integer"

		case "TypeList":
			attribute.AttributeType = "StringArray"

		case "TypeBool":
			attribute.AttributeType = "boolean"

		case "TypeSet":

			if va.Elem != nil {
				var schemaDataType = reflect.TypeOf(va.Elem).String()

				if schemaDataType == "*schema.Resource" {
					attribute.AttributeType = "JSON"
					attribute.SampleValue = getSampleValue(va.Elem, attribute)
				} else {
					attribute.AttributeType = "JSONArray"

				}

			} /* else {

			} */
		case "TypeMap":

			if va.Elem != nil {
				var schemaDataType = reflect.TypeOf(va.Elem).String()

				if schemaDataType == "*schema.Resource" {
					attribute.AttributeType = "JSON"
					attribute.SampleValue = getSampleValue(va.Elem, attribute)
				} else {
					attribute.AttributeType = "JSONArray"

				}
			}
		default:

			attribute.AttributeType = "Not found"
		}
		terraformAttributes = append(terraformAttributes, attribute)

	}
	resource.Attributes = terraformAttributes
	res2B, _ := json.MarshalIndent(resource, "", "\t")
	_ = ioutil.WriteFile(resource.Name+".json", res2B, 0644)

}
func getSampleValue(d interface{}, attribute TerraformAttribute) string {

	s := reflect.ValueOf(d).Elem()

	f := s.FieldByName("Schema")

	var str string = fmt.Sprintf("%v", f.Interface())
	fmt.Println("Update sample value of attribute name: ", attribute.Name)
	return str

}
