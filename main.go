package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kaytu-io/kaytu-azure-describer/describer"
)

func main() {
	lambda.Start(describer.DescribeHandler)
}
