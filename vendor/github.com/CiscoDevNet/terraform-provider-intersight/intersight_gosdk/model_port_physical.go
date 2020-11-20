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

// PortPhysical Abstract super class for all ports.
type PortPhysical struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Operational state of this port (enabled/disabled).
	OperState *string `json:"OperState,omitempty"`
	// Reason for this port's Operational state.
	OperStateQual *string `json:"OperStateQual,omitempty"`
	// Switch physical port identifier.
	PortId *int64 `json:"PortId,omitempty"`
	// The role assigned to this port.
	Role *string `json:"Role,omitempty"`
	// Switch expansion slot module identifier.
	SlotId *int64 `json:"SlotId,omitempty"`
	// Switch Identifier that is local to a cluster.
	SwitchId             *string `json:"SwitchId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortPhysical PortPhysical

// NewPortPhysical instantiates a new PortPhysical object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortPhysical(classId string, objectType string) *PortPhysical {
	this := PortPhysical{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPortPhysicalWithDefaults instantiates a new PortPhysical object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortPhysicalWithDefaults() *PortPhysical {
	this := PortPhysical{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PortPhysical) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PortPhysical) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PortPhysical) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PortPhysical) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PortPhysical) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PortPhysical) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *PortPhysical) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysical) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *PortPhysical) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *PortPhysical) SetOperState(v string) {
	o.OperState = &v
}

// GetOperStateQual returns the OperStateQual field value if set, zero value otherwise.
func (o *PortPhysical) GetOperStateQual() string {
	if o == nil || o.OperStateQual == nil {
		var ret string
		return ret
	}
	return *o.OperStateQual
}

// GetOperStateQualOk returns a tuple with the OperStateQual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysical) GetOperStateQualOk() (*string, bool) {
	if o == nil || o.OperStateQual == nil {
		return nil, false
	}
	return o.OperStateQual, true
}

// HasOperStateQual returns a boolean if a field has been set.
func (o *PortPhysical) HasOperStateQual() bool {
	if o != nil && o.OperStateQual != nil {
		return true
	}

	return false
}

// SetOperStateQual gets a reference to the given string and assigns it to the OperStateQual field.
func (o *PortPhysical) SetOperStateQual(v string) {
	o.OperStateQual = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *PortPhysical) GetPortId() int64 {
	if o == nil || o.PortId == nil {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysical) GetPortIdOk() (*int64, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *PortPhysical) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *PortPhysical) SetPortId(v int64) {
	o.PortId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PortPhysical) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysical) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PortPhysical) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *PortPhysical) SetRole(v string) {
	o.Role = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *PortPhysical) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysical) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *PortPhysical) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *PortPhysical) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *PortPhysical) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysical) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *PortPhysical) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *PortPhysical) SetSwitchId(v string) {
	o.SwitchId = &v
}

func (o PortPhysical) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.OperStateQual != nil {
		toSerialize["OperStateQual"] = o.OperStateQual
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PortPhysical) UnmarshalJSON(bytes []byte) (err error) {
	type PortPhysicalWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Operational state of this port (enabled/disabled).
		OperState *string `json:"OperState,omitempty"`
		// Reason for this port's Operational state.
		OperStateQual *string `json:"OperStateQual,omitempty"`
		// Switch physical port identifier.
		PortId *int64 `json:"PortId,omitempty"`
		// The role assigned to this port.
		Role *string `json:"Role,omitempty"`
		// Switch expansion slot module identifier.
		SlotId *int64 `json:"SlotId,omitempty"`
		// Switch Identifier that is local to a cluster.
		SwitchId *string `json:"SwitchId,omitempty"`
	}

	varPortPhysicalWithoutEmbeddedStruct := PortPhysicalWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPortPhysicalWithoutEmbeddedStruct)
	if err == nil {
		varPortPhysical := _PortPhysical{}
		varPortPhysical.ClassId = varPortPhysicalWithoutEmbeddedStruct.ClassId
		varPortPhysical.ObjectType = varPortPhysicalWithoutEmbeddedStruct.ObjectType
		varPortPhysical.OperState = varPortPhysicalWithoutEmbeddedStruct.OperState
		varPortPhysical.OperStateQual = varPortPhysicalWithoutEmbeddedStruct.OperStateQual
		varPortPhysical.PortId = varPortPhysicalWithoutEmbeddedStruct.PortId
		varPortPhysical.Role = varPortPhysicalWithoutEmbeddedStruct.Role
		varPortPhysical.SlotId = varPortPhysicalWithoutEmbeddedStruct.SlotId
		varPortPhysical.SwitchId = varPortPhysicalWithoutEmbeddedStruct.SwitchId
		*o = PortPhysical(varPortPhysical)
	} else {
		return err
	}

	varPortPhysical := _PortPhysical{}

	err = json.Unmarshal(bytes, &varPortPhysical)
	if err == nil {
		o.InventoryBase = varPortPhysical.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "OperStateQual")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "Role")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "SwitchId")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullablePortPhysical struct {
	value *PortPhysical
	isSet bool
}

func (v NullablePortPhysical) Get() *PortPhysical {
	return v.value
}

func (v *NullablePortPhysical) Set(val *PortPhysical) {
	v.value = val
	v.isSet = true
}

func (v NullablePortPhysical) IsSet() bool {
	return v.isSet
}

func (v *NullablePortPhysical) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortPhysical(val *PortPhysical) *NullablePortPhysical {
	return &NullablePortPhysical{value: val, isSet: true}
}

func (v NullablePortPhysical) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortPhysical) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
