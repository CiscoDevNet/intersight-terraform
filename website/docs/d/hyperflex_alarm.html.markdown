---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_alarm"
description: |-
  An alarm representing a fault in the HyperFlex cluster configuration or hardware.
---

# Data Source: intersight_hyperflex_alarm
An alarm representing a fault in the HyperFlex cluster configuration or hardware.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `acknowledged`:(bool) The acknowledgement state of the alarm. It is 'true' when the alarm is acknowledged and false otherwise. 
* `acknowledged_by`:(string) The username of the user who acknowledged the alarm. 
* `acknowledged_time`:(int) The time when the alarm was acknowledged, represented as a Unix timestamp. 
* `acknowledged_time_as_utc`:(string) The time when the alarm was acknowledged in ISO 6801 format. 
* `description`:(string) The description of the alarm which includes information about the fault that triggered the alarm. 
* `entity_data`:(string) The data pertaining to this entity. 
* `entity_name`:(string) The name of entity which triggered the alarm. 
* `entity_type`:(string) The type of entity which triggered the alarm. For example, this can be the cluster, a node, or VM running on the cluster.* `UNKNOWN` - The type of entity is not known.* `DISK` - The entity is a physical storage device.* `NODE` - The entity is a HyperFlex cluster node.* `CLUSTER` - The entity is the HyperFlex cluster itself.* `DATASTORE` - The entity is a logical datastore configured on the HyperFlex cluster.* `ZONE` - The entity is a logical or physical zone configured on the HyperFlex cluster.* `VIRTUALMACHINE` - The entity is a virtual machine running on the HyperFlex cluster. 
* `entity_uu_id`:(string) The unique identifier of the entity which triggered the alarm. 
* `message`:(string) The localized message displayed to the user which describes the alarm. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the alarm. This name identifies the type of alarm that was triggered. 
* `status`:(string) The severity of the alarm.* `UNKNOWN` - The alarm status is not known.* `CLEARED` - The event that triggered the alarm has been remedied and no longer requires the user's attention.* `INFO` - The alarm represents a message that does not require the user's immediate attention.* `WARNING` - The alarm represents a moderate fault.* `CRITICAL` - The alarm represents a critical fault. 
* `triggered_time`:(int) The time when alarm was triggered as a Unix timestamp. 
* `triggered_time_as_utc`:(string) The time when alarm was triggered in ISO 6801 UTC format. 
* `uuid`:(string) The unique identifier for this alarm instance. 
