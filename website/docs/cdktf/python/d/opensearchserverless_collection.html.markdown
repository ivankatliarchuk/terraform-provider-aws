---
subcategory: "OpenSearch Serverless"
layout: "aws"
page_title: "AWS: aws_opensearchserverless_collection"
description: |-
  Terraform data source for managing an AWS OpenSearch Serverless Collection.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_opensearchserverless_collection

Terraform data source for managing an AWS OpenSearch Serverless Collection.

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
from imports.aws.data_aws_opensearchserverless_collection import DataAwsOpensearchserverlessCollection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsOpensearchserverlessCollection(self, "example",
            name="example"
        )
```

## Argument Reference

The following arguments are required:

* `id` - (Required) ID of the collection. Either `id` or `name` must be provided.
* `name` - (Required) Name of the collection. Either `name` or `id` must be provided.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the collection.
* `collection_endpoint` - Collection-specific endpoint used to submit index, search, and data upload requests to an OpenSearch Serverless collection.
* `created_date` - Date the Collection was created.
* `dashboard_endpoint` - Collection-specific endpoint used to access OpenSearch Dashboards.
* `description` - Description of the collection.
* `failure_code` - A failure code associated with the collection.
* `failure_reason` - A failure reason associated with the collection.
* `kms_key_arn` - The ARN of the Amazon Web Services KMS key used to encrypt the collection.
* `last_modified_date` - Date the Collection was last modified.
* `standby_replicas` - Indicates whether standby replicas should be used for a collection.
* `tags` - A map of tags to assign to the collection.
* `type` - Type of collection.

<!-- cache-key: cdktf-0.20.8 input-773c4f4c82e5aba517ec43bce8ca59c03354cd73fbc2f9496ce75f162ef2ac0f -->