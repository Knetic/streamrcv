default: dockerPackage

export GOPATH=$(CURDIR)/
export GOBIN=$(CURDIR)/.output/bin/
export GOCACHE=$(CURDIR)/.output/cache/

build: clean
	go build -o ./.output/sauth .

clean:
	@rm -rf ./.output/
	@mkdir  ./.output/

containerized_build: dockerTest

	docker run \
		--rm \
		-v "$(CURDIR)":"/srv/build":rw \
		-u "$(shell id -u $(whoami)):$(shell id -g $(whoami))" \
		-e SAUTH_VERSION=$(SAUTH_VERSION) \
		golang:1.10 \
		bash -c \
		"cd /srv/build; make build"

dockerPackage: containerized_build
	docker build . -t sauth:latest

#dockerPublish: dockerPackage
#	docker tag sauth:$(SAUTH_VERSION) knetic/sauth:$(SAUTH_VERSION)
#	docker push knetic/sauth:$(SAUTH_VERSION)

dockerTest:
ifeq ($(shell which docker), )
	@echo "Docker is not installed."
	@exit 1
endif
