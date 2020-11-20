/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TamApiDataSource Data source using Intersight API to collect data regarding managed devices.
type TamApiDataSource struct {
	TamBaseDataSource
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of Intersight managed object used as data source.
	MoType               *string         `json:"MoType,omitempty"`
	Queries              []TamQueryEntry `json:"Queries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamApiDataSource TamApiDataSource

// NewTamApiDataSource instantiates a new TamApiDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamApiDataSource(classId string, objectType string) *TamApiDataSource {
	this := TamApiDataSource{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamApiDataSourceWithDefaults instantiates a new TamApiDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamApiDataSourceWithDefaults() *TamApiDataSource {
	this := TamApiDataSource{}
	var classId string = "tam.ApiDataSource"
	this.ClassId = classId
	var objectType string = "tam.ApiDataSource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamApiDataSource) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamApiDataSource) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamApiDataSource) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamApiDataSource) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamApiDataSource) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamApiDataSource) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *TamApiDataSource) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamApiDataSource) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *TamApiDataSource) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *TamApiDataSource) SetMoType(v string) {
	o.MoType = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamApiDataSource) GetQueries() []TamQueryEntry {
	if o == nil {
		var ret []TamQueryEntry
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamApiDataSource) GetQueriesOk() (*[]TamQueryEntry, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *TamApiDataSource) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []TamQueryEntry and assigns it to the Queries field.
func (o *TamApiDataSource) SetQueries(v []TamQueryEntry) {
	o.Queries = v
}

func (o TamApiDataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTamBaseDataSource, errTamBaseDataSource := json.Marshal(o.TamBaseDataSource)
	if errTamBaseDataSource != nil {
		return []byte{}, errTamBaseDataSource
	}
	errTamBaseDataSource = json.Unmarshal([]byte(serializedTamBaseDataSource), &toSerialize)
	if errTamBaseDataSource != nil {
		return []byte{}, errTamBaseDataSource
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}
	if o.Queries != nil {
		toSerialize["Queries"] = o.Queries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamApiDataSource) UnmarshalJSON(bytes []byte) (err error) {
	type TamApiDataSourceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of Intersight managed object used as data source.
		MoType  *string         `json:"MoType,omitempty"`
		Queries []TamQueryEntry `json:"Queries,omitempty"`
	}

	varTamApiDataSourceWithoutEmbeddedStruct := TamApiDataSourceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamApiDataSourceWithoutEmbeddedStruct)
	if err == nil {
		varTamApiDataSource := _TamApiDataSource{}
		varTamApiDataSource.ClassId = varTamApiDataSourceWithoutEmbeddedStruct.ClassId
		varTamApiDataSource.ObjectType = varTamApiDataSourceWithoutEmbeddedStruct.ObjectType
		varTamApiDataSource.MoType = varTamApiDataSourceWithoutEmbeddedStruct.MoType
		varTamApiDataSource.Queries = varTamApiDataSourceWithoutEmbeddedStruct.Queries
		*o = TamApiDataSource(varTamApiDataSource)
	} else {
		return err
	}

	varTamApiDataSource := _TamApiDataSource{}

	err = json.Unmarshal(bytes, &varTamApiDataSource)
	if err == nil {
		o.TamBaseDataSource = varTamApiDataSource.TamBaseDataSource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MoType")
		delete(additionalProperties, "Queries")

		// remove fields from embedded structs
		reflectTamBaseDataSource := reflect.ValueOf(o.TamBaseDataSource)
		for i := 0; i < reflectTamBaseDataSource.Type().NumField(); i++ {
			t := reflectTamBaseDataSource.Type().Field(i)

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

type NullableTamApiDataSource struct {
	value *TamApiDataSource
	isSet bool
}

func (v NullableTamApiDataSource) Get() *TamApiDataSource {
	return v.value
}

func (v *NullableTamApiDataSource) Set(val *TamApiDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTamApiDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTamApiDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamApiDataSource(val *TamApiDataSource) *NullableTamApiDataSource {
	return &NullableTamApiDataSource{value: val, isSet: true}
}

func (v NullableTamApiDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamApiDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
