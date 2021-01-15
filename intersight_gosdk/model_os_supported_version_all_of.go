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

// OsSupportedVersionAllOf Definition of the list of properties defined in 'os.SupportedVersion', excluding properties defined in parent classes.
type OsSupportedVersionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The OsInstall Supported Operating System Version Name.
	VersionName          *string                               `json:"VersionName,omitempty"`
	Vendor               *HclOperatingSystemVendorRelationship `json:"Vendor,omitempty"`
	Version              *HclOperatingSystemRelationship       `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsSupportedVersionAllOf OsSupportedVersionAllOf

// NewOsSupportedVersionAllOf instantiates a new OsSupportedVersionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsSupportedVersionAllOf(classId string, objectType string) *OsSupportedVersionAllOf {
	this := OsSupportedVersionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsSupportedVersionAllOfWithDefaults instantiates a new OsSupportedVersionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsSupportedVersionAllOfWithDefaults() *OsSupportedVersionAllOf {
	this := OsSupportedVersionAllOf{}
	var classId string = "os.SupportedVersion"
	this.ClassId = classId
	var objectType string = "os.SupportedVersion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsSupportedVersionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsSupportedVersionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsSupportedVersionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsSupportedVersionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsSupportedVersionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsSupportedVersionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *OsSupportedVersionAllOf) GetVersionName() string {
	if o == nil || o.VersionName == nil {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsSupportedVersionAllOf) GetVersionNameOk() (*string, bool) {
	if o == nil || o.VersionName == nil {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *OsSupportedVersionAllOf) HasVersionName() bool {
	if o != nil && o.VersionName != nil {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *OsSupportedVersionAllOf) SetVersionName(v string) {
	o.VersionName = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *OsSupportedVersionAllOf) GetVendor() HclOperatingSystemVendorRelationship {
	if o == nil || o.Vendor == nil {
		var ret HclOperatingSystemVendorRelationship
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsSupportedVersionAllOf) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *OsSupportedVersionAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given HclOperatingSystemVendorRelationship and assigns it to the Vendor field.
func (o *OsSupportedVersionAllOf) SetVendor(v HclOperatingSystemVendorRelationship) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OsSupportedVersionAllOf) GetVersion() HclOperatingSystemRelationship {
	if o == nil || o.Version == nil {
		var ret HclOperatingSystemRelationship
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsSupportedVersionAllOf) GetVersionOk() (*HclOperatingSystemRelationship, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OsSupportedVersionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given HclOperatingSystemRelationship and assigns it to the Version field.
func (o *OsSupportedVersionAllOf) SetVersion(v HclOperatingSystemRelationship) {
	o.Version = &v
}

func (o OsSupportedVersionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.VersionName != nil {
		toSerialize["VersionName"] = o.VersionName
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsSupportedVersionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsSupportedVersionAllOf := _OsSupportedVersionAllOf{}

	if err = json.Unmarshal(bytes, &varOsSupportedVersionAllOf); err == nil {
		*o = OsSupportedVersionAllOf(varOsSupportedVersionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VersionName")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsSupportedVersionAllOf struct {
	value *OsSupportedVersionAllOf
	isSet bool
}

func (v NullableOsSupportedVersionAllOf) Get() *OsSupportedVersionAllOf {
	return v.value
}

func (v *NullableOsSupportedVersionAllOf) Set(val *OsSupportedVersionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsSupportedVersionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsSupportedVersionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsSupportedVersionAllOf(val *OsSupportedVersionAllOf) *NullableOsSupportedVersionAllOf {
	return &NullableOsSupportedVersionAllOf{value: val, isSet: true}
}

func (v NullableOsSupportedVersionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsSupportedVersionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
