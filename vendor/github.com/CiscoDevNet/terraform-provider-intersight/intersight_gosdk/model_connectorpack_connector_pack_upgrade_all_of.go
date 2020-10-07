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

// ConnectorpackConnectorPackUpgradeAllOf Definition of the list of properties defined in 'connectorpack.ConnectorPackUpgrade', excluding properties defined in parent classes.
type ConnectorpackConnectorPackUpgradeAllOf struct {
	// The type of operation to be performed on UCS Director. * `Install` - Installs the requisite connector packs on UCS Director. * `Push` - Pushes the requisite connector packs to UCS Director.
	ConnectorPackOpType  *string                           `json:"ConnectorPackOpType,omitempty"`
	UcsdInfo             *IaasUcsdInfoRelationship         `json:"UcsdInfo,omitempty"`
	Workflow             *WorkflowWorkflowInfoRelationship `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorpackConnectorPackUpgradeAllOf ConnectorpackConnectorPackUpgradeAllOf

// NewConnectorpackConnectorPackUpgradeAllOf instantiates a new ConnectorpackConnectorPackUpgradeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorpackConnectorPackUpgradeAllOf() *ConnectorpackConnectorPackUpgradeAllOf {
	this := ConnectorpackConnectorPackUpgradeAllOf{}
	var connectorPackOpType string = "Install"
	this.ConnectorPackOpType = &connectorPackOpType
	return &this
}

// NewConnectorpackConnectorPackUpgradeAllOfWithDefaults instantiates a new ConnectorpackConnectorPackUpgradeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorpackConnectorPackUpgradeAllOfWithDefaults() *ConnectorpackConnectorPackUpgradeAllOf {
	this := ConnectorpackConnectorPackUpgradeAllOf{}
	var connectorPackOpType string = "Install"
	this.ConnectorPackOpType = &connectorPackOpType
	return &this
}

// GetConnectorPackOpType returns the ConnectorPackOpType field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpgradeAllOf) GetConnectorPackOpType() string {
	if o == nil || o.ConnectorPackOpType == nil {
		var ret string
		return ret
	}
	return *o.ConnectorPackOpType
}

// GetConnectorPackOpTypeOk returns a tuple with the ConnectorPackOpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpgradeAllOf) GetConnectorPackOpTypeOk() (*string, bool) {
	if o == nil || o.ConnectorPackOpType == nil {
		return nil, false
	}
	return o.ConnectorPackOpType, true
}

// HasConnectorPackOpType returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpgradeAllOf) HasConnectorPackOpType() bool {
	if o != nil && o.ConnectorPackOpType != nil {
		return true
	}

	return false
}

// SetConnectorPackOpType gets a reference to the given string and assigns it to the ConnectorPackOpType field.
func (o *ConnectorpackConnectorPackUpgradeAllOf) SetConnectorPackOpType(v string) {
	o.ConnectorPackOpType = &v
}

// GetUcsdInfo returns the UcsdInfo field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpgradeAllOf) GetUcsdInfo() IaasUcsdInfoRelationship {
	if o == nil || o.UcsdInfo == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.UcsdInfo
}

// GetUcsdInfoOk returns a tuple with the UcsdInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpgradeAllOf) GetUcsdInfoOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.UcsdInfo == nil {
		return nil, false
	}
	return o.UcsdInfo, true
}

// HasUcsdInfo returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpgradeAllOf) HasUcsdInfo() bool {
	if o != nil && o.UcsdInfo != nil {
		return true
	}

	return false
}

// SetUcsdInfo gets a reference to the given IaasUcsdInfoRelationship and assigns it to the UcsdInfo field.
func (o *ConnectorpackConnectorPackUpgradeAllOf) SetUcsdInfo(v IaasUcsdInfoRelationship) {
	o.UcsdInfo = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpgradeAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpgradeAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpgradeAllOf) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *ConnectorpackConnectorPackUpgradeAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o ConnectorpackConnectorPackUpgradeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectorPackOpType != nil {
		toSerialize["ConnectorPackOpType"] = o.ConnectorPackOpType
	}
	if o.UcsdInfo != nil {
		toSerialize["UcsdInfo"] = o.UcsdInfo
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorpackConnectorPackUpgradeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorpackConnectorPackUpgradeAllOf := _ConnectorpackConnectorPackUpgradeAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorpackConnectorPackUpgradeAllOf); err == nil {
		*o = ConnectorpackConnectorPackUpgradeAllOf(varConnectorpackConnectorPackUpgradeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ConnectorPackOpType")
		delete(additionalProperties, "UcsdInfo")
		delete(additionalProperties, "Workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorpackConnectorPackUpgradeAllOf struct {
	value *ConnectorpackConnectorPackUpgradeAllOf
	isSet bool
}

func (v NullableConnectorpackConnectorPackUpgradeAllOf) Get() *ConnectorpackConnectorPackUpgradeAllOf {
	return v.value
}

func (v *NullableConnectorpackConnectorPackUpgradeAllOf) Set(val *ConnectorpackConnectorPackUpgradeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorpackConnectorPackUpgradeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorpackConnectorPackUpgradeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorpackConnectorPackUpgradeAllOf(val *ConnectorpackConnectorPackUpgradeAllOf) *NullableConnectorpackConnectorPackUpgradeAllOf {
	return &NullableConnectorpackConnectorPackUpgradeAllOf{value: val, isSet: true}
}

func (v NullableConnectorpackConnectorPackUpgradeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorpackConnectorPackUpgradeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}