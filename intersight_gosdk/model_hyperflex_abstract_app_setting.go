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

// HyperflexAbstractAppSetting An application setting as a key/value pair.
type HyperflexAbstractAppSetting struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The application setting identifier.
	Name *string `json:"Name,omitempty"`
	// The application setting value.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAbstractAppSetting HyperflexAbstractAppSetting

// NewHyperflexAbstractAppSetting instantiates a new HyperflexAbstractAppSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAbstractAppSetting(classId string, objectType string) *HyperflexAbstractAppSetting {
	this := HyperflexAbstractAppSetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexAbstractAppSettingWithDefaults instantiates a new HyperflexAbstractAppSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAbstractAppSettingWithDefaults() *HyperflexAbstractAppSetting {
	this := HyperflexAbstractAppSetting{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAbstractAppSetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAbstractAppSetting) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAbstractAppSetting) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAbstractAppSetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAbstractAppSetting) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAbstractAppSetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexAbstractAppSetting) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAbstractAppSetting) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexAbstractAppSetting) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexAbstractAppSetting) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HyperflexAbstractAppSetting) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAbstractAppSetting) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HyperflexAbstractAppSetting) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HyperflexAbstractAppSetting) SetValue(v string) {
	o.Value = &v
}

func (o HyperflexAbstractAppSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAbstractAppSetting) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexAbstractAppSettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The application setting identifier.
		Name *string `json:"Name,omitempty"`
		// The application setting value.
		Value *string `json:"Value,omitempty"`
	}

	varHyperflexAbstractAppSettingWithoutEmbeddedStruct := HyperflexAbstractAppSettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexAbstractAppSettingWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexAbstractAppSetting := _HyperflexAbstractAppSetting{}
		varHyperflexAbstractAppSetting.ClassId = varHyperflexAbstractAppSettingWithoutEmbeddedStruct.ClassId
		varHyperflexAbstractAppSetting.ObjectType = varHyperflexAbstractAppSettingWithoutEmbeddedStruct.ObjectType
		varHyperflexAbstractAppSetting.Name = varHyperflexAbstractAppSettingWithoutEmbeddedStruct.Name
		varHyperflexAbstractAppSetting.Value = varHyperflexAbstractAppSettingWithoutEmbeddedStruct.Value
		*o = HyperflexAbstractAppSetting(varHyperflexAbstractAppSetting)
	} else {
		return err
	}

	varHyperflexAbstractAppSetting := _HyperflexAbstractAppSetting{}

	err = json.Unmarshal(bytes, &varHyperflexAbstractAppSetting)
	if err == nil {
		o.MoBaseComplexType = varHyperflexAbstractAppSetting.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableHyperflexAbstractAppSetting struct {
	value *HyperflexAbstractAppSetting
	isSet bool
}

func (v NullableHyperflexAbstractAppSetting) Get() *HyperflexAbstractAppSetting {
	return v.value
}

func (v *NullableHyperflexAbstractAppSetting) Set(val *HyperflexAbstractAppSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAbstractAppSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAbstractAppSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAbstractAppSetting(val *HyperflexAbstractAppSetting) *NullableHyperflexAbstractAppSetting {
	return &NullableHyperflexAbstractAppSetting{value: val, isSet: true}
}

func (v NullableHyperflexAbstractAppSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAbstractAppSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
