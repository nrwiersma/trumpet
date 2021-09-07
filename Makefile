include github.com/hamba/make/golang

# Build the docker image
docker:
	docker build -t trumpet .
.PHONY: docker
