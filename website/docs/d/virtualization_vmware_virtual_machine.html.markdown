
---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_virtual_machine"
sidebar_current: "docs-intersight-data-source-virtualization-vmware-virtual-machine"
description: |-
The VMware Virtual machine. It has details such as power state, IP address, resource consumption, etc. Basic elements come from the base class and VMware specific details are provided here.
---

# Data Source: intersight_virtualization._vmware_virtual_machine
The VMware Virtual machine. It has details such as power state, IP address, resource consumption, etc. Basic elements come from the base class and VMware specific details are provided here.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `annotation`:(string) List of annotations provided to this VM by user. Can be long. 
* `boot_time`:(string) Time when this VM booted up. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `config_name`:(string) The configuration name for this VM. This maybe the same as the guest hostname. 
* `connection_state`:(string) Shows if virtual machine is connected to vCenter. Values are Connected, Disconnected, Orphaned, Inaccessible, and Invalid. 
* `cpu_hot_add_enabled`:(bool) Indicates if the capability to add CPUs to a running VM is enabled. 
* `default_power_off_type`:(string) Indicates how the VM will be powered off (soft, hard etc.). 
* `dhcp_enabled`:(bool) Shows if DHCP is used for IP/DNS on this VM. 
* `folder`:(string) The folder name associated with this VM. 
* `guest_state`:(string) The state of the guest OS running on this VM. Could be running, not running etc.* `Unknown` - Indicates that the guest OS state cannot be determined.* `NotRunning` - Indicates that the guest OS is not running.* `Resetting` - Indicates that the guest OS is resetting.* `Running` - Indicates that the guest OS is running normally.* `ShuttingDown` - Indicates that the guest OS is shutting down.* `Standby` - Indicates that the guest OS is in standby mode. 
* `hypervisor_type`:(string) Type of hypervisor where the virtual machine is hosted for example ESXi.* `ESXi` - A Vmware ESXi hypervisor of any version.* `HXAP` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference). 
* `instance_uuid`:(string) UUID assigned by vCenter to every VM. 
* `is_template`:(bool) If true, indicates that the entity refers to a template of a virtual machine and not a real virtual machine. 
* `memory_hot_add_enabled`:(bool) Adding memory to a running VM. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the virtual machine. 
* `network_count`:(int) Indicates how many networks are used by this VM. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `power_state`:(string) Power state of the virtual machine.* `Unknown` - The entity's power state is unknown.* `PoweredOn` - The entity is powered on.* `PoweredOff` - The entity is powered down.* `StandBy` - The entity is in standby mode.* `Paused` - The entity is in pause state. 
* `protected_vm`:(bool) Shows if this is a protected VM. VMs can be in protection groups. 
* `remote_display_vnc_enabled`:(bool) Shows if support for a remote VNC access is enabled. 
* `resource_pool`:(string) Name of the resource pool to which this VM belongs (optional). 
* `resource_pool_owner`:(string) Who owns the resource pool. 
* `resource_pool_parent`:(string) The parent of the current resource pool to which this VM belongs. 
* `tool_running_status`:(string) Indicates if guest tools are running on this VM. Could be set to guestToolNotRunning or guestToolsRunning. 
* `tools_version`:(string) The version of the guest tools, usually not specified. 
* `uuid`:(string) The uuid of this virtual machine. The uuid is internally generated and not user specified. 
* `vm_disk_count`:(int) Shows the number of disks assigned to this VM. 
* `vm_overall_status`:(string) The operational state of the VM. Could be Available, Provisioned, Maintenance mode, Deleting, etc. 
* `vm_path`:(string) Example - [datastore3] VCSA-134/VCSA-134.vmx. 
* `vm_version`:(string) Information about the version of this VM (vmx-09, vmx-11 etc.). 
* `vm_vnic_count`:(int) How many vnics are present. 
* `vnic_device_config_id`:(string) Information related to the guest info's VNIC virtual device. It is a comma-separated list. 
