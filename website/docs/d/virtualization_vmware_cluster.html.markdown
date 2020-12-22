---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_cluster"
description: |-
  A real cluster of resources within a data center in VMware. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
---

# Data Source: intersight_virtualization_vmware_cluster
A real cluster of resources within a data center in VMware. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `datastore_count`:(int) Count of all datastores associated with this cluster. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user-provided name for this cluster to facilitate identification. 
* `status`:(string) Cluster health status as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `total_cores`:(int) Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster. 
