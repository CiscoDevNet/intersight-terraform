/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowPrimitiveDataType This data type is used to represent primitives like string, floats and integers.
type WorkflowPrimitiveDataType struct {
	WorkflowBaseDataType
	Properties           *WorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPrimitiveDataType WorkflowPrimitiveDataType

// NewWorkflowPrimitiveDataType instantiates a new WorkflowPrimitiveDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPrimitiveDataType() *WorkflowPrimitiveDataType {
	this := WorkflowPrimitiveDataType{}
	return &this
}

// NewWorkflowPrimitiveDataTypeWithDefaults instantiates a new WorkflowPrimitiveDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPrimitiveDataTypeWithDefaults() *WorkflowPrimitiveDataType {
	this := WorkflowPrimitiveDataType{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowPrimitiveDataType) GetProperties() WorkflowPrimitiveDataProperty {
	if o == nil || o.Properties == nil {
		var ret WorkflowPrimitiveDataProperty
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataType) GetPropertiesOk() (*WorkflowPrimitiveDataProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataType) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given WorkflowPrimitiveDataProperty and assigns it to the Properties field.
func (o *WorkflowPrimitiveDataType) SetProperties(v WorkflowPrimitiveDataProperty) {
	o.Properties = &v
}

func (o WorkflowPrimitiveDataType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowBaseDataType, errWorkflowBaseDataType := json.Marshal(o.WorkflowBaseDataType)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
	errWorkflowBaseDataType = json.Unmarshal([]byte(serializedWorkflowBaseDataType), &toSerialize)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPrimitiveDataType) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowPrimitiveDataTypeWithoutEmbeddedStruct struct {
		Properties *WorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
	}

	varWorkflowPrimitiveDataTypeWithoutEmbeddedStruct := WorkflowPrimitiveDataTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowPrimitiveDataTypeWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowPrimitiveDataType := _WorkflowPrimitiveDataType{}
		varWorkflowPrimitiveDataType.Properties = varWorkflowPrimitiveDataTypeWithoutEmbeddedStruct.Properties
		*o = WorkflowPrimitiveDataType(varWorkflowPrimitiveDataType)
	} else {
		return err
	}

	varWorkflowPrimitiveDataType := _WorkflowPrimitiveDataType{}

	err = json.Unmarshal(bytes, &varWorkflowPrimitiveDataType)
	if err == nil {
		o.WorkflowBaseDataType = varWorkflowPrimitiveDataType.WorkflowBaseDataType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Properties")

		// remove fields from embedded structs
		reflectWorkflowBaseDataType := reflect.ValueOf(o.WorkflowBaseDataType)
		for i := 0; i < reflectWorkflowBaseDataType.Type().NumField(); i++ {
			t := reflectWorkflowBaseDataType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowPrimitiveDataType struct {
	value *WorkflowPrimitiveDataType
	isSet bool
}

func (v NullableWorkflowPrimitiveDataType) Get() *WorkflowPrimitiveDataType {
	return v.value
}

func (v *NullableWorkflowPrimitiveDataType) Set(val *WorkflowPrimitiveDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPrimitiveDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPrimitiveDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPrimitiveDataType(val *WorkflowPrimitiveDataType) *NullableWorkflowPrimitiveDataType {
	return &NullableWorkflowPrimitiveDataType{value: val, isSet: true}
}

func (v NullableWorkflowPrimitiveDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPrimitiveDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}