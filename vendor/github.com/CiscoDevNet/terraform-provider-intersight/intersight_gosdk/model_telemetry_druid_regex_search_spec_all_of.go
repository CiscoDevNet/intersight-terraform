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
)

// TelemetryDruidRegexSearchSpecAllOf struct for TelemetryDruidRegexSearchSpecAllOf
type TelemetryDruidRegexSearchSpecAllOf struct {
	// The regular expression to match.  If any part of a dimension value contains the pattern specified in this search query a \"match\" occurs.
	Regex                string `json:"regex"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidRegexSearchSpecAllOf TelemetryDruidRegexSearchSpecAllOf

// NewTelemetryDruidRegexSearchSpecAllOf instantiates a new TelemetryDruidRegexSearchSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidRegexSearchSpecAllOf(regex string) *TelemetryDruidRegexSearchSpecAllOf {
	this := TelemetryDruidRegexSearchSpecAllOf{}
	this.Regex = regex
	return &this
}

// NewTelemetryDruidRegexSearchSpecAllOfWithDefaults instantiates a new TelemetryDruidRegexSearchSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidRegexSearchSpecAllOfWithDefaults() *TelemetryDruidRegexSearchSpecAllOf {
	this := TelemetryDruidRegexSearchSpecAllOf{}
	return &this
}

// GetRegex returns the Regex field value
func (o *TelemetryDruidRegexSearchSpecAllOf) GetRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Regex
}

// GetRegexOk returns a tuple with the Regex field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidRegexSearchSpecAllOf) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Regex, true
}

// SetRegex sets field value
func (o *TelemetryDruidRegexSearchSpecAllOf) SetRegex(v string) {
	o.Regex = v
}

func (o TelemetryDruidRegexSearchSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regex"] = o.Regex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidRegexSearchSpecAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidRegexSearchSpecAllOf := _TelemetryDruidRegexSearchSpecAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidRegexSearchSpecAllOf); err == nil {
		*o = TelemetryDruidRegexSearchSpecAllOf(varTelemetryDruidRegexSearchSpecAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "regex")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidRegexSearchSpecAllOf struct {
	value *TelemetryDruidRegexSearchSpecAllOf
	isSet bool
}

func (v NullableTelemetryDruidRegexSearchSpecAllOf) Get() *TelemetryDruidRegexSearchSpecAllOf {
	return v.value
}

func (v *NullableTelemetryDruidRegexSearchSpecAllOf) Set(val *TelemetryDruidRegexSearchSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidRegexSearchSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidRegexSearchSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidRegexSearchSpecAllOf(val *TelemetryDruidRegexSearchSpecAllOf) *NullableTelemetryDruidRegexSearchSpecAllOf {
	return &NullableTelemetryDruidRegexSearchSpecAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidRegexSearchSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidRegexSearchSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
