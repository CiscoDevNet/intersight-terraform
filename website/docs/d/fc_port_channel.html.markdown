---
subcategory: "fc"
layout: "intersight"
page_title: "Intersight: intersight_fc_port_channel"
description: |-
  Model contains the details of the ethernet port-channels configured on the FI.
---

# Data Source: intersight_fc_port_channel
Model contains the details of the ethernet port-channels configured on the FI.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_speed`:(string) Administrator configured Speed applied on the port channel. 
* `admin_state`:(string) Administratively configured state (enabled/disabled) for this portchannel. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `mode`:(string) Mode information N_proxy, F or E associated to the Fibre Channel portchannel. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_speed`:(string) Operational speed of this port-channel. 
* `oper_state`:(string) Operational state of this port-channel. 
* `oper_state_qual`:(string) Reason for this port-channel's Operational state. 
* `port_channel_id`:(int) Unique identifier for this port-channel on the FI. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `role`:(string) This port-channel's configured role (fcUplink, fcStorage, etc.). 
* `switch_id`:(string) Switch Identifier that is local to a cluster. 
* `vsan`:(int) Virtual San that is associated to the port-channel. 
