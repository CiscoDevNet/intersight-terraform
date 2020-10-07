/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VnicCompletionQueueSettings Completion Queue resource settings.
type VnicCompletionQueueSettings struct {
	MoBaseComplexType
	// The number of completion queue resources to allocate. In general, the number of completion queue resources to allocate is equal to the number of transmit queue resources plus the number of receive queue resources.
	Count *int64 `json:"Count,omitempty"`
	// The number of descriptors in each completion queue.
	RingSize             *int64 `json:"RingSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicCompletionQueueSettings VnicCompletionQueueSettings

// NewVnicCompletionQueueSettings instantiates a new VnicCompletionQueueSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicCompletionQueueSettings() *VnicCompletionQueueSettings {
	this := VnicCompletionQueueSettings{}
	return &this
}

// NewVnicCompletionQueueSettingsWithDefaults instantiates a new VnicCompletionQueueSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicCompletionQueueSettingsWithDefaults() *VnicCompletionQueueSettings {
	this := VnicCompletionQueueSettings{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicCompletionQueueSettings) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicCompletionQueueSettings) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicCompletionQueueSettings) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *VnicCompletionQueueSettings) SetCount(v int64) {
	o.Count = &v
}

// GetRingSize returns the RingSize field value if set, zero value otherwise.
func (o *VnicCompletionQueueSettings) GetRingSize() int64 {
	if o == nil || o.RingSize == nil {
		var ret int64
		return ret
	}
	return *o.RingSize
}

// GetRingSizeOk returns a tuple with the RingSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicCompletionQueueSettings) GetRingSizeOk() (*int64, bool) {
	if o == nil || o.RingSize == nil {
		return nil, false
	}
	return o.RingSize, true
}

// HasRingSize returns a boolean if a field has been set.
func (o *VnicCompletionQueueSettings) HasRingSize() bool {
	if o != nil && o.RingSize != nil {
		return true
	}

	return false
}

// SetRingSize gets a reference to the given int64 and assigns it to the RingSize field.
func (o *VnicCompletionQueueSettings) SetRingSize(v int64) {
	o.RingSize = &v
}

func (o VnicCompletionQueueSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.RingSize != nil {
		toSerialize["RingSize"] = o.RingSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicCompletionQueueSettings) UnmarshalJSON(bytes []byte) (err error) {
	type VnicCompletionQueueSettingsWithoutEmbeddedStruct struct {
		// The number of completion queue resources to allocate. In general, the number of completion queue resources to allocate is equal to the number of transmit queue resources plus the number of receive queue resources.
		Count *int64 `json:"Count,omitempty"`
		// The number of descriptors in each completion queue.
		RingSize *int64 `json:"RingSize,omitempty"`
	}

	varVnicCompletionQueueSettingsWithoutEmbeddedStruct := VnicCompletionQueueSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicCompletionQueueSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicCompletionQueueSettings := _VnicCompletionQueueSettings{}
		varVnicCompletionQueueSettings.Count = varVnicCompletionQueueSettingsWithoutEmbeddedStruct.Count
		varVnicCompletionQueueSettings.RingSize = varVnicCompletionQueueSettingsWithoutEmbeddedStruct.RingSize
		*o = VnicCompletionQueueSettings(varVnicCompletionQueueSettings)
	} else {
		return err
	}

	varVnicCompletionQueueSettings := _VnicCompletionQueueSettings{}

	err = json.Unmarshal(bytes, &varVnicCompletionQueueSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicCompletionQueueSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "RingSize")

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

type NullableVnicCompletionQueueSettings struct {
	value *VnicCompletionQueueSettings
	isSet bool
}

func (v NullableVnicCompletionQueueSettings) Get() *VnicCompletionQueueSettings {
	return v.value
}

func (v *NullableVnicCompletionQueueSettings) Set(val *VnicCompletionQueueSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicCompletionQueueSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicCompletionQueueSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicCompletionQueueSettings(val *VnicCompletionQueueSettings) *NullableVnicCompletionQueueSettings {
	return &NullableVnicCompletionQueueSettings{value: val, isSet: true}
}

func (v NullableVnicCompletionQueueSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicCompletionQueueSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}