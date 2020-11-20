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

// PkixRsaAlgorithm The key pair is generated using the RSA algorithm and specified parameters.
type PkixRsaAlgorithm struct {
	PkixKeyGenerationSpec
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The length of the RSA key, expressed in bits, for both public and private keys. * `2048` - A key length of 2048 bits. * `2560` - A key length of 2560 bits. * `3072` - A key length of 3072 bits. * `3584` - A key length of 3584 bits. * `4096` - A key length of 4096 bits.
	Modulus              *int32 `json:"Modulus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixRsaAlgorithm PkixRsaAlgorithm

// NewPkixRsaAlgorithm instantiates a new PkixRsaAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixRsaAlgorithm(classId string, objectType string) *PkixRsaAlgorithm {
	this := PkixRsaAlgorithm{}
	this.ClassId = classId
	this.ObjectType = objectType
	var modulus int32 = 2048
	this.Modulus = &modulus
	return &this
}

// NewPkixRsaAlgorithmWithDefaults instantiates a new PkixRsaAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixRsaAlgorithmWithDefaults() *PkixRsaAlgorithm {
	this := PkixRsaAlgorithm{}
	var classId string = "pkix.RsaAlgorithm"
	this.ClassId = classId
	var objectType string = "pkix.RsaAlgorithm"
	this.ObjectType = objectType
	var modulus int32 = 2048
	this.Modulus = &modulus
	return &this
}

// GetClassId returns the ClassId field value
func (o *PkixRsaAlgorithm) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PkixRsaAlgorithm) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PkixRsaAlgorithm) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PkixRsaAlgorithm) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PkixRsaAlgorithm) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PkixRsaAlgorithm) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModulus returns the Modulus field value if set, zero value otherwise.
func (o *PkixRsaAlgorithm) GetModulus() int32 {
	if o == nil || o.Modulus == nil {
		var ret int32
		return ret
	}
	return *o.Modulus
}

// GetModulusOk returns a tuple with the Modulus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixRsaAlgorithm) GetModulusOk() (*int32, bool) {
	if o == nil || o.Modulus == nil {
		return nil, false
	}
	return o.Modulus, true
}

// HasModulus returns a boolean if a field has been set.
func (o *PkixRsaAlgorithm) HasModulus() bool {
	if o != nil && o.Modulus != nil {
		return true
	}

	return false
}

// SetModulus gets a reference to the given int32 and assigns it to the Modulus field.
func (o *PkixRsaAlgorithm) SetModulus(v int32) {
	o.Modulus = &v
}

func (o PkixRsaAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPkixKeyGenerationSpec, errPkixKeyGenerationSpec := json.Marshal(o.PkixKeyGenerationSpec)
	if errPkixKeyGenerationSpec != nil {
		return []byte{}, errPkixKeyGenerationSpec
	}
	errPkixKeyGenerationSpec = json.Unmarshal([]byte(serializedPkixKeyGenerationSpec), &toSerialize)
	if errPkixKeyGenerationSpec != nil {
		return []byte{}, errPkixKeyGenerationSpec
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Modulus != nil {
		toSerialize["Modulus"] = o.Modulus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixRsaAlgorithm) UnmarshalJSON(bytes []byte) (err error) {
	type PkixRsaAlgorithmWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The length of the RSA key, expressed in bits, for both public and private keys. * `2048` - A key length of 2048 bits. * `2560` - A key length of 2560 bits. * `3072` - A key length of 3072 bits. * `3584` - A key length of 3584 bits. * `4096` - A key length of 4096 bits.
		Modulus *int32 `json:"Modulus,omitempty"`
	}

	varPkixRsaAlgorithmWithoutEmbeddedStruct := PkixRsaAlgorithmWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPkixRsaAlgorithmWithoutEmbeddedStruct)
	if err == nil {
		varPkixRsaAlgorithm := _PkixRsaAlgorithm{}
		varPkixRsaAlgorithm.ClassId = varPkixRsaAlgorithmWithoutEmbeddedStruct.ClassId
		varPkixRsaAlgorithm.ObjectType = varPkixRsaAlgorithmWithoutEmbeddedStruct.ObjectType
		varPkixRsaAlgorithm.Modulus = varPkixRsaAlgorithmWithoutEmbeddedStruct.Modulus
		*o = PkixRsaAlgorithm(varPkixRsaAlgorithm)
	} else {
		return err
	}

	varPkixRsaAlgorithm := _PkixRsaAlgorithm{}

	err = json.Unmarshal(bytes, &varPkixRsaAlgorithm)
	if err == nil {
		o.PkixKeyGenerationSpec = varPkixRsaAlgorithm.PkixKeyGenerationSpec
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Modulus")

		// remove fields from embedded structs
		reflectPkixKeyGenerationSpec := reflect.ValueOf(o.PkixKeyGenerationSpec)
		for i := 0; i < reflectPkixKeyGenerationSpec.Type().NumField(); i++ {
			t := reflectPkixKeyGenerationSpec.Type().Field(i)

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

type NullablePkixRsaAlgorithm struct {
	value *PkixRsaAlgorithm
	isSet bool
}

func (v NullablePkixRsaAlgorithm) Get() *PkixRsaAlgorithm {
	return v.value
}

func (v *NullablePkixRsaAlgorithm) Set(val *PkixRsaAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixRsaAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixRsaAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixRsaAlgorithm(val *PkixRsaAlgorithm) *NullablePkixRsaAlgorithm {
	return &NullablePkixRsaAlgorithm{value: val, isSet: true}
}

func (v NullablePkixRsaAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixRsaAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
