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

// AssetTargetAllOf Definition of the list of properties defined in 'asset.Target', excluding properties defined in parent classes.
type AssetTargetAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name or email id of the user who claimed the target.
	ClaimedByUserName *string           `json:"ClaimedByUserName,omitempty"`
	Connections       []AssetConnection `json:"Connections,omitempty"`
	// The Device Connector version for target types which are managed by via embedded Device Connector.
	ConnectorVersion *string `json:"ConnectorVersion,omitempty"`
	// ExternalIpAddress is applicable for targets which are managed via an Intersight Device Connector. The value is the IP Address of the target as seen from Intersight. It is either the IP Address of the managed target's interface which has a route to the internet or a NAT IP Addresss when the target is deployed in a private network.
	ExternalIpAddress *string  `json:"ExternalIpAddress,omitempty"`
	IpAddress         []string `json:"IpAddress,omitempty"`
	// A user provided name for the managed target.
	Name      *string  `json:"Name,omitempty"`
	ProductId []string `json:"ProductId,omitempty"`
	// For targets which are managed by an embedded Intersight Device Connector, this field indicates that an administrator of the target has disabled management operations of the Device Connector and only monitoring is permitted.
	ReadOnly *bool          `json:"ReadOnly,omitempty"`
	Services []AssetService `json:"Services,omitempty"`
	// Status indicates if Intersight can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target. * `` - The target details have been persisted but Intersight has not yet attempted to connect to the target. * `Connected` - Intersight is able to establish a connection to the target and initiate management activities. * `NotConnected` - Intersight is unable to establish a connection to the target. * `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established. * `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Custom Target details in Intersight. * `Claimed` - Custom Target is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target credentials are incorrect.
	Status *string `json:"Status,omitempty"`
	// StatusErrorReason provides additional context for the Status.
	StatusErrorReason *string  `json:"StatusErrorReason,omitempty"`
	TargetId          []string `json:"TargetId,omitempty"`
	// The type of the managed target. For example a UCS Server or VMware Vcenter target. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NetAppOntap` - A NetApp ONTAP storage system. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A Dynatrace controller that monitors applications. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its REST API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `CiscoCatalyst` - A Cisco Catalyst networking switch device.
	TargetType           *string                              `json:"TargetType,omitempty"`
	Account              *IamAccountRelationship              `json:"Account,omitempty"`
	Assist               *AssetTargetRelationship             `json:"Assist,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetTargetAllOf AssetTargetAllOf

// NewAssetTargetAllOf instantiates a new AssetTargetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTargetAllOf(classId string, objectType string) *AssetTargetAllOf {
	this := AssetTargetAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// NewAssetTargetAllOfWithDefaults instantiates a new AssetTargetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTargetAllOfWithDefaults() *AssetTargetAllOf {
	this := AssetTargetAllOf{}
	var classId string = "asset.Target"
	this.ClassId = classId
	var objectType string = "asset.Target"
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetTargetAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetTargetAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetTargetAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetTargetAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClaimedByUserName returns the ClaimedByUserName field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetClaimedByUserName() string {
	if o == nil || o.ClaimedByUserName == nil {
		var ret string
		return ret
	}
	return *o.ClaimedByUserName
}

// GetClaimedByUserNameOk returns a tuple with the ClaimedByUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetClaimedByUserNameOk() (*string, bool) {
	if o == nil || o.ClaimedByUserName == nil {
		return nil, false
	}
	return o.ClaimedByUserName, true
}

// HasClaimedByUserName returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasClaimedByUserName() bool {
	if o != nil && o.ClaimedByUserName != nil {
		return true
	}

	return false
}

// SetClaimedByUserName gets a reference to the given string and assigns it to the ClaimedByUserName field.
func (o *AssetTargetAllOf) SetClaimedByUserName(v string) {
	o.ClaimedByUserName = &v
}

// GetConnections returns the Connections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTargetAllOf) GetConnections() []AssetConnection {
	if o == nil {
		var ret []AssetConnection
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTargetAllOf) GetConnectionsOk() (*[]AssetConnection, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return &o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []AssetConnection and assigns it to the Connections field.
func (o *AssetTargetAllOf) SetConnections(v []AssetConnection) {
	o.Connections = v
}

// GetConnectorVersion returns the ConnectorVersion field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetConnectorVersion() string {
	if o == nil || o.ConnectorVersion == nil {
		var ret string
		return ret
	}
	return *o.ConnectorVersion
}

// GetConnectorVersionOk returns a tuple with the ConnectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetConnectorVersionOk() (*string, bool) {
	if o == nil || o.ConnectorVersion == nil {
		return nil, false
	}
	return o.ConnectorVersion, true
}

// HasConnectorVersion returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasConnectorVersion() bool {
	if o != nil && o.ConnectorVersion != nil {
		return true
	}

	return false
}

// SetConnectorVersion gets a reference to the given string and assigns it to the ConnectorVersion field.
func (o *AssetTargetAllOf) SetConnectorVersion(v string) {
	o.ConnectorVersion = &v
}

// GetExternalIpAddress returns the ExternalIpAddress field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetExternalIpAddress() string {
	if o == nil || o.ExternalIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ExternalIpAddress
}

// GetExternalIpAddressOk returns a tuple with the ExternalIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetExternalIpAddressOk() (*string, bool) {
	if o == nil || o.ExternalIpAddress == nil {
		return nil, false
	}
	return o.ExternalIpAddress, true
}

// HasExternalIpAddress returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasExternalIpAddress() bool {
	if o != nil && o.ExternalIpAddress != nil {
		return true
	}

	return false
}

// SetExternalIpAddress gets a reference to the given string and assigns it to the ExternalIpAddress field.
func (o *AssetTargetAllOf) SetExternalIpAddress(v string) {
	o.ExternalIpAddress = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTargetAllOf) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTargetAllOf) GetIpAddressOk() (*[]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *AssetTargetAllOf) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetTargetAllOf) SetName(v string) {
	o.Name = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTargetAllOf) GetProductId() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTargetAllOf) GetProductIdOk() (*[]string, bool) {
	if o == nil || o.ProductId == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasProductId() bool {
	if o != nil && o.ProductId != nil {
		return true
	}

	return false
}

// SetProductId gets a reference to the given []string and assigns it to the ProductId field.
func (o *AssetTargetAllOf) SetProductId(v []string) {
	o.ProductId = v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *AssetTargetAllOf) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTargetAllOf) GetServices() []AssetService {
	if o == nil {
		var ret []AssetService
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTargetAllOf) GetServicesOk() (*[]AssetService, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []AssetService and assigns it to the Services field.
func (o *AssetTargetAllOf) SetServices(v []AssetService) {
	o.Services = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetTargetAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusErrorReason returns the StatusErrorReason field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetStatusErrorReason() string {
	if o == nil || o.StatusErrorReason == nil {
		var ret string
		return ret
	}
	return *o.StatusErrorReason
}

// GetStatusErrorReasonOk returns a tuple with the StatusErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetStatusErrorReasonOk() (*string, bool) {
	if o == nil || o.StatusErrorReason == nil {
		return nil, false
	}
	return o.StatusErrorReason, true
}

// HasStatusErrorReason returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasStatusErrorReason() bool {
	if o != nil && o.StatusErrorReason != nil {
		return true
	}

	return false
}

// SetStatusErrorReason gets a reference to the given string and assigns it to the StatusErrorReason field.
func (o *AssetTargetAllOf) SetStatusErrorReason(v string) {
	o.StatusErrorReason = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTargetAllOf) GetTargetId() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTargetAllOf) GetTargetIdOk() (*[]string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given []string and assigns it to the TargetId field.
func (o *AssetTargetAllOf) SetTargetId(v []string) {
	o.TargetId = v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AssetTargetAllOf) SetTargetType(v string) {
	o.TargetType = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AssetTargetAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetAssist returns the Assist field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetAssist() AssetTargetRelationship {
	if o == nil || o.Assist == nil {
		var ret AssetTargetRelationship
		return ret
	}
	return *o.Assist
}

// GetAssistOk returns a tuple with the Assist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetAssistOk() (*AssetTargetRelationship, bool) {
	if o == nil || o.Assist == nil {
		return nil, false
	}
	return o.Assist, true
}

// HasAssist returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasAssist() bool {
	if o != nil && o.Assist != nil {
		return true
	}

	return false
}

// SetAssist gets a reference to the given AssetTargetRelationship and assigns it to the Assist field.
func (o *AssetTargetAllOf) SetAssist(v AssetTargetRelationship) {
	o.Assist = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AssetTargetAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AssetTargetAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AssetTargetAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AssetTargetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClaimedByUserName != nil {
		toSerialize["ClaimedByUserName"] = o.ClaimedByUserName
	}
	if o.Connections != nil {
		toSerialize["Connections"] = o.Connections
	}
	if o.ConnectorVersion != nil {
		toSerialize["ConnectorVersion"] = o.ConnectorVersion
	}
	if o.ExternalIpAddress != nil {
		toSerialize["ExternalIpAddress"] = o.ExternalIpAddress
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ProductId != nil {
		toSerialize["ProductId"] = o.ProductId
	}
	if o.ReadOnly != nil {
		toSerialize["ReadOnly"] = o.ReadOnly
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusErrorReason != nil {
		toSerialize["StatusErrorReason"] = o.StatusErrorReason
	}
	if o.TargetId != nil {
		toSerialize["TargetId"] = o.TargetId
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Assist != nil {
		toSerialize["Assist"] = o.Assist
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTargetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetTargetAllOf := _AssetTargetAllOf{}

	if err = json.Unmarshal(bytes, &varAssetTargetAllOf); err == nil {
		*o = AssetTargetAllOf(varAssetTargetAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClaimedByUserName")
		delete(additionalProperties, "Connections")
		delete(additionalProperties, "ConnectorVersion")
		delete(additionalProperties, "ExternalIpAddress")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ProductId")
		delete(additionalProperties, "ReadOnly")
		delete(additionalProperties, "Services")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusErrorReason")
		delete(additionalProperties, "TargetId")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Assist")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetTargetAllOf struct {
	value *AssetTargetAllOf
	isSet bool
}

func (v NullableAssetTargetAllOf) Get() *AssetTargetAllOf {
	return v.value
}

func (v *NullableAssetTargetAllOf) Set(val *AssetTargetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTargetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTargetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTargetAllOf(val *AssetTargetAllOf) *NullableAssetTargetAllOf {
	return &NullableAssetTargetAllOf{value: val, isSet: true}
}

func (v NullableAssetTargetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTargetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
