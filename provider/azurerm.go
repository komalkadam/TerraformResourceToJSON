
package azurerm
	import (
	"github.com/hashicorp/terraform/helper/schema"
	)
	func GetResourceSchema() *schema.Resource {
	return resourceArmSecurityCenterWorkspace()
	}
