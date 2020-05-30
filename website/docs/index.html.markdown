---
layout: "mxroute"
page_title: "Provider: MXroute"
sidebar_current: "docs-mxroute-index"
description: |-
  The MXroute provider is used to interact with resources supported by MXroute. The provider needs to be configured with the proper credentials before it can be used.
---

# MXroute Provider

The MXroute provider is used to interact with resources supported by
MXroute (DIRECTADMIN ONLY). The provider needs to be configured with the proper credentials
before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the MXroute provider.
provider "mxroute" {
  username = "${var.mxroute_username}"
  password = "${var.mxroute_password}"
}

# Connect domain
resource "mxroute_domain" "email" {
  # ...
}

# Create email account
resource "mxroute_email_account" "email" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Optional) The email associated with the DirectAdmin account (NOT MXROUTE ACCOUNT). This can also be
  specified with the `MXROUTE_USERNAME` shell environment variable.
* `password` - (Optional) The DirectAdmin password or API key. This can also be specified
  with the `MXROUTE_PASSWORD` shell environment variable.