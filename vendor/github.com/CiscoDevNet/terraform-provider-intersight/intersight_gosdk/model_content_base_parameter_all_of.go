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

// ContentBaseParameterAllOf Definition of the list of properties defined in 'content.BaseParameter', excluding properties defined in parent classes.
type ContentBaseParameterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The flag that allows single values in content to be extracted as a single element collection in case the parameter is of Collection type. This flag is applicable for parameters of type Collection only.
	AcceptSingleValue *bool `json:"AcceptSingleValue,omitempty"`
	// The name of the complex type definition in case this is a complex parameter. The content.Grammar object must have a complex type, content.ComplexType, defined with the specified name in types collection property.
	ComplexType *string `json:"ComplexType,omitempty"`
	// The type of the collection item in case this is a collection parameter. * `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum. * `string` - The parameter value to be extracted is of the string type. * `integer` - The parameter value to be extracted is of the number type. * `float` - The parameter value to be extracted is of the float number type. * `boolean` - The parameter value to be extracted is of the boolean type. * `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type. * `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters. * `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type.
	ItemType *string `json:"ItemType,omitempty"`
	// The name of the parameter.
	Name *string `json:"Name,omitempty"`
	// The content specific path information that identifies the parameter value within the content. The value is usually a XPath or JSONPath or a regular expression in case of text content.
	Path *string `json:"Path,omitempty"`
	// The type of the parameter. Accepted values are simple, complex, collection. * `simple` - The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum. * `string` - The parameter value to be extracted is of the string type. * `integer` - The parameter value to be extracted is of the number type. * `float` - The parameter value to be extracted is of the float number type. * `boolean` - The parameter value to be extracted is of the boolean type. * `json` - The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type. * `complex` - The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters. * `collection` - The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContentBaseParameterAllOf ContentBaseParameterAllOf

// NewContentBaseParameterAllOf instantiates a new ContentBaseParameterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentBaseParameterAllOf(classId string, objectType string) *ContentBaseParameterAllOf {
	this := ContentBaseParameterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var itemType string = "simple"
	this.ItemType = &itemType
	var type_ string = "simple"
	this.Type = &type_
	return &this
}

// NewContentBaseParameterAllOfWithDefaults instantiates a new ContentBaseParameterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentBaseParameterAllOfWithDefaults() *ContentBaseParameterAllOf {
	this := ContentBaseParameterAllOf{}
	var itemType string = "simple"
	this.ItemType = &itemType
	var type_ string = "simple"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *ContentBaseParameterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ContentBaseParameterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ContentBaseParameterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ContentBaseParameterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ContentBaseParameterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ContentBaseParameterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAcceptSingleValue returns the AcceptSingleValue field value if set, zero value otherwise.
func (o *ContentBaseParameterAllOf) GetAcceptSingleValue() bool {
	if o == nil || o.AcceptSingleValue == nil {
		var ret bool
		return ret
	}
	return *o.AcceptSingleValue
}

// GetAcceptSingleValueOk returns a tuple with the AcceptSingleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentBaseParameterAllOf) GetAcceptSingleValueOk() (*bool, bool) {
	if o == nil || o.AcceptSingleValue == nil {
		return nil, false
	}
	return o.AcceptSingleValue, true
}

// HasAcceptSingleValue returns a boolean if a field has been set.
func (o *ContentBaseParameterAllOf) HasAcceptSingleValue() bool {
	if o != nil && o.AcceptSingleValue != nil {
		return true
	}

	return false
}

// SetAcceptSingleValue gets a reference to the given bool and assigns it to the AcceptSingleValue field.
func (o *ContentBaseParameterAllOf) SetAcceptSingleValue(v bool) {
	o.AcceptSingleValue = &v
}

// GetComplexType returns the ComplexType field value if set, zero value otherwise.
func (o *ContentBaseParameterAllOf) GetComplexType() string {
	if o == nil || o.ComplexType == nil {
		var ret string
		return ret
	}
	return *o.ComplexType
}

// GetComplexTypeOk returns a tuple with the ComplexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentBaseParameterAllOf) GetComplexTypeOk() (*string, bool) {
	if o == nil || o.ComplexType == nil {
		return nil, false
	}
	return o.ComplexType, true
}

// HasComplexType returns a boolean if a field has been set.
func (o *ContentBaseParameterAllOf) HasComplexType() bool {
	if o != nil && o.ComplexType != nil {
		return true
	}

	return false
}

// SetComplexType gets a reference to the given string and assigns it to the ComplexType field.
func (o *ContentBaseParameterAllOf) SetComplexType(v string) {
	o.ComplexType = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *ContentBaseParameterAllOf) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentBaseParameterAllOf) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *ContentBaseParameterAllOf) HasItemType() bool {
	if o != nil && o.ItemType != nil {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *ContentBaseParameterAllOf) SetItemType(v string) {
	o.ItemType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContentBaseParameterAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentBaseParameterAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContentBaseParameterAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContentBaseParameterAllOf) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ContentBaseParameterAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentBaseParameterAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ContentBaseParameterAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ContentBaseParameterAllOf) SetPath(v string) {
	o.Path = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContentBaseParameterAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentBaseParameterAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContentBaseParameterAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContentBaseParameterAllOf) SetType(v string) {
	o.Type = &v
}

func (o ContentBaseParameterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AcceptSingleValue != nil {
		toSerialize["AcceptSingleValue"] = o.AcceptSingleValue
	}
	if o.ComplexType != nil {
		toSerialize["ComplexType"] = o.ComplexType
	}
	if o.ItemType != nil {
		toSerialize["ItemType"] = o.ItemType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContentBaseParameterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varContentBaseParameterAllOf := _ContentBaseParameterAllOf{}

	if err = json.Unmarshal(bytes, &varContentBaseParameterAllOf); err == nil {
		*o = ContentBaseParameterAllOf(varContentBaseParameterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AcceptSingleValue")
		delete(additionalProperties, "ComplexType")
		delete(additionalProperties, "ItemType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContentBaseParameterAllOf struct {
	value *ContentBaseParameterAllOf
	isSet bool
}

func (v NullableContentBaseParameterAllOf) Get() *ContentBaseParameterAllOf {
	return v.value
}

func (v *NullableContentBaseParameterAllOf) Set(val *ContentBaseParameterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContentBaseParameterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContentBaseParameterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentBaseParameterAllOf(val *ContentBaseParameterAllOf) *NullableContentBaseParameterAllOf {
	return &NullableContentBaseParameterAllOf{value: val, isSet: true}
}

func (v NullableContentBaseParameterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentBaseParameterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
