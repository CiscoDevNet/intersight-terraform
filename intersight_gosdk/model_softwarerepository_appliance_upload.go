/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// SoftwarerepositoryApplianceUpload The user's local machine is being used as the source of an image. This upload operation is for an airgapped appliance.
type SoftwarerepositoryApplianceUpload struct {
	SoftwarerepositoryFileServer
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryApplianceUpload SoftwarerepositoryApplianceUpload

// NewSoftwarerepositoryApplianceUpload instantiates a new SoftwarerepositoryApplianceUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryApplianceUpload() *SoftwarerepositoryApplianceUpload {
	this := SoftwarerepositoryApplianceUpload{}
	return &this
}

// NewSoftwarerepositoryApplianceUploadWithDefaults instantiates a new SoftwarerepositoryApplianceUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryApplianceUploadWithDefaults() *SoftwarerepositoryApplianceUpload {
	this := SoftwarerepositoryApplianceUpload{}
	return &this
}

func (o SoftwarerepositoryApplianceUpload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSoftwarerepositoryFileServer, errSoftwarerepositoryFileServer := json.Marshal(o.SoftwarerepositoryFileServer)
	if errSoftwarerepositoryFileServer != nil {
		return []byte{}, errSoftwarerepositoryFileServer
	}
	errSoftwarerepositoryFileServer = json.Unmarshal([]byte(serializedSoftwarerepositoryFileServer), &toSerialize)
	if errSoftwarerepositoryFileServer != nil {
		return []byte{}, errSoftwarerepositoryFileServer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryApplianceUpload) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwarerepositoryApplianceUploadWithoutEmbeddedStruct struct {
	}

	varSoftwarerepositoryApplianceUploadWithoutEmbeddedStruct := SoftwarerepositoryApplianceUploadWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryApplianceUploadWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryApplianceUpload := _SoftwarerepositoryApplianceUpload{}
		*o = SoftwarerepositoryApplianceUpload(varSoftwarerepositoryApplianceUpload)
	} else {
		return err
	}

	varSoftwarerepositoryApplianceUpload := _SoftwarerepositoryApplianceUpload{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryApplianceUpload)
	if err == nil {
		o.SoftwarerepositoryFileServer = varSoftwarerepositoryApplianceUpload.SoftwarerepositoryFileServer
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectSoftwarerepositoryFileServer := reflect.ValueOf(o.SoftwarerepositoryFileServer)
		for i := 0; i < reflectSoftwarerepositoryFileServer.Type().NumField(); i++ {
			t := reflectSoftwarerepositoryFileServer.Type().Field(i)

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

type NullableSoftwarerepositoryApplianceUpload struct {
	value *SoftwarerepositoryApplianceUpload
	isSet bool
}

func (v NullableSoftwarerepositoryApplianceUpload) Get() *SoftwarerepositoryApplianceUpload {
	return v.value
}

func (v *NullableSoftwarerepositoryApplianceUpload) Set(val *SoftwarerepositoryApplianceUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryApplianceUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryApplianceUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryApplianceUpload(val *SoftwarerepositoryApplianceUpload) *NullableSoftwarerepositoryApplianceUpload {
	return &NullableSoftwarerepositoryApplianceUpload{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryApplianceUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryApplianceUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
