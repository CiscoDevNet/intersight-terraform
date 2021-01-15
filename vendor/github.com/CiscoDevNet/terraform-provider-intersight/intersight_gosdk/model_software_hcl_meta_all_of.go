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

// SoftwareHclMetaAllOf Definition of the list of properties defined in 'software.HclMeta', excluding properties defined in parent classes.
type SoftwareHclMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The type of content that the Json file holds (Incremental or full dump). * `Full` - Indicates that the JSON File does have full content for HCL metadata. * `Incremental` - Indicates that the JSON File does have only the diff of the Hcl meta to be uploaded.
	ContentType          *string                                `json:"ContentType,omitempty"`
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareHclMetaAllOf SoftwareHclMetaAllOf

// NewSoftwareHclMetaAllOf instantiates a new SoftwareHclMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareHclMetaAllOf(classId string, objectType string) *SoftwareHclMetaAllOf {
	this := SoftwareHclMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var contentType string = "Full"
	this.ContentType = &contentType
	return &this
}

// NewSoftwareHclMetaAllOfWithDefaults instantiates a new SoftwareHclMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareHclMetaAllOfWithDefaults() *SoftwareHclMetaAllOf {
	this := SoftwareHclMetaAllOf{}
	var classId string = "software.HclMeta"
	this.ClassId = classId
	var objectType string = "software.HclMeta"
	this.ObjectType = objectType
	var contentType string = "Full"
	this.ContentType = &contentType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareHclMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareHclMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareHclMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareHclMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareHclMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareHclMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *SoftwareHclMetaAllOf) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareHclMetaAllOf) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *SoftwareHclMetaAllOf) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *SoftwareHclMetaAllOf) SetContentType(v string) {
	o.ContentType = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwareHclMetaAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareHclMetaAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareHclMetaAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareHclMetaAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o SoftwareHclMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ContentType != nil {
		toSerialize["ContentType"] = o.ContentType
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwareHclMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwareHclMetaAllOf := _SoftwareHclMetaAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwareHclMetaAllOf); err == nil {
		*o = SoftwareHclMetaAllOf(varSoftwareHclMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContentType")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwareHclMetaAllOf struct {
	value *SoftwareHclMetaAllOf
	isSet bool
}

func (v NullableSoftwareHclMetaAllOf) Get() *SoftwareHclMetaAllOf {
	return v.value
}

func (v *NullableSoftwareHclMetaAllOf) Set(val *SoftwareHclMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareHclMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareHclMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareHclMetaAllOf(val *SoftwareHclMetaAllOf) *NullableSoftwareHclMetaAllOf {
	return &NullableSoftwareHclMetaAllOf{value: val, isSet: true}
}

func (v NullableSoftwareHclMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareHclMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
