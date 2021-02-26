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
	"time"
)

// MoBaseMo The base abstract class for all Cisco Intersight managed objects.
type MoBaseMo struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The Account ID for this managed object.
	AccountMoid *string `json:"AccountMoid,omitempty"`
	// The time when this managed object was created.
	CreateTime *time.Time `json:"CreateTime,omitempty"`
	// The DomainGroup ID for this managed object.
	DomainGroupMoid *string `json:"DomainGroupMoid,omitempty"`
	// The time when this managed object was last modified.
	ModTime *time.Time `json:"ModTime,omitempty"`
	// The unique identifier of this Managed Object instance.
	Moid   *string  `json:"Moid,omitempty"`
	Owners []string `json:"Owners,omitempty"`
	// Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.
	SharedScope                         *string                                               `json:"SharedScope,omitempty"`
	Tags                                []MoTag                                               `json:"Tags,omitempty"`
	VersionContext                      NullableMoVersionContext                              `json:"VersionContext,omitempty"`
	Var0ClusterReplicationNetworkPolicy *HyperflexClusterReplicationNetworkPolicyRelationship `json:"_0_ClusterReplicationNetworkPolicy,omitempty"`
	// An array of relationships to moBaseMo resources.
	Ancestors []MoBaseMoRelationship `json:"Ancestors,omitempty"`
	Parent    *MoBaseMoRelationship  `json:"Parent,omitempty"`
	// An array of relationships to moBaseMo resources.
	PermissionResources []MoBaseMoRelationship `json:"PermissionResources,omitempty"`
	// A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a 'Name' property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently 'short' and 'hierarchical'. The value is an array and clients should use the first element of the array.
	DisplayNames         map[string][]string `json:"DisplayNames,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MoBaseMo MoBaseMo

// NewMoBaseMo instantiates a new MoBaseMo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoBaseMo(classId string, objectType string) *MoBaseMo {
	this := MoBaseMo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMoBaseMoWithDefaults instantiates a new MoBaseMo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoBaseMoWithDefaults() *MoBaseMo {
	this := MoBaseMo{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *MoBaseMo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MoBaseMo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MoBaseMo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MoBaseMo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountMoid returns the AccountMoid field value if set, zero value otherwise.
func (o *MoBaseMo) GetAccountMoid() string {
	if o == nil || o.AccountMoid == nil {
		var ret string
		return ret
	}
	return *o.AccountMoid
}

// GetAccountMoidOk returns a tuple with the AccountMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetAccountMoidOk() (*string, bool) {
	if o == nil || o.AccountMoid == nil {
		return nil, false
	}
	return o.AccountMoid, true
}

// HasAccountMoid returns a boolean if a field has been set.
func (o *MoBaseMo) HasAccountMoid() bool {
	if o != nil && o.AccountMoid != nil {
		return true
	}

	return false
}

// SetAccountMoid gets a reference to the given string and assigns it to the AccountMoid field.
func (o *MoBaseMo) SetAccountMoid(v string) {
	o.AccountMoid = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *MoBaseMo) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *MoBaseMo) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *MoBaseMo) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetDomainGroupMoid returns the DomainGroupMoid field value if set, zero value otherwise.
func (o *MoBaseMo) GetDomainGroupMoid() string {
	if o == nil || o.DomainGroupMoid == nil {
		var ret string
		return ret
	}
	return *o.DomainGroupMoid
}

// GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetDomainGroupMoidOk() (*string, bool) {
	if o == nil || o.DomainGroupMoid == nil {
		return nil, false
	}
	return o.DomainGroupMoid, true
}

// HasDomainGroupMoid returns a boolean if a field has been set.
func (o *MoBaseMo) HasDomainGroupMoid() bool {
	if o != nil && o.DomainGroupMoid != nil {
		return true
	}

	return false
}

// SetDomainGroupMoid gets a reference to the given string and assigns it to the DomainGroupMoid field.
func (o *MoBaseMo) SetDomainGroupMoid(v string) {
	o.DomainGroupMoid = &v
}

// GetModTime returns the ModTime field value if set, zero value otherwise.
func (o *MoBaseMo) GetModTime() time.Time {
	if o == nil || o.ModTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ModTime
}

// GetModTimeOk returns a tuple with the ModTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetModTimeOk() (*time.Time, bool) {
	if o == nil || o.ModTime == nil {
		return nil, false
	}
	return o.ModTime, true
}

// HasModTime returns a boolean if a field has been set.
func (o *MoBaseMo) HasModTime() bool {
	if o != nil && o.ModTime != nil {
		return true
	}

	return false
}

// SetModTime gets a reference to the given time.Time and assigns it to the ModTime field.
func (o *MoBaseMo) SetModTime(v time.Time) {
	o.ModTime = &v
}

// GetMoid returns the Moid field value if set, zero value otherwise.
func (o *MoBaseMo) GetMoid() string {
	if o == nil || o.Moid == nil {
		var ret string
		return ret
	}
	return *o.Moid
}

// GetMoidOk returns a tuple with the Moid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetMoidOk() (*string, bool) {
	if o == nil || o.Moid == nil {
		return nil, false
	}
	return o.Moid, true
}

// HasMoid returns a boolean if a field has been set.
func (o *MoBaseMo) HasMoid() bool {
	if o != nil && o.Moid != nil {
		return true
	}

	return false
}

// SetMoid gets a reference to the given string and assigns it to the Moid field.
func (o *MoBaseMo) SetMoid(v string) {
	o.Moid = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoBaseMo) GetOwners() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoBaseMo) GetOwnersOk() (*[]string, bool) {
	if o == nil || o.Owners == nil {
		return nil, false
	}
	return &o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *MoBaseMo) HasOwners() bool {
	if o != nil && o.Owners != nil {
		return true
	}

	return false
}

// SetOwners gets a reference to the given []string and assigns it to the Owners field.
func (o *MoBaseMo) SetOwners(v []string) {
	o.Owners = v
}

// GetSharedScope returns the SharedScope field value if set, zero value otherwise.
func (o *MoBaseMo) GetSharedScope() string {
	if o == nil || o.SharedScope == nil {
		var ret string
		return ret
	}
	return *o.SharedScope
}

// GetSharedScopeOk returns a tuple with the SharedScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetSharedScopeOk() (*string, bool) {
	if o == nil || o.SharedScope == nil {
		return nil, false
	}
	return o.SharedScope, true
}

// HasSharedScope returns a boolean if a field has been set.
func (o *MoBaseMo) HasSharedScope() bool {
	if o != nil && o.SharedScope != nil {
		return true
	}

	return false
}

// SetSharedScope gets a reference to the given string and assigns it to the SharedScope field.
func (o *MoBaseMo) SetSharedScope(v string) {
	o.SharedScope = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoBaseMo) GetTags() []MoTag {
	if o == nil {
		var ret []MoTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoBaseMo) GetTagsOk() (*[]MoTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MoBaseMo) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []MoTag and assigns it to the Tags field.
func (o *MoBaseMo) SetTags(v []MoTag) {
	o.Tags = v
}

// GetVersionContext returns the VersionContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoBaseMo) GetVersionContext() MoVersionContext {
	if o == nil || o.VersionContext.Get() == nil {
		var ret MoVersionContext
		return ret
	}
	return *o.VersionContext.Get()
}

// GetVersionContextOk returns a tuple with the VersionContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoBaseMo) GetVersionContextOk() (*MoVersionContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.VersionContext.Get(), o.VersionContext.IsSet()
}

// HasVersionContext returns a boolean if a field has been set.
func (o *MoBaseMo) HasVersionContext() bool {
	if o != nil && o.VersionContext.IsSet() {
		return true
	}

	return false
}

// SetVersionContext gets a reference to the given NullableMoVersionContext and assigns it to the VersionContext field.
func (o *MoBaseMo) SetVersionContext(v MoVersionContext) {
	o.VersionContext.Set(&v)
}

// SetVersionContextNil sets the value for VersionContext to be an explicit nil
func (o *MoBaseMo) SetVersionContextNil() {
	o.VersionContext.Set(nil)
}

// UnsetVersionContext ensures that no value is present for VersionContext, not even an explicit nil
func (o *MoBaseMo) UnsetVersionContext() {
	o.VersionContext.Unset()
}

// GetVar0ClusterReplicationNetworkPolicy returns the Var0ClusterReplicationNetworkPolicy field value if set, zero value otherwise.
func (o *MoBaseMo) GetVar0ClusterReplicationNetworkPolicy() HyperflexClusterReplicationNetworkPolicyRelationship {
	if o == nil || o.Var0ClusterReplicationNetworkPolicy == nil {
		var ret HyperflexClusterReplicationNetworkPolicyRelationship
		return ret
	}
	return *o.Var0ClusterReplicationNetworkPolicy
}

// GetVar0ClusterReplicationNetworkPolicyOk returns a tuple with the Var0ClusterReplicationNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetVar0ClusterReplicationNetworkPolicyOk() (*HyperflexClusterReplicationNetworkPolicyRelationship, bool) {
	if o == nil || o.Var0ClusterReplicationNetworkPolicy == nil {
		return nil, false
	}
	return o.Var0ClusterReplicationNetworkPolicy, true
}

// HasVar0ClusterReplicationNetworkPolicy returns a boolean if a field has been set.
func (o *MoBaseMo) HasVar0ClusterReplicationNetworkPolicy() bool {
	if o != nil && o.Var0ClusterReplicationNetworkPolicy != nil {
		return true
	}

	return false
}

// SetVar0ClusterReplicationNetworkPolicy gets a reference to the given HyperflexClusterReplicationNetworkPolicyRelationship and assigns it to the Var0ClusterReplicationNetworkPolicy field.
func (o *MoBaseMo) SetVar0ClusterReplicationNetworkPolicy(v HyperflexClusterReplicationNetworkPolicyRelationship) {
	o.Var0ClusterReplicationNetworkPolicy = &v
}

// GetAncestors returns the Ancestors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoBaseMo) GetAncestors() []MoBaseMoRelationship {
	if o == nil {
		var ret []MoBaseMoRelationship
		return ret
	}
	return o.Ancestors
}

// GetAncestorsOk returns a tuple with the Ancestors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoBaseMo) GetAncestorsOk() (*[]MoBaseMoRelationship, bool) {
	if o == nil || o.Ancestors == nil {
		return nil, false
	}
	return &o.Ancestors, true
}

// HasAncestors returns a boolean if a field has been set.
func (o *MoBaseMo) HasAncestors() bool {
	if o != nil && o.Ancestors != nil {
		return true
	}

	return false
}

// SetAncestors gets a reference to the given []MoBaseMoRelationship and assigns it to the Ancestors field.
func (o *MoBaseMo) SetAncestors(v []MoBaseMoRelationship) {
	o.Ancestors = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *MoBaseMo) GetParent() MoBaseMoRelationship {
	if o == nil || o.Parent == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoBaseMo) GetParentOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *MoBaseMo) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given MoBaseMoRelationship and assigns it to the Parent field.
func (o *MoBaseMo) SetParent(v MoBaseMoRelationship) {
	o.Parent = &v
}

// GetPermissionResources returns the PermissionResources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoBaseMo) GetPermissionResources() []MoBaseMoRelationship {
	if o == nil {
		var ret []MoBaseMoRelationship
		return ret
	}
	return o.PermissionResources
}

// GetPermissionResourcesOk returns a tuple with the PermissionResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoBaseMo) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool) {
	if o == nil || o.PermissionResources == nil {
		return nil, false
	}
	return &o.PermissionResources, true
}

// HasPermissionResources returns a boolean if a field has been set.
func (o *MoBaseMo) HasPermissionResources() bool {
	if o != nil && o.PermissionResources != nil {
		return true
	}

	return false
}

// SetPermissionResources gets a reference to the given []MoBaseMoRelationship and assigns it to the PermissionResources field.
func (o *MoBaseMo) SetPermissionResources(v []MoBaseMoRelationship) {
	o.PermissionResources = v
}

// GetDisplayNames returns the DisplayNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoBaseMo) GetDisplayNames() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.DisplayNames
}

// GetDisplayNamesOk returns a tuple with the DisplayNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoBaseMo) GetDisplayNamesOk() (*map[string][]string, bool) {
	if o == nil || o.DisplayNames == nil {
		return nil, false
	}
	return &o.DisplayNames, true
}

// HasDisplayNames returns a boolean if a field has been set.
func (o *MoBaseMo) HasDisplayNames() bool {
	if o != nil && o.DisplayNames != nil {
		return true
	}

	return false
}

// SetDisplayNames gets a reference to the given map[string][]string and assigns it to the DisplayNames field.
func (o *MoBaseMo) SetDisplayNames(v map[string][]string) {
	o.DisplayNames = v
}

func (o MoBaseMo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccountMoid != nil {
		toSerialize["AccountMoid"] = o.AccountMoid
	}
	if o.CreateTime != nil {
		toSerialize["CreateTime"] = o.CreateTime
	}
	if o.DomainGroupMoid != nil {
		toSerialize["DomainGroupMoid"] = o.DomainGroupMoid
	}
	if o.ModTime != nil {
		toSerialize["ModTime"] = o.ModTime
	}
	if o.Moid != nil {
		toSerialize["Moid"] = o.Moid
	}
	if o.Owners != nil {
		toSerialize["Owners"] = o.Owners
	}
	if o.SharedScope != nil {
		toSerialize["SharedScope"] = o.SharedScope
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.VersionContext.IsSet() {
		toSerialize["VersionContext"] = o.VersionContext.Get()
	}
	if o.Var0ClusterReplicationNetworkPolicy != nil {
		toSerialize["_0_ClusterReplicationNetworkPolicy"] = o.Var0ClusterReplicationNetworkPolicy
	}
	if o.Ancestors != nil {
		toSerialize["Ancestors"] = o.Ancestors
	}
	if o.Parent != nil {
		toSerialize["Parent"] = o.Parent
	}
	if o.PermissionResources != nil {
		toSerialize["PermissionResources"] = o.PermissionResources
	}
	if o.DisplayNames != nil {
		toSerialize["DisplayNames"] = o.DisplayNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MoBaseMo) UnmarshalJSON(bytes []byte) (err error) {
	varMoBaseMo := _MoBaseMo{}

	if err = json.Unmarshal(bytes, &varMoBaseMo); err == nil {
		*o = MoBaseMo(varMoBaseMo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountMoid")
		delete(additionalProperties, "CreateTime")
		delete(additionalProperties, "DomainGroupMoid")
		delete(additionalProperties, "ModTime")
		delete(additionalProperties, "Moid")
		delete(additionalProperties, "Owners")
		delete(additionalProperties, "SharedScope")
		delete(additionalProperties, "Tags")
		delete(additionalProperties, "VersionContext")
		delete(additionalProperties, "_0_ClusterReplicationNetworkPolicy")
		delete(additionalProperties, "Ancestors")
		delete(additionalProperties, "Parent")
		delete(additionalProperties, "PermissionResources")
		delete(additionalProperties, "DisplayNames")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMoBaseMo struct {
	value *MoBaseMo
	isSet bool
}

func (v NullableMoBaseMo) Get() *MoBaseMo {
	return v.value
}

func (v *NullableMoBaseMo) Set(val *MoBaseMo) {
	v.value = val
	v.isSet = true
}

func (v NullableMoBaseMo) IsSet() bool {
	return v.isSet
}

func (v *NullableMoBaseMo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoBaseMo(val *MoBaseMo) *NullableMoBaseMo {
	return &NullableMoBaseMo{value: val, isSet: true}
}

func (v NullableMoBaseMo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoBaseMo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
