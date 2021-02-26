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
	"reflect"
	"strings"
)

// IamPermissionToRoles Complex type representing a permission and its associated roles.
type IamPermissionToRoles struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string           `json:"ObjectType"`
	Permission           NullableCmrfCmRf `json:"Permission,omitempty"`
	Roles                []CmrfCmRf       `json:"Roles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPermissionToRoles IamPermissionToRoles

// NewIamPermissionToRoles instantiates a new IamPermissionToRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPermissionToRoles(classId string, objectType string) *IamPermissionToRoles {
	this := IamPermissionToRoles{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamPermissionToRolesWithDefaults instantiates a new IamPermissionToRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPermissionToRolesWithDefaults() *IamPermissionToRoles {
	this := IamPermissionToRoles{}
	var classId string = "iam.PermissionToRoles"
	this.ClassId = classId
	var objectType string = "iam.PermissionToRoles"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamPermissionToRoles) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamPermissionToRoles) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamPermissionToRoles) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamPermissionToRoles) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamPermissionToRoles) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamPermissionToRoles) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPermission returns the Permission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermissionToRoles) GetPermission() CmrfCmRf {
	if o == nil || o.Permission.Get() == nil {
		var ret CmrfCmRf
		return ret
	}
	return *o.Permission.Get()
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermissionToRoles) GetPermissionOk() (*CmrfCmRf, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permission.Get(), o.Permission.IsSet()
}

// HasPermission returns a boolean if a field has been set.
func (o *IamPermissionToRoles) HasPermission() bool {
	if o != nil && o.Permission.IsSet() {
		return true
	}

	return false
}

// SetPermission gets a reference to the given NullableCmrfCmRf and assigns it to the Permission field.
func (o *IamPermissionToRoles) SetPermission(v CmrfCmRf) {
	o.Permission.Set(&v)
}

// SetPermissionNil sets the value for Permission to be an explicit nil
func (o *IamPermissionToRoles) SetPermissionNil() {
	o.Permission.Set(nil)
}

// UnsetPermission ensures that no value is present for Permission, not even an explicit nil
func (o *IamPermissionToRoles) UnsetPermission() {
	o.Permission.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermissionToRoles) GetRoles() []CmrfCmRf {
	if o == nil {
		var ret []CmrfCmRf
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermissionToRoles) GetRolesOk() (*[]CmrfCmRf, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamPermissionToRoles) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []CmrfCmRf and assigns it to the Roles field.
func (o *IamPermissionToRoles) SetRoles(v []CmrfCmRf) {
	o.Roles = v
}

func (o IamPermissionToRoles) MarshalJSON() ([]byte, error) {
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
	if o.Permission.IsSet() {
		toSerialize["Permission"] = o.Permission.Get()
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPermissionToRoles) UnmarshalJSON(bytes []byte) (err error) {
	type IamPermissionToRolesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string           `json:"ObjectType"`
		Permission NullableCmrfCmRf `json:"Permission,omitempty"`
		Roles      []CmrfCmRf       `json:"Roles,omitempty"`
	}

	varIamPermissionToRolesWithoutEmbeddedStruct := IamPermissionToRolesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamPermissionToRolesWithoutEmbeddedStruct)
	if err == nil {
		varIamPermissionToRoles := _IamPermissionToRoles{}
		varIamPermissionToRoles.ClassId = varIamPermissionToRolesWithoutEmbeddedStruct.ClassId
		varIamPermissionToRoles.ObjectType = varIamPermissionToRolesWithoutEmbeddedStruct.ObjectType
		varIamPermissionToRoles.Permission = varIamPermissionToRolesWithoutEmbeddedStruct.Permission
		varIamPermissionToRoles.Roles = varIamPermissionToRolesWithoutEmbeddedStruct.Roles
		*o = IamPermissionToRoles(varIamPermissionToRoles)
	} else {
		return err
	}

	varIamPermissionToRoles := _IamPermissionToRoles{}

	err = json.Unmarshal(bytes, &varIamPermissionToRoles)
	if err == nil {
		o.MoBaseComplexType = varIamPermissionToRoles.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "Roles")

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

type NullableIamPermissionToRoles struct {
	value *IamPermissionToRoles
	isSet bool
}

func (v NullableIamPermissionToRoles) Get() *IamPermissionToRoles {
	return v.value
}

func (v *NullableIamPermissionToRoles) Set(val *IamPermissionToRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPermissionToRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPermissionToRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPermissionToRoles(val *IamPermissionToRoles) *NullableIamPermissionToRoles {
	return &NullableIamPermissionToRoles{value: val, isSet: true}
}

func (v NullableIamPermissionToRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPermissionToRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
