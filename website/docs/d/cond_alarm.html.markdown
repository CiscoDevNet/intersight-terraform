---
subcategory: "cond"
layout: "intersight"
page_title: "Intersight: intersight_cond_alarm"
description: |-
  A state-full entity representing a found problem. Alarms can be reported by the managed system itself or can be determined by Intersight.
---

# Data Source: intersight_cond_alarm
A state-full entity representing a found problem. Alarms can be reported by the managed system itself or can be determined by Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `acknowledge`:(string) Alarm acknowledgment state. Default value is None.* `None` - The Enum value None represents that the alarm is not acknowledged and is included as part of health status and overall alarm count.* `Acknowledge` - The Enum value Acknowledge represents that the alarm is acknowledged by user. The alarm will be ignored from the health status and overall alarm count. 
* `acknowledge_by`:(string) User who acknowledged the alarm. 
* `acknowledge_time`:(string) Time at which the alarm was acknowledged by the user. 
* `affected_mo_display_name`:(string) Display name of the affected object on which the alarm is raised. The name uniquely identifies an alarm target such that it can be distinguished from similar objects managed by Intersight. 
* `affected_mo_id`:(string) MoId of the affected object from the managed system's point of view. 
* `affected_mo_type`:(string) Managed system affected object type. For example Adaptor, FI, CIMC. 
* `ancestor_mo_id`:(string) Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault. 
* `ancestor_mo_type`:(string) Parent MO type of the fault from managed system. For example, Blade for adaptor fault. 
* `code`:(string) A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code. 
* `creation_time`:(string) The time the alarm was created. 
* `description`:(string) A longer description of the alarm than the name. The description contains details of the component reporting the issue. 
* `last_transition_time`:(string) The time the alarm last had a change in severity. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `ms_affected_object`:(string) A unique key for the alarm from the managed system's point of view. For example, in the case of UCS, this is the fault's dn. 
* `name`:(string) Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code. 
* `orig_severity`:(string) The original severity when the alarm was first created.* `None` - The Enum value None represents that there is no severity.* `Info` - The Enum value Info represents the Informational level of severity.* `Critical` - The Enum value Critical represents the Critical level of severity.* `Warning` - The Enum value Warning represents the Warning level of severity.* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared. 
* `severity`:(string) The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared.* `None` - The Enum value None represents that there is no severity.* `Info` - The Enum value Info represents the Informational level of severity.* `Critical` - The Enum value Critical represents the Critical level of severity.* `Warning` - The Enum value Warning represents the Warning level of severity.* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared. 
