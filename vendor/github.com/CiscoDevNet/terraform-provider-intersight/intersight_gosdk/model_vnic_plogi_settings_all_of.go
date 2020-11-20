/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VnicPlogiSettingsAllOf Definition of the list of properties defined in 'vnic.PlogiSettings', excluding properties defined in parent classes.
type VnicPlogiSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of times that the system tries to log in to a port after the first failure.
	Retries *int64 `json:"Retries,omitempty"`
	// The number of milliseconds that the system waits before it tries to log in again.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicPlogiSettingsAllOf VnicPlogiSettingsAllOf

// NewVnicPlogiSettingsAllOf instantiates a new VnicPlogiSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicPlogiSettingsAllOf(classId string, objectType string) *VnicPlogiSettingsAllOf {
	this := VnicPlogiSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicPlogiSettingsAllOfWithDefaults instantiates a new VnicPlogiSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicPlogiSettingsAllOfWithDefaults() *VnicPlogiSettingsAllOf {
	this := VnicPlogiSettingsAllOf{}
	var classId string = "vnic.PlogiSettings"
	this.ClassId = classId
	var objectType string = "vnic.PlogiSettings"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicPlogiSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicPlogiSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicPlogiSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicPlogiSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicPlogiSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicPlogiSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *VnicPlogiSettingsAllOf) GetRetries() int64 {
	if o == nil || o.Retries == nil {
		var ret int64
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlogiSettingsAllOf) GetRetriesOk() (*int64, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *VnicPlogiSettingsAllOf) HasRetries() bool {
	if o != nil && o.Retries != nil {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int64 and assigns it to the Retries field.
func (o *VnicPlogiSettingsAllOf) SetRetries(v int64) {
	o.Retries = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *VnicPlogiSettingsAllOf) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlogiSettingsAllOf) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *VnicPlogiSettingsAllOf) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *VnicPlogiSettingsAllOf) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o VnicPlogiSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Retries != nil {
		toSerialize["Retries"] = o.Retries
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicPlogiSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicPlogiSettingsAllOf := _VnicPlogiSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicPlogiSettingsAllOf); err == nil {
		*o = VnicPlogiSettingsAllOf(varVnicPlogiSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Retries")
		delete(additionalProperties, "Timeout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicPlogiSettingsAllOf struct {
	value *VnicPlogiSettingsAllOf
	isSet bool
}

func (v NullableVnicPlogiSettingsAllOf) Get() *VnicPlogiSettingsAllOf {
	return v.value
}

func (v *NullableVnicPlogiSettingsAllOf) Set(val *VnicPlogiSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicPlogiSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicPlogiSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicPlogiSettingsAllOf(val *VnicPlogiSettingsAllOf) *NullableVnicPlogiSettingsAllOf {
	return &NullableVnicPlogiSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicPlogiSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicPlogiSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
