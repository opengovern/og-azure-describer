.PHONY: install build clean

build:
	GOPRIVATE="github.com/opengovern" CC=/usr/bin/musl-gcc GOPRIVATE="github.com/opengovern" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -o ./build/steampipe-plugin-azuread.plugin *.go

docker-build:
	docker build -f docker/Dockerfile . -t steampipe-plugin-azuread:0.0.1
	docker build -f docker/Dockerfile . -t steampipe-plugin-azuread:0.0.1-$CI_PIPELINE_IID

docker-push:
	docker push steampipe-plugin-azuread:0.0.1
	docker push steampipe-plugin-azuread:0.0.1-$CI_PIPELINE_IID

podman-build:
	podman build -f docker/Dockerfile . -t steampipe-plugin-azuread:0.0.1
	podman build -f docker/Dockerfile . -t steampipe-plugin-azuread:0.0.1-$CI_PIPELINE_IID

podman-push:
	podman push steampipe-plugin-azuread:0.0.1
	podman push steampipe-plugin-azuread:0.0.1-$CI_PIPELINE_IID

clean:
	rm -r ./build

install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/azuread@latest/steampipe-plugin-azuread.plugin *.go
