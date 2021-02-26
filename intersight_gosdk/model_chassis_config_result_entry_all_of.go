/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-17T05:06:15Z.
 *
 * API version: 1.0.9-3714
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ChassisConfigResultEntryAllOf Definition of the list of properties defined in 'chassis.ConfigResultEntry', excluding properties defined in parent classes.
type ChassisConfigResultEntryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                           `json:"ObjectType"`
	ConfigResult         *ChassisConfigResultRelationship `json:"ConfigResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisConfigResultEntryAllOf ChassisConfigResultEntryAllOf

// NewChassisConfigResultEntryAllOf instantiates a new ChassisConfigResultEntryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisConfigResultEntryAllOf(classId string, objectType string) *ChassisConfigResultEntryAllOf {
	this := ChassisConfigResultEntryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewChassisConfigResultEntryAllOfWithDefaults instantiates a new ChassisConfigResultEntryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisConfigResultEntryAllOfWithDefaults() *ChassisConfigResultEntryAllOf {
	this := ChassisConfigResultEntryAllOf{}
	var classId string = "chassis.ConfigResultEntry"
	this.ClassId = classId
	var objectType string = "chassis.ConfigResultEntry"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisConfigResultEntryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigResultEntryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisConfigResultEntryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ChassisConfigResultEntryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigResultEntryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisConfigResultEntryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *ChassisConfigResultEntryAllOf) GetConfigResult() ChassisConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret ChassisConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigResultEntryAllOf) GetConfigResultOk() (*ChassisConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *ChassisConfigResultEntryAllOf) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given ChassisConfigResultRelationship and assigns it to the ConfigResult field.
func (o *ChassisConfigResultEntryAllOf) SetConfigResult(v ChassisConfigResultRelationship) {
	o.ConfigResult = &v
}

func (o ChassisConfigResultEntryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChassisConfigResultEntryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varChassisConfigResultEntryAllOf := _ChassisConfigResultEntryAllOf{}

	if err = json.Unmarshal(bytes, &varChassisConfigResultEntryAllOf); err == nil {
		*o = ChassisConfigResultEntryAllOf(varChassisConfigResultEntryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigResult")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChassisConfigResultEntryAllOf struct {
	value *ChassisConfigResultEntryAllOf
	isSet bool
}

func (v NullableChassisConfigResultEntryAllOf) Get() *ChassisConfigResultEntryAllOf {
	return v.value
}

func (v *NullableChassisConfigResultEntryAllOf) Set(val *ChassisConfigResultEntryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisConfigResultEntryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisConfigResultEntryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisConfigResultEntryAllOf(val *ChassisConfigResultEntryAllOf) *NullableChassisConfigResultEntryAllOf {
	return &NullableChassisConfigResultEntryAllOf{value: val, isSet: true}
}

func (v NullableChassisConfigResultEntryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisConfigResultEntryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
