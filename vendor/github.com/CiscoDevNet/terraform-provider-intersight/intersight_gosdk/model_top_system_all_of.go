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

// TopSystemAllOf Definition of the list of properties defined in 'top.System', excluding properties defined in parent classes.
type TopSystemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The IPv4 address of system.
	Ipv4Address *string `json:"Ipv4Address,omitempty"`
	// The IPv6 address of system.
	Ipv6Address *string `json:"Ipv6Address,omitempty"`
	// The current mode of the system.
	Mode *string `json:"Mode,omitempty"`
	// The admin configured name of the system.
	Name *string `json:"Name,omitempty"`
	// The operational timezone of the system, empty indicates no timezone has been set specifically.
	TimeZone *string `json:"TimeZone,omitempty"`
	// An array of relationships to computeBlade resources.
	ComputeBlades []ComputeBladeRelationship `json:"ComputeBlades,omitempty"`
	// An array of relationships to computeRackUnit resources.
	ComputeRackUnits     []ComputeRackUnitRelationship     `json:"ComputeRackUnits,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship  `json:"InventoryDeviceInfo,omitempty"`
	ManagementController *ManagementControllerRelationship `json:"ManagementController,omitempty"`
	// An array of relationships to networkElement resources.
	NetworkElements      []NetworkElementRelationship         `json:"NetworkElements,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TopSystemAllOf TopSystemAllOf

// NewTopSystemAllOf instantiates a new TopSystemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopSystemAllOf(classId string, objectType string) *TopSystemAllOf {
	this := TopSystemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTopSystemAllOfWithDefaults instantiates a new TopSystemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopSystemAllOfWithDefaults() *TopSystemAllOf {
	this := TopSystemAllOf{}
	var classId string = "top.System"
	this.ClassId = classId
	var objectType string = "top.System"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TopSystemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TopSystemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TopSystemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TopSystemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *TopSystemAllOf) GetIpv4Address() string {
	if o == nil || o.Ipv4Address == nil {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetIpv4AddressOk() (*string, bool) {
	if o == nil || o.Ipv4Address == nil {
		return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasIpv4Address() bool {
	if o != nil && o.Ipv4Address != nil {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *TopSystemAllOf) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *TopSystemAllOf) GetIpv6Address() string {
	if o == nil || o.Ipv6Address == nil {
		var ret string
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetIpv6AddressOk() (*string, bool) {
	if o == nil || o.Ipv6Address == nil {
		return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasIpv6Address() bool {
	if o != nil && o.Ipv6Address != nil {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given string and assigns it to the Ipv6Address field.
func (o *TopSystemAllOf) SetIpv6Address(v string) {
	o.Ipv6Address = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *TopSystemAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *TopSystemAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TopSystemAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TopSystemAllOf) SetName(v string) {
	o.Name = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *TopSystemAllOf) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *TopSystemAllOf) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetComputeBlades returns the ComputeBlades field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopSystemAllOf) GetComputeBlades() []ComputeBladeRelationship {
	if o == nil {
		var ret []ComputeBladeRelationship
		return ret
	}
	return o.ComputeBlades
}

// GetComputeBladesOk returns a tuple with the ComputeBlades field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopSystemAllOf) GetComputeBladesOk() (*[]ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlades == nil {
		return nil, false
	}
	return &o.ComputeBlades, true
}

// HasComputeBlades returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasComputeBlades() bool {
	if o != nil && o.ComputeBlades != nil {
		return true
	}

	return false
}

// SetComputeBlades gets a reference to the given []ComputeBladeRelationship and assigns it to the ComputeBlades field.
func (o *TopSystemAllOf) SetComputeBlades(v []ComputeBladeRelationship) {
	o.ComputeBlades = v
}

// GetComputeRackUnits returns the ComputeRackUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopSystemAllOf) GetComputeRackUnits() []ComputeRackUnitRelationship {
	if o == nil {
		var ret []ComputeRackUnitRelationship
		return ret
	}
	return o.ComputeRackUnits
}

// GetComputeRackUnitsOk returns a tuple with the ComputeRackUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopSystemAllOf) GetComputeRackUnitsOk() (*[]ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnits == nil {
		return nil, false
	}
	return &o.ComputeRackUnits, true
}

// HasComputeRackUnits returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasComputeRackUnits() bool {
	if o != nil && o.ComputeRackUnits != nil {
		return true
	}

	return false
}

// SetComputeRackUnits gets a reference to the given []ComputeRackUnitRelationship and assigns it to the ComputeRackUnits field.
func (o *TopSystemAllOf) SetComputeRackUnits(v []ComputeRackUnitRelationship) {
	o.ComputeRackUnits = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *TopSystemAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *TopSystemAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetManagementController returns the ManagementController field value if set, zero value otherwise.
func (o *TopSystemAllOf) GetManagementController() ManagementControllerRelationship {
	if o == nil || o.ManagementController == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.ManagementController
}

// GetManagementControllerOk returns a tuple with the ManagementController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetManagementControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.ManagementController == nil {
		return nil, false
	}
	return o.ManagementController, true
}

// HasManagementController returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasManagementController() bool {
	if o != nil && o.ManagementController != nil {
		return true
	}

	return false
}

// SetManagementController gets a reference to the given ManagementControllerRelationship and assigns it to the ManagementController field.
func (o *TopSystemAllOf) SetManagementController(v ManagementControllerRelationship) {
	o.ManagementController = &v
}

// GetNetworkElements returns the NetworkElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopSystemAllOf) GetNetworkElements() []NetworkElementRelationship {
	if o == nil {
		var ret []NetworkElementRelationship
		return ret
	}
	return o.NetworkElements
}

// GetNetworkElementsOk returns a tuple with the NetworkElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopSystemAllOf) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElements == nil {
		return nil, false
	}
	return &o.NetworkElements, true
}

// HasNetworkElements returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasNetworkElements() bool {
	if o != nil && o.NetworkElements != nil {
		return true
	}

	return false
}

// SetNetworkElements gets a reference to the given []NetworkElementRelationship and assigns it to the NetworkElements field.
func (o *TopSystemAllOf) SetNetworkElements(v []NetworkElementRelationship) {
	o.NetworkElements = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *TopSystemAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSystemAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *TopSystemAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *TopSystemAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o TopSystemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Ipv4Address != nil {
		toSerialize["Ipv4Address"] = o.Ipv4Address
	}
	if o.Ipv6Address != nil {
		toSerialize["Ipv6Address"] = o.Ipv6Address
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.TimeZone != nil {
		toSerialize["TimeZone"] = o.TimeZone
	}
	if o.ComputeBlades != nil {
		toSerialize["ComputeBlades"] = o.ComputeBlades
	}
	if o.ComputeRackUnits != nil {
		toSerialize["ComputeRackUnits"] = o.ComputeRackUnits
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.ManagementController != nil {
		toSerialize["ManagementController"] = o.ManagementController
	}
	if o.NetworkElements != nil {
		toSerialize["NetworkElements"] = o.NetworkElements
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TopSystemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTopSystemAllOf := _TopSystemAllOf{}

	if err = json.Unmarshal(bytes, &varTopSystemAllOf); err == nil {
		*o = TopSystemAllOf(varTopSystemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ipv4Address")
		delete(additionalProperties, "Ipv6Address")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "TimeZone")
		delete(additionalProperties, "ComputeBlades")
		delete(additionalProperties, "ComputeRackUnits")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "ManagementController")
		delete(additionalProperties, "NetworkElements")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTopSystemAllOf struct {
	value *TopSystemAllOf
	isSet bool
}

func (v NullableTopSystemAllOf) Get() *TopSystemAllOf {
	return v.value
}

func (v *NullableTopSystemAllOf) Set(val *TopSystemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTopSystemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTopSystemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopSystemAllOf(val *TopSystemAllOf) *NullableTopSystemAllOf {
	return &NullableTopSystemAllOf{value: val, isSet: true}
}

func (v NullableTopSystemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopSystemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
