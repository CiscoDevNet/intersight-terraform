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
	"reflect"
	"strings"
)

// WorkflowTargetContext Target details like moid, type and name of the Intersight managed object.
type WorkflowTargetContext struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Moid of the target Intersight managed object.
	TargetMoid *string `json:"TargetMoid,omitempty"`
	// Name of the target instance.
	TargetName *string `json:"TargetName,omitempty"`
	// Object type of the target Intersight managed object.
	TargetType           *string `json:"TargetType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTargetContext WorkflowTargetContext

// NewWorkflowTargetContext instantiates a new WorkflowTargetContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetContext(classId string, objectType string) *WorkflowTargetContext {
	this := WorkflowTargetContext{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTargetContextWithDefaults instantiates a new WorkflowTargetContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetContextWithDefaults() *WorkflowTargetContext {
	this := WorkflowTargetContext{}
	var classId string = "workflow.TargetContext"
	this.ClassId = classId
	var objectType string = "workflow.TargetContext"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTargetContext) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContext) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTargetContext) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTargetContext) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContext) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTargetContext) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *WorkflowTargetContext) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContext) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *WorkflowTargetContext) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *WorkflowTargetContext) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *WorkflowTargetContext) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContext) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *WorkflowTargetContext) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *WorkflowTargetContext) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *WorkflowTargetContext) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContext) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *WorkflowTargetContext) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *WorkflowTargetContext) SetTargetType(v string) {
	o.TargetType = &v
}

func (o WorkflowTargetContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetMoid != nil {
		toSerialize["TargetMoid"] = o.TargetMoid
	}
	if o.TargetName != nil {
		toSerialize["TargetName"] = o.TargetName
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTargetContext) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowTargetContextWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Moid of the target Intersight managed object.
		TargetMoid *string `json:"TargetMoid,omitempty"`
		// Name of the target instance.
		TargetName *string `json:"TargetName,omitempty"`
		// Object type of the target Intersight managed object.
		TargetType *string `json:"TargetType,omitempty"`
	}

	varWorkflowTargetContextWithoutEmbeddedStruct := WorkflowTargetContextWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowTargetContextWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowTargetContext := _WorkflowTargetContext{}
		varWorkflowTargetContext.ClassId = varWorkflowTargetContextWithoutEmbeddedStruct.ClassId
		varWorkflowTargetContext.ObjectType = varWorkflowTargetContextWithoutEmbeddedStruct.ObjectType
		varWorkflowTargetContext.TargetMoid = varWorkflowTargetContextWithoutEmbeddedStruct.TargetMoid
		varWorkflowTargetContext.TargetName = varWorkflowTargetContextWithoutEmbeddedStruct.TargetName
		varWorkflowTargetContext.TargetType = varWorkflowTargetContextWithoutEmbeddedStruct.TargetType
		*o = WorkflowTargetContext(varWorkflowTargetContext)
	} else {
		return err
	}

	varWorkflowTargetContext := _WorkflowTargetContext{}

	err = json.Unmarshal(bytes, &varWorkflowTargetContext)
	if err == nil {
		o.MoBaseComplexType = varWorkflowTargetContext.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetMoid")
		delete(additionalProperties, "TargetName")
		delete(additionalProperties, "TargetType")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableWorkflowTargetContext struct {
	value *WorkflowTargetContext
	isSet bool
}

func (v NullableWorkflowTargetContext) Get() *WorkflowTargetContext {
	return v.value
}

func (v *NullableWorkflowTargetContext) Set(val *WorkflowTargetContext) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetContext) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetContext(val *WorkflowTargetContext) *NullableWorkflowTargetContext {
	return &NullableWorkflowTargetContext{value: val, isSet: true}
}

func (v NullableWorkflowTargetContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
