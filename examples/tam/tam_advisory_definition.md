### Resource Creation

```hcl
resource "intersight_tam_advisory_definition" "tam_advisory_definition" {
  name           = "tam_advisory_definition"
  operation_type = "create"
  state          = "ready"
  severity = {
    object_type = "tam.SecurityAdvisoryDetails"
  }
  actions = {
    object_type = "tam.SecurityAdvisoryDetails"
    alert_type  = "psirt"
  }
  type = "securityAdvisory"
  api_data_sources {
    object_type = "tam.ApiDataSource"
    name        = "api_data_source_1"
    type        = "intersightApi"

  }
  advisory_details = {
    object_type = "tam.SecurityAdvisoryDetails"
    description = "tam security advisory"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```