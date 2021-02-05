# NiatelemetryNiaInventoryFabricAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NiaInventoryFabric"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NiaInventoryFabric"]
**AnycastGwMac** | Pointer to **string** | Returns the aycast gateway mac. | [optional] 
**DcnmtrackerEnabled** | Pointer to **bool** | Returns the value of the dcnmtrackerEnabled field. | [optional] 
**FabricId** | Pointer to **string** | Uniquely identifies a fabric. | [optional] 
**FabricName** | Pointer to **string** | Returns the value of the Name of a fabric. | [optional] 
**IsNgoamEnabled** | Pointer to **string** | Returns if ngoam is enabled. | [optional] 
**IsScheduledBackUpEnabled** | Pointer to **string** | Returns if the scheduled backup is enabled. | [optional] 
**LeafCount** | Pointer to **int64** | Returns total number of leafs in the fabric. | [optional] 
**LogicalLinks** | Pointer to [**[]NiatelemetryLogicalLink**](NiatelemetryLogicalLink.md) |  | [optional] 
**NxosVrfCount** | Pointer to **int64** | Returns the value of the nxosVrfCount field. | [optional] 
**SpineCount** | Pointer to **int64** | Returns total number of spines in the fabric. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaInventoryFabricAllOf

`func NewNiatelemetryNiaInventoryFabricAllOf(classId string, objectType string, ) *NiatelemetryNiaInventoryFabricAllOf`

NewNiatelemetryNiaInventoryFabricAllOf instantiates a new NiatelemetryNiaInventoryFabricAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryFabricAllOfWithDefaults

`func NewNiatelemetryNiaInventoryFabricAllOfWithDefaults() *NiatelemetryNiaInventoryFabricAllOf`

NewNiatelemetryNiaInventoryFabricAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryFabricAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetAnycastGwMac() string`

GetAnycastGwMac returns the AnycastGwMac field if non-nil, zero value otherwise.

### GetAnycastGwMacOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetAnycastGwMacOk() (*string, bool)`

GetAnycastGwMacOk returns a tuple with the AnycastGwMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetAnycastGwMac(v string)`

SetAnycastGwMac sets AnycastGwMac field to given value.

### HasAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasAnycastGwMac() bool`

HasAnycastGwMac returns a boolean if a field has been set.

### GetDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetDcnmtrackerEnabled() bool`

GetDcnmtrackerEnabled returns the DcnmtrackerEnabled field if non-nil, zero value otherwise.

### GetDcnmtrackerEnabledOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetDcnmtrackerEnabledOk() (*bool, bool)`

GetDcnmtrackerEnabledOk returns a tuple with the DcnmtrackerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetDcnmtrackerEnabled(v bool)`

SetDcnmtrackerEnabled sets DcnmtrackerEnabled field to given value.

### HasDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasDcnmtrackerEnabled() bool`

HasDcnmtrackerEnabled returns a boolean if a field has been set.

### GetFabricId

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsNgoamEnabled() string`

GetIsNgoamEnabled returns the IsNgoamEnabled field if non-nil, zero value otherwise.

### GetIsNgoamEnabledOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsNgoamEnabledOk() (*string, bool)`

GetIsNgoamEnabledOk returns a tuple with the IsNgoamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsNgoamEnabled(v string)`

SetIsNgoamEnabled sets IsNgoamEnabled field to given value.

### HasIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsNgoamEnabled() bool`

HasIsNgoamEnabled returns a boolean if a field has been set.

### GetIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsScheduledBackUpEnabled() string`

GetIsScheduledBackUpEnabled returns the IsScheduledBackUpEnabled field if non-nil, zero value otherwise.

### GetIsScheduledBackUpEnabledOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetIsScheduledBackUpEnabledOk() (*string, bool)`

GetIsScheduledBackUpEnabledOk returns a tuple with the IsScheduledBackUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetIsScheduledBackUpEnabled(v string)`

SetIsScheduledBackUpEnabled sets IsScheduledBackUpEnabled field to given value.

### HasIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasIsScheduledBackUpEnabled() bool`

HasIsScheduledBackUpEnabled returns a boolean if a field has been set.

### GetLeafCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetLeafCount() int64`

GetLeafCount returns the LeafCount field if non-nil, zero value otherwise.

### GetLeafCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetLeafCountOk() (*int64, bool)`

GetLeafCountOk returns a tuple with the LeafCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetLeafCount(v int64)`

SetLeafCount sets LeafCount field to given value.

### HasLeafCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasLeafCount() bool`

HasLeafCount returns a boolean if a field has been set.

### GetLogicalLinks

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetLogicalLinks() []NiatelemetryLogicalLink`

GetLogicalLinks returns the LogicalLinks field if non-nil, zero value otherwise.

### GetLogicalLinksOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetLogicalLinksOk() (*[]NiatelemetryLogicalLink, bool)`

GetLogicalLinksOk returns a tuple with the LogicalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalLinks

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetLogicalLinks(v []NiatelemetryLogicalLink)`

SetLogicalLinks sets LogicalLinks field to given value.

### HasLogicalLinks

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasLogicalLinks() bool`

HasLogicalLinks returns a boolean if a field has been set.

### SetLogicalLinksNil

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetLogicalLinksNil(b bool)`

 SetLogicalLinksNil sets the value for LogicalLinks to be an explicit nil

### UnsetLogicalLinks
`func (o *NiatelemetryNiaInventoryFabricAllOf) UnsetLogicalLinks()`

UnsetLogicalLinks ensures that no value is present for LogicalLinks, not even an explicit nil
### GetNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfCount() int64`

GetNxosVrfCount returns the NxosVrfCount field if non-nil, zero value otherwise.

### GetNxosVrfCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetNxosVrfCountOk() (*int64, bool)`

GetNxosVrfCountOk returns a tuple with the NxosVrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetNxosVrfCount(v int64)`

SetNxosVrfCount sets NxosVrfCount field to given value.

### HasNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasNxosVrfCount() bool`

HasNxosVrfCount returns a boolean if a field has been set.

### GetSpineCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetSpineCount() int64`

GetSpineCount returns the SpineCount field if non-nil, zero value otherwise.

### GetSpineCountOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetSpineCountOk() (*int64, bool)`

GetSpineCountOk returns a tuple with the SpineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpineCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetSpineCount(v int64)`

SetSpineCount sets SpineCount field to given value.

### HasSpineCount

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasSpineCount() bool`

HasSpineCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaInventoryFabricAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabricAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabricAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


