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

// StoragePhysicalDiskUsage Has usage map between physical disks and virtual drives.
type StoragePhysicalDiskUsage struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of blocks that are a part of the virtual drive.
	NumberOfBlocks *string `json:"NumberOfBlocks,omitempty"`
	// The physical disk for which the usage is reported.
	PhysicalDrive *string `json:"PhysicalDrive,omitempty"`
	// The span of the physical disk.
	Span *string `json:"Span,omitempty"`
	// The starting block id of the virtual drive within the physical drive.
	StartingBlock *string `json:"StartingBlock,omitempty"`
	// The current state of the physical disk usage.
	State *string `json:"State,omitempty"`
	// The virtual drive corresponding to the physical disk.
	VirtualDrive         *string                              `json:"VirtualDrive,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePhysicalDiskUsage StoragePhysicalDiskUsage

// NewStoragePhysicalDiskUsage instantiates a new StoragePhysicalDiskUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePhysicalDiskUsage(classId string, objectType string) *StoragePhysicalDiskUsage {
	this := StoragePhysicalDiskUsage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePhysicalDiskUsageWithDefaults instantiates a new StoragePhysicalDiskUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePhysicalDiskUsageWithDefaults() *StoragePhysicalDiskUsage {
	this := StoragePhysicalDiskUsage{}
	var classId string = "storage.PhysicalDiskUsage"
	this.ClassId = classId
	var objectType string = "storage.PhysicalDiskUsage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePhysicalDiskUsage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePhysicalDiskUsage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePhysicalDiskUsage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePhysicalDiskUsage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNumberOfBlocks returns the NumberOfBlocks field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetNumberOfBlocks() string {
	if o == nil || o.NumberOfBlocks == nil {
		var ret string
		return ret
	}
	return *o.NumberOfBlocks
}

// GetNumberOfBlocksOk returns a tuple with the NumberOfBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetNumberOfBlocksOk() (*string, bool) {
	if o == nil || o.NumberOfBlocks == nil {
		return nil, false
	}
	return o.NumberOfBlocks, true
}

// HasNumberOfBlocks returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasNumberOfBlocks() bool {
	if o != nil && o.NumberOfBlocks != nil {
		return true
	}

	return false
}

// SetNumberOfBlocks gets a reference to the given string and assigns it to the NumberOfBlocks field.
func (o *StoragePhysicalDiskUsage) SetNumberOfBlocks(v string) {
	o.NumberOfBlocks = &v
}

// GetPhysicalDrive returns the PhysicalDrive field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetPhysicalDrive() string {
	if o == nil || o.PhysicalDrive == nil {
		var ret string
		return ret
	}
	return *o.PhysicalDrive
}

// GetPhysicalDriveOk returns a tuple with the PhysicalDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetPhysicalDriveOk() (*string, bool) {
	if o == nil || o.PhysicalDrive == nil {
		return nil, false
	}
	return o.PhysicalDrive, true
}

// HasPhysicalDrive returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasPhysicalDrive() bool {
	if o != nil && o.PhysicalDrive != nil {
		return true
	}

	return false
}

// SetPhysicalDrive gets a reference to the given string and assigns it to the PhysicalDrive field.
func (o *StoragePhysicalDiskUsage) SetPhysicalDrive(v string) {
	o.PhysicalDrive = &v
}

// GetSpan returns the Span field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetSpan() string {
	if o == nil || o.Span == nil {
		var ret string
		return ret
	}
	return *o.Span
}

// GetSpanOk returns a tuple with the Span field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetSpanOk() (*string, bool) {
	if o == nil || o.Span == nil {
		return nil, false
	}
	return o.Span, true
}

// HasSpan returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasSpan() bool {
	if o != nil && o.Span != nil {
		return true
	}

	return false
}

// SetSpan gets a reference to the given string and assigns it to the Span field.
func (o *StoragePhysicalDiskUsage) SetSpan(v string) {
	o.Span = &v
}

// GetStartingBlock returns the StartingBlock field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetStartingBlock() string {
	if o == nil || o.StartingBlock == nil {
		var ret string
		return ret
	}
	return *o.StartingBlock
}

// GetStartingBlockOk returns a tuple with the StartingBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetStartingBlockOk() (*string, bool) {
	if o == nil || o.StartingBlock == nil {
		return nil, false
	}
	return o.StartingBlock, true
}

// HasStartingBlock returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasStartingBlock() bool {
	if o != nil && o.StartingBlock != nil {
		return true
	}

	return false
}

// SetStartingBlock gets a reference to the given string and assigns it to the StartingBlock field.
func (o *StoragePhysicalDiskUsage) SetStartingBlock(v string) {
	o.StartingBlock = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StoragePhysicalDiskUsage) SetState(v string) {
	o.State = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetVirtualDrive() string {
	if o == nil || o.VirtualDrive == nil {
		var ret string
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetVirtualDriveOk() (*string, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given string and assigns it to the VirtualDrive field.
func (o *StoragePhysicalDiskUsage) SetVirtualDrive(v string) {
	o.VirtualDrive = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StoragePhysicalDiskUsage) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePhysicalDiskUsage) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePhysicalDiskUsage) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePhysicalDiskUsage) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePhysicalDiskUsage) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePhysicalDiskUsage) MarshalJSON() ([]byte, error) {
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
	if o.NumberOfBlocks != nil {
		toSerialize["NumberOfBlocks"] = o.NumberOfBlocks
	}
	if o.PhysicalDrive != nil {
		toSerialize["PhysicalDrive"] = o.PhysicalDrive
	}
	if o.Span != nil {
		toSerialize["Span"] = o.Span
	}
	if o.StartingBlock != nil {
		toSerialize["StartingBlock"] = o.StartingBlock
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePhysicalDiskUsage) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePhysicalDiskUsageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of blocks that are a part of the virtual drive.
		NumberOfBlocks *string `json:"NumberOfBlocks,omitempty"`
		// The physical disk for which the usage is reported.
		PhysicalDrive *string `json:"PhysicalDrive,omitempty"`
		// The span of the physical disk.
		Span *string `json:"Span,omitempty"`
		// The starting block id of the virtual drive within the physical drive.
		StartingBlock *string `json:"StartingBlock,omitempty"`
		// The current state of the physical disk usage.
		State *string `json:"State,omitempty"`
		// The virtual drive corresponding to the physical disk.
		VirtualDrive        *string                              `json:"VirtualDrive,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStoragePhysicalDiskUsageWithoutEmbeddedStruct := StoragePhysicalDiskUsageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePhysicalDiskUsageWithoutEmbeddedStruct)
	if err == nil {
		varStoragePhysicalDiskUsage := _StoragePhysicalDiskUsage{}
		varStoragePhysicalDiskUsage.ClassId = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.ClassId
		varStoragePhysicalDiskUsage.ObjectType = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.ObjectType
		varStoragePhysicalDiskUsage.NumberOfBlocks = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.NumberOfBlocks
		varStoragePhysicalDiskUsage.PhysicalDrive = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.PhysicalDrive
		varStoragePhysicalDiskUsage.Span = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.Span
		varStoragePhysicalDiskUsage.StartingBlock = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.StartingBlock
		varStoragePhysicalDiskUsage.State = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.State
		varStoragePhysicalDiskUsage.VirtualDrive = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.VirtualDrive
		varStoragePhysicalDiskUsage.InventoryDeviceInfo = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.InventoryDeviceInfo
		varStoragePhysicalDiskUsage.RegisteredDevice = varStoragePhysicalDiskUsageWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePhysicalDiskUsage(varStoragePhysicalDiskUsage)
	} else {
		return err
	}

	varStoragePhysicalDiskUsage := _StoragePhysicalDiskUsage{}

	err = json.Unmarshal(bytes, &varStoragePhysicalDiskUsage)
	if err == nil {
		o.InventoryBase = varStoragePhysicalDiskUsage.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NumberOfBlocks")
		delete(additionalProperties, "PhysicalDrive")
		delete(additionalProperties, "Span")
		delete(additionalProperties, "StartingBlock")
		delete(additionalProperties, "State")
		delete(additionalProperties, "VirtualDrive")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableStoragePhysicalDiskUsage struct {
	value *StoragePhysicalDiskUsage
	isSet bool
}

func (v NullableStoragePhysicalDiskUsage) Get() *StoragePhysicalDiskUsage {
	return v.value
}

func (v *NullableStoragePhysicalDiskUsage) Set(val *StoragePhysicalDiskUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePhysicalDiskUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePhysicalDiskUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePhysicalDiskUsage(val *StoragePhysicalDiskUsage) *NullableStoragePhysicalDiskUsage {
	return &NullableStoragePhysicalDiskUsage{value: val, isSet: true}
}

func (v NullableStoragePhysicalDiskUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePhysicalDiskUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
