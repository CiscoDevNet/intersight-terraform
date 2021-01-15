/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-06T06:42:37Z.
 *
 * API version: 1.0.9-3181
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowPrimitiveDataType This data type is used to represent primitives like string, floats and integers.
type WorkflowPrimitiveDataType struct {
	WorkflowBaseDataType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                `json:"ObjectType"`
	Properties           NullableWorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPrimitiveDataType WorkflowPrimitiveDataType

// NewWorkflowPrimitiveDataType instantiates a new WorkflowPrimitiveDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPrimitiveDataType(classId string, objectType string) *WorkflowPrimitiveDataType {
	this := WorkflowPrimitiveDataType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowPrimitiveDataTypeWithDefaults instantiates a new WorkflowPrimitiveDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPrimitiveDataTypeWithDefaults() *WorkflowPrimitiveDataType {
	this := WorkflowPrimitiveDataType{}
	var classId string = "workflow.PrimitiveDataType"
	this.ClassId = classId
	var objectType string = "workflow.PrimitiveDataType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowPrimitiveDataType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowPrimitiveDataType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowPrimitiveDataType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowPrimitiveDataType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPrimitiveDataType) GetProperties() WorkflowPrimitiveDataProperty {
	if o == nil || o.Properties.Get() == nil {
		var ret WorkflowPrimitiveDataProperty
		return ret
	}
	return *o.Properties.Get()
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPrimitiveDataType) GetPropertiesOk() (*WorkflowPrimitiveDataProperty, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Get(), o.Properties.IsSet()
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataType) HasProperties() bool {
	if o != nil && o.Properties.IsSet() {
		return true
	}

	return false
}

// SetProperties gets a reference to the given NullableWorkflowPrimitiveDataProperty and assigns it to the Properties field.
func (o *WorkflowPrimitiveDataType) SetProperties(v WorkflowPrimitiveDataProperty) {
	o.Properties.Set(&v)
}

// SetPropertiesNil sets the value for Properties to be an explicit nil
func (o *WorkflowPrimitiveDataType) SetPropertiesNil() {
	o.Properties.Set(nil)
}

// UnsetProperties ensures that no value is present for Properties, not even an explicit nil
func (o *WorkflowPrimitiveDataType) UnsetProperties() {
	o.Properties.Unset()
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
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Properties.IsSet() {
		toSerialize["Properties"] = o.Properties.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPrimitiveDataType) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowPrimitiveDataTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                `json:"ObjectType"`
		Properties NullableWorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
	}

	varWorkflowPrimitiveDataTypeWithoutEmbeddedStruct := WorkflowPrimitiveDataTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowPrimitiveDataTypeWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowPrimitiveDataType := _WorkflowPrimitiveDataType{}
		varWorkflowPrimitiveDataType.ClassId = varWorkflowPrimitiveDataTypeWithoutEmbeddedStruct.ClassId
		varWorkflowPrimitiveDataType.ObjectType = varWorkflowPrimitiveDataTypeWithoutEmbeddedStruct.ObjectType
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
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
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
