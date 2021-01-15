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

// InventoryRequest Request MO allows the inventory of specific devices to be collected on demand. The inventory can be collected in three levels - all the MOs of a specific device, MOs of specific MO types for a given device or specific MO instances of specific MO types for a given device. These MO instances are used just to collect the requests and not persisted.
type InventoryRequest struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	Mos                  []InventoryInventoryMo               `json:"Mos,omitempty"`
	Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryRequest InventoryRequest

// NewInventoryRequest instantiates a new InventoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryRequest(classId string, objectType string) *InventoryRequest {
	this := InventoryRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryRequestWithDefaults instantiates a new InventoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryRequestWithDefaults() *InventoryRequest {
	this := InventoryRequest{}
	var classId string = "inventory.Request"
	this.ClassId = classId
	var objectType string = "inventory.Request"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMos returns the Mos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryRequest) GetMos() []InventoryInventoryMo {
	if o == nil {
		var ret []InventoryInventoryMo
		return ret
	}
	return o.Mos
}

// GetMosOk returns a tuple with the Mos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryRequest) GetMosOk() (*[]InventoryInventoryMo, bool) {
	if o == nil || o.Mos == nil {
		return nil, false
	}
	return &o.Mos, true
}

// HasMos returns a boolean if a field has been set.
func (o *InventoryRequest) HasMos() bool {
	if o != nil && o.Mos != nil {
		return true
	}

	return false
}

// SetMos gets a reference to the given []InventoryInventoryMo and assigns it to the Mos field.
func (o *InventoryRequest) SetMos(v []InventoryInventoryMo) {
	o.Mos = v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InventoryRequest) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRequest) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InventoryRequest) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *InventoryRequest) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

func (o InventoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Mos != nil {
		toSerialize["Mos"] = o.Mos
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryRequest) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                               `json:"ObjectType"`
		Mos        []InventoryInventoryMo               `json:"Mos,omitempty"`
		Device     *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	}

	varInventoryRequestWithoutEmbeddedStruct := InventoryRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryRequestWithoutEmbeddedStruct)
	if err == nil {
		varInventoryRequest := _InventoryRequest{}
		varInventoryRequest.ClassId = varInventoryRequestWithoutEmbeddedStruct.ClassId
		varInventoryRequest.ObjectType = varInventoryRequestWithoutEmbeddedStruct.ObjectType
		varInventoryRequest.Mos = varInventoryRequestWithoutEmbeddedStruct.Mos
		varInventoryRequest.Device = varInventoryRequestWithoutEmbeddedStruct.Device
		*o = InventoryRequest(varInventoryRequest)
	} else {
		return err
	}

	varInventoryRequest := _InventoryRequest{}

	err = json.Unmarshal(bytes, &varInventoryRequest)
	if err == nil {
		o.MoBaseMo = varInventoryRequest.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Mos")
		delete(additionalProperties, "Device")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableInventoryRequest struct {
	value *InventoryRequest
	isSet bool
}

func (v NullableInventoryRequest) Get() *InventoryRequest {
	return v.value
}

func (v *NullableInventoryRequest) Set(val *InventoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryRequest(val *InventoryRequest) *NullableInventoryRequest {
	return &NullableInventoryRequest{value: val, isSet: true}
}

func (v NullableInventoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
