GoToJson
========


Steps to use the utility
------------------------
* Import task_example in eclipse
* Find out the markdown file for your resource from your respective provider and released tag e.g. for `aws_ami_from_instance` markdown file is `ami_from_instance.html.markdown` and update attributes `markdown_file_name` from `MarkDownParser.java`
* Update `provider_type` in `MarkDownParser.java`
* Update `released_version` in `MarkDownParser.java`
* Install Go
* Execute `go get github.com/terraform-providers/terraform-provider-aws/aws`
* Chane the method in the file to that resource's method for getting it's attribute e.g. `resourceAwsInstance()` for `aws_instance`
* Copy this file `<GO_HOME>/src/github.com/terraform-providers/terraform-provider-aws/aws/`
* Change your resource name in file `main.go`
* Run following command `go run attributes.go main.go <resource_name>`
