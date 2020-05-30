---
layout: "mxroute"
page_title: "MXroute: mxroute_email_account"
sidebar_current: "docs-mxroute-resource-email-account"
description: |-
  The mxroute_domain resource allows a MXroute domain to be added.
---

# mxroute\_email\_account

The email account resource allows an email account to be create on MXroute server.

## Example Usage

```hcl
resource "mxroute_domain" "email_at_domain_com" {
	domain = mxroute_domain.domain_com.name
	username = "email"
	password = "password"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Domain of email account (this domain must be added to MXroute account)
* `username` - (Required) Email account username
* `password` - (Required) Email account password

## Import

Existing email addresses can be imported using their full address

```shell
$ terraform import mxroute_domain.email_at_domain_com email@domain.com
```

NOTE: Because password of email account cannot be downloaded, when importing email account password in state file will be set to "" (empty string).