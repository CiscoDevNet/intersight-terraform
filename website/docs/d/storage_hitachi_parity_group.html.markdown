---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_parity_group"
description: |-
  A parity group in Hitachi storage array.
---

# Data Source: intersight_storage_hitachi_parity_group
A parity group in Hitachi storage array.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `disk_speed`:(string) Speed (rpm) of the disk belonging to the parity group. 
* `disk_type`:(string) Type of the disk belonging to the parity group. 
* `is_accelerated_compression_enabled`:(bool) Value of the accelerated compression of the parity group. true, Accelerated compression for the parity group is enabled. false, Accelerated compression for the parity group is disabled. 
* `is_copy_back_mode_enabled`:(bool) Value of the copy back mode setting of the parity group. true, Copy back mode is enabled. false, Copy back mode is disabled. 
* `is_encryption_enabled`:(bool) Value of the encryption setting of the parity group. true, Encryption is enabled. false, Encryption is disabled. 
* `level`:(string) The RAID level associated with the group of disks.* `Unknown` - Default unknown RAID type.* `RAID0` - RAID0 splits (\ stripes\ ) data evenly across two or more disks, without parity information.* `RAID1` - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover.* `RAID4` - RAID4 stripes block level data and dedicates a disk to parity.* `RAID5` - RAID5  distributes striping and parity at a block level.* `RAID6` - RAID6 level operates like RAID5 with distributed parity and striping.* `RAID10` - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy.* `RAIDDP` - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group.* `RAIDTEC` - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Human readable name of the RAID group. 
