.PHONY: build

build:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/kaytu-azure-describer ./command/lambda/main.go
	cd build && zip ./kaytu-azure-describer.zip ./kaytu-azure-describer
	aws s3 cp ./build/kaytu-azure-describer.zip s3://lambda-describe-binary/kaytu-azure-describer.zip
	aws lambda update-function-code --function-name DescribeAzure --s3-bucket lambda-describe-binary --s3-key kaytu-azure-describer.zip --no-cli-pager --no-cli-auto-prompt

build-cli:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/kaytu-azure-cli ./command/main.go
