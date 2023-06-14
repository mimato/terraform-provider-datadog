---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datadog_team_link Resource - terraform-provider-datadog"
subcategory: ""
description: |-
  Provides a Datadog TeamLink resource. This can be used to create and manage Datadog team_link.
---

# datadog_team_link (Resource)

Provides a Datadog TeamLink resource. This can be used to create and manage Datadog team_link.

## Example Usage

```terraform
resource "datadog_team" "foo" {
  description = "Example team"
  handle      = "example-team-updated"
  name        = "Example Team-updated"
}

# Create new team_link resource

resource "datadog_team_link" "foo" {
  team_id  = datadog_team.foo.id
  label    = "Link label"
  position = "Example link"
  url      = "https://example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `label` (String) The link's label.
- `team_id` (String) ID of the team the link is associated with.
- `url` (String) The URL for the link.

### Optional

- `position` (Number) The link's position, used to sort links for the team.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import datadog_team_link.new_list "${team_id}:${resource_id}"
```