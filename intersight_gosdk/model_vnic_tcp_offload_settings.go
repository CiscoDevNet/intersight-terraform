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
	"reflect"
	"strings"
)

// VnicTcpOffloadSettings The TCP offload settings decide whether to offload the TCP related network functions from the CPU to the network hardware or not. These options help reduce the CPU overhead and increase the network throughput.
type VnicTcpOffloadSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables the reassembly of segmented packets in hardware before sending them to the CPU.
	LargeReceive *bool `json:"LargeReceive,omitempty"`
	// Enables the CPU to send large packets to the hardware for segmentation.
	LargeSend *bool `json:"LargeSend,omitempty"`
	// When enabled, the CPU sends all packet checksums to the hardware for validation.
	RxChecksum *bool `json:"RxChecksum,omitempty"`
	// When enabled, the CPU sends all packets to the hardware so that the checksum can be calculated.
	TxChecksum           *bool `json:"TxChecksum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicTcpOffloadSettings VnicTcpOffloadSettings

// NewVnicTcpOffloadSettings instantiates a new VnicTcpOffloadSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicTcpOffloadSettings(classId string, objectType string) *VnicTcpOffloadSettings {
	this := VnicTcpOffloadSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicTcpOffloadSettingsWithDefaults instantiates a new VnicTcpOffloadSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicTcpOffloadSettingsWithDefaults() *VnicTcpOffloadSettings {
	this := VnicTcpOffloadSettings{}
	var classId string = "vnic.TcpOffloadSettings"
	this.ClassId = classId
	var objectType string = "vnic.TcpOffloadSettings"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicTcpOffloadSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicTcpOffloadSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicTcpOffloadSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicTcpOffloadSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLargeReceive returns the LargeReceive field value if set, zero value otherwise.
func (o *VnicTcpOffloadSettings) GetLargeReceive() bool {
	if o == nil || o.LargeReceive == nil {
		var ret bool
		return ret
	}
	return *o.LargeReceive
}

// GetLargeReceiveOk returns a tuple with the LargeReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettings) GetLargeReceiveOk() (*bool, bool) {
	if o == nil || o.LargeReceive == nil {
		return nil, false
	}
	return o.LargeReceive, true
}

// HasLargeReceive returns a boolean if a field has been set.
func (o *VnicTcpOffloadSettings) HasLargeReceive() bool {
	if o != nil && o.LargeReceive != nil {
		return true
	}

	return false
}

// SetLargeReceive gets a reference to the given bool and assigns it to the LargeReceive field.
func (o *VnicTcpOffloadSettings) SetLargeReceive(v bool) {
	o.LargeReceive = &v
}

// GetLargeSend returns the LargeSend field value if set, zero value otherwise.
func (o *VnicTcpOffloadSettings) GetLargeSend() bool {
	if o == nil || o.LargeSend == nil {
		var ret bool
		return ret
	}
	return *o.LargeSend
}

// GetLargeSendOk returns a tuple with the LargeSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettings) GetLargeSendOk() (*bool, bool) {
	if o == nil || o.LargeSend == nil {
		return nil, false
	}
	return o.LargeSend, true
}

// HasLargeSend returns a boolean if a field has been set.
func (o *VnicTcpOffloadSettings) HasLargeSend() bool {
	if o != nil && o.LargeSend != nil {
		return true
	}

	return false
}

// SetLargeSend gets a reference to the given bool and assigns it to the LargeSend field.
func (o *VnicTcpOffloadSettings) SetLargeSend(v bool) {
	o.LargeSend = &v
}

// GetRxChecksum returns the RxChecksum field value if set, zero value otherwise.
func (o *VnicTcpOffloadSettings) GetRxChecksum() bool {
	if o == nil || o.RxChecksum == nil {
		var ret bool
		return ret
	}
	return *o.RxChecksum
}

// GetRxChecksumOk returns a tuple with the RxChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettings) GetRxChecksumOk() (*bool, bool) {
	if o == nil || o.RxChecksum == nil {
		return nil, false
	}
	return o.RxChecksum, true
}

// HasRxChecksum returns a boolean if a field has been set.
func (o *VnicTcpOffloadSettings) HasRxChecksum() bool {
	if o != nil && o.RxChecksum != nil {
		return true
	}

	return false
}

// SetRxChecksum gets a reference to the given bool and assigns it to the RxChecksum field.
func (o *VnicTcpOffloadSettings) SetRxChecksum(v bool) {
	o.RxChecksum = &v
}

// GetTxChecksum returns the TxChecksum field value if set, zero value otherwise.
func (o *VnicTcpOffloadSettings) GetTxChecksum() bool {
	if o == nil || o.TxChecksum == nil {
		var ret bool
		return ret
	}
	return *o.TxChecksum
}

// GetTxChecksumOk returns a tuple with the TxChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicTcpOffloadSettings) GetTxChecksumOk() (*bool, bool) {
	if o == nil || o.TxChecksum == nil {
		return nil, false
	}
	return o.TxChecksum, true
}

// HasTxChecksum returns a boolean if a field has been set.
func (o *VnicTcpOffloadSettings) HasTxChecksum() bool {
	if o != nil && o.TxChecksum != nil {
		return true
	}

	return false
}

// SetTxChecksum gets a reference to the given bool and assigns it to the TxChecksum field.
func (o *VnicTcpOffloadSettings) SetTxChecksum(v bool) {
	o.TxChecksum = &v
}

func (o VnicTcpOffloadSettings) MarshalJSON() ([]byte, error) {
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
	if o.LargeReceive != nil {
		toSerialize["LargeReceive"] = o.LargeReceive
	}
	if o.LargeSend != nil {
		toSerialize["LargeSend"] = o.LargeSend
	}
	if o.RxChecksum != nil {
		toSerialize["RxChecksum"] = o.RxChecksum
	}
	if o.TxChecksum != nil {
		toSerialize["TxChecksum"] = o.TxChecksum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicTcpOffloadSettings) UnmarshalJSON(bytes []byte) (err error) {
	type VnicTcpOffloadSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enables the reassembly of segmented packets in hardware before sending them to the CPU.
		LargeReceive *bool `json:"LargeReceive,omitempty"`
		// Enables the CPU to send large packets to the hardware for segmentation.
		LargeSend *bool `json:"LargeSend,omitempty"`
		// When enabled, the CPU sends all packet checksums to the hardware for validation.
		RxChecksum *bool `json:"RxChecksum,omitempty"`
		// When enabled, the CPU sends all packets to the hardware so that the checksum can be calculated.
		TxChecksum *bool `json:"TxChecksum,omitempty"`
	}

	varVnicTcpOffloadSettingsWithoutEmbeddedStruct := VnicTcpOffloadSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicTcpOffloadSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicTcpOffloadSettings := _VnicTcpOffloadSettings{}
		varVnicTcpOffloadSettings.ClassId = varVnicTcpOffloadSettingsWithoutEmbeddedStruct.ClassId
		varVnicTcpOffloadSettings.ObjectType = varVnicTcpOffloadSettingsWithoutEmbeddedStruct.ObjectType
		varVnicTcpOffloadSettings.LargeReceive = varVnicTcpOffloadSettingsWithoutEmbeddedStruct.LargeReceive
		varVnicTcpOffloadSettings.LargeSend = varVnicTcpOffloadSettingsWithoutEmbeddedStruct.LargeSend
		varVnicTcpOffloadSettings.RxChecksum = varVnicTcpOffloadSettingsWithoutEmbeddedStruct.RxChecksum
		varVnicTcpOffloadSettings.TxChecksum = varVnicTcpOffloadSettingsWithoutEmbeddedStruct.TxChecksum
		*o = VnicTcpOffloadSettings(varVnicTcpOffloadSettings)
	} else {
		return err
	}

	varVnicTcpOffloadSettings := _VnicTcpOffloadSettings{}

	err = json.Unmarshal(bytes, &varVnicTcpOffloadSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicTcpOffloadSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LargeReceive")
		delete(additionalProperties, "LargeSend")
		delete(additionalProperties, "RxChecksum")
		delete(additionalProperties, "TxChecksum")

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

type NullableVnicTcpOffloadSettings struct {
	value *VnicTcpOffloadSettings
	isSet bool
}

func (v NullableVnicTcpOffloadSettings) Get() *VnicTcpOffloadSettings {
	return v.value
}

func (v *NullableVnicTcpOffloadSettings) Set(val *VnicTcpOffloadSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicTcpOffloadSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicTcpOffloadSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicTcpOffloadSettings(val *VnicTcpOffloadSettings) *NullableVnicTcpOffloadSettings {
	return &NullableVnicTcpOffloadSettings{value: val, isSet: true}
}

func (v NullableVnicTcpOffloadSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicTcpOffloadSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
