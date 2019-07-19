#!/usr/bin/env bash
set -e
home=$(pwd)
INPUT=provider/aws.csv
OLDIFS=$IFS
IFS=","
[ ! -f $INPUT ] && { echo "$INPUT file not found"; exit 99; }
while IFS=',' read -r  markdownFileName functionName resourceName
do
	echo "markdown_file_name : $markdownFileName"
	echo "Resource function Name : $functionName"
	echo "Resource Name : $resourceName"
	echo > provider/azurerm.go
	echo "package aws
	import (
	\"github.com/hashicorp/terraform/helper/schema\"
	)
	func GetResourceSchema() *schema.Resource {
	return $functionName()
	}">> provider/aws.go

	cp provider/aws.go ~/go/src/github.com/terraform-providers/terraform-provider-aws/aws/
	cd $home/task_example
	mvn spring-boot:run -Dspring-boot.run.arguments=$markdownFileName,aws,v1.60.0
	cd ..
	go run attributes.go main.go $resourceName
done < "$INPUT"
