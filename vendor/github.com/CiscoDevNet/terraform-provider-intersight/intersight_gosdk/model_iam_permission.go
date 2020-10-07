/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IamPermission Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.
type IamPermission struct {
	MoBaseMo
	// The informative description about each permission.
	Description *string `json:"Description,omitempty"`
	// The name of the permission which has to be granted to user.
	Name    *string                 `json:"Name,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to iamEndPointRole resources.
	EndPointRoles []IamEndPointRoleRelationship `json:"EndPointRoles,omitempty"`
	// An array of relationships to iamResourceRoles resources.
	ResourceRoles []IamResourceRolesRelationship `json:"ResourceRoles,omitempty"`
	// An array of relationships to iamRole resources.
	Roles         []IamRoleRelationship         `json:"Roles,omitempty"`
	SessionLimits *IamSessionLimitsRelationship `json:"SessionLimits,omitempty"`
	// An array of relationships to iamUserGroup resources.
	UserGroups []IamUserGroupRelationship `json:"UserGroups,omitempty"`
	// An array of relationships to iamUser resources.
	Users                []IamUserRelationship `json:"Users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPermission IamPermission

// NewIamPermission instantiates a new IamPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPermission() *IamPermission {
	this := IamPermission{}
	return &this
}

// NewIamPermissionWithDefaults instantiates a new IamPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPermissionWithDefaults() *IamPermission {
	this := IamPermission{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamPermission) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamPermission) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamPermission) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamPermission) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamPermission) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamPermission) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamPermission) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamPermission) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamPermission) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetEndPointRoles returns the EndPointRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermission) GetEndPointRoles() []IamEndPointRoleRelationship {
	if o == nil {
		var ret []IamEndPointRoleRelationship
		return ret
	}
	return o.EndPointRoles
}

// GetEndPointRolesOk returns a tuple with the EndPointRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermission) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool) {
	if o == nil || o.EndPointRoles == nil {
		return nil, false
	}
	return &o.EndPointRoles, true
}

// HasEndPointRoles returns a boolean if a field has been set.
func (o *IamPermission) HasEndPointRoles() bool {
	if o != nil && o.EndPointRoles != nil {
		return true
	}

	return false
}

// SetEndPointRoles gets a reference to the given []IamEndPointRoleRelationship and assigns it to the EndPointRoles field.
func (o *IamPermission) SetEndPointRoles(v []IamEndPointRoleRelationship) {
	o.EndPointRoles = v
}

// GetResourceRoles returns the ResourceRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermission) GetResourceRoles() []IamResourceRolesRelationship {
	if o == nil {
		var ret []IamResourceRolesRelationship
		return ret
	}
	return o.ResourceRoles
}

// GetResourceRolesOk returns a tuple with the ResourceRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermission) GetResourceRolesOk() (*[]IamResourceRolesRelationship, bool) {
	if o == nil || o.ResourceRoles == nil {
		return nil, false
	}
	return &o.ResourceRoles, true
}

// HasResourceRoles returns a boolean if a field has been set.
func (o *IamPermission) HasResourceRoles() bool {
	if o != nil && o.ResourceRoles != nil {
		return true
	}

	return false
}

// SetResourceRoles gets a reference to the given []IamResourceRolesRelationship and assigns it to the ResourceRoles field.
func (o *IamPermission) SetResourceRoles(v []IamResourceRolesRelationship) {
	o.ResourceRoles = v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermission) GetRoles() []IamRoleRelationship {
	if o == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermission) GetRolesOk() (*[]IamRoleRelationship, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamPermission) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IamRoleRelationship and assigns it to the Roles field.
func (o *IamPermission) SetRoles(v []IamRoleRelationship) {
	o.Roles = v
}

// GetSessionLimits returns the SessionLimits field value if set, zero value otherwise.
func (o *IamPermission) GetSessionLimits() IamSessionLimitsRelationship {
	if o == nil || o.SessionLimits == nil {
		var ret IamSessionLimitsRelationship
		return ret
	}
	return *o.SessionLimits
}

// GetSessionLimitsOk returns a tuple with the SessionLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetSessionLimitsOk() (*IamSessionLimitsRelationship, bool) {
	if o == nil || o.SessionLimits == nil {
		return nil, false
	}
	return o.SessionLimits, true
}

// HasSessionLimits returns a boolean if a field has been set.
func (o *IamPermission) HasSessionLimits() bool {
	if o != nil && o.SessionLimits != nil {
		return true
	}

	return false
}

// SetSessionLimits gets a reference to the given IamSessionLimitsRelationship and assigns it to the SessionLimits field.
func (o *IamPermission) SetSessionLimits(v IamSessionLimitsRelationship) {
	o.SessionLimits = &v
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermission) GetUserGroups() []IamUserGroupRelationship {
	if o == nil {
		var ret []IamUserGroupRelationship
		return ret
	}
	return o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermission) GetUserGroupsOk() (*[]IamUserGroupRelationship, bool) {
	if o == nil || o.UserGroups == nil {
		return nil, false
	}
	return &o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *IamPermission) HasUserGroups() bool {
	if o != nil && o.UserGroups != nil {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given []IamUserGroupRelationship and assigns it to the UserGroups field.
func (o *IamPermission) SetUserGroups(v []IamUserGroupRelationship) {
	o.UserGroups = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPermission) GetUsers() []IamUserRelationship {
	if o == nil {
		var ret []IamUserRelationship
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPermission) GetUsersOk() (*[]IamUserRelationship, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IamPermission) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []IamUserRelationship and assigns it to the Users field.
func (o *IamPermission) SetUsers(v []IamUserRelationship) {
	o.Users = v
}

func (o IamPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.EndPointRoles != nil {
		toSerialize["EndPointRoles"] = o.EndPointRoles
	}
	if o.ResourceRoles != nil {
		toSerialize["ResourceRoles"] = o.ResourceRoles
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}
	if o.SessionLimits != nil {
		toSerialize["SessionLimits"] = o.SessionLimits
	}
	if o.UserGroups != nil {
		toSerialize["UserGroups"] = o.UserGroups
	}
	if o.Users != nil {
		toSerialize["Users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPermission) UnmarshalJSON(bytes []byte) (err error) {
	type IamPermissionWithoutEmbeddedStruct struct {
		// The informative description about each permission.
		Description *string `json:"Description,omitempty"`
		// The name of the permission which has to be granted to user.
		Name    *string                 `json:"Name,omitempty"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
		// An array of relationships to iamEndPointRole resources.
		EndPointRoles []IamEndPointRoleRelationship `json:"EndPointRoles,omitempty"`
		// An array of relationships to iamResourceRoles resources.
		ResourceRoles []IamResourceRolesRelationship `json:"ResourceRoles,omitempty"`
		// An array of relationships to iamRole resources.
		Roles         []IamRoleRelationship         `json:"Roles,omitempty"`
		SessionLimits *IamSessionLimitsRelationship `json:"SessionLimits,omitempty"`
		// An array of relationships to iamUserGroup resources.
		UserGroups []IamUserGroupRelationship `json:"UserGroups,omitempty"`
		// An array of relationships to iamUser resources.
		Users []IamUserRelationship `json:"Users,omitempty"`
	}

	varIamPermissionWithoutEmbeddedStruct := IamPermissionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamPermissionWithoutEmbeddedStruct)
	if err == nil {
		varIamPermission := _IamPermission{}
		varIamPermission.Description = varIamPermissionWithoutEmbeddedStruct.Description
		varIamPermission.Name = varIamPermissionWithoutEmbeddedStruct.Name
		varIamPermission.Account = varIamPermissionWithoutEmbeddedStruct.Account
		varIamPermission.EndPointRoles = varIamPermissionWithoutEmbeddedStruct.EndPointRoles
		varIamPermission.ResourceRoles = varIamPermissionWithoutEmbeddedStruct.ResourceRoles
		varIamPermission.Roles = varIamPermissionWithoutEmbeddedStruct.Roles
		varIamPermission.SessionLimits = varIamPermissionWithoutEmbeddedStruct.SessionLimits
		varIamPermission.UserGroups = varIamPermissionWithoutEmbeddedStruct.UserGroups
		varIamPermission.Users = varIamPermissionWithoutEmbeddedStruct.Users
		*o = IamPermission(varIamPermission)
	} else {
		return err
	}

	varIamPermission := _IamPermission{}

	err = json.Unmarshal(bytes, &varIamPermission)
	if err == nil {
		o.MoBaseMo = varIamPermission.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "EndPointRoles")
		delete(additionalProperties, "ResourceRoles")
		delete(additionalProperties, "Roles")
		delete(additionalProperties, "SessionLimits")
		delete(additionalProperties, "UserGroups")
		delete(additionalProperties, "Users")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIamPermission struct {
	value *IamPermission
	isSet bool
}

func (v NullableIamPermission) Get() *IamPermission {
	return v.value
}

func (v *NullableIamPermission) Set(val *IamPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPermission(val *IamPermission) *NullableIamPermission {
	return &NullableIamPermission{value: val, isSet: true}
}

func (v NullableIamPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}