#!/usr/bin/env bash
set -e
home=$(pwd)
INPUT=provider/azurerm.csv
OLDIFS=$IFS
IFS=","
[ ! -f $INPUT ] && { echo "$INPUT file not found"; exit 99; }
while IFS=',' read -r  markdownFileName functionName resourceName
do
	echo "markdown_file_name : $markdownFileName"
	echo "Resource function Name : $functionName"
	echo "Resource Name : $resourceName"
	echo > provider/azurerm.go
	echo "package azurerm
	import (
	\"github.com/hashicorp/terraform/helper/schema\"
	)
	func GetResourceSchema() *schema.Resource {
	return $functionName()
	}">> provider/azurerm.go

	cp provider/azurerm.go ~/go/src/github.com/terraform-providers/terraform-provider-azurerm/azurerm/
	cd $home/task_example
	mvn spring-boot:run -Dspring-boot.run.arguments=$markdownFileName,azurerm,v1.30.1
	cd ..
	go run attributes.go azurermResourceGenerator.go $resourceName
done < "$INPUT"
