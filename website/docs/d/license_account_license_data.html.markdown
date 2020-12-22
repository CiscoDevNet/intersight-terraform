---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_account_license_data"
description: |-
  License information for an account.
---

# Data Source: intersight_license_account_license_data
License information for an account.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_id`:(string) Root user's ID of the account. 
* `agent_data`:(string) Agent trusted store data. 
* `auth_expire_time`:(string) Authorization expiration time. 
* `auth_initial_time`:(string) Intial authorization time. 
* `auth_next_time`:(string) Next time for the authorization. 
* `category`:(string) Account license data category name. 
* `default_license_type`:(string) Default license tier set by user.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type. 
* `error_desc`:(string) The detailed error message when there is any error related to license sync of this account. 
* `group`:(string) Account license data group name. 
* `highest_compliant_license_tier`:(string) The highest license tier which is in compliant of this account.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type. 
* `last_sync`:(string) Specifies last sync time with SA. 
* `last_updated_time`:(string) Record's last update datetime. 
* `license_state`:(string) Aggregrated mode for the agent. 
* `license_tech_support_info`:(string) Tech-support info of a smart-agent. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `register_expire_time`:(string) Registration exipiration time. 
* `register_initial_time`:(string) Initial time of registration. 
* `register_next_time`:(string) Next time for the license registration. 
* `registration_status`:(string) Registration status of a smart-agent. 
* `renew_failure_string`:(string) License renewal failure message. 
* `smart_account`:(string) Name of the smart account. 
* `sync_status`:(string) Current sync status for the account. 
* `virtual_account`:(string) Name of the virtual account. 
