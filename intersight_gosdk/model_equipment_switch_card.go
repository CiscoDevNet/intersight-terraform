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

// EquipmentSwitchCard Fixed / Removable module on a Fabric Interconnect / Switch.
type EquipmentSwitchCard struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Detailed description of this switch hardware.
	Description *string `json:"Description,omitempty"`
	// The user configured Ethernet switching mode for this switch (End-Host or Switch). * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
	EthernetSwitchingMode *string `json:"EthernetSwitchingMode,omitempty"`
	// The user configured FC switching mode for this switch (End-Host or Switch). * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
	FcSwitchingMode *string `json:"FcSwitchingMode,omitempty"`
	// Number of ports present in this switch hardware.
	NumPorts *int64 `json:"NumPorts,omitempty"`
	// Field specifies this Switch's Out-of-band IP address.
	OutOfBandIpAddress *string `json:"OutOfBandIpAddress,omitempty"`
	// Field specifies this Switch's default gateway for the out-of-band management interface.
	OutOfBandIpGateway *string `json:"OutOfBandIpGateway,omitempty"`
	// Presence for this switch hardware.
	Presence *string `json:"Presence,omitempty"`
	// Slot identifier of the local Switch slot Interface.
	SlotId *int64 `json:"SlotId,omitempty"`
	// Operational state of the switch hardware.
	State *string `json:"State,omitempty"`
	// Switch Identifier that is local to a cluster.
	SwitchId *string `json:"SwitchId,omitempty"`
	// The Thermal status of the fabric interconnect. * `unknown` - The default state of the sensor (in case no data is received). * `ok` - State of the sensor indicating the sensor's temperature range is okay. * `upper-non-recoverable` - State of the sensor indicating that the temperature is extremely high above normal range. * `upper-critical` - State of the sensor indicating that the temperature is above normal range. * `upper-non-critical` - State of the sensor indicating that the temperature is a little above the normal range. * `lower-non-critical` - State of the sensor indicating that the temperature is a little below the normal range. * `lower-critical` - State of the sensor indicating that the temperature is below normal range. * `lower-non-recoverable` - State of the sensor indicating that the temperature is extremely below normal range.
	Thermal *string `json:"Thermal,omitempty"`
	// An array of relationships to fcPortChannel resources.
	FcPortChannels      []FcPortChannelRelationship      `json:"FcPortChannels,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement      *NetworkElementRelationship      `json:"NetworkElement,omitempty"`
	// An array of relationships to etherPortChannel resources.
	PortChannels []EtherPortChannelRelationship `json:"PortChannels,omitempty"`
	// An array of relationships to portGroup resources.
	PortGroups           []PortGroupRelationship              `json:"PortGroups,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentSwitchCard EquipmentSwitchCard

// NewEquipmentSwitchCard instantiates a new EquipmentSwitchCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSwitchCard(classId string, objectType string) *EquipmentSwitchCard {
	this := EquipmentSwitchCard{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ethernetSwitchingMode string = "end-host"
	this.EthernetSwitchingMode = &ethernetSwitchingMode
	var fcSwitchingMode string = "end-host"
	this.FcSwitchingMode = &fcSwitchingMode
	var thermal string = "unknown"
	this.Thermal = &thermal
	return &this
}

// NewEquipmentSwitchCardWithDefaults instantiates a new EquipmentSwitchCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSwitchCardWithDefaults() *EquipmentSwitchCard {
	this := EquipmentSwitchCard{}
	var classId string = "equipment.SwitchCard"
	this.ClassId = classId
	var objectType string = "equipment.SwitchCard"
	this.ObjectType = objectType
	var ethernetSwitchingMode string = "end-host"
	this.EthernetSwitchingMode = &ethernetSwitchingMode
	var fcSwitchingMode string = "end-host"
	this.FcSwitchingMode = &fcSwitchingMode
	var thermal string = "unknown"
	this.Thermal = &thermal
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentSwitchCard) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentSwitchCard) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentSwitchCard) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentSwitchCard) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentSwitchCard) SetDescription(v string) {
	o.Description = &v
}

// GetEthernetSwitchingMode returns the EthernetSwitchingMode field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetEthernetSwitchingMode() string {
	if o == nil || o.EthernetSwitchingMode == nil {
		var ret string
		return ret
	}
	return *o.EthernetSwitchingMode
}

// GetEthernetSwitchingModeOk returns a tuple with the EthernetSwitchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetEthernetSwitchingModeOk() (*string, bool) {
	if o == nil || o.EthernetSwitchingMode == nil {
		return nil, false
	}
	return o.EthernetSwitchingMode, true
}

// HasEthernetSwitchingMode returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasEthernetSwitchingMode() bool {
	if o != nil && o.EthernetSwitchingMode != nil {
		return true
	}

	return false
}

// SetEthernetSwitchingMode gets a reference to the given string and assigns it to the EthernetSwitchingMode field.
func (o *EquipmentSwitchCard) SetEthernetSwitchingMode(v string) {
	o.EthernetSwitchingMode = &v
}

// GetFcSwitchingMode returns the FcSwitchingMode field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetFcSwitchingMode() string {
	if o == nil || o.FcSwitchingMode == nil {
		var ret string
		return ret
	}
	return *o.FcSwitchingMode
}

// GetFcSwitchingModeOk returns a tuple with the FcSwitchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetFcSwitchingModeOk() (*string, bool) {
	if o == nil || o.FcSwitchingMode == nil {
		return nil, false
	}
	return o.FcSwitchingMode, true
}

// HasFcSwitchingMode returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasFcSwitchingMode() bool {
	if o != nil && o.FcSwitchingMode != nil {
		return true
	}

	return false
}

// SetFcSwitchingMode gets a reference to the given string and assigns it to the FcSwitchingMode field.
func (o *EquipmentSwitchCard) SetFcSwitchingMode(v string) {
	o.FcSwitchingMode = &v
}

// GetNumPorts returns the NumPorts field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetNumPorts() int64 {
	if o == nil || o.NumPorts == nil {
		var ret int64
		return ret
	}
	return *o.NumPorts
}

// GetNumPortsOk returns a tuple with the NumPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetNumPortsOk() (*int64, bool) {
	if o == nil || o.NumPorts == nil {
		return nil, false
	}
	return o.NumPorts, true
}

// HasNumPorts returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasNumPorts() bool {
	if o != nil && o.NumPorts != nil {
		return true
	}

	return false
}

// SetNumPorts gets a reference to the given int64 and assigns it to the NumPorts field.
func (o *EquipmentSwitchCard) SetNumPorts(v int64) {
	o.NumPorts = &v
}

// GetOutOfBandIpAddress returns the OutOfBandIpAddress field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetOutOfBandIpAddress() string {
	if o == nil || o.OutOfBandIpAddress == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpAddress
}

// GetOutOfBandIpAddressOk returns a tuple with the OutOfBandIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetOutOfBandIpAddressOk() (*string, bool) {
	if o == nil || o.OutOfBandIpAddress == nil {
		return nil, false
	}
	return o.OutOfBandIpAddress, true
}

// HasOutOfBandIpAddress returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasOutOfBandIpAddress() bool {
	if o != nil && o.OutOfBandIpAddress != nil {
		return true
	}

	return false
}

// SetOutOfBandIpAddress gets a reference to the given string and assigns it to the OutOfBandIpAddress field.
func (o *EquipmentSwitchCard) SetOutOfBandIpAddress(v string) {
	o.OutOfBandIpAddress = &v
}

// GetOutOfBandIpGateway returns the OutOfBandIpGateway field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetOutOfBandIpGateway() string {
	if o == nil || o.OutOfBandIpGateway == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpGateway
}

// GetOutOfBandIpGatewayOk returns a tuple with the OutOfBandIpGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetOutOfBandIpGatewayOk() (*string, bool) {
	if o == nil || o.OutOfBandIpGateway == nil {
		return nil, false
	}
	return o.OutOfBandIpGateway, true
}

// HasOutOfBandIpGateway returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasOutOfBandIpGateway() bool {
	if o != nil && o.OutOfBandIpGateway != nil {
		return true
	}

	return false
}

// SetOutOfBandIpGateway gets a reference to the given string and assigns it to the OutOfBandIpGateway field.
func (o *EquipmentSwitchCard) SetOutOfBandIpGateway(v string) {
	o.OutOfBandIpGateway = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *EquipmentSwitchCard) SetPresence(v string) {
	o.Presence = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *EquipmentSwitchCard) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EquipmentSwitchCard) SetState(v string) {
	o.State = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *EquipmentSwitchCard) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetThermal returns the Thermal field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetThermal() string {
	if o == nil || o.Thermal == nil {
		var ret string
		return ret
	}
	return *o.Thermal
}

// GetThermalOk returns a tuple with the Thermal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetThermalOk() (*string, bool) {
	if o == nil || o.Thermal == nil {
		return nil, false
	}
	return o.Thermal, true
}

// HasThermal returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasThermal() bool {
	if o != nil && o.Thermal != nil {
		return true
	}

	return false
}

// SetThermal gets a reference to the given string and assigns it to the Thermal field.
func (o *EquipmentSwitchCard) SetThermal(v string) {
	o.Thermal = &v
}

// GetFcPortChannels returns the FcPortChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSwitchCard) GetFcPortChannels() []FcPortChannelRelationship {
	if o == nil {
		var ret []FcPortChannelRelationship
		return ret
	}
	return o.FcPortChannels
}

// GetFcPortChannelsOk returns a tuple with the FcPortChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSwitchCard) GetFcPortChannelsOk() (*[]FcPortChannelRelationship, bool) {
	if o == nil || o.FcPortChannels == nil {
		return nil, false
	}
	return &o.FcPortChannels, true
}

// HasFcPortChannels returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasFcPortChannels() bool {
	if o != nil && o.FcPortChannels != nil {
		return true
	}

	return false
}

// SetFcPortChannels gets a reference to the given []FcPortChannelRelationship and assigns it to the FcPortChannels field.
func (o *EquipmentSwitchCard) SetFcPortChannels(v []FcPortChannelRelationship) {
	o.FcPortChannels = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentSwitchCard) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EquipmentSwitchCard) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetPortChannels returns the PortChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSwitchCard) GetPortChannels() []EtherPortChannelRelationship {
	if o == nil {
		var ret []EtherPortChannelRelationship
		return ret
	}
	return o.PortChannels
}

// GetPortChannelsOk returns a tuple with the PortChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSwitchCard) GetPortChannelsOk() (*[]EtherPortChannelRelationship, bool) {
	if o == nil || o.PortChannels == nil {
		return nil, false
	}
	return &o.PortChannels, true
}

// HasPortChannels returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasPortChannels() bool {
	if o != nil && o.PortChannels != nil {
		return true
	}

	return false
}

// SetPortChannels gets a reference to the given []EtherPortChannelRelationship and assigns it to the PortChannels field.
func (o *EquipmentSwitchCard) SetPortChannels(v []EtherPortChannelRelationship) {
	o.PortChannels = v
}

// GetPortGroups returns the PortGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSwitchCard) GetPortGroups() []PortGroupRelationship {
	if o == nil {
		var ret []PortGroupRelationship
		return ret
	}
	return o.PortGroups
}

// GetPortGroupsOk returns a tuple with the PortGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSwitchCard) GetPortGroupsOk() (*[]PortGroupRelationship, bool) {
	if o == nil || o.PortGroups == nil {
		return nil, false
	}
	return &o.PortGroups, true
}

// HasPortGroups returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasPortGroups() bool {
	if o != nil && o.PortGroups != nil {
		return true
	}

	return false
}

// SetPortGroups gets a reference to the given []PortGroupRelationship and assigns it to the PortGroups field.
func (o *EquipmentSwitchCard) SetPortGroups(v []PortGroupRelationship) {
	o.PortGroups = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentSwitchCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentSwitchCard) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentSwitchCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentSwitchCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EthernetSwitchingMode != nil {
		toSerialize["EthernetSwitchingMode"] = o.EthernetSwitchingMode
	}
	if o.FcSwitchingMode != nil {
		toSerialize["FcSwitchingMode"] = o.FcSwitchingMode
	}
	if o.NumPorts != nil {
		toSerialize["NumPorts"] = o.NumPorts
	}
	if o.OutOfBandIpAddress != nil {
		toSerialize["OutOfBandIpAddress"] = o.OutOfBandIpAddress
	}
	if o.OutOfBandIpGateway != nil {
		toSerialize["OutOfBandIpGateway"] = o.OutOfBandIpGateway
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Thermal != nil {
		toSerialize["Thermal"] = o.Thermal
	}
	if o.FcPortChannels != nil {
		toSerialize["FcPortChannels"] = o.FcPortChannels
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.PortChannels != nil {
		toSerialize["PortChannels"] = o.PortChannels
	}
	if o.PortGroups != nil {
		toSerialize["PortGroups"] = o.PortGroups
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentSwitchCard) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentSwitchCardWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Detailed description of this switch hardware.
		Description *string `json:"Description,omitempty"`
		// The user configured Ethernet switching mode for this switch (End-Host or Switch). * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
		EthernetSwitchingMode *string `json:"EthernetSwitchingMode,omitempty"`
		// The user configured FC switching mode for this switch (End-Host or Switch). * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
		FcSwitchingMode *string `json:"FcSwitchingMode,omitempty"`
		// Number of ports present in this switch hardware.
		NumPorts *int64 `json:"NumPorts,omitempty"`
		// Field specifies this Switch's Out-of-band IP address.
		OutOfBandIpAddress *string `json:"OutOfBandIpAddress,omitempty"`
		// Field specifies this Switch's default gateway for the out-of-band management interface.
		OutOfBandIpGateway *string `json:"OutOfBandIpGateway,omitempty"`
		// Presence for this switch hardware.
		Presence *string `json:"Presence,omitempty"`
		// Slot identifier of the local Switch slot Interface.
		SlotId *int64 `json:"SlotId,omitempty"`
		// Operational state of the switch hardware.
		State *string `json:"State,omitempty"`
		// Switch Identifier that is local to a cluster.
		SwitchId *string `json:"SwitchId,omitempty"`
		// The Thermal status of the fabric interconnect. * `unknown` - The default state of the sensor (in case no data is received). * `ok` - State of the sensor indicating the sensor's temperature range is okay. * `upper-non-recoverable` - State of the sensor indicating that the temperature is extremely high above normal range. * `upper-critical` - State of the sensor indicating that the temperature is above normal range. * `upper-non-critical` - State of the sensor indicating that the temperature is a little above the normal range. * `lower-non-critical` - State of the sensor indicating that the temperature is a little below the normal range. * `lower-critical` - State of the sensor indicating that the temperature is below normal range. * `lower-non-recoverable` - State of the sensor indicating that the temperature is extremely below normal range.
		Thermal *string `json:"Thermal,omitempty"`
		// An array of relationships to fcPortChannel resources.
		FcPortChannels      []FcPortChannelRelationship      `json:"FcPortChannels,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		NetworkElement      *NetworkElementRelationship      `json:"NetworkElement,omitempty"`
		// An array of relationships to etherPortChannel resources.
		PortChannels []EtherPortChannelRelationship `json:"PortChannels,omitempty"`
		// An array of relationships to portGroup resources.
		PortGroups       []PortGroupRelationship              `json:"PortGroups,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentSwitchCardWithoutEmbeddedStruct := EquipmentSwitchCardWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentSwitchCardWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentSwitchCard := _EquipmentSwitchCard{}
		varEquipmentSwitchCard.ClassId = varEquipmentSwitchCardWithoutEmbeddedStruct.ClassId
		varEquipmentSwitchCard.ObjectType = varEquipmentSwitchCardWithoutEmbeddedStruct.ObjectType
		varEquipmentSwitchCard.Description = varEquipmentSwitchCardWithoutEmbeddedStruct.Description
		varEquipmentSwitchCard.EthernetSwitchingMode = varEquipmentSwitchCardWithoutEmbeddedStruct.EthernetSwitchingMode
		varEquipmentSwitchCard.FcSwitchingMode = varEquipmentSwitchCardWithoutEmbeddedStruct.FcSwitchingMode
		varEquipmentSwitchCard.NumPorts = varEquipmentSwitchCardWithoutEmbeddedStruct.NumPorts
		varEquipmentSwitchCard.OutOfBandIpAddress = varEquipmentSwitchCardWithoutEmbeddedStruct.OutOfBandIpAddress
		varEquipmentSwitchCard.OutOfBandIpGateway = varEquipmentSwitchCardWithoutEmbeddedStruct.OutOfBandIpGateway
		varEquipmentSwitchCard.Presence = varEquipmentSwitchCardWithoutEmbeddedStruct.Presence
		varEquipmentSwitchCard.SlotId = varEquipmentSwitchCardWithoutEmbeddedStruct.SlotId
		varEquipmentSwitchCard.State = varEquipmentSwitchCardWithoutEmbeddedStruct.State
		varEquipmentSwitchCard.SwitchId = varEquipmentSwitchCardWithoutEmbeddedStruct.SwitchId
		varEquipmentSwitchCard.Thermal = varEquipmentSwitchCardWithoutEmbeddedStruct.Thermal
		varEquipmentSwitchCard.FcPortChannels = varEquipmentSwitchCardWithoutEmbeddedStruct.FcPortChannels
		varEquipmentSwitchCard.InventoryDeviceInfo = varEquipmentSwitchCardWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentSwitchCard.NetworkElement = varEquipmentSwitchCardWithoutEmbeddedStruct.NetworkElement
		varEquipmentSwitchCard.PortChannels = varEquipmentSwitchCardWithoutEmbeddedStruct.PortChannels
		varEquipmentSwitchCard.PortGroups = varEquipmentSwitchCardWithoutEmbeddedStruct.PortGroups
		varEquipmentSwitchCard.RegisteredDevice = varEquipmentSwitchCardWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentSwitchCard(varEquipmentSwitchCard)
	} else {
		return err
	}

	varEquipmentSwitchCard := _EquipmentSwitchCard{}

	err = json.Unmarshal(bytes, &varEquipmentSwitchCard)
	if err == nil {
		o.EquipmentBase = varEquipmentSwitchCard.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EthernetSwitchingMode")
		delete(additionalProperties, "FcSwitchingMode")
		delete(additionalProperties, "NumPorts")
		delete(additionalProperties, "OutOfBandIpAddress")
		delete(additionalProperties, "OutOfBandIpGateway")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "State")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Thermal")
		delete(additionalProperties, "FcPortChannels")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "PortChannels")
		delete(additionalProperties, "PortGroups")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableEquipmentSwitchCard struct {
	value *EquipmentSwitchCard
	isSet bool
}

func (v NullableEquipmentSwitchCard) Get() *EquipmentSwitchCard {
	return v.value
}

func (v *NullableEquipmentSwitchCard) Set(val *EquipmentSwitchCard) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSwitchCard) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSwitchCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSwitchCard(val *EquipmentSwitchCard) *NullableEquipmentSwitchCard {
	return &NullableEquipmentSwitchCard{value: val, isSet: true}
}

func (v NullableEquipmentSwitchCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSwitchCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
