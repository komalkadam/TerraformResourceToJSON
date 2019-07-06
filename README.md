GoToJson
========


Steps to use the utility
------------------------
* Install Go
* Execute go get github.com/terraform-providers/terraform-provider-aws/aws
* Chane the method in the file to that resource's method for getting it's attribute e.g. resourceAwsInstance() for aws_instance
* Copy this file <GO_HOME>/src/github.com/terraform-providers/terraform-provider-aws/aws/
* Change your resource name in file main.go
* Run following command go run main.go
