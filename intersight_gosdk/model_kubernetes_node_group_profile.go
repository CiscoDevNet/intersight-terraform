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

// KubernetesNodeGroupProfile A configuration profile for a node group in a Kubernetes cluster.
type KubernetesNodeGroupProfile struct {
	PolicyAbstractConfigProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Current number of nodes in this node group at any given point in time.
	Currentsize *int64 `json:"Currentsize,omitempty"`
	// Desired number of nodes in this node group, same as minsize initially and is updated by the auto-scaler.
	Desiredsize *int64                     `json:"Desiredsize,omitempty"`
	Labels      []KubernetesNodeGroupLabel `json:"Labels,omitempty"`
	// Maximum number of nodes desired in this node group.
	Maxsize *int64 `json:"Maxsize,omitempty"`
	// Minimum number of nodes desired in this node group.
	Minsize *int64 `json:"Minsize,omitempty"`
	// The node type Master, Worker or EmbeddedMaster. * `Worker` - Node will be marked as a worker node. * `Master` - Node will be marked as a master node. * `EmbeddedMaster` - Node will be both a master and a worker.
	NodeType       *string                                       `json:"NodeType,omitempty"`
	Taints         []KubernetesNodeGroupTaint                    `json:"Taints,omitempty"`
	ClusterProfile *KubernetesClusterProfileRelationship         `json:"ClusterProfile,omitempty"`
	InfraProvider  *KubernetesInfrastructureProviderRelationship `json:"InfraProvider,omitempty"`
	// An array of relationships to ippoolPool resources.
	IpPools           []IppoolPoolRelationship             `json:"IpPools,omitempty"`
	KubernetesVersion *KubernetesVersionPolicyRelationship `json:"KubernetesVersion,omitempty"`
	// An array of relationships to kubernetesNodeProfile resources.
	Nodes                []KubernetesNodeProfileRelationship `json:"Nodes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeGroupProfile KubernetesNodeGroupProfile

// NewKubernetesNodeGroupProfile instantiates a new KubernetesNodeGroupProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeGroupProfile(classId string, objectType string) *KubernetesNodeGroupProfile {
	this := KubernetesNodeGroupProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var desiredsize int64 = 3
	this.Desiredsize = &desiredsize
	var nodeType string = "Worker"
	this.NodeType = &nodeType
	return &this
}

// NewKubernetesNodeGroupProfileWithDefaults instantiates a new KubernetesNodeGroupProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeGroupProfileWithDefaults() *KubernetesNodeGroupProfile {
	this := KubernetesNodeGroupProfile{}
	var classId string = "kubernetes.NodeGroupProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.NodeGroupProfile"
	this.ObjectType = objectType
	var desiredsize int64 = 3
	this.Desiredsize = &desiredsize
	var nodeType string = "Worker"
	this.NodeType = &nodeType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeGroupProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeGroupProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeGroupProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeGroupProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrentsize returns the Currentsize field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfile) GetCurrentsize() int64 {
	if o == nil || o.Currentsize == nil {
		var ret int64
		return ret
	}
	return *o.Currentsize
}

// GetCurrentsizeOk returns a tuple with the Currentsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetCurrentsizeOk() (*int64, bool) {
	if o == nil || o.Currentsize == nil {
		return nil, false
	}
	return o.Currentsize, true
}

// HasCurrentsize returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasCurrentsize() bool {
	if o != nil && o.Currentsize != nil {
		return true
	}

	return false
}

// SetCurrentsize gets a reference to the given int64 and assigns it to the Currentsize field.
func (o *KubernetesNodeGroupProfile) SetCurrentsize(v int64) {
	o.Currentsize = &v
}

// GetDesiredsize returns the Desiredsize field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfile) GetDesiredsize() int64 {
	if o == nil || o.Desiredsize == nil {
		var ret int64
		return ret
	}
	return *o.Desiredsize
}

// GetDesiredsizeOk returns a tuple with the Desiredsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetDesiredsizeOk() (*int64, bool) {
	if o == nil || o.Desiredsize == nil {
		return nil, false
	}
	return o.Desiredsize, true
}

// HasDesiredsize returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasDesiredsize() bool {
	if o != nil && o.Desiredsize != nil {
		return true
	}

	return false
}

// SetDesiredsize gets a reference to the given int64 and assigns it to the Desiredsize field.
func (o *KubernetesNodeGroupProfile) SetDesiredsize(v int64) {
	o.Desiredsize = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeGroupProfile) GetLabels() []KubernetesNodeGroupLabel {
	if o == nil {
		var ret []KubernetesNodeGroupLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeGroupProfile) GetLabelsOk() (*[]KubernetesNodeGroupLabel, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []KubernetesNodeGroupLabel and assigns it to the Labels field.
func (o *KubernetesNodeGroupProfile) SetLabels(v []KubernetesNodeGroupLabel) {
	o.Labels = v
}

// GetMaxsize returns the Maxsize field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfile) GetMaxsize() int64 {
	if o == nil || o.Maxsize == nil {
		var ret int64
		return ret
	}
	return *o.Maxsize
}

// GetMaxsizeOk returns a tuple with the Maxsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetMaxsizeOk() (*int64, bool) {
	if o == nil || o.Maxsize == nil {
		return nil, false
	}
	return o.Maxsize, true
}

// HasMaxsize returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasMaxsize() bool {
	if o != nil && o.Maxsize != nil {
		return true
	}

	return false
}

// SetMaxsize gets a reference to the given int64 and assigns it to the Maxsize field.
func (o *KubernetesNodeGroupProfile) SetMaxsize(v int64) {
	o.Maxsize = &v
}

// GetMinsize returns the Minsize field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfile) GetMinsize() int64 {
	if o == nil || o.Minsize == nil {
		var ret int64
		return ret
	}
	return *o.Minsize
}

// GetMinsizeOk returns a tuple with the Minsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetMinsizeOk() (*int64, bool) {
	if o == nil || o.Minsize == nil {
		return nil, false
	}
	return o.Minsize, true
}

// HasMinsize returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasMinsize() bool {
	if o != nil && o.Minsize != nil {
		return true
	}

	return false
}

// SetMinsize gets a reference to the given int64 and assigns it to the Minsize field.
func (o *KubernetesNodeGroupProfile) SetMinsize(v int64) {
	o.Minsize = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfile) GetNodeType() string {
	if o == nil || o.NodeType == nil {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetNodeTypeOk() (*string, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasNodeType() bool {
	if o != nil && o.NodeType != nil {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *KubernetesNodeGroupProfile) SetNodeType(v string) {
	o.NodeType = &v
}

// GetTaints returns the Taints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeGroupProfile) GetTaints() []KubernetesNodeGroupTaint {
	if o == nil {
		var ret []KubernetesNodeGroupTaint
		return ret
	}
	return o.Taints
}

// GetTaintsOk returns a tuple with the Taints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeGroupProfile) GetTaintsOk() (*[]KubernetesNodeGroupTaint, bool) {
	if o == nil || o.Taints == nil {
		return nil, false
	}
	return &o.Taints, true
}

// HasTaints returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasTaints() bool {
	if o != nil && o.Taints != nil {
		return true
	}

	return false
}

// SetTaints gets a reference to the given []KubernetesNodeGroupTaint and assigns it to the Taints field.
func (o *KubernetesNodeGroupProfile) SetTaints(v []KubernetesNodeGroupTaint) {
	o.Taints = v
}

// GetClusterProfile returns the ClusterProfile field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfile) GetClusterProfile() KubernetesClusterProfileRelationship {
	if o == nil || o.ClusterProfile == nil {
		var ret KubernetesClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfile
}

// GetClusterProfileOk returns a tuple with the ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetClusterProfileOk() (*KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfile == nil {
		return nil, false
	}
	return o.ClusterProfile, true
}

// HasClusterProfile returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasClusterProfile() bool {
	if o != nil && o.ClusterProfile != nil {
		return true
	}

	return false
}

// SetClusterProfile gets a reference to the given KubernetesClusterProfileRelationship and assigns it to the ClusterProfile field.
func (o *KubernetesNodeGroupProfile) SetClusterProfile(v KubernetesClusterProfileRelationship) {
	o.ClusterProfile = &v
}

// GetInfraProvider returns the InfraProvider field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfile) GetInfraProvider() KubernetesInfrastructureProviderRelationship {
	if o == nil || o.InfraProvider == nil {
		var ret KubernetesInfrastructureProviderRelationship
		return ret
	}
	return *o.InfraProvider
}

// GetInfraProviderOk returns a tuple with the InfraProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetInfraProviderOk() (*KubernetesInfrastructureProviderRelationship, bool) {
	if o == nil || o.InfraProvider == nil {
		return nil, false
	}
	return o.InfraProvider, true
}

// HasInfraProvider returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasInfraProvider() bool {
	if o != nil && o.InfraProvider != nil {
		return true
	}

	return false
}

// SetInfraProvider gets a reference to the given KubernetesInfrastructureProviderRelationship and assigns it to the InfraProvider field.
func (o *KubernetesNodeGroupProfile) SetInfraProvider(v KubernetesInfrastructureProviderRelationship) {
	o.InfraProvider = &v
}

// GetIpPools returns the IpPools field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeGroupProfile) GetIpPools() []IppoolPoolRelationship {
	if o == nil {
		var ret []IppoolPoolRelationship
		return ret
	}
	return o.IpPools
}

// GetIpPoolsOk returns a tuple with the IpPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeGroupProfile) GetIpPoolsOk() (*[]IppoolPoolRelationship, bool) {
	if o == nil || o.IpPools == nil {
		return nil, false
	}
	return &o.IpPools, true
}

// HasIpPools returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasIpPools() bool {
	if o != nil && o.IpPools != nil {
		return true
	}

	return false
}

// SetIpPools gets a reference to the given []IppoolPoolRelationship and assigns it to the IpPools field.
func (o *KubernetesNodeGroupProfile) SetIpPools(v []IppoolPoolRelationship) {
	o.IpPools = v
}

// GetKubernetesVersion returns the KubernetesVersion field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfile) GetKubernetesVersion() KubernetesVersionPolicyRelationship {
	if o == nil || o.KubernetesVersion == nil {
		var ret KubernetesVersionPolicyRelationship
		return ret
	}
	return *o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfile) GetKubernetesVersionOk() (*KubernetesVersionPolicyRelationship, bool) {
	if o == nil || o.KubernetesVersion == nil {
		return nil, false
	}
	return o.KubernetesVersion, true
}

// HasKubernetesVersion returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasKubernetesVersion() bool {
	if o != nil && o.KubernetesVersion != nil {
		return true
	}

	return false
}

// SetKubernetesVersion gets a reference to the given KubernetesVersionPolicyRelationship and assigns it to the KubernetesVersion field.
func (o *KubernetesNodeGroupProfile) SetKubernetesVersion(v KubernetesVersionPolicyRelationship) {
	o.KubernetesVersion = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeGroupProfile) GetNodes() []KubernetesNodeProfileRelationship {
	if o == nil {
		var ret []KubernetesNodeProfileRelationship
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeGroupProfile) GetNodesOk() (*[]KubernetesNodeProfileRelationship, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfile) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []KubernetesNodeProfileRelationship and assigns it to the Nodes field.
func (o *KubernetesNodeGroupProfile) SetNodes(v []KubernetesNodeProfileRelationship) {
	o.Nodes = v
}

func (o KubernetesNodeGroupProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigProfile, errPolicyAbstractConfigProfile := json.Marshal(o.PolicyAbstractConfigProfile)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	errPolicyAbstractConfigProfile = json.Unmarshal([]byte(serializedPolicyAbstractConfigProfile), &toSerialize)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Currentsize != nil {
		toSerialize["Currentsize"] = o.Currentsize
	}
	if o.Desiredsize != nil {
		toSerialize["Desiredsize"] = o.Desiredsize
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.Maxsize != nil {
		toSerialize["Maxsize"] = o.Maxsize
	}
	if o.Minsize != nil {
		toSerialize["Minsize"] = o.Minsize
	}
	if o.NodeType != nil {
		toSerialize["NodeType"] = o.NodeType
	}
	if o.Taints != nil {
		toSerialize["Taints"] = o.Taints
	}
	if o.ClusterProfile != nil {
		toSerialize["ClusterProfile"] = o.ClusterProfile
	}
	if o.InfraProvider != nil {
		toSerialize["InfraProvider"] = o.InfraProvider
	}
	if o.IpPools != nil {
		toSerialize["IpPools"] = o.IpPools
	}
	if o.KubernetesVersion != nil {
		toSerialize["KubernetesVersion"] = o.KubernetesVersion
	}
	if o.Nodes != nil {
		toSerialize["Nodes"] = o.Nodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNodeGroupProfile) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesNodeGroupProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Current number of nodes in this node group at any given point in time.
		Currentsize *int64 `json:"Currentsize,omitempty"`
		// Desired number of nodes in this node group, same as minsize initially and is updated by the auto-scaler.
		Desiredsize *int64                     `json:"Desiredsize,omitempty"`
		Labels      []KubernetesNodeGroupLabel `json:"Labels,omitempty"`
		// Maximum number of nodes desired in this node group.
		Maxsize *int64 `json:"Maxsize,omitempty"`
		// Minimum number of nodes desired in this node group.
		Minsize *int64 `json:"Minsize,omitempty"`
		// The node type Master, Worker or EmbeddedMaster. * `Worker` - Node will be marked as a worker node. * `Master` - Node will be marked as a master node. * `EmbeddedMaster` - Node will be both a master and a worker.
		NodeType       *string                                       `json:"NodeType,omitempty"`
		Taints         []KubernetesNodeGroupTaint                    `json:"Taints,omitempty"`
		ClusterProfile *KubernetesClusterProfileRelationship         `json:"ClusterProfile,omitempty"`
		InfraProvider  *KubernetesInfrastructureProviderRelationship `json:"InfraProvider,omitempty"`
		// An array of relationships to ippoolPool resources.
		IpPools           []IppoolPoolRelationship             `json:"IpPools,omitempty"`
		KubernetesVersion *KubernetesVersionPolicyRelationship `json:"KubernetesVersion,omitempty"`
		// An array of relationships to kubernetesNodeProfile resources.
		Nodes []KubernetesNodeProfileRelationship `json:"Nodes,omitempty"`
	}

	varKubernetesNodeGroupProfileWithoutEmbeddedStruct := KubernetesNodeGroupProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesNodeGroupProfileWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNodeGroupProfile := _KubernetesNodeGroupProfile{}
		varKubernetesNodeGroupProfile.ClassId = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.ClassId
		varKubernetesNodeGroupProfile.ObjectType = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.ObjectType
		varKubernetesNodeGroupProfile.Currentsize = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.Currentsize
		varKubernetesNodeGroupProfile.Desiredsize = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.Desiredsize
		varKubernetesNodeGroupProfile.Labels = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.Labels
		varKubernetesNodeGroupProfile.Maxsize = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.Maxsize
		varKubernetesNodeGroupProfile.Minsize = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.Minsize
		varKubernetesNodeGroupProfile.NodeType = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.NodeType
		varKubernetesNodeGroupProfile.Taints = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.Taints
		varKubernetesNodeGroupProfile.ClusterProfile = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.ClusterProfile
		varKubernetesNodeGroupProfile.InfraProvider = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.InfraProvider
		varKubernetesNodeGroupProfile.IpPools = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.IpPools
		varKubernetesNodeGroupProfile.KubernetesVersion = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.KubernetesVersion
		varKubernetesNodeGroupProfile.Nodes = varKubernetesNodeGroupProfileWithoutEmbeddedStruct.Nodes
		*o = KubernetesNodeGroupProfile(varKubernetesNodeGroupProfile)
	} else {
		return err
	}

	varKubernetesNodeGroupProfile := _KubernetesNodeGroupProfile{}

	err = json.Unmarshal(bytes, &varKubernetesNodeGroupProfile)
	if err == nil {
		o.PolicyAbstractConfigProfile = varKubernetesNodeGroupProfile.PolicyAbstractConfigProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Currentsize")
		delete(additionalProperties, "Desiredsize")
		delete(additionalProperties, "Labels")
		delete(additionalProperties, "Maxsize")
		delete(additionalProperties, "Minsize")
		delete(additionalProperties, "NodeType")
		delete(additionalProperties, "Taints")
		delete(additionalProperties, "ClusterProfile")
		delete(additionalProperties, "InfraProvider")
		delete(additionalProperties, "IpPools")
		delete(additionalProperties, "KubernetesVersion")
		delete(additionalProperties, "Nodes")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigProfile := reflect.ValueOf(o.PolicyAbstractConfigProfile)
		for i := 0; i < reflectPolicyAbstractConfigProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigProfile.Type().Field(i)

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

type NullableKubernetesNodeGroupProfile struct {
	value *KubernetesNodeGroupProfile
	isSet bool
}

func (v NullableKubernetesNodeGroupProfile) Get() *KubernetesNodeGroupProfile {
	return v.value
}

func (v *NullableKubernetesNodeGroupProfile) Set(val *KubernetesNodeGroupProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeGroupProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeGroupProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeGroupProfile(val *KubernetesNodeGroupProfile) *NullableKubernetesNodeGroupProfile {
	return &NullableKubernetesNodeGroupProfile{value: val, isSet: true}
}

func (v NullableKubernetesNodeGroupProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeGroupProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
