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
)

// SoftwarerepositoryCategorySupportConstraintAllOf Definition of the list of properties defined in 'softwarerepository.CategorySupportConstraint', excluding properties defined in parent classes.
type SoftwarerepositoryCategorySupportConstraintAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Identifier for this managed object.
	ConstraintId   *string                              `json:"ConstraintId,omitempty"`
	FilteredModels []SoftwarerepositoryConstraintModels `json:"FilteredModels,omitempty"`
	// Cisco software repository image category identifier.
	MdfId *string `json:"MdfId,omitempty"`
	// Minimum image version from where the models can be supported.
	MinVersion *string `json:"MinVersion,omitempty"`
	// Fields which tells if the constraint is based on image name parsing.
	ParseFromImageName   *bool    `json:"ParseFromImageName,omitempty"`
	SupportedModels      []string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCategorySupportConstraintAllOf SoftwarerepositoryCategorySupportConstraintAllOf

// NewSoftwarerepositoryCategorySupportConstraintAllOf instantiates a new SoftwarerepositoryCategorySupportConstraintAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCategorySupportConstraintAllOf(classId string, objectType string) *SoftwarerepositoryCategorySupportConstraintAllOf {
	this := SoftwarerepositoryCategorySupportConstraintAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryCategorySupportConstraintAllOfWithDefaults instantiates a new SoftwarerepositoryCategorySupportConstraintAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCategorySupportConstraintAllOfWithDefaults() *SoftwarerepositoryCategorySupportConstraintAllOf {
	this := SoftwarerepositoryCategorySupportConstraintAllOf{}
	var classId string = "softwarerepository.CategorySupportConstraint"
	this.ClassId = classId
	var objectType string = "softwarerepository.CategorySupportConstraint"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConstraintId returns the ConstraintId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetConstraintId() string {
	if o == nil || o.ConstraintId == nil {
		var ret string
		return ret
	}
	return *o.ConstraintId
}

// GetConstraintIdOk returns a tuple with the ConstraintId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetConstraintIdOk() (*string, bool) {
	if o == nil || o.ConstraintId == nil {
		return nil, false
	}
	return o.ConstraintId, true
}

// HasConstraintId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) HasConstraintId() bool {
	if o != nil && o.ConstraintId != nil {
		return true
	}

	return false
}

// SetConstraintId gets a reference to the given string and assigns it to the ConstraintId field.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) SetConstraintId(v string) {
	o.ConstraintId = &v
}

// GetFilteredModels returns the FilteredModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetFilteredModels() []SoftwarerepositoryConstraintModels {
	if o == nil {
		var ret []SoftwarerepositoryConstraintModels
		return ret
	}
	return o.FilteredModels
}

// GetFilteredModelsOk returns a tuple with the FilteredModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetFilteredModelsOk() (*[]SoftwarerepositoryConstraintModels, bool) {
	if o == nil || o.FilteredModels == nil {
		return nil, false
	}
	return &o.FilteredModels, true
}

// HasFilteredModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) HasFilteredModels() bool {
	if o != nil && o.FilteredModels != nil {
		return true
	}

	return false
}

// SetFilteredModels gets a reference to the given []SoftwarerepositoryConstraintModels and assigns it to the FilteredModels field.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) SetFilteredModels(v []SoftwarerepositoryConstraintModels) {
	o.FilteredModels = v
}

// GetMdfId returns the MdfId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetMdfId() string {
	if o == nil || o.MdfId == nil {
		var ret string
		return ret
	}
	return *o.MdfId
}

// GetMdfIdOk returns a tuple with the MdfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetMdfIdOk() (*string, bool) {
	if o == nil || o.MdfId == nil {
		return nil, false
	}
	return o.MdfId, true
}

// HasMdfId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) HasMdfId() bool {
	if o != nil && o.MdfId != nil {
		return true
	}

	return false
}

// SetMdfId gets a reference to the given string and assigns it to the MdfId field.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) SetMdfId(v string) {
	o.MdfId = &v
}

// GetMinVersion returns the MinVersion field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetMinVersion() string {
	if o == nil || o.MinVersion == nil {
		var ret string
		return ret
	}
	return *o.MinVersion
}

// GetMinVersionOk returns a tuple with the MinVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetMinVersionOk() (*string, bool) {
	if o == nil || o.MinVersion == nil {
		return nil, false
	}
	return o.MinVersion, true
}

// HasMinVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) HasMinVersion() bool {
	if o != nil && o.MinVersion != nil {
		return true
	}

	return false
}

// SetMinVersion gets a reference to the given string and assigns it to the MinVersion field.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) SetMinVersion(v string) {
	o.MinVersion = &v
}

// GetParseFromImageName returns the ParseFromImageName field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetParseFromImageName() bool {
	if o == nil || o.ParseFromImageName == nil {
		var ret bool
		return ret
	}
	return *o.ParseFromImageName
}

// GetParseFromImageNameOk returns a tuple with the ParseFromImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetParseFromImageNameOk() (*bool, bool) {
	if o == nil || o.ParseFromImageName == nil {
		return nil, false
	}
	return o.ParseFromImageName, true
}

// HasParseFromImageName returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) HasParseFromImageName() bool {
	if o != nil && o.ParseFromImageName != nil {
		return true
	}

	return false
}

// SetParseFromImageName gets a reference to the given bool and assigns it to the ParseFromImageName field.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) SetParseFromImageName(v bool) {
	o.ParseFromImageName = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return &o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryCategorySupportConstraintAllOf) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o SoftwarerepositoryCategorySupportConstraintAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConstraintId != nil {
		toSerialize["ConstraintId"] = o.ConstraintId
	}
	if o.FilteredModels != nil {
		toSerialize["FilteredModels"] = o.FilteredModels
	}
	if o.MdfId != nil {
		toSerialize["MdfId"] = o.MdfId
	}
	if o.MinVersion != nil {
		toSerialize["MinVersion"] = o.MinVersion
	}
	if o.ParseFromImageName != nil {
		toSerialize["ParseFromImageName"] = o.ParseFromImageName
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryCategorySupportConstraintAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryCategorySupportConstraintAllOf := _SoftwarerepositoryCategorySupportConstraintAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryCategorySupportConstraintAllOf); err == nil {
		*o = SoftwarerepositoryCategorySupportConstraintAllOf(varSoftwarerepositoryCategorySupportConstraintAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConstraintId")
		delete(additionalProperties, "FilteredModels")
		delete(additionalProperties, "MdfId")
		delete(additionalProperties, "MinVersion")
		delete(additionalProperties, "ParseFromImageName")
		delete(additionalProperties, "SupportedModels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryCategorySupportConstraintAllOf struct {
	value *SoftwarerepositoryCategorySupportConstraintAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryCategorySupportConstraintAllOf) Get() *SoftwarerepositoryCategorySupportConstraintAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryCategorySupportConstraintAllOf) Set(val *SoftwarerepositoryCategorySupportConstraintAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCategorySupportConstraintAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCategorySupportConstraintAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCategorySupportConstraintAllOf(val *SoftwarerepositoryCategorySupportConstraintAllOf) *NullableSoftwarerepositoryCategorySupportConstraintAllOf {
	return &NullableSoftwarerepositoryCategorySupportConstraintAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCategorySupportConstraintAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCategorySupportConstraintAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
