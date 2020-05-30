---
layout: "mxroute"
page_title: "MXroute: mxroute_domain"
sidebar_current: "docs-mxroute-resource-domain"
description: |-
  The mxroute_domain resource allows a MXroute domain to be added.
---

# mxroute\_domain

The domain resource allows a domain to be connected to MXroute server.

## Example Usage

```hcl
resource "mxroute_domain" "domain_com" {
  name = "domain.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A domain to be added to MXroute dashboard

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `dkim` - DKIM value that should be added to DNS record ([MXroute documentation](https://mxroutehelp.com/index.php/2019/08/25/set-up-dkim/))

## Import

Existing domains can be imported using their name

```shell
$ terraform import mxroute_domain.domain_com domain.com
```