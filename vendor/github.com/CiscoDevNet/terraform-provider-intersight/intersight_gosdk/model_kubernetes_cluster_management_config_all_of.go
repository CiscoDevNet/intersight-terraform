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

// KubernetesClusterManagementConfigAllOf Definition of the list of properties defined in 'kubernetes.ClusterManagementConfig', excluding properties defined in parent classes.
type KubernetesClusterManagementConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Encrypt ETCD data at rest using the etcdEncryptionKey specified in the cluster Kubernetes configuration.
	EncryptedEtcd *bool `json:"EncryptedEtcd,omitempty"`
	// Private key for internal mgmt of the cluster via SSH.
	InternalMgmtPrivateKey *string `json:"InternalMgmtPrivateKey,omitempty"`
	// Public key for internal mgmt of the cluster via SSH.
	InternalMgmtPublicKey *string `json:"InternalMgmtPublicKey,omitempty"`
	// Number of IP addresses to reserve for load balancer services.
	LoadBalancerCount *int64   `json:"LoadBalancerCount,omitempty"`
	LoadBalancers     []string `json:"LoadBalancers,omitempty"`
	// VIP for the cluster Kubernetes API server. If this is empty and a cluster IP pool is specified, it will be allocated from the IP pool.
	MasterVip *string  `json:"MasterVip,omitempty"`
	SshKeys   []string `json:"SshKeys,omitempty"`
	// Name of the user to SSH to nodes in a cluster.
	SshUser              *string `json:"SshUser,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterManagementConfigAllOf KubernetesClusterManagementConfigAllOf

// NewKubernetesClusterManagementConfigAllOf instantiates a new KubernetesClusterManagementConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterManagementConfigAllOf(classId string, objectType string) *KubernetesClusterManagementConfigAllOf {
	this := KubernetesClusterManagementConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var sshUser string = "iksadmin"
	this.SshUser = &sshUser
	return &this
}

// NewKubernetesClusterManagementConfigAllOfWithDefaults instantiates a new KubernetesClusterManagementConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterManagementConfigAllOfWithDefaults() *KubernetesClusterManagementConfigAllOf {
	this := KubernetesClusterManagementConfigAllOf{}
	var classId string = "kubernetes.ClusterManagementConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.ClusterManagementConfig"
	this.ObjectType = objectType
	var sshUser string = "iksadmin"
	this.SshUser = &sshUser
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesClusterManagementConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesClusterManagementConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesClusterManagementConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesClusterManagementConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEncryptedEtcd returns the EncryptedEtcd field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfigAllOf) GetEncryptedEtcd() bool {
	if o == nil || o.EncryptedEtcd == nil {
		var ret bool
		return ret
	}
	return *o.EncryptedEtcd
}

// GetEncryptedEtcdOk returns a tuple with the EncryptedEtcd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfigAllOf) GetEncryptedEtcdOk() (*bool, bool) {
	if o == nil || o.EncryptedEtcd == nil {
		return nil, false
	}
	return o.EncryptedEtcd, true
}

// HasEncryptedEtcd returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfigAllOf) HasEncryptedEtcd() bool {
	if o != nil && o.EncryptedEtcd != nil {
		return true
	}

	return false
}

// SetEncryptedEtcd gets a reference to the given bool and assigns it to the EncryptedEtcd field.
func (o *KubernetesClusterManagementConfigAllOf) SetEncryptedEtcd(v bool) {
	o.EncryptedEtcd = &v
}

// GetInternalMgmtPrivateKey returns the InternalMgmtPrivateKey field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfigAllOf) GetInternalMgmtPrivateKey() string {
	if o == nil || o.InternalMgmtPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.InternalMgmtPrivateKey
}

// GetInternalMgmtPrivateKeyOk returns a tuple with the InternalMgmtPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfigAllOf) GetInternalMgmtPrivateKeyOk() (*string, bool) {
	if o == nil || o.InternalMgmtPrivateKey == nil {
		return nil, false
	}
	return o.InternalMgmtPrivateKey, true
}

// HasInternalMgmtPrivateKey returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfigAllOf) HasInternalMgmtPrivateKey() bool {
	if o != nil && o.InternalMgmtPrivateKey != nil {
		return true
	}

	return false
}

// SetInternalMgmtPrivateKey gets a reference to the given string and assigns it to the InternalMgmtPrivateKey field.
func (o *KubernetesClusterManagementConfigAllOf) SetInternalMgmtPrivateKey(v string) {
	o.InternalMgmtPrivateKey = &v
}

// GetInternalMgmtPublicKey returns the InternalMgmtPublicKey field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfigAllOf) GetInternalMgmtPublicKey() string {
	if o == nil || o.InternalMgmtPublicKey == nil {
		var ret string
		return ret
	}
	return *o.InternalMgmtPublicKey
}

// GetInternalMgmtPublicKeyOk returns a tuple with the InternalMgmtPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfigAllOf) GetInternalMgmtPublicKeyOk() (*string, bool) {
	if o == nil || o.InternalMgmtPublicKey == nil {
		return nil, false
	}
	return o.InternalMgmtPublicKey, true
}

// HasInternalMgmtPublicKey returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfigAllOf) HasInternalMgmtPublicKey() bool {
	if o != nil && o.InternalMgmtPublicKey != nil {
		return true
	}

	return false
}

// SetInternalMgmtPublicKey gets a reference to the given string and assigns it to the InternalMgmtPublicKey field.
func (o *KubernetesClusterManagementConfigAllOf) SetInternalMgmtPublicKey(v string) {
	o.InternalMgmtPublicKey = &v
}

// GetLoadBalancerCount returns the LoadBalancerCount field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfigAllOf) GetLoadBalancerCount() int64 {
	if o == nil || o.LoadBalancerCount == nil {
		var ret int64
		return ret
	}
	return *o.LoadBalancerCount
}

// GetLoadBalancerCountOk returns a tuple with the LoadBalancerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfigAllOf) GetLoadBalancerCountOk() (*int64, bool) {
	if o == nil || o.LoadBalancerCount == nil {
		return nil, false
	}
	return o.LoadBalancerCount, true
}

// HasLoadBalancerCount returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfigAllOf) HasLoadBalancerCount() bool {
	if o != nil && o.LoadBalancerCount != nil {
		return true
	}

	return false
}

// SetLoadBalancerCount gets a reference to the given int64 and assigns it to the LoadBalancerCount field.
func (o *KubernetesClusterManagementConfigAllOf) SetLoadBalancerCount(v int64) {
	o.LoadBalancerCount = &v
}

// GetLoadBalancers returns the LoadBalancers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterManagementConfigAllOf) GetLoadBalancers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.LoadBalancers
}

// GetLoadBalancersOk returns a tuple with the LoadBalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterManagementConfigAllOf) GetLoadBalancersOk() (*[]string, bool) {
	if o == nil || o.LoadBalancers == nil {
		return nil, false
	}
	return &o.LoadBalancers, true
}

// HasLoadBalancers returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfigAllOf) HasLoadBalancers() bool {
	if o != nil && o.LoadBalancers != nil {
		return true
	}

	return false
}

// SetLoadBalancers gets a reference to the given []string and assigns it to the LoadBalancers field.
func (o *KubernetesClusterManagementConfigAllOf) SetLoadBalancers(v []string) {
	o.LoadBalancers = v
}

// GetMasterVip returns the MasterVip field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfigAllOf) GetMasterVip() string {
	if o == nil || o.MasterVip == nil {
		var ret string
		return ret
	}
	return *o.MasterVip
}

// GetMasterVipOk returns a tuple with the MasterVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfigAllOf) GetMasterVipOk() (*string, bool) {
	if o == nil || o.MasterVip == nil {
		return nil, false
	}
	return o.MasterVip, true
}

// HasMasterVip returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfigAllOf) HasMasterVip() bool {
	if o != nil && o.MasterVip != nil {
		return true
	}

	return false
}

// SetMasterVip gets a reference to the given string and assigns it to the MasterVip field.
func (o *KubernetesClusterManagementConfigAllOf) SetMasterVip(v string) {
	o.MasterVip = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterManagementConfigAllOf) GetSshKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterManagementConfigAllOf) GetSshKeysOk() (*[]string, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return &o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfigAllOf) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *KubernetesClusterManagementConfigAllOf) SetSshKeys(v []string) {
	o.SshKeys = v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfigAllOf) GetSshUser() string {
	if o == nil || o.SshUser == nil {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfigAllOf) GetSshUserOk() (*string, bool) {
	if o == nil || o.SshUser == nil {
		return nil, false
	}
	return o.SshUser, true
}

// HasSshUser returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfigAllOf) HasSshUser() bool {
	if o != nil && o.SshUser != nil {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *KubernetesClusterManagementConfigAllOf) SetSshUser(v string) {
	o.SshUser = &v
}

func (o KubernetesClusterManagementConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EncryptedEtcd != nil {
		toSerialize["EncryptedEtcd"] = o.EncryptedEtcd
	}
	if o.InternalMgmtPrivateKey != nil {
		toSerialize["InternalMgmtPrivateKey"] = o.InternalMgmtPrivateKey
	}
	if o.InternalMgmtPublicKey != nil {
		toSerialize["InternalMgmtPublicKey"] = o.InternalMgmtPublicKey
	}
	if o.LoadBalancerCount != nil {
		toSerialize["LoadBalancerCount"] = o.LoadBalancerCount
	}
	if o.LoadBalancers != nil {
		toSerialize["LoadBalancers"] = o.LoadBalancers
	}
	if o.MasterVip != nil {
		toSerialize["MasterVip"] = o.MasterVip
	}
	if o.SshKeys != nil {
		toSerialize["SshKeys"] = o.SshKeys
	}
	if o.SshUser != nil {
		toSerialize["SshUser"] = o.SshUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterManagementConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesClusterManagementConfigAllOf := _KubernetesClusterManagementConfigAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesClusterManagementConfigAllOf); err == nil {
		*o = KubernetesClusterManagementConfigAllOf(varKubernetesClusterManagementConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EncryptedEtcd")
		delete(additionalProperties, "InternalMgmtPrivateKey")
		delete(additionalProperties, "InternalMgmtPublicKey")
		delete(additionalProperties, "LoadBalancerCount")
		delete(additionalProperties, "LoadBalancers")
		delete(additionalProperties, "MasterVip")
		delete(additionalProperties, "SshKeys")
		delete(additionalProperties, "SshUser")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesClusterManagementConfigAllOf struct {
	value *KubernetesClusterManagementConfigAllOf
	isSet bool
}

func (v NullableKubernetesClusterManagementConfigAllOf) Get() *KubernetesClusterManagementConfigAllOf {
	return v.value
}

func (v *NullableKubernetesClusterManagementConfigAllOf) Set(val *KubernetesClusterManagementConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterManagementConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterManagementConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterManagementConfigAllOf(val *KubernetesClusterManagementConfigAllOf) *NullableKubernetesClusterManagementConfigAllOf {
	return &NullableKubernetesClusterManagementConfigAllOf{value: val, isSet: true}
}

func (v NullableKubernetesClusterManagementConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterManagementConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
