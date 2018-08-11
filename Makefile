REPO_PATH=github.com/lstoll/dex-oidc-staticgroups
export PATH := $(PWD)/bin:$(PATH)

VERSION ?= $(shell ./hack/git-version)

DOCKER_REPO=lstoll/dex-oidc-staticgroups
DOCKER_IMAGE=$(DOCKER_REPO):$(VERSION)

$( shell mkdir -p bin )

user=$(shell id -u -n)
group=$(shell id -g -n)

export GOBIN=$(PWD)/bin

LD_FLAGS="-w -X $(REPO_PATH)/vendor/github.com/coreos/dex/version.Version=$(VERSION)"

build: bin/dex bin/example-app bin/grpc-client

.PHONY: bin/dex
bin/dex:
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/dex

.PHONY: bin/example-app
bin/example-app:
	@go install -v -ldflags $(LD_FLAGS) github.com/lstoll/dex-oidc-staticgroups/vendor/github.com/coreos/dex/cmd/example-app

.PHONY: bin/grpc-client
bin/grpc-client:
	@go install -v -ldflags $(LD_FLAGS) github.com/lstoll/dex-oidc-staticgroups/vendor/github.com/coreos/dex/examples/grpc-client

.PHONY: release-binary
release-binary:
	@go build -o /go/bin/dex -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/dex

test:
	@go test -v ./...

testrace:
	@go test -v --race ./...

.PHONY: docker-image
docker-image:
	@docker build -t $(DOCKER_IMAGE) .

,PHONY: minikube-deploy
minikube-deploy:
	@eval $$(minikube docker-env) ;\
	docker build -t $(DOCKER_REPO):latest -f Dockerfile .
	#kubectl create -f kubernetes/deploy/

clean:
	@rm -rf bin/

testall: testrace vet fmt lint

FORCE:

.PHONY: test testrace vet fmt lint testall