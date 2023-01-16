default: dockerPackage

dockerPackage:
	docker build . -t sauth:latest

#dockerPublish: dockerPackage
#	docker tag sauth:$(SAUTH_VERSION) knetic/sauth:$(SAUTH_VERSION)
#	docker push knetic/sauth:$(SAUTH_VERSION)
