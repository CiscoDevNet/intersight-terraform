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
)

// NiatelemetryNexusDashboardDetailsAllOf Definition of the list of properties defined in 'niatelemetry.NexusDashboardDetails', excluding properties defined in parent classes.
type NiatelemetryNexusDashboardDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the NexusDashboard.
	NexusDashboardName *string `json:"NexusDashboardName,omitempty"`
	// Serial number of NexusDashboard.
	NexusDashboardSerialNumber *string                              `json:"NexusDashboardSerialNumber,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _NiatelemetryNexusDashboardDetailsAllOf NiatelemetryNexusDashboardDetailsAllOf

// NewNiatelemetryNexusDashboardDetailsAllOf instantiates a new NiatelemetryNexusDashboardDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusDashboardDetailsAllOf(classId string, objectType string) *NiatelemetryNexusDashboardDetailsAllOf {
	this := NiatelemetryNexusDashboardDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusDashboardDetailsAllOfWithDefaults instantiates a new NiatelemetryNexusDashboardDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusDashboardDetailsAllOfWithDefaults() *NiatelemetryNexusDashboardDetailsAllOf {
	this := NiatelemetryNexusDashboardDetailsAllOf{}
	var classId string = "niatelemetry.NexusDashboardDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusDashboardDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusDashboardDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusDashboardDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNexusDashboardName returns the NexusDashboardName field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetNexusDashboardName() string {
	if o == nil || o.NexusDashboardName == nil {
		var ret string
		return ret
	}
	return *o.NexusDashboardName
}

// GetNexusDashboardNameOk returns a tuple with the NexusDashboardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetNexusDashboardNameOk() (*string, bool) {
	if o == nil || o.NexusDashboardName == nil {
		return nil, false
	}
	return o.NexusDashboardName, true
}

// HasNexusDashboardName returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardDetailsAllOf) HasNexusDashboardName() bool {
	if o != nil && o.NexusDashboardName != nil {
		return true
	}

	return false
}

// SetNexusDashboardName gets a reference to the given string and assigns it to the NexusDashboardName field.
func (o *NiatelemetryNexusDashboardDetailsAllOf) SetNexusDashboardName(v string) {
	o.NexusDashboardName = &v
}

// GetNexusDashboardSerialNumber returns the NexusDashboardSerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetNexusDashboardSerialNumber() string {
	if o == nil || o.NexusDashboardSerialNumber == nil {
		var ret string
		return ret
	}
	return *o.NexusDashboardSerialNumber
}

// GetNexusDashboardSerialNumberOk returns a tuple with the NexusDashboardSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetNexusDashboardSerialNumberOk() (*string, bool) {
	if o == nil || o.NexusDashboardSerialNumber == nil {
		return nil, false
	}
	return o.NexusDashboardSerialNumber, true
}

// HasNexusDashboardSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardDetailsAllOf) HasNexusDashboardSerialNumber() bool {
	if o != nil && o.NexusDashboardSerialNumber != nil {
		return true
	}

	return false
}

// SetNexusDashboardSerialNumber gets a reference to the given string and assigns it to the NexusDashboardSerialNumber field.
func (o *NiatelemetryNexusDashboardDetailsAllOf) SetNexusDashboardSerialNumber(v string) {
	o.NexusDashboardSerialNumber = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNexusDashboardDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNexusDashboardDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NexusDashboardName != nil {
		toSerialize["NexusDashboardName"] = o.NexusDashboardName
	}
	if o.NexusDashboardSerialNumber != nil {
		toSerialize["NexusDashboardSerialNumber"] = o.NexusDashboardSerialNumber
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNexusDashboardDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNexusDashboardDetailsAllOf := _NiatelemetryNexusDashboardDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardDetailsAllOf); err == nil {
		*o = NiatelemetryNexusDashboardDetailsAllOf(varNiatelemetryNexusDashboardDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NexusDashboardName")
		delete(additionalProperties, "NexusDashboardSerialNumber")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNexusDashboardDetailsAllOf struct {
	value *NiatelemetryNexusDashboardDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryNexusDashboardDetailsAllOf) Get() *NiatelemetryNexusDashboardDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryNexusDashboardDetailsAllOf) Set(val *NiatelemetryNexusDashboardDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusDashboardDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusDashboardDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusDashboardDetailsAllOf(val *NiatelemetryNexusDashboardDetailsAllOf) *NullableNiatelemetryNexusDashboardDetailsAllOf {
	return &NullableNiatelemetryNexusDashboardDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusDashboardDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusDashboardDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
