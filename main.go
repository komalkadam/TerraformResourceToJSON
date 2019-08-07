package main

import (
	//"github.com/hashicorp/terraform/helper/schema"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/terraform-providers/terraform-provider-aws/aws"
)

type TerraformStyle struct {
	Body              StyleBody         `json:".body"`
	ResourceNode_Text ResourceNode_text `json:".resourceNode-text"`
}
type StyleBody struct {
	FillObject FillObject1 `json:"fill"`
}
type FillObject1 struct {
	ObjectType string               `json:"type"`
	Stops      [2]map[string]string `json:"stops"`
	Attributes map[string]string    `json:"attrs"`
}
type ResourceNode_text struct {
	Fill string `json:"fill"`
}

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
	Style       TerraformStyle       `json:"style"`
	Image       string               `json:"image"`
	Provider    string               `json:"provider"`
	Attributes  []TerraformAttribute `json:"attributes"`
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	var v = aws.GetResourceSchema()
	//Update the resource name here
	var resourceName string = os.Args[1]
	var providerType string = os.Args[2]

	var resource TerraformResource = TerraformResource{Name: resourceName}
	var terraformAttributes []TerraformAttribute = []TerraformAttribute{}
	resource.Attributes = terraformAttributes
	resource.ShortName = strings.Title(strings.ReplaceAll(resourceName, "_", " "))
	resource.DisplayName = strings.Title(strings.ReplaceAll(resourceName, "_", " "))

	resource.Provider = providerType

	if providerType == "aws" {
		resource.Image = "/images/aws/ec2/Compute_AmazonEC2_instance.png"
	} else if providerType == "vsphere" {
		resource.Image = "/images/vsphere/vm.png"
	}

	resource.Style = TerraformStyle{}
	resource.Style.Body = StyleBody{}
	resource.Style.Body.FillObject = FillObject1{}
	resource.Style.ResourceNode_Text = ResourceNode_text{}
	resource.Style.ResourceNode_Text.Fill = "#333"
	resource.Style.Body.FillObject.ObjectType = "linearGradient"
	resource.Style.Body.FillObject.Attributes = make(map[string]string)
	resource.Style.Body.FillObject.Attributes["x1"] = "0%"
	resource.Style.Body.FillObject.Attributes["x2"] = "0%"
	resource.Style.Body.FillObject.Attributes["y1"] = "0%"
	resource.Style.Body.FillObject.Attributes["y2"] = "100%"

	var stops [2]map[string]string
	stops[0] = make(map[string]string)
	stops[0]["offset"] = "0%"
	stops[0]["color"] = "#F2F2F2"

	stops[1] = make(map[string]string)
	stops[1]["offset"] = "100%"
	stops[1]["color"] = "#F2F2F2"

	resource.Style.Body.FillObject.Stops = stops
	fmt.Println("attributes  ::", Attributes_array)

	for k, va := range v.Schema {

		/* if !va.Optional && va.Computed {
			continue
		}

		if va.Optional && va.Computed {
			fmt.Println("Verfiy attribute its marked as Optional and computed as well ::", k)
		} */
		//fmt.Println(k)

		if !Contains(Attributes_array, k) {
			//fmt.Println("======" + k)
			continue
		}

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
			//fmt.Println("Default value:", va.Default)
		}

		if k == "tags" {
			resource.TagsSupport = true
			attribute.AttributeType = "JSON"
			terraformAttributes = append(terraformAttributes, attribute)
			continue
		}

		switch dataTypeStr {
		case "TypeString":
			attribute.AttributeType = "String"

		case "TypeInt":
			attribute.AttributeType = "Integer"

		case "TypeList":
			//attribute.AttributeType = "StringArray"
			if va.Elem != nil {
				var schemaDataType = reflect.TypeOf(va.Elem).String()

				if schemaDataType == "*schema.Resource" {
					attribute.AttributeType = "JSON"
					attribute.SampleValue = getSampleValue(va.Elem, attribute)
				} else {
					attribute.AttributeType = "JSONArray"

				}

			}

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
