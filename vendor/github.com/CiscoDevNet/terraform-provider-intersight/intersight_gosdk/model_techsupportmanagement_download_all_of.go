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

// TechsupportmanagementDownloadAllOf Definition of the list of properties defined in 'techsupportmanagement.Download', excluding properties defined in parent classes.
type TechsupportmanagementDownloadAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                              `json:"ObjectType"`
	TechSupportStatus    *TechsupportmanagementTechSupportStatusRelationship `json:"TechSupportStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TechsupportmanagementDownloadAllOf TechsupportmanagementDownloadAllOf

// NewTechsupportmanagementDownloadAllOf instantiates a new TechsupportmanagementDownloadAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementDownloadAllOf(classId string, objectType string) *TechsupportmanagementDownloadAllOf {
	this := TechsupportmanagementDownloadAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTechsupportmanagementDownloadAllOfWithDefaults instantiates a new TechsupportmanagementDownloadAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementDownloadAllOfWithDefaults() *TechsupportmanagementDownloadAllOf {
	this := TechsupportmanagementDownloadAllOf{}
	var classId string = "techsupportmanagement.Download"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.Download"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementDownloadAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementDownloadAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementDownloadAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementDownloadAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementDownloadAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementDownloadAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTechSupportStatus returns the TechSupportStatus field value if set, zero value otherwise.
func (o *TechsupportmanagementDownloadAllOf) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship {
	if o == nil || o.TechSupportStatus == nil {
		var ret TechsupportmanagementTechSupportStatusRelationship
		return ret
	}
	return *o.TechSupportStatus
}

// GetTechSupportStatusOk returns a tuple with the TechSupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementDownloadAllOf) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool) {
	if o == nil || o.TechSupportStatus == nil {
		return nil, false
	}
	return o.TechSupportStatus, true
}

// HasTechSupportStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementDownloadAllOf) HasTechSupportStatus() bool {
	if o != nil && o.TechSupportStatus != nil {
		return true
	}

	return false
}

// SetTechSupportStatus gets a reference to the given TechsupportmanagementTechSupportStatusRelationship and assigns it to the TechSupportStatus field.
func (o *TechsupportmanagementDownloadAllOf) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship) {
	o.TechSupportStatus = &v
}

func (o TechsupportmanagementDownloadAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TechSupportStatus != nil {
		toSerialize["TechSupportStatus"] = o.TechSupportStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TechsupportmanagementDownloadAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTechsupportmanagementDownloadAllOf := _TechsupportmanagementDownloadAllOf{}

	if err = json.Unmarshal(bytes, &varTechsupportmanagementDownloadAllOf); err == nil {
		*o = TechsupportmanagementDownloadAllOf(varTechsupportmanagementDownloadAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TechSupportStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTechsupportmanagementDownloadAllOf struct {
	value *TechsupportmanagementDownloadAllOf
	isSet bool
}

func (v NullableTechsupportmanagementDownloadAllOf) Get() *TechsupportmanagementDownloadAllOf {
	return v.value
}

func (v *NullableTechsupportmanagementDownloadAllOf) Set(val *TechsupportmanagementDownloadAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementDownloadAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementDownloadAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementDownloadAllOf(val *TechsupportmanagementDownloadAllOf) *NullableTechsupportmanagementDownloadAllOf {
	return &NullableTechsupportmanagementDownloadAllOf{value: val, isSet: true}
}

func (v NullableTechsupportmanagementDownloadAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementDownloadAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
