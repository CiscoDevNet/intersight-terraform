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

// HyperflexHxapDatacenterAllOf Definition of the list of properties defined in 'hyperflex.HxapDatacenter', excluding properties defined in parent classes.
type HyperflexHxapDatacenterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                       `json:"ObjectType"`
	Account              *IamAccountRelationship                      `json:"Account,omitempty"`
	HypervisorManager    *HyperflexCiscoHypervisorManagerRelationship `json:"HypervisorManager,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapDatacenterAllOf HyperflexHxapDatacenterAllOf

// NewHyperflexHxapDatacenterAllOf instantiates a new HyperflexHxapDatacenterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapDatacenterAllOf(classId string, objectType string) *HyperflexHxapDatacenterAllOf {
	this := HyperflexHxapDatacenterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxapDatacenterAllOfWithDefaults instantiates a new HyperflexHxapDatacenterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapDatacenterAllOfWithDefaults() *HyperflexHxapDatacenterAllOf {
	this := HyperflexHxapDatacenterAllOf{}
	var classId string = "hyperflex.HxapDatacenter"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapDatacenter"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapDatacenterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDatacenterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapDatacenterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapDatacenterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDatacenterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapDatacenterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *HyperflexHxapDatacenterAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDatacenterAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *HyperflexHxapDatacenterAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *HyperflexHxapDatacenterAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetHypervisorManager returns the HypervisorManager field value if set, zero value otherwise.
func (o *HyperflexHxapDatacenterAllOf) GetHypervisorManager() HyperflexCiscoHypervisorManagerRelationship {
	if o == nil || o.HypervisorManager == nil {
		var ret HyperflexCiscoHypervisorManagerRelationship
		return ret
	}
	return *o.HypervisorManager
}

// GetHypervisorManagerOk returns a tuple with the HypervisorManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDatacenterAllOf) GetHypervisorManagerOk() (*HyperflexCiscoHypervisorManagerRelationship, bool) {
	if o == nil || o.HypervisorManager == nil {
		return nil, false
	}
	return o.HypervisorManager, true
}

// HasHypervisorManager returns a boolean if a field has been set.
func (o *HyperflexHxapDatacenterAllOf) HasHypervisorManager() bool {
	if o != nil && o.HypervisorManager != nil {
		return true
	}

	return false
}

// SetHypervisorManager gets a reference to the given HyperflexCiscoHypervisorManagerRelationship and assigns it to the HypervisorManager field.
func (o *HyperflexHxapDatacenterAllOf) SetHypervisorManager(v HyperflexCiscoHypervisorManagerRelationship) {
	o.HypervisorManager = &v
}

func (o HyperflexHxapDatacenterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.HypervisorManager != nil {
		toSerialize["HypervisorManager"] = o.HypervisorManager
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapDatacenterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxapDatacenterAllOf := _HyperflexHxapDatacenterAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxapDatacenterAllOf); err == nil {
		*o = HyperflexHxapDatacenterAllOf(varHyperflexHxapDatacenterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "HypervisorManager")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxapDatacenterAllOf struct {
	value *HyperflexHxapDatacenterAllOf
	isSet bool
}

func (v NullableHyperflexHxapDatacenterAllOf) Get() *HyperflexHxapDatacenterAllOf {
	return v.value
}

func (v *NullableHyperflexHxapDatacenterAllOf) Set(val *HyperflexHxapDatacenterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapDatacenterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapDatacenterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapDatacenterAllOf(val *HyperflexHxapDatacenterAllOf) *NullableHyperflexHxapDatacenterAllOf {
	return &NullableHyperflexHxapDatacenterAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxapDatacenterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapDatacenterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
