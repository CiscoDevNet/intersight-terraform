---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_health_check_execution"
description: |-
  Health check execution result for a health check definition on a HyperFlex device.
---

# Data Source: intersight_hyperflex_health_check_execution
Health check execution result for a health check definition on a HyperFlex device.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `category`:(string) Category that the HyperFlex health check Definition belongs to. 
* `cause`:(string) Information detailing the possible cause of the healthcheck failure, if the check fails. 
* `completion_time`:(string) Health check execution completion time. 
* `health_check_details`:(string) Details of the health check execution result. 
* `health_check_execution_error_details`:(string) Error details of a script execution failure. 
* `health_check_execution_error_summary`:(string) Error summary of a script execution failure. 
* `health_check_execution_status`:(string) Status of the health check execution.* `UNKNOWN` - Indicates that the health heck execution results are unknown.* `SUCCEEDED` - Indicates that the health check execution succeeded.* `FAILED` - Indicates that the health check execution failed.* `TIMED_OUT` - Indicates that the health check execution timed out before completion. 
* `health_check_result`:(string) Health check execution result. Valid only if HealthCheckExecutionStatus is SUCCEEDED.* `UNKNOWN` - Indicates that the health check results could not be determined.* `PASS` - Indicates that the health check passed.* `FAIL` - Indicates that the health check failed.* `WARN` - Indicates that the health check completed with a warning.* `NOT_APPLICABLE` - Indicates that the health check is either unsupported, or not applicable on the Cluster. 
* `health_check_summary`:(string) A brief summary of health check results. 
* `hx_device_name`:(string) HyperFlex Device Name where the healthcheck is executed. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `uuid`:(string) UUID of an instance of health check execution. 
