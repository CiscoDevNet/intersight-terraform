---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_cluster_profile"
description: |-
  Cluster profile specifies the config profile for a Kubernetes cluster. It also depicts operations to control the life cycle of a Kubernetes cluster.
---

# Resource: intersight_kubernetes_cluster_profile
Cluster profile specifies the config profile for a Kubernetes cluster. It also depicts operations to control the life cycle of a Kubernetes cluster.
## Argument Reference
The following arguments are supported:
* `aci_cni_profile`:(HashMap) -(Computed) A reference to a kubernetesAciCniProfile resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `action_info`:(HashMap) -(Computed) Action details such as name of the action performed, status, failure reason, etc. 
This complex property has following sub-properties:
  + `failure_reason`:(string)(Computed) Description of failure i.e. derived from the workflow failure message. 
  + `name`:(string)(Computed) Name of the Action performed on a resource like VM, Disk etc. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `status`:(string)(Computed) Status of the Action like InProgress, Success, Failure etc.* `None` - A place holder for the default value.* `InProgress` - Action triggered on the resource is still running.* `Success` - Action triggered on the resource is completed successfully.* `Failure` - Action triggered on the resource is failed. 
* `addons`:(Array) An array of relationships to kubernetesAddonPolicy resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `associated_cluster`:(HashMap) -(Computed) A reference to a kubernetesCluster resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `cert_config`:(HashMap) - Certificates and keys that are used to configure a Kubernetescluster. If user does not specify any Certificates or Keys,system generated certificates will be used to configure the cluster. 
This complex property has following sub-properties:
  + `ca_cert`:(string) Certificate for the Kubernetes API server. 
  + `ca_key`:(string) Private Key for the Kubernetes API server. 
  + `etcd_cert`:(string) Certificate for the etcd cluster. 
  + `etcd_encryption_key`:
                (Array of schema.TypeString) -
  + `etcd_key`:(string) Private key for the etcd cluster. 
  + `front_proxy_cert`:(string) Certificate for the front proxy to support Kubernetes API aggregators. 
  + `front_proxy_key`:(string) Private key for the front proxy to support Kubernetes API aggregators. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `sa_private_key`:(string) Service account private key used by Kubernetes Token Controller to sign generated service account tokens. 
  + `sa_public_key`:(string) Service account public key used by Kubernetes API Server to validate service account tokens generated by the Token Controller. 
* `cluster_ip_pools`:(Array) An array of relationships to ippoolPool resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `config_context`:(HashMap) - The configuration state and results of the last configuration operation. 
This complex property has following sub-properties:
  + `config_state`:(string)(Computed) Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Validating, Configuring, Failed. 
  + `control_action`:(string) System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind. 
  + `error_state`:(string) Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error). 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `oper_state`:(string)(Computed) Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed. 
* `container_runtime_config`:(HashMap) - A reference to a kubernetesContainerRuntimePolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) Description of the profile. 
* `kube_config`:(HashMap) -(Computed) Certificates and keys that are used to access a Kubernetescluster. 
This complex property has following sub-properties:
  + `kube_config`:(string)(Computed) Kubernetes configuration to connect to the cluster as an system administrator. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `loadbalancer_ip_leases`:(Array) An array of relationships to ippoolIpLease resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `managed_mode`:(string) Management mode for the cluster. In some cases Intersight kubernetes service is not requiredto provision and manage the management entities and endpoints (for e.g. EKS). In most other casesit will be required to provision and manage these entities and endpoints.* `Provided` - Cluster management entities and endpoints are provided by the infrastructure platform.* `Managed` - Cluster management entities and endpoints are provisioned and managed by Intersight kubernetes service. 
* `management_config`:(HashMap) - Configuration required for provisioning and management of cluster management entities. Required if'managedMode' is set to 'Managed'. 
This complex property has following sub-properties:
  + `encrypted_etcd`:(bool) Encrypt ETCD data at rest using the etcdEncryptionKey specified in the cluster Kubernetes configuration. 
  + `internal_mgmt_public_key`:(string) Public key for internal mgmt of the cluster via SSH. 
  + `load_balancer_count`:(int) Number of IP addresses to reserve for load balancer services. 
  + `load_balancers`:
                (Array of schema.TypeString) -
  + `master_vip`:(string) VIP for the cluster Kubernetes API server. If this is empty and a cluster IP pool is specified, it will be allocated from the IP pool. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ssh_keys`:
                (Array of schema.TypeString) -
  + `ssh_user`:(string)(Computed) Name of the user to SSH to nodes in a cluster. 
* `master_vip_lease`:(HashMap) - A reference to a ippoolIpLease resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `net_config`:(HashMap) - A reference to a kubernetesNetworkPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `node_groups`:(Array) An array of relationships to kubernetesNodeGroupProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `src_template`:(HashMap) - A reference to a policyAbstractProfile resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `status`:(string) Status of the Kubernetes cluster and its nodes.* `Configuring` - The cluster is being configured.* `Deploying` - The cluster is being deployed.* `Undeploying` - The cluster is being undeployed.* `DeployFailed` - The cluster deployment failed.* `Upgrading` - The cluster is being upgraded.* `Deleting` - The cluster is being deleted.* `DeleteFailed` - The cluster delete failed.* `Ready` - The cluster is ready for use.* `Active` - The cluster is being active.* `Shutdown` - All the nodes in the cluster are powered off.* `Terminated` - The cluster is terminated.* `Deployed` - The cluster is deployed. The cluster may not yet be ready for use.* `UndeployFailed` - The cluster undeploy action failed.* `NotReady` - The cluster is created and some nodes are not ready. 
* `sys_config`:(HashMap) - A reference to a kubernetesSysConfigPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `trusted_registries`:(HashMap) - A reference to a kubernetesTrustedRegistriesPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `workflow_info`:(HashMap) -(Computed) A reference to a workflowWorkflowInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 


## Import
`intersight_kubernetes_cluster_profile` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_kubernetes_cluster_profile.example 1234567890987654321abcde
``` 