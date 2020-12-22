---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_node_group_profile"
description: |-
  A configuration profile for a node group in a Kubernetes cluster.
---

# Data Source: intersight_kubernetes_node_group_profile
A configuration profile for a node group in a Kubernetes cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `currentsize`:(int) Current number of nodes in this node group at any given point in time. 
* `description`:(string) Description of the profile. 
* `desiredsize`:(int) Desired number of nodes in this node group, same as minsize initially and is updated by the auto-scaler. 
* `maxsize`:(int) Maximum number of nodes desired in this node group. 
* `minsize`:(int) Minimum number of nodes desired in this node group. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `node_type`:(string) The node type Master, Worker or EmbeddedMaster.* `Worker` - Node will be marked as a worker node.* `Master` - Node will be marked as a master node.* `EmbeddedMaster` - Node will be both a master and a worker. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
