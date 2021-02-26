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

// HyperflexAppSettingConstraint A constraint that can be used to qualify an application setting.
type HyperflexAppSettingConstraint struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The supported HyperFlex Data Platform version in regex format.
	HxdpVersion *string `json:"HxdpVersion,omitempty"`
	// The hypervisor type for the HyperFlex cluster. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The supported management platform for the HyperFlex Cluster. * `FI` - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect. * `EDGE` - The host servers used in the cluster deployment are standalone severs.
	MgmtPlatform *string `json:"MgmtPlatform,omitempty"`
	// The supported server models in regex format.
	ServerModel          *string `json:"ServerModel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAppSettingConstraint HyperflexAppSettingConstraint

// NewHyperflexAppSettingConstraint instantiates a new HyperflexAppSettingConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAppSettingConstraint(classId string, objectType string) *HyperflexAppSettingConstraint {
	this := HyperflexAppSettingConstraint{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var mgmtPlatform string = "FI"
	this.MgmtPlatform = &mgmtPlatform
	return &this
}

// NewHyperflexAppSettingConstraintWithDefaults instantiates a new HyperflexAppSettingConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAppSettingConstraintWithDefaults() *HyperflexAppSettingConstraint {
	this := HyperflexAppSettingConstraint{}
	var classId string = "hyperflex.AppSettingConstraint"
	this.ClassId = classId
	var objectType string = "hyperflex.AppSettingConstraint"
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var mgmtPlatform string = "FI"
	this.MgmtPlatform = &mgmtPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAppSettingConstraint) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAppSettingConstraint) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAppSettingConstraint) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAppSettingConstraint) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHxdpVersion returns the HxdpVersion field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraint) GetHxdpVersion() string {
	if o == nil || o.HxdpVersion == nil {
		var ret string
		return ret
	}
	return *o.HxdpVersion
}

// GetHxdpVersionOk returns a tuple with the HxdpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetHxdpVersionOk() (*string, bool) {
	if o == nil || o.HxdpVersion == nil {
		return nil, false
	}
	return o.HxdpVersion, true
}

// HasHxdpVersion returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraint) HasHxdpVersion() bool {
	if o != nil && o.HxdpVersion != nil {
		return true
	}

	return false
}

// SetHxdpVersion gets a reference to the given string and assigns it to the HxdpVersion field.
func (o *HyperflexAppSettingConstraint) SetHxdpVersion(v string) {
	o.HxdpVersion = &v
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraint) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraint) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *HyperflexAppSettingConstraint) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetMgmtPlatform returns the MgmtPlatform field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraint) GetMgmtPlatform() string {
	if o == nil || o.MgmtPlatform == nil {
		var ret string
		return ret
	}
	return *o.MgmtPlatform
}

// GetMgmtPlatformOk returns a tuple with the MgmtPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetMgmtPlatformOk() (*string, bool) {
	if o == nil || o.MgmtPlatform == nil {
		return nil, false
	}
	return o.MgmtPlatform, true
}

// HasMgmtPlatform returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraint) HasMgmtPlatform() bool {
	if o != nil && o.MgmtPlatform != nil {
		return true
	}

	return false
}

// SetMgmtPlatform gets a reference to the given string and assigns it to the MgmtPlatform field.
func (o *HyperflexAppSettingConstraint) SetMgmtPlatform(v string) {
	o.MgmtPlatform = &v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HyperflexAppSettingConstraint) GetServerModel() string {
	if o == nil || o.ServerModel == nil {
		var ret string
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppSettingConstraint) GetServerModelOk() (*string, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HyperflexAppSettingConstraint) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given string and assigns it to the ServerModel field.
func (o *HyperflexAppSettingConstraint) SetServerModel(v string) {
	o.ServerModel = &v
}

func (o HyperflexAppSettingConstraint) MarshalJSON() ([]byte, error) {
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
	if o.HxdpVersion != nil {
		toSerialize["HxdpVersion"] = o.HxdpVersion
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.MgmtPlatform != nil {
		toSerialize["MgmtPlatform"] = o.MgmtPlatform
	}
	if o.ServerModel != nil {
		toSerialize["ServerModel"] = o.ServerModel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAppSettingConstraint) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexAppSettingConstraintWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The supported HyperFlex Data Platform version in regex format.
		HxdpVersion *string `json:"HxdpVersion,omitempty"`
		// The hypervisor type for the HyperFlex cluster. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
		HypervisorType *string `json:"HypervisorType,omitempty"`
		// The supported management platform for the HyperFlex Cluster. * `FI` - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect. * `EDGE` - The host servers used in the cluster deployment are standalone severs.
		MgmtPlatform *string `json:"MgmtPlatform,omitempty"`
		// The supported server models in regex format.
		ServerModel *string `json:"ServerModel,omitempty"`
	}

	varHyperflexAppSettingConstraintWithoutEmbeddedStruct := HyperflexAppSettingConstraintWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexAppSettingConstraintWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexAppSettingConstraint := _HyperflexAppSettingConstraint{}
		varHyperflexAppSettingConstraint.ClassId = varHyperflexAppSettingConstraintWithoutEmbeddedStruct.ClassId
		varHyperflexAppSettingConstraint.ObjectType = varHyperflexAppSettingConstraintWithoutEmbeddedStruct.ObjectType
		varHyperflexAppSettingConstraint.HxdpVersion = varHyperflexAppSettingConstraintWithoutEmbeddedStruct.HxdpVersion
		varHyperflexAppSettingConstraint.HypervisorType = varHyperflexAppSettingConstraintWithoutEmbeddedStruct.HypervisorType
		varHyperflexAppSettingConstraint.MgmtPlatform = varHyperflexAppSettingConstraintWithoutEmbeddedStruct.MgmtPlatform
		varHyperflexAppSettingConstraint.ServerModel = varHyperflexAppSettingConstraintWithoutEmbeddedStruct.ServerModel
		*o = HyperflexAppSettingConstraint(varHyperflexAppSettingConstraint)
	} else {
		return err
	}

	varHyperflexAppSettingConstraint := _HyperflexAppSettingConstraint{}

	err = json.Unmarshal(bytes, &varHyperflexAppSettingConstraint)
	if err == nil {
		o.MoBaseComplexType = varHyperflexAppSettingConstraint.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HxdpVersion")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "MgmtPlatform")
		delete(additionalProperties, "ServerModel")

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

type NullableHyperflexAppSettingConstraint struct {
	value *HyperflexAppSettingConstraint
	isSet bool
}

func (v NullableHyperflexAppSettingConstraint) Get() *HyperflexAppSettingConstraint {
	return v.value
}

func (v *NullableHyperflexAppSettingConstraint) Set(val *HyperflexAppSettingConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAppSettingConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAppSettingConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAppSettingConstraint(val *HyperflexAppSettingConstraint) *NullableHyperflexAppSettingConstraint {
	return &NullableHyperflexAppSettingConstraint{value: val, isSet: true}
}

func (v NullableHyperflexAppSettingConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAppSettingConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
