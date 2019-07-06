package aws

import (
   "github.com/hashicorp/terraform/helper/schema"
)

func GetResourceSchema() *schema.Resource {
    return  dataSourceAwsDbInstance()
}
