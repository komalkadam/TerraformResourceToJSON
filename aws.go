package aws

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func GetResourceSchema() *schema.Resource {
	//Change the method name here
	return dataSourceAwsDbInstance()
}
