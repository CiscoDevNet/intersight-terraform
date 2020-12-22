---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_chassis_identity"
description: |-
  ChassisIdentity Object conatains connectivity information about IOMs of the chassis. ChassisID is uniquely allocated for the combination of vendor, model and serial number of the chassis.
---

# Data Source: intersight_equipment_chassis_identity
ChassisIdentity Object conatains connectivity information about IOMs of the chassis. ChassisID is uniquely allocated for the combination of vendor, model and serial number of the chassis.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_action`:(string) Updated by UI/API to trigger specific chassis action type.* `None` - No operation value for maintenance actions on an equipment.* `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight.* `Recommission` - Recommission the equipment.* `Reack` - Reacknowledge the equipment and discover it again.* `Remove` - Remove the equipment permanently from Intersight management. 
* `admin_action_state`:(string) The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `identifier`:(int) Numeric Identifier assigned by the management system to the equipment. 
* `nr_lifecycle`:(string) The equipment's lifecycle status.* `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment.* `Active` - Default Lifecycle State for a physical entity.* `Decommissioned` - Decommission Lifecycle state.* `DecommissionInProgress` - Decommission Inprogress Lifecycle state.* `RecommissionInProgress` - Recommission Inprogress Lifecycle state.* `OperationFailed` - Failed Operation Lifecycle state.* `ReackInProgress` - ReackInProgress Lifecycle state.* `RemoveInProgress` - RemoveInProgress Lifecycle state.* `Discovered` - Discovered Lifecycle state.* `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state.* `DiscoveryFailed` - DiscoveryFailed Lifecycle state.* `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity. 
* `model`:(string) The vendor provided model name for the equipment. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `serial`:(string) The serial number of the equipment. 
* `vendor`:(string) The manufacturer of the equipment. 
