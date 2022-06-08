.PHONY: docker-build
docker-build:
	docker build -t habada .

.PHONY: build

build:
	./build-darwin-arm64.sh

amd64:
	./build-linux-amd64.sh

