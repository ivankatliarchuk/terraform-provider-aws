---
subcategory: "Shield"
layout: "aws"
page_title: "AWS: aws_shield_protection"
description: |-
  Terraform data source for managing an AWS Shield Protection.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_shield_protection

Terraform data source for managing an AWS Shield Protection.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_shield_protection import DataAwsShieldProtection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsShieldProtection(self, "example",
            protection_id="abc123"
        )
```

### By Resource ARN

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_shield_protection import DataAwsShieldProtection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsShieldProtection(self, "example",
            resource_arn="arn:aws:globalaccelerator::123456789012:accelerator/1234abcd-abcd-1234-abcd-1234abcdefgh"
        )
```

## Argument Reference

~> Exactly one of `protection_id` or `resource_arn` is required.

The following arguments are optional:

* `protection_id` - (Optional) Unique identifier for the protection.
* `resource_arn` - (Optional) ARN (Amazon Resource Name) of the resource being protected.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `name` - Name of the protection.
* `protection_arn` - ARN of the protection.

<!-- cache-key: cdktf-0.20.8 input-59de3a79156bc1c7c56349a69033dc1dcd795ab4a1f865d892748664811bc987 -->