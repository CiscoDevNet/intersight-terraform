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

// MetadataSkuVolumeTypeAllOf Definition of the list of properties defined in 'metadata.SkuVolumeType', excluding properties defined in parent classes.
type MetadataSkuVolumeTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The units to measure IOPS.
	IopsUnit *string `json:"IopsUnit,omitempty"`
	// Indicates if this volume can be used as a boot volume.
	IsBootable *bool `json:"IsBootable,omitempty"`
	// Flag to indicate if this is a default volume. Default volumes will be used when an instance type is launched unless another volume type is specified.
	IsDefault *bool `json:"IsDefault,omitempty"`
	// Indicates if this volume type supports provisioned IOPS.
	IsProvisionedIops *bool `json:"IsProvisionedIops,omitempty"`
	// The max I/O operations per second that this volume type supports. Read or write numbers does not go beyond this value.
	MaxIops *float64 `json:"MaxIops,omitempty"`
	// The maximum read IOPS that this volume type supports.
	MaxReadIops *float64 `json:"MaxReadIops,omitempty"`
	// The maximum read throughput limit for this volume type.
	MaxReadThroughput *float64 `json:"MaxReadThroughput,omitempty"`
	// The maximum throughput limit for this volume type.
	MaxThroughput *float64 `json:"MaxThroughput,omitempty"`
	// The maximum storage size that this volume type supports.
	MaxVolumeSize *float64 `json:"MaxVolumeSize,omitempty"`
	// The maximum write IOPS that this volume type supports.
	MaxWriteIops *float64 `json:"MaxWriteIops,omitempty"`
	// The maximum write throughput limit for this volume type.
	MaxWriteThroughput *float64 `json:"MaxWriteThroughput,omitempty"`
	// The minimum storage size that this volume type supports.
	MinVolumeSize *float64 `json:"MinVolumeSize,omitempty"`
	// The units for measuring throughput.
	ThroughputUnit *string `json:"ThroughputUnit,omitempty"`
	// The volume type like gp2 or st1.
	Type *string `json:"Type,omitempty"`
	// The units for measuring volume size.
	VolumeSizeUnit       *string `json:"VolumeSizeUnit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetadataSkuVolumeTypeAllOf MetadataSkuVolumeTypeAllOf

// NewMetadataSkuVolumeTypeAllOf instantiates a new MetadataSkuVolumeTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataSkuVolumeTypeAllOf(classId string, objectType string) *MetadataSkuVolumeTypeAllOf {
	this := MetadataSkuVolumeTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMetadataSkuVolumeTypeAllOfWithDefaults instantiates a new MetadataSkuVolumeTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataSkuVolumeTypeAllOfWithDefaults() *MetadataSkuVolumeTypeAllOf {
	this := MetadataSkuVolumeTypeAllOf{}
	var classId string = "metadata.SkuVolumeType"
	this.ClassId = classId
	var objectType string = "metadata.SkuVolumeType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetadataSkuVolumeTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetadataSkuVolumeTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MetadataSkuVolumeTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetadataSkuVolumeTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIopsUnit returns the IopsUnit field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetIopsUnit() string {
	if o == nil || o.IopsUnit == nil {
		var ret string
		return ret
	}
	return *o.IopsUnit
}

// GetIopsUnitOk returns a tuple with the IopsUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetIopsUnitOk() (*string, bool) {
	if o == nil || o.IopsUnit == nil {
		return nil, false
	}
	return o.IopsUnit, true
}

// HasIopsUnit returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasIopsUnit() bool {
	if o != nil && o.IopsUnit != nil {
		return true
	}

	return false
}

// SetIopsUnit gets a reference to the given string and assigns it to the IopsUnit field.
func (o *MetadataSkuVolumeTypeAllOf) SetIopsUnit(v string) {
	o.IopsUnit = &v
}

// GetIsBootable returns the IsBootable field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetIsBootable() bool {
	if o == nil || o.IsBootable == nil {
		var ret bool
		return ret
	}
	return *o.IsBootable
}

// GetIsBootableOk returns a tuple with the IsBootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetIsBootableOk() (*bool, bool) {
	if o == nil || o.IsBootable == nil {
		return nil, false
	}
	return o.IsBootable, true
}

// HasIsBootable returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasIsBootable() bool {
	if o != nil && o.IsBootable != nil {
		return true
	}

	return false
}

// SetIsBootable gets a reference to the given bool and assigns it to the IsBootable field.
func (o *MetadataSkuVolumeTypeAllOf) SetIsBootable(v bool) {
	o.IsBootable = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *MetadataSkuVolumeTypeAllOf) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsProvisionedIops returns the IsProvisionedIops field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetIsProvisionedIops() bool {
	if o == nil || o.IsProvisionedIops == nil {
		var ret bool
		return ret
	}
	return *o.IsProvisionedIops
}

// GetIsProvisionedIopsOk returns a tuple with the IsProvisionedIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetIsProvisionedIopsOk() (*bool, bool) {
	if o == nil || o.IsProvisionedIops == nil {
		return nil, false
	}
	return o.IsProvisionedIops, true
}

// HasIsProvisionedIops returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasIsProvisionedIops() bool {
	if o != nil && o.IsProvisionedIops != nil {
		return true
	}

	return false
}

// SetIsProvisionedIops gets a reference to the given bool and assigns it to the IsProvisionedIops field.
func (o *MetadataSkuVolumeTypeAllOf) SetIsProvisionedIops(v bool) {
	o.IsProvisionedIops = &v
}

// GetMaxIops returns the MaxIops field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxIops() float64 {
	if o == nil || o.MaxIops == nil {
		var ret float64
		return ret
	}
	return *o.MaxIops
}

// GetMaxIopsOk returns a tuple with the MaxIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxIopsOk() (*float64, bool) {
	if o == nil || o.MaxIops == nil {
		return nil, false
	}
	return o.MaxIops, true
}

// HasMaxIops returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasMaxIops() bool {
	if o != nil && o.MaxIops != nil {
		return true
	}

	return false
}

// SetMaxIops gets a reference to the given float64 and assigns it to the MaxIops field.
func (o *MetadataSkuVolumeTypeAllOf) SetMaxIops(v float64) {
	o.MaxIops = &v
}

// GetMaxReadIops returns the MaxReadIops field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxReadIops() float64 {
	if o == nil || o.MaxReadIops == nil {
		var ret float64
		return ret
	}
	return *o.MaxReadIops
}

// GetMaxReadIopsOk returns a tuple with the MaxReadIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxReadIopsOk() (*float64, bool) {
	if o == nil || o.MaxReadIops == nil {
		return nil, false
	}
	return o.MaxReadIops, true
}

// HasMaxReadIops returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasMaxReadIops() bool {
	if o != nil && o.MaxReadIops != nil {
		return true
	}

	return false
}

// SetMaxReadIops gets a reference to the given float64 and assigns it to the MaxReadIops field.
func (o *MetadataSkuVolumeTypeAllOf) SetMaxReadIops(v float64) {
	o.MaxReadIops = &v
}

// GetMaxReadThroughput returns the MaxReadThroughput field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxReadThroughput() float64 {
	if o == nil || o.MaxReadThroughput == nil {
		var ret float64
		return ret
	}
	return *o.MaxReadThroughput
}

// GetMaxReadThroughputOk returns a tuple with the MaxReadThroughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxReadThroughputOk() (*float64, bool) {
	if o == nil || o.MaxReadThroughput == nil {
		return nil, false
	}
	return o.MaxReadThroughput, true
}

// HasMaxReadThroughput returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasMaxReadThroughput() bool {
	if o != nil && o.MaxReadThroughput != nil {
		return true
	}

	return false
}

// SetMaxReadThroughput gets a reference to the given float64 and assigns it to the MaxReadThroughput field.
func (o *MetadataSkuVolumeTypeAllOf) SetMaxReadThroughput(v float64) {
	o.MaxReadThroughput = &v
}

// GetMaxThroughput returns the MaxThroughput field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxThroughput() float64 {
	if o == nil || o.MaxThroughput == nil {
		var ret float64
		return ret
	}
	return *o.MaxThroughput
}

// GetMaxThroughputOk returns a tuple with the MaxThroughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxThroughputOk() (*float64, bool) {
	if o == nil || o.MaxThroughput == nil {
		return nil, false
	}
	return o.MaxThroughput, true
}

// HasMaxThroughput returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasMaxThroughput() bool {
	if o != nil && o.MaxThroughput != nil {
		return true
	}

	return false
}

// SetMaxThroughput gets a reference to the given float64 and assigns it to the MaxThroughput field.
func (o *MetadataSkuVolumeTypeAllOf) SetMaxThroughput(v float64) {
	o.MaxThroughput = &v
}

// GetMaxVolumeSize returns the MaxVolumeSize field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxVolumeSize() float64 {
	if o == nil || o.MaxVolumeSize == nil {
		var ret float64
		return ret
	}
	return *o.MaxVolumeSize
}

// GetMaxVolumeSizeOk returns a tuple with the MaxVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxVolumeSizeOk() (*float64, bool) {
	if o == nil || o.MaxVolumeSize == nil {
		return nil, false
	}
	return o.MaxVolumeSize, true
}

// HasMaxVolumeSize returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasMaxVolumeSize() bool {
	if o != nil && o.MaxVolumeSize != nil {
		return true
	}

	return false
}

// SetMaxVolumeSize gets a reference to the given float64 and assigns it to the MaxVolumeSize field.
func (o *MetadataSkuVolumeTypeAllOf) SetMaxVolumeSize(v float64) {
	o.MaxVolumeSize = &v
}

// GetMaxWriteIops returns the MaxWriteIops field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxWriteIops() float64 {
	if o == nil || o.MaxWriteIops == nil {
		var ret float64
		return ret
	}
	return *o.MaxWriteIops
}

// GetMaxWriteIopsOk returns a tuple with the MaxWriteIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxWriteIopsOk() (*float64, bool) {
	if o == nil || o.MaxWriteIops == nil {
		return nil, false
	}
	return o.MaxWriteIops, true
}

// HasMaxWriteIops returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasMaxWriteIops() bool {
	if o != nil && o.MaxWriteIops != nil {
		return true
	}

	return false
}

// SetMaxWriteIops gets a reference to the given float64 and assigns it to the MaxWriteIops field.
func (o *MetadataSkuVolumeTypeAllOf) SetMaxWriteIops(v float64) {
	o.MaxWriteIops = &v
}

// GetMaxWriteThroughput returns the MaxWriteThroughput field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxWriteThroughput() float64 {
	if o == nil || o.MaxWriteThroughput == nil {
		var ret float64
		return ret
	}
	return *o.MaxWriteThroughput
}

// GetMaxWriteThroughputOk returns a tuple with the MaxWriteThroughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetMaxWriteThroughputOk() (*float64, bool) {
	if o == nil || o.MaxWriteThroughput == nil {
		return nil, false
	}
	return o.MaxWriteThroughput, true
}

// HasMaxWriteThroughput returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasMaxWriteThroughput() bool {
	if o != nil && o.MaxWriteThroughput != nil {
		return true
	}

	return false
}

// SetMaxWriteThroughput gets a reference to the given float64 and assigns it to the MaxWriteThroughput field.
func (o *MetadataSkuVolumeTypeAllOf) SetMaxWriteThroughput(v float64) {
	o.MaxWriteThroughput = &v
}

// GetMinVolumeSize returns the MinVolumeSize field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetMinVolumeSize() float64 {
	if o == nil || o.MinVolumeSize == nil {
		var ret float64
		return ret
	}
	return *o.MinVolumeSize
}

// GetMinVolumeSizeOk returns a tuple with the MinVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetMinVolumeSizeOk() (*float64, bool) {
	if o == nil || o.MinVolumeSize == nil {
		return nil, false
	}
	return o.MinVolumeSize, true
}

// HasMinVolumeSize returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasMinVolumeSize() bool {
	if o != nil && o.MinVolumeSize != nil {
		return true
	}

	return false
}

// SetMinVolumeSize gets a reference to the given float64 and assigns it to the MinVolumeSize field.
func (o *MetadataSkuVolumeTypeAllOf) SetMinVolumeSize(v float64) {
	o.MinVolumeSize = &v
}

// GetThroughputUnit returns the ThroughputUnit field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetThroughputUnit() string {
	if o == nil || o.ThroughputUnit == nil {
		var ret string
		return ret
	}
	return *o.ThroughputUnit
}

// GetThroughputUnitOk returns a tuple with the ThroughputUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetThroughputUnitOk() (*string, bool) {
	if o == nil || o.ThroughputUnit == nil {
		return nil, false
	}
	return o.ThroughputUnit, true
}

// HasThroughputUnit returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasThroughputUnit() bool {
	if o != nil && o.ThroughputUnit != nil {
		return true
	}

	return false
}

// SetThroughputUnit gets a reference to the given string and assigns it to the ThroughputUnit field.
func (o *MetadataSkuVolumeTypeAllOf) SetThroughputUnit(v string) {
	o.ThroughputUnit = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MetadataSkuVolumeTypeAllOf) SetType(v string) {
	o.Type = &v
}

// GetVolumeSizeUnit returns the VolumeSizeUnit field value if set, zero value otherwise.
func (o *MetadataSkuVolumeTypeAllOf) GetVolumeSizeUnit() string {
	if o == nil || o.VolumeSizeUnit == nil {
		var ret string
		return ret
	}
	return *o.VolumeSizeUnit
}

// GetVolumeSizeUnitOk returns a tuple with the VolumeSizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSkuVolumeTypeAllOf) GetVolumeSizeUnitOk() (*string, bool) {
	if o == nil || o.VolumeSizeUnit == nil {
		return nil, false
	}
	return o.VolumeSizeUnit, true
}

// HasVolumeSizeUnit returns a boolean if a field has been set.
func (o *MetadataSkuVolumeTypeAllOf) HasVolumeSizeUnit() bool {
	if o != nil && o.VolumeSizeUnit != nil {
		return true
	}

	return false
}

// SetVolumeSizeUnit gets a reference to the given string and assigns it to the VolumeSizeUnit field.
func (o *MetadataSkuVolumeTypeAllOf) SetVolumeSizeUnit(v string) {
	o.VolumeSizeUnit = &v
}

func (o MetadataSkuVolumeTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IopsUnit != nil {
		toSerialize["IopsUnit"] = o.IopsUnit
	}
	if o.IsBootable != nil {
		toSerialize["IsBootable"] = o.IsBootable
	}
	if o.IsDefault != nil {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if o.IsProvisionedIops != nil {
		toSerialize["IsProvisionedIops"] = o.IsProvisionedIops
	}
	if o.MaxIops != nil {
		toSerialize["MaxIops"] = o.MaxIops
	}
	if o.MaxReadIops != nil {
		toSerialize["MaxReadIops"] = o.MaxReadIops
	}
	if o.MaxReadThroughput != nil {
		toSerialize["MaxReadThroughput"] = o.MaxReadThroughput
	}
	if o.MaxThroughput != nil {
		toSerialize["MaxThroughput"] = o.MaxThroughput
	}
	if o.MaxVolumeSize != nil {
		toSerialize["MaxVolumeSize"] = o.MaxVolumeSize
	}
	if o.MaxWriteIops != nil {
		toSerialize["MaxWriteIops"] = o.MaxWriteIops
	}
	if o.MaxWriteThroughput != nil {
		toSerialize["MaxWriteThroughput"] = o.MaxWriteThroughput
	}
	if o.MinVolumeSize != nil {
		toSerialize["MinVolumeSize"] = o.MinVolumeSize
	}
	if o.ThroughputUnit != nil {
		toSerialize["ThroughputUnit"] = o.ThroughputUnit
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VolumeSizeUnit != nil {
		toSerialize["VolumeSizeUnit"] = o.VolumeSizeUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MetadataSkuVolumeTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMetadataSkuVolumeTypeAllOf := _MetadataSkuVolumeTypeAllOf{}

	if err = json.Unmarshal(bytes, &varMetadataSkuVolumeTypeAllOf); err == nil {
		*o = MetadataSkuVolumeTypeAllOf(varMetadataSkuVolumeTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IopsUnit")
		delete(additionalProperties, "IsBootable")
		delete(additionalProperties, "IsDefault")
		delete(additionalProperties, "IsProvisionedIops")
		delete(additionalProperties, "MaxIops")
		delete(additionalProperties, "MaxReadIops")
		delete(additionalProperties, "MaxReadThroughput")
		delete(additionalProperties, "MaxThroughput")
		delete(additionalProperties, "MaxVolumeSize")
		delete(additionalProperties, "MaxWriteIops")
		delete(additionalProperties, "MaxWriteThroughput")
		delete(additionalProperties, "MinVolumeSize")
		delete(additionalProperties, "ThroughputUnit")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VolumeSizeUnit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetadataSkuVolumeTypeAllOf struct {
	value *MetadataSkuVolumeTypeAllOf
	isSet bool
}

func (v NullableMetadataSkuVolumeTypeAllOf) Get() *MetadataSkuVolumeTypeAllOf {
	return v.value
}

func (v *NullableMetadataSkuVolumeTypeAllOf) Set(val *MetadataSkuVolumeTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataSkuVolumeTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataSkuVolumeTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataSkuVolumeTypeAllOf(val *MetadataSkuVolumeTypeAllOf) *NullableMetadataSkuVolumeTypeAllOf {
	return &NullableMetadataSkuVolumeTypeAllOf{value: val, isSet: true}
}

func (v NullableMetadataSkuVolumeTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataSkuVolumeTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
