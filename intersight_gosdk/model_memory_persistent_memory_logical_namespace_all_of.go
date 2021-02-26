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

// MemoryPersistentMemoryLogicalNamespaceAllOf Definition of the list of properties defined in 'memory.PersistentMemoryLogicalNamespace', excluding properties defined in parent classes.
type MemoryPersistentMemoryLogicalNamespaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Capacity of this Namespace that is created or modified.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Mode of this Namespace that is created or modified. * `raw` - The raw mode of Persistent Memory Namespace. * `block` - The block mode of Persistent Memory Namespace.
	Mode *string `json:"Mode,omitempty"`
	// Name of this Namespace to be created on the server.
	Name *string `json:"Name,omitempty"`
	// Socket ID of the region on which this Namespace has to be created or modified. * `1` - The first CPU socket in a server. * `2` - The second CPU socket in a server. * `3` - The third CPU socket in a server. * `4` - The fourth CPU socket in a server.
	SocketId *int32 `json:"SocketId,omitempty"`
	// Socket Memory ID of the region on which this Namespace has to be created or modified. * `Not Applicable` - The socket memory ID is not applicable if app-direct persistent memory type is selected in the goal. * `2` - The second socket memory ID within a socket in a server. * `4` - The fourth socket memory ID within a socket in a server. * `6` - The sixth socket memory ID within a socket in a server. * `8` - The eighth socket memory ID within a socket in a server. * `10` - The tenth socket memory ID within a socket in a server. * `12` - The twelfth socket memory ID within a socket in a server.
	SocketMemoryId       *string `json:"SocketMemoryId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryPersistentMemoryLogicalNamespaceAllOf MemoryPersistentMemoryLogicalNamespaceAllOf

// NewMemoryPersistentMemoryLogicalNamespaceAllOf instantiates a new MemoryPersistentMemoryLogicalNamespaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryLogicalNamespaceAllOf(classId string, objectType string) *MemoryPersistentMemoryLogicalNamespaceAllOf {
	this := MemoryPersistentMemoryLogicalNamespaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mode string = "raw"
	this.Mode = &mode
	var socketId int32 = 1
	this.SocketId = &socketId
	var socketMemoryId string = "Not Applicable"
	this.SocketMemoryId = &socketMemoryId
	return &this
}

// NewMemoryPersistentMemoryLogicalNamespaceAllOfWithDefaults instantiates a new MemoryPersistentMemoryLogicalNamespaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryLogicalNamespaceAllOfWithDefaults() *MemoryPersistentMemoryLogicalNamespaceAllOf {
	this := MemoryPersistentMemoryLogicalNamespaceAllOf{}
	var classId string = "memory.PersistentMemoryLogicalNamespace"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryLogicalNamespace"
	this.ObjectType = objectType
	var mode string = "raw"
	this.Mode = &mode
	var socketId int32 = 1
	this.SocketId = &socketId
	var socketMemoryId string = "Not Applicable"
	this.SocketMemoryId = &socketMemoryId
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetName(v string) {
	o.Name = &v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetSocketId() int32 {
	if o == nil || o.SocketId == nil {
		var ret int32
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetSocketIdOk() (*int32, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given int32 and assigns it to the SocketId field.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetSocketId(v int32) {
	o.SocketId = &v
}

// GetSocketMemoryId returns the SocketMemoryId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetSocketMemoryId() string {
	if o == nil || o.SocketMemoryId == nil {
		var ret string
		return ret
	}
	return *o.SocketMemoryId
}

// GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetSocketMemoryIdOk() (*string, bool) {
	if o == nil || o.SocketMemoryId == nil {
		return nil, false
	}
	return o.SocketMemoryId, true
}

// HasSocketMemoryId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasSocketMemoryId() bool {
	if o != nil && o.SocketMemoryId != nil {
		return true
	}

	return false
}

// SetSocketMemoryId gets a reference to the given string and assigns it to the SocketMemoryId field.
func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetSocketMemoryId(v string) {
	o.SocketMemoryId = &v
}

func (o MemoryPersistentMemoryLogicalNamespaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}
	if o.SocketMemoryId != nil {
		toSerialize["SocketMemoryId"] = o.SocketMemoryId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMemoryPersistentMemoryLogicalNamespaceAllOf := _MemoryPersistentMemoryLogicalNamespaceAllOf{}

	if err = json.Unmarshal(bytes, &varMemoryPersistentMemoryLogicalNamespaceAllOf); err == nil {
		*o = MemoryPersistentMemoryLogicalNamespaceAllOf(varMemoryPersistentMemoryLogicalNamespaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SocketId")
		delete(additionalProperties, "SocketMemoryId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemoryPersistentMemoryLogicalNamespaceAllOf struct {
	value *MemoryPersistentMemoryLogicalNamespaceAllOf
	isSet bool
}

func (v NullableMemoryPersistentMemoryLogicalNamespaceAllOf) Get() *MemoryPersistentMemoryLogicalNamespaceAllOf {
	return v.value
}

func (v *NullableMemoryPersistentMemoryLogicalNamespaceAllOf) Set(val *MemoryPersistentMemoryLogicalNamespaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryLogicalNamespaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryLogicalNamespaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryLogicalNamespaceAllOf(val *MemoryPersistentMemoryLogicalNamespaceAllOf) *NullableMemoryPersistentMemoryLogicalNamespaceAllOf {
	return &NullableMemoryPersistentMemoryLogicalNamespaceAllOf{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryLogicalNamespaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryLogicalNamespaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
