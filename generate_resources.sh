#!/bin/sh
PROVIDER_TYPE="aws"
PROVIDER_VERSION="v1.60.0"
COPY_FILE_LOCATION=""

echo "Enter provider type:"
read PROVIDER_TYPE
echo "Enter provider version:"
read PROVIDER_VERSION
echo "Enter resouces CSV"
read input_csv
PROVIDER_PATH=""

if [ $PROVIDER_TYPE == "aws" ]
then
   COPY_FILE_LOCATION=$HOME"/go/src/github.com/terraform-providers/terraform-provider-aws/aws/deploy.go"
   PROVIDER_PATH="github.com/terraform-providers/terraform-provider-aws/aws"
   echo $PROVIDER_PATH
elif [ $PROVIDER_TYPE == "azurerm" ]
then
   COPY_FILE_LOCATION="~/go/src/github.com/terraform-providers/terraform-provider-azurerm/aws/deploy.go"
   PROVIDER_PATH="github.com/terraform-providers/terraform-provider-azurerm/azurerm"
fi

sed "s+terraform_provider_path+${PROVIDER_PATH}+g" main_template.go.tpl > main_template.go

INPUT=$input_csv
OLDIFS=$IFS
IFS=,
[ ! -f $INPUT ] && { echo "$INPUT file not found"; exit 99; }
while read resourcename methodname markdownfile
do
	echo "Name : $resourcename"
	echo "method : $methodname"
	echo "markdown : $markdownfile"
	cd task_example && mvn exec:java -Dexec.mainClass="com.komal.taskexecutor.task_example.MarkDownParser" -Dexec.args="$PROVIDER_TYPE $PROVIDER_VERSION $markdownfile"
	markdown_generation=$!
	#wait $markdown_generation
	cd -
	sed "s/provider_type/${PROVIDER_TYPE}/g" resource_template.go > deploy.go
	sed "s/method_name/${methodname}/g" deploy.go > deploy1.go
	
	
	
	cp deploy1.go $COPY_FILE_LOCATION
	go run attributes.go main_template.go $resourcename
	rm deploy1.go
	rm deploy.go
done < $INPUT
IFS=$OLDIFS

rm main_template.go
