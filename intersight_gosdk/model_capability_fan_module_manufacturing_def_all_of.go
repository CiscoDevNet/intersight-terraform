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

// CapabilityFanModuleManufacturingDefAllOf Definition of the list of properties defined in 'capability.FanModuleManufacturingDef', excluding properties defined in parent classes.
type CapabilityFanModuleManufacturingDefAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Caption for a fan module.
	Caption *string `json:"Caption,omitempty"`
	// Description for a fan module.
	Description *string `json:"Description,omitempty"`
	// Product Identifier for a fan module.
	Pid *string `json:"Pid,omitempty"`
	// Product Name for Fan Module Unit.
	ProductName *string `json:"ProductName,omitempty"`
	// SKU information for a fan module.
	Sku *string `json:"Sku,omitempty"`
	// VID information for a fan module.
	Vid                  *string `json:"Vid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityFanModuleManufacturingDefAllOf CapabilityFanModuleManufacturingDefAllOf

// NewCapabilityFanModuleManufacturingDefAllOf instantiates a new CapabilityFanModuleManufacturingDefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityFanModuleManufacturingDefAllOf(classId string, objectType string) *CapabilityFanModuleManufacturingDefAllOf {
	this := CapabilityFanModuleManufacturingDefAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityFanModuleManufacturingDefAllOfWithDefaults instantiates a new CapabilityFanModuleManufacturingDefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityFanModuleManufacturingDefAllOfWithDefaults() *CapabilityFanModuleManufacturingDefAllOf {
	this := CapabilityFanModuleManufacturingDefAllOf{}
	var classId string = "capability.FanModuleManufacturingDef"
	this.ClassId = classId
	var objectType string = "capability.FanModuleManufacturingDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityFanModuleManufacturingDefAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityFanModuleManufacturingDefAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityFanModuleManufacturingDefAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityFanModuleManufacturingDefAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetCaption() string {
	if o == nil || o.Caption == nil {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetCaptionOk() (*string, bool) {
	if o == nil || o.Caption == nil {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) HasCaption() bool {
	if o != nil && o.Caption != nil {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *CapabilityFanModuleManufacturingDefAllOf) SetCaption(v string) {
	o.Caption = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilityFanModuleManufacturingDefAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *CapabilityFanModuleManufacturingDefAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *CapabilityFanModuleManufacturingDefAllOf) SetProductName(v string) {
	o.ProductName = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *CapabilityFanModuleManufacturingDefAllOf) SetSku(v string) {
	o.Sku = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *CapabilityFanModuleManufacturingDefAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *CapabilityFanModuleManufacturingDefAllOf) SetVid(v string) {
	o.Vid = &v
}

func (o CapabilityFanModuleManufacturingDefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Caption != nil {
		toSerialize["Caption"] = o.Caption
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityFanModuleManufacturingDefAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityFanModuleManufacturingDefAllOf := _CapabilityFanModuleManufacturingDefAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityFanModuleManufacturingDefAllOf); err == nil {
		*o = CapabilityFanModuleManufacturingDefAllOf(varCapabilityFanModuleManufacturingDefAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Caption")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "Vid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityFanModuleManufacturingDefAllOf struct {
	value *CapabilityFanModuleManufacturingDefAllOf
	isSet bool
}

func (v NullableCapabilityFanModuleManufacturingDefAllOf) Get() *CapabilityFanModuleManufacturingDefAllOf {
	return v.value
}

func (v *NullableCapabilityFanModuleManufacturingDefAllOf) Set(val *CapabilityFanModuleManufacturingDefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityFanModuleManufacturingDefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityFanModuleManufacturingDefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityFanModuleManufacturingDefAllOf(val *CapabilityFanModuleManufacturingDefAllOf) *NullableCapabilityFanModuleManufacturingDefAllOf {
	return &NullableCapabilityFanModuleManufacturingDefAllOf{value: val, isSet: true}
}

func (v NullableCapabilityFanModuleManufacturingDefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityFanModuleManufacturingDefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
