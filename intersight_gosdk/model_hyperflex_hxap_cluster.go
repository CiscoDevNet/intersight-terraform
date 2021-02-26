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

// HyperflexHxapCluster A HyperFlex Application Platform compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.
type HyperflexHxapCluster struct {
	VirtualizationBaseCluster
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Datacenter to which the cluster belongs.
	DatacenterName *string `json:"DatacenterName,omitempty"`
	// Reason of the failure when cluster is in failed state.
	FailureReason *string `json:"FailureReason,omitempty"`
	// Management IP Address of the cluster.
	ManagementIpAddress *string `json:"ManagementIpAddress,omitempty"`
	// Product version of HyperFlex compute cluster.
	Version              *string                              `json:"Version,omitempty"`
	HxCluster            *HyperflexClusterRelationship        `json:"HxCluster,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapCluster HyperflexHxapCluster

// NewHyperflexHxapCluster instantiates a new HyperflexHxapCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapCluster(classId string, objectType string) *HyperflexHxapCluster {
	this := HyperflexHxapCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxapClusterWithDefaults instantiates a new HyperflexHxapCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapClusterWithDefaults() *HyperflexHxapCluster {
	this := HyperflexHxapCluster{}
	var classId string = "hyperflex.HxapCluster"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapCluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDatacenterName returns the DatacenterName field value if set, zero value otherwise.
func (o *HyperflexHxapCluster) GetDatacenterName() string {
	if o == nil || o.DatacenterName == nil {
		var ret string
		return ret
	}
	return *o.DatacenterName
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapCluster) GetDatacenterNameOk() (*string, bool) {
	if o == nil || o.DatacenterName == nil {
		return nil, false
	}
	return o.DatacenterName, true
}

// HasDatacenterName returns a boolean if a field has been set.
func (o *HyperflexHxapCluster) HasDatacenterName() bool {
	if o != nil && o.DatacenterName != nil {
		return true
	}

	return false
}

// SetDatacenterName gets a reference to the given string and assigns it to the DatacenterName field.
func (o *HyperflexHxapCluster) SetDatacenterName(v string) {
	o.DatacenterName = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *HyperflexHxapCluster) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapCluster) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *HyperflexHxapCluster) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *HyperflexHxapCluster) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetManagementIpAddress returns the ManagementIpAddress field value if set, zero value otherwise.
func (o *HyperflexHxapCluster) GetManagementIpAddress() string {
	if o == nil || o.ManagementIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementIpAddress
}

// GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapCluster) GetManagementIpAddressOk() (*string, bool) {
	if o == nil || o.ManagementIpAddress == nil {
		return nil, false
	}
	return o.ManagementIpAddress, true
}

// HasManagementIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapCluster) HasManagementIpAddress() bool {
	if o != nil && o.ManagementIpAddress != nil {
		return true
	}

	return false
}

// SetManagementIpAddress gets a reference to the given string and assigns it to the ManagementIpAddress field.
func (o *HyperflexHxapCluster) SetManagementIpAddress(v string) {
	o.ManagementIpAddress = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexHxapCluster) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapCluster) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexHxapCluster) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexHxapCluster) SetVersion(v string) {
	o.Version = &v
}

// GetHxCluster returns the HxCluster field value if set, zero value otherwise.
func (o *HyperflexHxapCluster) GetHxCluster() HyperflexClusterRelationship {
	if o == nil || o.HxCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.HxCluster
}

// GetHxClusterOk returns a tuple with the HxCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapCluster) GetHxClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.HxCluster == nil {
		return nil, false
	}
	return o.HxCluster, true
}

// HasHxCluster returns a boolean if a field has been set.
func (o *HyperflexHxapCluster) HasHxCluster() bool {
	if o != nil && o.HxCluster != nil {
		return true
	}

	return false
}

// SetHxCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the HxCluster field.
func (o *HyperflexHxapCluster) SetHxCluster(v HyperflexClusterRelationship) {
	o.HxCluster = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexHxapCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexHxapCluster) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexHxapCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o HyperflexHxapCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseCluster, errVirtualizationBaseCluster := json.Marshal(o.VirtualizationBaseCluster)
	if errVirtualizationBaseCluster != nil {
		return []byte{}, errVirtualizationBaseCluster
	}
	errVirtualizationBaseCluster = json.Unmarshal([]byte(serializedVirtualizationBaseCluster), &toSerialize)
	if errVirtualizationBaseCluster != nil {
		return []byte{}, errVirtualizationBaseCluster
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DatacenterName != nil {
		toSerialize["DatacenterName"] = o.DatacenterName
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.ManagementIpAddress != nil {
		toSerialize["ManagementIpAddress"] = o.ManagementIpAddress
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.HxCluster != nil {
		toSerialize["HxCluster"] = o.HxCluster
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapCluster) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxapClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Datacenter to which the cluster belongs.
		DatacenterName *string `json:"DatacenterName,omitempty"`
		// Reason of the failure when cluster is in failed state.
		FailureReason *string `json:"FailureReason,omitempty"`
		// Management IP Address of the cluster.
		ManagementIpAddress *string `json:"ManagementIpAddress,omitempty"`
		// Product version of HyperFlex compute cluster.
		Version          *string                              `json:"Version,omitempty"`
		HxCluster        *HyperflexClusterRelationship        `json:"HxCluster,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varHyperflexHxapClusterWithoutEmbeddedStruct := HyperflexHxapClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxapClusterWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxapCluster := _HyperflexHxapCluster{}
		varHyperflexHxapCluster.ClassId = varHyperflexHxapClusterWithoutEmbeddedStruct.ClassId
		varHyperflexHxapCluster.ObjectType = varHyperflexHxapClusterWithoutEmbeddedStruct.ObjectType
		varHyperflexHxapCluster.DatacenterName = varHyperflexHxapClusterWithoutEmbeddedStruct.DatacenterName
		varHyperflexHxapCluster.FailureReason = varHyperflexHxapClusterWithoutEmbeddedStruct.FailureReason
		varHyperflexHxapCluster.ManagementIpAddress = varHyperflexHxapClusterWithoutEmbeddedStruct.ManagementIpAddress
		varHyperflexHxapCluster.Version = varHyperflexHxapClusterWithoutEmbeddedStruct.Version
		varHyperflexHxapCluster.HxCluster = varHyperflexHxapClusterWithoutEmbeddedStruct.HxCluster
		varHyperflexHxapCluster.RegisteredDevice = varHyperflexHxapClusterWithoutEmbeddedStruct.RegisteredDevice
		*o = HyperflexHxapCluster(varHyperflexHxapCluster)
	} else {
		return err
	}

	varHyperflexHxapCluster := _HyperflexHxapCluster{}

	err = json.Unmarshal(bytes, &varHyperflexHxapCluster)
	if err == nil {
		o.VirtualizationBaseCluster = varHyperflexHxapCluster.VirtualizationBaseCluster
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DatacenterName")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "ManagementIpAddress")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "HxCluster")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectVirtualizationBaseCluster := reflect.ValueOf(o.VirtualizationBaseCluster)
		for i := 0; i < reflectVirtualizationBaseCluster.Type().NumField(); i++ {
			t := reflectVirtualizationBaseCluster.Type().Field(i)

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

type NullableHyperflexHxapCluster struct {
	value *HyperflexHxapCluster
	isSet bool
}

func (v NullableHyperflexHxapCluster) Get() *HyperflexHxapCluster {
	return v.value
}

func (v *NullableHyperflexHxapCluster) Set(val *HyperflexHxapCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapCluster(val *HyperflexHxapCluster) *NullableHyperflexHxapCluster {
	return &NullableHyperflexHxapCluster{value: val, isSet: true}
}

func (v NullableHyperflexHxapCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
