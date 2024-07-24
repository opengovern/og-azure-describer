.PHONY: build

lambda-build:
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/kaytu-io" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl,lambda.norpc -o ./build/kaytu-azure-describer ./lambda/main.go

functions-build:
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/kaytu-io" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -o ./azfunction/kaytu-azure-describer ./azfunction/main.go

local-build:
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/kaytu-io" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -o ./local/kaytu-azure-describer ./local/main/main.go

docker:
	docker build -t 435670955331.dkr.ecr.us-east-2.amazonaws.com/kaytu-azure-describer:latest .
	docker push 435670955331.dkr.ecr.us-east-2.amazonaws.com/kaytu-azure-describer:latest

aws-update:
	aws lambda update-function-code --function-name kaytu-azure-describer --image-uri 435670955331.dkr.ecr.us-east-2.amazonaws.com/kaytu-azure-describer:latest --region us-east-2


build-cli:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/kaytu-azure-cli ./command/main.go
