default: dockerPackage

dockerTest:
ifeq ($(shell which docker), )
	@echo "Docker is not installed."
	@exit 1
endif

build: clean
	go build -o ./sauth .

clean:
	@rm -f ./sauth

dockerPackage: dockerTest
	docker build . -t sauth:latest

#dockerPublish: dockerPackage
#	docker tag sauth:$(SAUTH_VERSION) knetic/sauth:$(SAUTH_VERSION)
#	docker push knetic/sauth:$(SAUTH_VERSION)
