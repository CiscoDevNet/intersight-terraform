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
	"reflect"
	"strings"
)

// HyperflexMapClusterIdToProtectionInfo A map from Cluster to Protection Info.
type HyperflexMapClusterIdToProtectionInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Cluster Id we are using to map to ProtectionInfo.
	ClusterId            *string                         `json:"ClusterId,omitempty"`
	ProtectionInfo       NullableHyperflexProtectionInfo `json:"ProtectionInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexMapClusterIdToProtectionInfo HyperflexMapClusterIdToProtectionInfo

// NewHyperflexMapClusterIdToProtectionInfo instantiates a new HyperflexMapClusterIdToProtectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexMapClusterIdToProtectionInfo(classId string, objectType string) *HyperflexMapClusterIdToProtectionInfo {
	this := HyperflexMapClusterIdToProtectionInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexMapClusterIdToProtectionInfoWithDefaults instantiates a new HyperflexMapClusterIdToProtectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexMapClusterIdToProtectionInfoWithDefaults() *HyperflexMapClusterIdToProtectionInfo {
	this := HyperflexMapClusterIdToProtectionInfo{}
	var classId string = "hyperflex.MapClusterIdToProtectionInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.MapClusterIdToProtectionInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexMapClusterIdToProtectionInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexMapClusterIdToProtectionInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexMapClusterIdToProtectionInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexMapClusterIdToProtectionInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexMapClusterIdToProtectionInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexMapClusterIdToProtectionInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *HyperflexMapClusterIdToProtectionInfo) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexMapClusterIdToProtectionInfo) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *HyperflexMapClusterIdToProtectionInfo) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *HyperflexMapClusterIdToProtectionInfo) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetProtectionInfo returns the ProtectionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexMapClusterIdToProtectionInfo) GetProtectionInfo() HyperflexProtectionInfo {
	if o == nil || o.ProtectionInfo.Get() == nil {
		var ret HyperflexProtectionInfo
		return ret
	}
	return *o.ProtectionInfo.Get()
}

// GetProtectionInfoOk returns a tuple with the ProtectionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexMapClusterIdToProtectionInfo) GetProtectionInfoOk() (*HyperflexProtectionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionInfo.Get(), o.ProtectionInfo.IsSet()
}

// HasProtectionInfo returns a boolean if a field has been set.
func (o *HyperflexMapClusterIdToProtectionInfo) HasProtectionInfo() bool {
	if o != nil && o.ProtectionInfo.IsSet() {
		return true
	}

	return false
}

// SetProtectionInfo gets a reference to the given NullableHyperflexProtectionInfo and assigns it to the ProtectionInfo field.
func (o *HyperflexMapClusterIdToProtectionInfo) SetProtectionInfo(v HyperflexProtectionInfo) {
	o.ProtectionInfo.Set(&v)
}

// SetProtectionInfoNil sets the value for ProtectionInfo to be an explicit nil
func (o *HyperflexMapClusterIdToProtectionInfo) SetProtectionInfoNil() {
	o.ProtectionInfo.Set(nil)
}

// UnsetProtectionInfo ensures that no value is present for ProtectionInfo, not even an explicit nil
func (o *HyperflexMapClusterIdToProtectionInfo) UnsetProtectionInfo() {
	o.ProtectionInfo.Unset()
}

func (o HyperflexMapClusterIdToProtectionInfo) MarshalJSON() ([]byte, error) {
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
	if o.ClusterId != nil {
		toSerialize["ClusterId"] = o.ClusterId
	}
	if o.ProtectionInfo.IsSet() {
		toSerialize["ProtectionInfo"] = o.ProtectionInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexMapClusterIdToProtectionInfo) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexMapClusterIdToProtectionInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Cluster Id we are using to map to ProtectionInfo.
		ClusterId      *string                         `json:"ClusterId,omitempty"`
		ProtectionInfo NullableHyperflexProtectionInfo `json:"ProtectionInfo,omitempty"`
	}

	varHyperflexMapClusterIdToProtectionInfoWithoutEmbeddedStruct := HyperflexMapClusterIdToProtectionInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexMapClusterIdToProtectionInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexMapClusterIdToProtectionInfo := _HyperflexMapClusterIdToProtectionInfo{}
		varHyperflexMapClusterIdToProtectionInfo.ClassId = varHyperflexMapClusterIdToProtectionInfoWithoutEmbeddedStruct.ClassId
		varHyperflexMapClusterIdToProtectionInfo.ObjectType = varHyperflexMapClusterIdToProtectionInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexMapClusterIdToProtectionInfo.ClusterId = varHyperflexMapClusterIdToProtectionInfoWithoutEmbeddedStruct.ClusterId
		varHyperflexMapClusterIdToProtectionInfo.ProtectionInfo = varHyperflexMapClusterIdToProtectionInfoWithoutEmbeddedStruct.ProtectionInfo
		*o = HyperflexMapClusterIdToProtectionInfo(varHyperflexMapClusterIdToProtectionInfo)
	} else {
		return err
	}

	varHyperflexMapClusterIdToProtectionInfo := _HyperflexMapClusterIdToProtectionInfo{}

	err = json.Unmarshal(bytes, &varHyperflexMapClusterIdToProtectionInfo)
	if err == nil {
		o.MoBaseComplexType = varHyperflexMapClusterIdToProtectionInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterId")
		delete(additionalProperties, "ProtectionInfo")

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

type NullableHyperflexMapClusterIdToProtectionInfo struct {
	value *HyperflexMapClusterIdToProtectionInfo
	isSet bool
}

func (v NullableHyperflexMapClusterIdToProtectionInfo) Get() *HyperflexMapClusterIdToProtectionInfo {
	return v.value
}

func (v *NullableHyperflexMapClusterIdToProtectionInfo) Set(val *HyperflexMapClusterIdToProtectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexMapClusterIdToProtectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexMapClusterIdToProtectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexMapClusterIdToProtectionInfo(val *HyperflexMapClusterIdToProtectionInfo) *NullableHyperflexMapClusterIdToProtectionInfo {
	return &NullableHyperflexMapClusterIdToProtectionInfo{value: val, isSet: true}
}

func (v NullableHyperflexMapClusterIdToProtectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexMapClusterIdToProtectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
