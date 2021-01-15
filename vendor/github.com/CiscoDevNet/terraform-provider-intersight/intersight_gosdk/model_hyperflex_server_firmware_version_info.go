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

// HyperflexServerFirmwareVersionInfo The firmware version details for UCS servers.
type HyperflexServerFirmwareVersionInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The platform type for UCS server. * `M5` - M5 generation of UCS server. * `M4` - M4 generation of UCS server.
	ServerPlatform *string `json:"ServerPlatform,omitempty"`
	// The server firmware bundle version.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexServerFirmwareVersionInfo HyperflexServerFirmwareVersionInfo

// NewHyperflexServerFirmwareVersionInfo instantiates a new HyperflexServerFirmwareVersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexServerFirmwareVersionInfo(classId string, objectType string) *HyperflexServerFirmwareVersionInfo {
	this := HyperflexServerFirmwareVersionInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// NewHyperflexServerFirmwareVersionInfoWithDefaults instantiates a new HyperflexServerFirmwareVersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexServerFirmwareVersionInfoWithDefaults() *HyperflexServerFirmwareVersionInfo {
	this := HyperflexServerFirmwareVersionInfo{}
	var classId string = "hyperflex.ServerFirmwareVersionInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.ServerFirmwareVersionInfo"
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexServerFirmwareVersionInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexServerFirmwareVersionInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexServerFirmwareVersionInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexServerFirmwareVersionInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetServerPlatform returns the ServerPlatform field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionInfo) GetServerPlatform() string {
	if o == nil || o.ServerPlatform == nil {
		var ret string
		return ret
	}
	return *o.ServerPlatform
}

// GetServerPlatformOk returns a tuple with the ServerPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfo) GetServerPlatformOk() (*string, bool) {
	if o == nil || o.ServerPlatform == nil {
		return nil, false
	}
	return o.ServerPlatform, true
}

// HasServerPlatform returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionInfo) HasServerPlatform() bool {
	if o != nil && o.ServerPlatform != nil {
		return true
	}

	return false
}

// SetServerPlatform gets a reference to the given string and assigns it to the ServerPlatform field.
func (o *HyperflexServerFirmwareVersionInfo) SetServerPlatform(v string) {
	o.ServerPlatform = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexServerFirmwareVersionInfo) SetVersion(v string) {
	o.Version = &v
}

func (o HyperflexServerFirmwareVersionInfo) MarshalJSON() ([]byte, error) {
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
	if o.ServerPlatform != nil {
		toSerialize["ServerPlatform"] = o.ServerPlatform
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexServerFirmwareVersionInfo) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The platform type for UCS server. * `M5` - M5 generation of UCS server. * `M4` - M4 generation of UCS server.
		ServerPlatform *string `json:"ServerPlatform,omitempty"`
		// The server firmware bundle version.
		Version *string `json:"Version,omitempty"`
	}

	varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct := HyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexServerFirmwareVersionInfo := _HyperflexServerFirmwareVersionInfo{}
		varHyperflexServerFirmwareVersionInfo.ClassId = varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct.ClassId
		varHyperflexServerFirmwareVersionInfo.ObjectType = varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexServerFirmwareVersionInfo.ServerPlatform = varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct.ServerPlatform
		varHyperflexServerFirmwareVersionInfo.Version = varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct.Version
		*o = HyperflexServerFirmwareVersionInfo(varHyperflexServerFirmwareVersionInfo)
	} else {
		return err
	}

	varHyperflexServerFirmwareVersionInfo := _HyperflexServerFirmwareVersionInfo{}

	err = json.Unmarshal(bytes, &varHyperflexServerFirmwareVersionInfo)
	if err == nil {
		o.MoBaseComplexType = varHyperflexServerFirmwareVersionInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ServerPlatform")
		delete(additionalProperties, "Version")

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

type NullableHyperflexServerFirmwareVersionInfo struct {
	value *HyperflexServerFirmwareVersionInfo
	isSet bool
}

func (v NullableHyperflexServerFirmwareVersionInfo) Get() *HyperflexServerFirmwareVersionInfo {
	return v.value
}

func (v *NullableHyperflexServerFirmwareVersionInfo) Set(val *HyperflexServerFirmwareVersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerFirmwareVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerFirmwareVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerFirmwareVersionInfo(val *HyperflexServerFirmwareVersionInfo) *NullableHyperflexServerFirmwareVersionInfo {
	return &NullableHyperflexServerFirmwareVersionInfo{value: val, isSet: true}
}

func (v NullableHyperflexServerFirmwareVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerFirmwareVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
