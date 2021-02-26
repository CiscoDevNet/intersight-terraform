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

// TelemetryDruidInvertedTopNMetricSpec Sort dimension values in inverted order, i.e inverts the order of the delegate metric spec. It can be used to sort the values in ascending order.
type TelemetryDruidInvertedTopNMetricSpec struct {
	// The dimension spec type.
	Type                 string                       `json:"type"`
	Metric               TelemetryDruidTopNMetricSpec `json:"metric"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidInvertedTopNMetricSpec TelemetryDruidInvertedTopNMetricSpec

// NewTelemetryDruidInvertedTopNMetricSpec instantiates a new TelemetryDruidInvertedTopNMetricSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidInvertedTopNMetricSpec(type_ string, metric TelemetryDruidTopNMetricSpec) *TelemetryDruidInvertedTopNMetricSpec {
	this := TelemetryDruidInvertedTopNMetricSpec{}
	this.Type = type_
	this.Metric = metric
	return &this
}

// NewTelemetryDruidInvertedTopNMetricSpecWithDefaults instantiates a new TelemetryDruidInvertedTopNMetricSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidInvertedTopNMetricSpecWithDefaults() *TelemetryDruidInvertedTopNMetricSpec {
	this := TelemetryDruidInvertedTopNMetricSpec{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidInvertedTopNMetricSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInvertedTopNMetricSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidInvertedTopNMetricSpec) SetType(v string) {
	o.Type = v
}

// GetMetric returns the Metric field value
func (o *TelemetryDruidInvertedTopNMetricSpec) GetMetric() TelemetryDruidTopNMetricSpec {
	if o == nil {
		var ret TelemetryDruidTopNMetricSpec
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInvertedTopNMetricSpec) GetMetricOk() (*TelemetryDruidTopNMetricSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *TelemetryDruidInvertedTopNMetricSpec) SetMetric(v TelemetryDruidTopNMetricSpec) {
	o.Metric = v
}

func (o TelemetryDruidInvertedTopNMetricSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["metric"] = o.Metric
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidInvertedTopNMetricSpec) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidInvertedTopNMetricSpec := _TelemetryDruidInvertedTopNMetricSpec{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidInvertedTopNMetricSpec); err == nil {
		*o = TelemetryDruidInvertedTopNMetricSpec(varTelemetryDruidInvertedTopNMetricSpec)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "metric")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidInvertedTopNMetricSpec struct {
	value *TelemetryDruidInvertedTopNMetricSpec
	isSet bool
}

func (v NullableTelemetryDruidInvertedTopNMetricSpec) Get() *TelemetryDruidInvertedTopNMetricSpec {
	return v.value
}

func (v *NullableTelemetryDruidInvertedTopNMetricSpec) Set(val *TelemetryDruidInvertedTopNMetricSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidInvertedTopNMetricSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidInvertedTopNMetricSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidInvertedTopNMetricSpec(val *TelemetryDruidInvertedTopNMetricSpec) *NullableTelemetryDruidInvertedTopNMetricSpec {
	return &NullableTelemetryDruidInvertedTopNMetricSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidInvertedTopNMetricSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidInvertedTopNMetricSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
