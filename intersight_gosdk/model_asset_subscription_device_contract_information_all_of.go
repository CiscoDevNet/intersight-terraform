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

// AssetSubscriptionDeviceContractInformationAllOf Definition of the list of properties defined in 'asset.SubscriptionDeviceContractInformation', excluding properties defined in parent classes.
type AssetSubscriptionDeviceContractInformationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identifier of the Cisco device.
	DeviceId          *string                        `json:"DeviceId,omitempty"`
	DeviceInformation NullableAssetDeviceInformation `json:"DeviceInformation,omitempty"`
	// Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices.
	DevicePid *string `json:"DevicePid,omitempty"`
	// Identifies the consumption-based subscription.
	SubscriptionRefId         *string                                     `json:"SubscriptionRefId,omitempty"`
	DeviceContractInformation *AssetDeviceContractInformationRelationship `json:"DeviceContractInformation,omitempty"`
	RegisteredDevice          *AssetDeviceRegistrationRelationship        `json:"RegisteredDevice,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _AssetSubscriptionDeviceContractInformationAllOf AssetSubscriptionDeviceContractInformationAllOf

// NewAssetSubscriptionDeviceContractInformationAllOf instantiates a new AssetSubscriptionDeviceContractInformationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetSubscriptionDeviceContractInformationAllOf(classId string, objectType string) *AssetSubscriptionDeviceContractInformationAllOf {
	this := AssetSubscriptionDeviceContractInformationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetSubscriptionDeviceContractInformationAllOfWithDefaults instantiates a new AssetSubscriptionDeviceContractInformationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetSubscriptionDeviceContractInformationAllOfWithDefaults() *AssetSubscriptionDeviceContractInformationAllOf {
	this := AssetSubscriptionDeviceContractInformationAllOf{}
	var classId string = "asset.SubscriptionDeviceContractInformation"
	this.ClassId = classId
	var objectType string = "asset.SubscriptionDeviceContractInformation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetSubscriptionDeviceContractInformationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetSubscriptionDeviceContractInformationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *AssetSubscriptionDeviceContractInformationAllOf) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceInformation returns the DeviceInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetDeviceInformation() AssetDeviceInformation {
	if o == nil || o.DeviceInformation.Get() == nil {
		var ret AssetDeviceInformation
		return ret
	}
	return *o.DeviceInformation.Get()
}

// GetDeviceInformationOk returns a tuple with the DeviceInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetDeviceInformationOk() (*AssetDeviceInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceInformation.Get(), o.DeviceInformation.IsSet()
}

// HasDeviceInformation returns a boolean if a field has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) HasDeviceInformation() bool {
	if o != nil && o.DeviceInformation.IsSet() {
		return true
	}

	return false
}

// SetDeviceInformation gets a reference to the given NullableAssetDeviceInformation and assigns it to the DeviceInformation field.
func (o *AssetSubscriptionDeviceContractInformationAllOf) SetDeviceInformation(v AssetDeviceInformation) {
	o.DeviceInformation.Set(&v)
}

// SetDeviceInformationNil sets the value for DeviceInformation to be an explicit nil
func (o *AssetSubscriptionDeviceContractInformationAllOf) SetDeviceInformationNil() {
	o.DeviceInformation.Set(nil)
}

// UnsetDeviceInformation ensures that no value is present for DeviceInformation, not even an explicit nil
func (o *AssetSubscriptionDeviceContractInformationAllOf) UnsetDeviceInformation() {
	o.DeviceInformation.Unset()
}

// GetDevicePid returns the DevicePid field value if set, zero value otherwise.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetDevicePid() string {
	if o == nil || o.DevicePid == nil {
		var ret string
		return ret
	}
	return *o.DevicePid
}

// GetDevicePidOk returns a tuple with the DevicePid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetDevicePidOk() (*string, bool) {
	if o == nil || o.DevicePid == nil {
		return nil, false
	}
	return o.DevicePid, true
}

// HasDevicePid returns a boolean if a field has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) HasDevicePid() bool {
	if o != nil && o.DevicePid != nil {
		return true
	}

	return false
}

// SetDevicePid gets a reference to the given string and assigns it to the DevicePid field.
func (o *AssetSubscriptionDeviceContractInformationAllOf) SetDevicePid(v string) {
	o.DevicePid = &v
}

// GetSubscriptionRefId returns the SubscriptionRefId field value if set, zero value otherwise.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetSubscriptionRefId() string {
	if o == nil || o.SubscriptionRefId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionRefId
}

// GetSubscriptionRefIdOk returns a tuple with the SubscriptionRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetSubscriptionRefIdOk() (*string, bool) {
	if o == nil || o.SubscriptionRefId == nil {
		return nil, false
	}
	return o.SubscriptionRefId, true
}

// HasSubscriptionRefId returns a boolean if a field has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) HasSubscriptionRefId() bool {
	if o != nil && o.SubscriptionRefId != nil {
		return true
	}

	return false
}

// SetSubscriptionRefId gets a reference to the given string and assigns it to the SubscriptionRefId field.
func (o *AssetSubscriptionDeviceContractInformationAllOf) SetSubscriptionRefId(v string) {
	o.SubscriptionRefId = &v
}

// GetDeviceContractInformation returns the DeviceContractInformation field value if set, zero value otherwise.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetDeviceContractInformation() AssetDeviceContractInformationRelationship {
	if o == nil || o.DeviceContractInformation == nil {
		var ret AssetDeviceContractInformationRelationship
		return ret
	}
	return *o.DeviceContractInformation
}

// GetDeviceContractInformationOk returns a tuple with the DeviceContractInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetDeviceContractInformationOk() (*AssetDeviceContractInformationRelationship, bool) {
	if o == nil || o.DeviceContractInformation == nil {
		return nil, false
	}
	return o.DeviceContractInformation, true
}

// HasDeviceContractInformation returns a boolean if a field has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) HasDeviceContractInformation() bool {
	if o != nil && o.DeviceContractInformation != nil {
		return true
	}

	return false
}

// SetDeviceContractInformation gets a reference to the given AssetDeviceContractInformationRelationship and assigns it to the DeviceContractInformation field.
func (o *AssetSubscriptionDeviceContractInformationAllOf) SetDeviceContractInformation(v AssetDeviceContractInformationRelationship) {
	o.DeviceContractInformation = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AssetSubscriptionDeviceContractInformationAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AssetSubscriptionDeviceContractInformationAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AssetSubscriptionDeviceContractInformationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.DeviceInformation.IsSet() {
		toSerialize["DeviceInformation"] = o.DeviceInformation.Get()
	}
	if o.DevicePid != nil {
		toSerialize["DevicePid"] = o.DevicePid
	}
	if o.SubscriptionRefId != nil {
		toSerialize["SubscriptionRefId"] = o.SubscriptionRefId
	}
	if o.DeviceContractInformation != nil {
		toSerialize["DeviceContractInformation"] = o.DeviceContractInformation
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetSubscriptionDeviceContractInformationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetSubscriptionDeviceContractInformationAllOf := _AssetSubscriptionDeviceContractInformationAllOf{}

	if err = json.Unmarshal(bytes, &varAssetSubscriptionDeviceContractInformationAllOf); err == nil {
		*o = AssetSubscriptionDeviceContractInformationAllOf(varAssetSubscriptionDeviceContractInformationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "DeviceInformation")
		delete(additionalProperties, "DevicePid")
		delete(additionalProperties, "SubscriptionRefId")
		delete(additionalProperties, "DeviceContractInformation")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetSubscriptionDeviceContractInformationAllOf struct {
	value *AssetSubscriptionDeviceContractInformationAllOf
	isSet bool
}

func (v NullableAssetSubscriptionDeviceContractInformationAllOf) Get() *AssetSubscriptionDeviceContractInformationAllOf {
	return v.value
}

func (v *NullableAssetSubscriptionDeviceContractInformationAllOf) Set(val *AssetSubscriptionDeviceContractInformationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSubscriptionDeviceContractInformationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSubscriptionDeviceContractInformationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSubscriptionDeviceContractInformationAllOf(val *AssetSubscriptionDeviceContractInformationAllOf) *NullableAssetSubscriptionDeviceContractInformationAllOf {
	return &NullableAssetSubscriptionDeviceContractInformationAllOf{value: val, isSet: true}
}

func (v NullableAssetSubscriptionDeviceContractInformationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetSubscriptionDeviceContractInformationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
