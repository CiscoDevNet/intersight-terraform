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

// FcpoolFcBlockAllOf Definition of the list of properties defined in 'fcpool.FcBlock', excluding properties defined in parent classes.
type FcpoolFcBlockAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                  `json:"ObjectType"`
	IdBlock              *FcpoolBlock            `json:"IdBlock,omitempty"`
	Pool                 *FcpoolPoolRelationship `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcpoolFcBlockAllOf FcpoolFcBlockAllOf

// NewFcpoolFcBlockAllOf instantiates a new FcpoolFcBlockAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolFcBlockAllOf(classId string, objectType string) *FcpoolFcBlockAllOf {
	this := FcpoolFcBlockAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFcpoolFcBlockAllOfWithDefaults instantiates a new FcpoolFcBlockAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolFcBlockAllOfWithDefaults() *FcpoolFcBlockAllOf {
	this := FcpoolFcBlockAllOf{}
	var classId string = "fcpool.FcBlock"
	this.ClassId = classId
	var objectType string = "fcpool.FcBlock"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcpoolFcBlockAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlockAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcpoolFcBlockAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcpoolFcBlockAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlockAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcpoolFcBlockAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdBlock returns the IdBlock field value if set, zero value otherwise.
func (o *FcpoolFcBlockAllOf) GetIdBlock() FcpoolBlock {
	if o == nil || o.IdBlock == nil {
		var ret FcpoolBlock
		return ret
	}
	return *o.IdBlock
}

// GetIdBlockOk returns a tuple with the IdBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlockAllOf) GetIdBlockOk() (*FcpoolBlock, bool) {
	if o == nil || o.IdBlock == nil {
		return nil, false
	}
	return o.IdBlock, true
}

// HasIdBlock returns a boolean if a field has been set.
func (o *FcpoolFcBlockAllOf) HasIdBlock() bool {
	if o != nil && o.IdBlock != nil {
		return true
	}

	return false
}

// SetIdBlock gets a reference to the given FcpoolBlock and assigns it to the IdBlock field.
func (o *FcpoolFcBlockAllOf) SetIdBlock(v FcpoolBlock) {
	o.IdBlock = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *FcpoolFcBlockAllOf) GetPool() FcpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlockAllOf) GetPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *FcpoolFcBlockAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given FcpoolPoolRelationship and assigns it to the Pool field.
func (o *FcpoolFcBlockAllOf) SetPool(v FcpoolPoolRelationship) {
	o.Pool = &v
}

func (o FcpoolFcBlockAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IdBlock != nil {
		toSerialize["IdBlock"] = o.IdBlock
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FcpoolFcBlockAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFcpoolFcBlockAllOf := _FcpoolFcBlockAllOf{}

	if err = json.Unmarshal(bytes, &varFcpoolFcBlockAllOf); err == nil {
		*o = FcpoolFcBlockAllOf(varFcpoolFcBlockAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IdBlock")
		delete(additionalProperties, "Pool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFcpoolFcBlockAllOf struct {
	value *FcpoolFcBlockAllOf
	isSet bool
}

func (v NullableFcpoolFcBlockAllOf) Get() *FcpoolFcBlockAllOf {
	return v.value
}

func (v *NullableFcpoolFcBlockAllOf) Set(val *FcpoolFcBlockAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolFcBlockAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolFcBlockAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolFcBlockAllOf(val *FcpoolFcBlockAllOf) *NullableFcpoolFcBlockAllOf {
	return &NullableFcpoolFcBlockAllOf{value: val, isSet: true}
}

func (v NullableFcpoolFcBlockAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolFcBlockAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
