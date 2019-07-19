GoToJson
========


Steps to use the utility
------------------------

* Install Go
* Execute `go get github.com/terraform-providers/terraform-provider-aws/aws`
* Prepare your resources csv
   ```aws_default_vpc,resourceAwsDefaultVpc,default_vpc.html
    aws_ami_from_instance,resourceAwsAmiFromInstance,ami_from_instance.html```
Fist Column: Terraform resource name
Second Column: Get the method name from terraform go code
Third Column: Find out the markdown file for your resource from your respective provider and released tag e.g. for `aws_ami_from_instance` markdown file is `ami_from_instance.html.markdown` sp entry would be `ami_from_instance.html`

* Run following command `sh generate_resources.sh`
* Enter your provider type
    e.g. For AWS it should be `aws`
* Enter provider version:
    e.g. For AWS v1.60.0
* Enter resouces CSV : Enter your file name

You will get your json generated for the resources mention in the csv file with resource_name.json i.e `aws_default_vpc.json, aws_ami_from_instance`

