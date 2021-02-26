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

// AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerAmazonWebServicesBillingOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of S3 bucket that contains the Amazon web service Cost and Usage report to get cloud service spend.
	CostAndUsageReportBucket *string `json:"CostAndUsageReportBucket,omitempty"`
	// Report path prefix for the Amazon web service cost and usage report to get cloud service spend.
	CostAndUsageReportPath *string `json:"CostAndUsageReportPath,omitempty"`
	// Region for S3 bucket that contains the Amazon web service Cost and Usage report to get cloud service spend.
	CostAndUsageReportRegion *string `json:"CostAndUsageReportRegion,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf

// NewAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf instantiates a new AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf {
	this := AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOfWithDefaults() *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf {
	this := AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerAmazonWebServicesBillingOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerAmazonWebServicesBillingOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCostAndUsageReportBucket returns the CostAndUsageReportBucket field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetCostAndUsageReportBucket() string {
	if o == nil || o.CostAndUsageReportBucket == nil {
		var ret string
		return ret
	}
	return *o.CostAndUsageReportBucket
}

// GetCostAndUsageReportBucketOk returns a tuple with the CostAndUsageReportBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetCostAndUsageReportBucketOk() (*string, bool) {
	if o == nil || o.CostAndUsageReportBucket == nil {
		return nil, false
	}
	return o.CostAndUsageReportBucket, true
}

// HasCostAndUsageReportBucket returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) HasCostAndUsageReportBucket() bool {
	if o != nil && o.CostAndUsageReportBucket != nil {
		return true
	}

	return false
}

// SetCostAndUsageReportBucket gets a reference to the given string and assigns it to the CostAndUsageReportBucket field.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) SetCostAndUsageReportBucket(v string) {
	o.CostAndUsageReportBucket = &v
}

// GetCostAndUsageReportPath returns the CostAndUsageReportPath field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetCostAndUsageReportPath() string {
	if o == nil || o.CostAndUsageReportPath == nil {
		var ret string
		return ret
	}
	return *o.CostAndUsageReportPath
}

// GetCostAndUsageReportPathOk returns a tuple with the CostAndUsageReportPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetCostAndUsageReportPathOk() (*string, bool) {
	if o == nil || o.CostAndUsageReportPath == nil {
		return nil, false
	}
	return o.CostAndUsageReportPath, true
}

// HasCostAndUsageReportPath returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) HasCostAndUsageReportPath() bool {
	if o != nil && o.CostAndUsageReportPath != nil {
		return true
	}

	return false
}

// SetCostAndUsageReportPath gets a reference to the given string and assigns it to the CostAndUsageReportPath field.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) SetCostAndUsageReportPath(v string) {
	o.CostAndUsageReportPath = &v
}

// GetCostAndUsageReportRegion returns the CostAndUsageReportRegion field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetCostAndUsageReportRegion() string {
	if o == nil || o.CostAndUsageReportRegion == nil {
		var ret string
		return ret
	}
	return *o.CostAndUsageReportRegion
}

// GetCostAndUsageReportRegionOk returns a tuple with the CostAndUsageReportRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) GetCostAndUsageReportRegionOk() (*string, bool) {
	if o == nil || o.CostAndUsageReportRegion == nil {
		return nil, false
	}
	return o.CostAndUsageReportRegion, true
}

// HasCostAndUsageReportRegion returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) HasCostAndUsageReportRegion() bool {
	if o != nil && o.CostAndUsageReportRegion != nil {
		return true
	}

	return false
}

// SetCostAndUsageReportRegion gets a reference to the given string and assigns it to the CostAndUsageReportRegion field.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) SetCostAndUsageReportRegion(v string) {
	o.CostAndUsageReportRegion = &v
}

func (o AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CostAndUsageReportBucket != nil {
		toSerialize["CostAndUsageReportBucket"] = o.CostAndUsageReportBucket
	}
	if o.CostAndUsageReportPath != nil {
		toSerialize["CostAndUsageReportPath"] = o.CostAndUsageReportPath
	}
	if o.CostAndUsageReportRegion != nil {
		toSerialize["CostAndUsageReportRegion"] = o.CostAndUsageReportRegion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf := _AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf(varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CostAndUsageReportBucket")
		delete(additionalProperties, "CostAndUsageReportPath")
		delete(additionalProperties, "CostAndUsageReportRegion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf struct {
	value *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) Get() *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) Set(val *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf(val *AssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) *NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf {
	return &NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
