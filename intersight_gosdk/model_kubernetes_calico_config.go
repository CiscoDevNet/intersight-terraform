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

// KubernetesCalicoConfig Configuration of Calico CNI plugin.
type KubernetesCalicoConfig struct {
	KubernetesCniConfig
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP version that can take on values v4 or v6. * `v4` - This refers to the IPv4 address. * `v6` - This refers to the IPv6 address.
	IpVersion *string `json:"IpVersion,omitempty"`
	// Workload interface maximum transmission unit (MTU).
	Mtu                  *int64 `json:"Mtu,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesCalicoConfig KubernetesCalicoConfig

// NewKubernetesCalicoConfig instantiates a new KubernetesCalicoConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCalicoConfig(classId string, objectType string) *KubernetesCalicoConfig {
	this := KubernetesCalicoConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipVersion string = "v4"
	this.IpVersion = &ipVersion
	return &this
}

// NewKubernetesCalicoConfigWithDefaults instantiates a new KubernetesCalicoConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesCalicoConfigWithDefaults() *KubernetesCalicoConfig {
	this := KubernetesCalicoConfig{}
	var classId string = "kubernetes.CalicoConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.CalicoConfig"
	this.ObjectType = objectType
	var ipVersion string = "v4"
	this.IpVersion = &ipVersion
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesCalicoConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesCalicoConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesCalicoConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesCalicoConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesCalicoConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesCalicoConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *KubernetesCalicoConfig) GetIpVersion() string {
	if o == nil || o.IpVersion == nil {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCalicoConfig) GetIpVersionOk() (*string, bool) {
	if o == nil || o.IpVersion == nil {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *KubernetesCalicoConfig) HasIpVersion() bool {
	if o != nil && o.IpVersion != nil {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *KubernetesCalicoConfig) SetIpVersion(v string) {
	o.IpVersion = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *KubernetesCalicoConfig) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCalicoConfig) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *KubernetesCalicoConfig) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *KubernetesCalicoConfig) SetMtu(v int64) {
	o.Mtu = &v
}

func (o KubernetesCalicoConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesCniConfig, errKubernetesCniConfig := json.Marshal(o.KubernetesCniConfig)
	if errKubernetesCniConfig != nil {
		return []byte{}, errKubernetesCniConfig
	}
	errKubernetesCniConfig = json.Unmarshal([]byte(serializedKubernetesCniConfig), &toSerialize)
	if errKubernetesCniConfig != nil {
		return []byte{}, errKubernetesCniConfig
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpVersion != nil {
		toSerialize["IpVersion"] = o.IpVersion
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesCalicoConfig) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesCalicoConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP version that can take on values v4 or v6. * `v4` - This refers to the IPv4 address. * `v6` - This refers to the IPv6 address.
		IpVersion *string `json:"IpVersion,omitempty"`
		// Workload interface maximum transmission unit (MTU).
		Mtu *int64 `json:"Mtu,omitempty"`
	}

	varKubernetesCalicoConfigWithoutEmbeddedStruct := KubernetesCalicoConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesCalicoConfigWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesCalicoConfig := _KubernetesCalicoConfig{}
		varKubernetesCalicoConfig.ClassId = varKubernetesCalicoConfigWithoutEmbeddedStruct.ClassId
		varKubernetesCalicoConfig.ObjectType = varKubernetesCalicoConfigWithoutEmbeddedStruct.ObjectType
		varKubernetesCalicoConfig.IpVersion = varKubernetesCalicoConfigWithoutEmbeddedStruct.IpVersion
		varKubernetesCalicoConfig.Mtu = varKubernetesCalicoConfigWithoutEmbeddedStruct.Mtu
		*o = KubernetesCalicoConfig(varKubernetesCalicoConfig)
	} else {
		return err
	}

	varKubernetesCalicoConfig := _KubernetesCalicoConfig{}

	err = json.Unmarshal(bytes, &varKubernetesCalicoConfig)
	if err == nil {
		o.KubernetesCniConfig = varKubernetesCalicoConfig.KubernetesCniConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpVersion")
		delete(additionalProperties, "Mtu")

		// remove fields from embedded structs
		reflectKubernetesCniConfig := reflect.ValueOf(o.KubernetesCniConfig)
		for i := 0; i < reflectKubernetesCniConfig.Type().NumField(); i++ {
			t := reflectKubernetesCniConfig.Type().Field(i)

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

type NullableKubernetesCalicoConfig struct {
	value *KubernetesCalicoConfig
	isSet bool
}

func (v NullableKubernetesCalicoConfig) Get() *KubernetesCalicoConfig {
	return v.value
}

func (v *NullableKubernetesCalicoConfig) Set(val *KubernetesCalicoConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCalicoConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCalicoConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCalicoConfig(val *KubernetesCalicoConfig) *NullableKubernetesCalicoConfig {
	return &NullableKubernetesCalicoConfig{value: val, isSet: true}
}

func (v NullableKubernetesCalicoConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCalicoConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
