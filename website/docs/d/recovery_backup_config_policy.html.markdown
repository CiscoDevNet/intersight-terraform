
---
layout: "intersight"
page_title: "Intersight: intersight_recovery_backup_config_policy"
sidebar_current: "docs-intersight-data-source-recovery-backup-config-policy"
description: |-
Backup config policy which contains all the required inputs to do backup on a local or remote server.
---

# Data Source: intersight_recovery._backup_config_policy
Backup config policy which contains all the required inputs to do backup on a local or remote server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the policy. 
* `file_name_prefix`:(string) The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `location_type`:(string) Specifies whether the backup will be stored locally or remotely.* `Network Share` - The backup is stored remotely on a separate server.* `Local Storage` - The backup is stored locally on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `password`:(string) Password of Backup server. 
* `path`:(string) The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/. 
* `protocol`:(string) Protocol for transferring the backup image to the network share location.* `SCP` - Secure Copy Protocol (SCP) to access the file server.* `SFTP` - SSH File Transfer Protocol (SFTP) to access file server.* `FTP` - File Transfer Protocol (FTP) to access file server. 
* `retention_count`:(int) Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner. 
* `user_name`:(string) Username for the backup server. 
