---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_lun"
description: |-
  NetApp LUN (logical unit number) is an identifier for a device called a logical unit addressed by a SAN protocol.
---

# Data Source: intersight_storage_net_app_lun
NetApp LUN (logical unit number) is an identifier for a device called a logical unit addressed by a SAN protocol.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Short description about the volume. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `naa_id`:(string) NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. 
* `name`:(string) Named entity of the volume. 
* `os_type`:(string) The operating system (OS) type for this LUN.* `Linux` - Family of open source Unix-like operating systems based on the Linux kernel.* `AIX` - Advanced Interactive Executive (AIX).* `HP-UX` - HP-UX is implementation of the Unix operating system, based on Unix System V.* `Hyper-V` - Windows Server 2008 or Windows Server 2012 Hyper-V.* `OpenVMS` - OpenVMS is multi-user, multiprocessing virtual memory-based operating system.* `Solaris` - Solaris is a Unix operating system.* `NetWare` - NetWare is a computer network operating system.* `VMware` - An enterprise-class, type-1 hypervisor developed by VMware for deploying and serving virtual computers.* `Windows` - Single-partition Windows disk using the Master Boot Record (MBR) partitioning style.* `Xen` - Xen is a type-1 hypervisor, providing services that allow multiple computer operating systems to execute on the same computer hardware concurrently. 
* `serial`:(string) Serial number for the provisioned LUN. 
* `size`:(int) User provisioned volume size. It is the size exposed to host. 
* `state`:(string) The administrative state of a LUN.* `offline` - The LUN is administratively offline, or a more detailed offline reason is not available.* `online` - The LUN is online. 
* `uuid`:(string) UUID of the LUN. 
