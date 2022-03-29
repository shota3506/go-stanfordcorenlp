IMAGE_TAG ?= latest
DOCKER_CORENLP_NAME ?= go-stanfordcorenlp/corenlp
STANFORD_CORENLP_URL ?= http://localhost:9000

.PHONY: all
all: vet

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	STANFORD_CORENLP_URL=$(STANFORD_CORENLP_URL) go test -v ./...

.PHONY: build_docker_corenlp
build_docker_corenlp:
	docker build -f Dockerfile-corenlp -t $(DOCKER_CORENLP_NAME):$(IMAGE_TAG) .

.PHONY: run_docker_corenlp
run_docker_corenlp:
	docker run -itd -p 9000:9000 $(DOCKER_CORENLP_NAME)

