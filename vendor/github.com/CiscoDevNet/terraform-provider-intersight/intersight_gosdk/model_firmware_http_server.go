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

// FirmwareHttpServer An external HTTP file server.
type FirmwareHttpServer struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices.
	LocationLink *string `json:"LocationLink,omitempty"`
	// Mount option as configured on the HTTP server. Empty if nothing is configured.
	MountOptions         *string `json:"MountOptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareHttpServer FirmwareHttpServer

// NewFirmwareHttpServer instantiates a new FirmwareHttpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareHttpServer(classId string, objectType string) *FirmwareHttpServer {
	this := FirmwareHttpServer{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareHttpServerWithDefaults instantiates a new FirmwareHttpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareHttpServerWithDefaults() *FirmwareHttpServer {
	this := FirmwareHttpServer{}
	var classId string = "firmware.HttpServer"
	this.ClassId = classId
	var objectType string = "firmware.HttpServer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareHttpServer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareHttpServer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareHttpServer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareHttpServer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLocationLink returns the LocationLink field value if set, zero value otherwise.
func (o *FirmwareHttpServer) GetLocationLink() string {
	if o == nil || o.LocationLink == nil {
		var ret string
		return ret
	}
	return *o.LocationLink
}

// GetLocationLinkOk returns a tuple with the LocationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServer) GetLocationLinkOk() (*string, bool) {
	if o == nil || o.LocationLink == nil {
		return nil, false
	}
	return o.LocationLink, true
}

// HasLocationLink returns a boolean if a field has been set.
func (o *FirmwareHttpServer) HasLocationLink() bool {
	if o != nil && o.LocationLink != nil {
		return true
	}

	return false
}

// SetLocationLink gets a reference to the given string and assigns it to the LocationLink field.
func (o *FirmwareHttpServer) SetLocationLink(v string) {
	o.LocationLink = &v
}

// GetMountOptions returns the MountOptions field value if set, zero value otherwise.
func (o *FirmwareHttpServer) GetMountOptions() string {
	if o == nil || o.MountOptions == nil {
		var ret string
		return ret
	}
	return *o.MountOptions
}

// GetMountOptionsOk returns a tuple with the MountOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServer) GetMountOptionsOk() (*string, bool) {
	if o == nil || o.MountOptions == nil {
		return nil, false
	}
	return o.MountOptions, true
}

// HasMountOptions returns a boolean if a field has been set.
func (o *FirmwareHttpServer) HasMountOptions() bool {
	if o != nil && o.MountOptions != nil {
		return true
	}

	return false
}

// SetMountOptions gets a reference to the given string and assigns it to the MountOptions field.
func (o *FirmwareHttpServer) SetMountOptions(v string) {
	o.MountOptions = &v
}

func (o FirmwareHttpServer) MarshalJSON() ([]byte, error) {
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
	if o.LocationLink != nil {
		toSerialize["LocationLink"] = o.LocationLink
	}
	if o.MountOptions != nil {
		toSerialize["MountOptions"] = o.MountOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareHttpServer) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareHttpServerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices.
		LocationLink *string `json:"LocationLink,omitempty"`
		// Mount option as configured on the HTTP server. Empty if nothing is configured.
		MountOptions *string `json:"MountOptions,omitempty"`
	}

	varFirmwareHttpServerWithoutEmbeddedStruct := FirmwareHttpServerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareHttpServerWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareHttpServer := _FirmwareHttpServer{}
		varFirmwareHttpServer.ClassId = varFirmwareHttpServerWithoutEmbeddedStruct.ClassId
		varFirmwareHttpServer.ObjectType = varFirmwareHttpServerWithoutEmbeddedStruct.ObjectType
		varFirmwareHttpServer.LocationLink = varFirmwareHttpServerWithoutEmbeddedStruct.LocationLink
		varFirmwareHttpServer.MountOptions = varFirmwareHttpServerWithoutEmbeddedStruct.MountOptions
		*o = FirmwareHttpServer(varFirmwareHttpServer)
	} else {
		return err
	}

	varFirmwareHttpServer := _FirmwareHttpServer{}

	err = json.Unmarshal(bytes, &varFirmwareHttpServer)
	if err == nil {
		o.MoBaseComplexType = varFirmwareHttpServer.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LocationLink")
		delete(additionalProperties, "MountOptions")

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

type NullableFirmwareHttpServer struct {
	value *FirmwareHttpServer
	isSet bool
}

func (v NullableFirmwareHttpServer) Get() *FirmwareHttpServer {
	return v.value
}

func (v *NullableFirmwareHttpServer) Set(val *FirmwareHttpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareHttpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareHttpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareHttpServer(val *FirmwareHttpServer) *NullableFirmwareHttpServer {
	return &NullableFirmwareHttpServer{value: val, isSet: true}
}

func (v NullableFirmwareHttpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareHttpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
