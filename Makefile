.PHONY: build

build:
	export GOOS=linux
	export GOARCH=amd64
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/kaytu-io" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -o ./build/kaytu-azure-describer ./main.go

docker:
	docker build -t kaytu-azure-describer:latest .

build-cli:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/kaytu-azure-cli ./command/main.go
