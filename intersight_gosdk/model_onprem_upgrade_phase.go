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
	"time"
)

// OnpremUpgradePhase UpgradePhase represents a phase of the Intersight Appliance software upgrade process. This data structure is shared by both the Intersight upgrade service and the Intersight Appliance's upgrade service.
type OnpremUpgradePhase struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Elapsed time in seconds to complete the current upgrade phase.
	ElapsedTime *int64 `json:"ElapsedTime,omitempty"`
	// End date of the software upgrade phase.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Indicates if the upgrade phase has failed or not.
	Failed *bool `json:"Failed,omitempty"`
	// Status message set during the upgrade phase.
	Message *string `json:"Message,omitempty"`
	// Name of the upgrade phase. * `init` - Upgrade service initialization phase. * `Prepare` - Upgrade service prepares folders and templated files. * `ServiceLoad` - Upgrade service loads the service images into the local docker cache. * `UiLoad` - Upgrade service loads the UI packages into the local cache. * `GenerateConfig` - Upgrade service generates the Kubernetes configuration files. * `DeployService` - Upgrade service deploys the Kubernetes services. * `Success` - Upgrade completed successfully. * `Fail` - Indicates that the upgrade process has failed. * `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance. * `Telemetry` - Upgrade service sends basic telemetry data to the Intersight.
	Name *string `json:"Name,omitempty"`
	// Start date of the software upgrade phase.
	StartTime            *time.Time `json:"StartTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnpremUpgradePhase OnpremUpgradePhase

// NewOnpremUpgradePhase instantiates a new OnpremUpgradePhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnpremUpgradePhase(classId string, objectType string) *OnpremUpgradePhase {
	this := OnpremUpgradePhase{}
	this.ClassId = classId
	this.ObjectType = objectType
	var name string = "init"
	this.Name = &name
	return &this
}

// NewOnpremUpgradePhaseWithDefaults instantiates a new OnpremUpgradePhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnpremUpgradePhaseWithDefaults() *OnpremUpgradePhase {
	this := OnpremUpgradePhase{}
	var classId string = "onprem.UpgradePhase"
	this.ClassId = classId
	var objectType string = "onprem.UpgradePhase"
	this.ObjectType = objectType
	var name string = "init"
	this.Name = &name
	return &this
}

// GetClassId returns the ClassId field value
func (o *OnpremUpgradePhase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OnpremUpgradePhase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OnpremUpgradePhase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OnpremUpgradePhase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *OnpremUpgradePhase) GetElapsedTime() int64 {
	if o == nil || o.ElapsedTime == nil {
		var ret int64
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhase) GetElapsedTimeOk() (*int64, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *OnpremUpgradePhase) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int64 and assigns it to the ElapsedTime field.
func (o *OnpremUpgradePhase) SetElapsedTime(v int64) {
	o.ElapsedTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *OnpremUpgradePhase) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhase) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *OnpremUpgradePhase) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *OnpremUpgradePhase) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *OnpremUpgradePhase) GetFailed() bool {
	if o == nil || o.Failed == nil {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhase) GetFailedOk() (*bool, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *OnpremUpgradePhase) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *OnpremUpgradePhase) SetFailed(v bool) {
	o.Failed = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OnpremUpgradePhase) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhase) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OnpremUpgradePhase) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OnpremUpgradePhase) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OnpremUpgradePhase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OnpremUpgradePhase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OnpremUpgradePhase) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *OnpremUpgradePhase) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhase) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *OnpremUpgradePhase) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *OnpremUpgradePhase) SetStartTime(v time.Time) {
	o.StartTime = &v
}

func (o OnpremUpgradePhase) MarshalJSON() ([]byte, error) {
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
	if o.ElapsedTime != nil {
		toSerialize["ElapsedTime"] = o.ElapsedTime
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Failed != nil {
		toSerialize["Failed"] = o.Failed
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OnpremUpgradePhase) UnmarshalJSON(bytes []byte) (err error) {
	type OnpremUpgradePhaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Elapsed time in seconds to complete the current upgrade phase.
		ElapsedTime *int64 `json:"ElapsedTime,omitempty"`
		// End date of the software upgrade phase.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// Indicates if the upgrade phase has failed or not.
		Failed *bool `json:"Failed,omitempty"`
		// Status message set during the upgrade phase.
		Message *string `json:"Message,omitempty"`
		// Name of the upgrade phase. * `init` - Upgrade service initialization phase. * `Prepare` - Upgrade service prepares folders and templated files. * `ServiceLoad` - Upgrade service loads the service images into the local docker cache. * `UiLoad` - Upgrade service loads the UI packages into the local cache. * `GenerateConfig` - Upgrade service generates the Kubernetes configuration files. * `DeployService` - Upgrade service deploys the Kubernetes services. * `Success` - Upgrade completed successfully. * `Fail` - Indicates that the upgrade process has failed. * `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance. * `Telemetry` - Upgrade service sends basic telemetry data to the Intersight.
		Name *string `json:"Name,omitempty"`
		// Start date of the software upgrade phase.
		StartTime *time.Time `json:"StartTime,omitempty"`
	}

	varOnpremUpgradePhaseWithoutEmbeddedStruct := OnpremUpgradePhaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOnpremUpgradePhaseWithoutEmbeddedStruct)
	if err == nil {
		varOnpremUpgradePhase := _OnpremUpgradePhase{}
		varOnpremUpgradePhase.ClassId = varOnpremUpgradePhaseWithoutEmbeddedStruct.ClassId
		varOnpremUpgradePhase.ObjectType = varOnpremUpgradePhaseWithoutEmbeddedStruct.ObjectType
		varOnpremUpgradePhase.ElapsedTime = varOnpremUpgradePhaseWithoutEmbeddedStruct.ElapsedTime
		varOnpremUpgradePhase.EndTime = varOnpremUpgradePhaseWithoutEmbeddedStruct.EndTime
		varOnpremUpgradePhase.Failed = varOnpremUpgradePhaseWithoutEmbeddedStruct.Failed
		varOnpremUpgradePhase.Message = varOnpremUpgradePhaseWithoutEmbeddedStruct.Message
		varOnpremUpgradePhase.Name = varOnpremUpgradePhaseWithoutEmbeddedStruct.Name
		varOnpremUpgradePhase.StartTime = varOnpremUpgradePhaseWithoutEmbeddedStruct.StartTime
		*o = OnpremUpgradePhase(varOnpremUpgradePhase)
	} else {
		return err
	}

	varOnpremUpgradePhase := _OnpremUpgradePhase{}

	err = json.Unmarshal(bytes, &varOnpremUpgradePhase)
	if err == nil {
		o.MoBaseComplexType = varOnpremUpgradePhase.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ElapsedTime")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Failed")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StartTime")

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

type NullableOnpremUpgradePhase struct {
	value *OnpremUpgradePhase
	isSet bool
}

func (v NullableOnpremUpgradePhase) Get() *OnpremUpgradePhase {
	return v.value
}

func (v *NullableOnpremUpgradePhase) Set(val *OnpremUpgradePhase) {
	v.value = val
	v.isSet = true
}

func (v NullableOnpremUpgradePhase) IsSet() bool {
	return v.isSet
}

func (v *NullableOnpremUpgradePhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnpremUpgradePhase(val *OnpremUpgradePhase) *NullableOnpremUpgradePhase {
	return &NullableOnpremUpgradePhase{value: val, isSet: true}
}

func (v NullableOnpremUpgradePhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnpremUpgradePhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
