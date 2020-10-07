/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// LicenseSmartlicenseTokenAllOf Definition of the list of properties defined in 'license.SmartlicenseToken', excluding properties defined in parent classes.
type LicenseSmartlicenseTokenAllOf struct {
	// Smart license registration token.
	Token                *string                                `json:"Token,omitempty"`
	AccountLicenseData   *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseSmartlicenseTokenAllOf LicenseSmartlicenseTokenAllOf

// NewLicenseSmartlicenseTokenAllOf instantiates a new LicenseSmartlicenseTokenAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseSmartlicenseTokenAllOf() *LicenseSmartlicenseTokenAllOf {
	this := LicenseSmartlicenseTokenAllOf{}
	return &this
}

// NewLicenseSmartlicenseTokenAllOfWithDefaults instantiates a new LicenseSmartlicenseTokenAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseSmartlicenseTokenAllOfWithDefaults() *LicenseSmartlicenseTokenAllOf {
	this := LicenseSmartlicenseTokenAllOf{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LicenseSmartlicenseTokenAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSmartlicenseTokenAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LicenseSmartlicenseTokenAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LicenseSmartlicenseTokenAllOf) SetToken(v string) {
	o.Token = &v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise.
func (o *LicenseSmartlicenseTokenAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || o.AccountLicenseData == nil {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSmartlicenseTokenAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil || o.AccountLicenseData == nil {
		return nil, false
	}
	return o.AccountLicenseData, true
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseSmartlicenseTokenAllOf) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData != nil {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given LicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseSmartlicenseTokenAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData = &v
}

func (o LicenseSmartlicenseTokenAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["Token"] = o.Token
	}
	if o.AccountLicenseData != nil {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseSmartlicenseTokenAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLicenseSmartlicenseTokenAllOf := _LicenseSmartlicenseTokenAllOf{}

	if err = json.Unmarshal(bytes, &varLicenseSmartlicenseTokenAllOf); err == nil {
		*o = LicenseSmartlicenseTokenAllOf(varLicenseSmartlicenseTokenAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Token")
		delete(additionalProperties, "AccountLicenseData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseSmartlicenseTokenAllOf struct {
	value *LicenseSmartlicenseTokenAllOf
	isSet bool
}

func (v NullableLicenseSmartlicenseTokenAllOf) Get() *LicenseSmartlicenseTokenAllOf {
	return v.value
}

func (v *NullableLicenseSmartlicenseTokenAllOf) Set(val *LicenseSmartlicenseTokenAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseSmartlicenseTokenAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseSmartlicenseTokenAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseSmartlicenseTokenAllOf(val *LicenseSmartlicenseTokenAllOf) *NullableLicenseSmartlicenseTokenAllOf {
	return &NullableLicenseSmartlicenseTokenAllOf{value: val, isSet: true}
}

func (v NullableLicenseSmartlicenseTokenAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseSmartlicenseTokenAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}