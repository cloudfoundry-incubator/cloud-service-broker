SHELL = /bin/bash
GO-VER = go1.13

OSFAMILY=$(shell uname)
ifeq ($(OSFAMILY),Darwin)
OSFAMILY=darwin
else
OSFAMILY=linux
endif

ifeq ($(USE_GO_CONTAINERS),)
GO=go
GOIMPORTS=goimports
else
UID:=$(shell id -u)
DOCKER_OPTS=--rm -u $(UID) -v $(HOME):$(HOME) -e HOME -e USER=$(USER) -e USERNAME=$(USER) -w $(PWD)
GO=docker run $(DOCKER_OPTS) -e GOARCH -e GOOS golang go
GOTOOLS=docker run $(DOCKER_OPTS) jare/go-tools
GOIMPORTS=$(GOTOOLS) goimports
HAS_GO_IMPORTS=true
endif

SRC = $(shell find . -name "*.go" | grep -v "_test\." )

VERSION := $(or $(VERSION), dev)

LDFLAGS="-X github.com/pivotal/cloud-service-broker/utils.Version=$(VERSION)"

.PHONY: deps-go-binary
deps-go-binary:
	echo "Expect: $(GO-VER)" && \
		echo "Actual: $$($(GO) version)" && \
	 	$(GO) version | grep $(GO-VER) > /dev/null

HAS_GO_IMPORTS := $(shell command -v goimports;)

deps-goimports: deps-go-binary
ifndef HAS_GO_IMPORTS
	go get -u golang.org/x/tools/cmd/goimports
endif

.PHONY: test-units
test-units: deps-go-binary
	$(GO) test -v ./... -tags=service_broker

.PHONY: test-acceptance 
test-acceptance: ./build/cloud-service-broker.$(OSFAMILY) security-user-name security-user-password
	./build/cloud-service-broker.$(OSFAMILY) client run-examples

./build/cloud-service-broker.linux: $(SRC)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GO) build -o ./build/cloud-service-broker.linux -ldflags ${LDFLAGS}

./build/cloud-service-broker.darwin: $(SRC)
	GOARCH=amd64 GOOS=darwin $(GO) build -o ./build/cloud-service-broker.darwin -ldflags ${LDFLAGS}

.PHONY: build
build: deps-go-binary ./build/cloud-service-broker.linux ./build/cloud-service-broker.darwin

.PHONY: package
package: ./build/cloud-service-broker.$(OSFAMILY) ./tile.yml ./manifest.yml docs/customization.md

./tile.yml:
	./build/cloud-service-broker.$(OSFAMILY) generate tile > ./tile.yml

./manifest.yml:
	./build/cloud-service-broker.$(OSFAMILY) generate manifest > ./manifest.yml

docs/customization.md:
	./build/cloud-service-broker.$(OSFAMILY) generate customization > docs/customization.md

.PHONY: clean
clean: deps-go-binary
	$(GO) clean --modcache
	rm -rf ./build
	rm gcp-brokerpak/*.brokerpak
	rm azure-brokerpak/*.brokerpak

.PHONY: lint
lint: deps-goimports
	git ls-files | grep '.go$$' | xargs $(GOIMPORTS) -l -w

# GCP broker

.PHONY:	lint-brokerpak-gcp
lint-brokerpak-gcp:
	docker run --rm -it -v $(PWD):/broker upstreamable/yamlint /usr/local/bin/yamllint -c /broker/yamllint.conf /broker/gcp-brokerpak

gcp-brokerpak/*.brokerpak: lint-brokerpak-gcp ./build/cloud-service-broker.$(OSFAMILY) ./gcp-brokerpak/*.yml
	cd ./gcp-brokerpak && ../build/cloud-service-broker.$(OSFAMILY) pak build

.PHONY: run-broker-gcp
run-broker-gcp: check-gcp-env-vars ./build/cloud-service-broker.$(OSFAMILY) gcp-brokerpak/*.brokerpak
	GSB_BROKERPAK_BUILTIN_PATH=./gcp-brokerpak ./build/cloud-service-broker.$(OSFAMILY) serve

.PHONY: push-broker-gcp
push-broker-gcp: check-gcp-env-vars ./build/cloud-service-broker.$(OSFAMILY) gcp-brokerpak/*.brokerpak
	GSB_BROKERPAK_BUILTIN_PATH=./gcp-brokerpak ./scripts/push-broker.sh	

.PHONY: run-broker-gcp-docker
run-broker-gcp-docker: check-gcp-env-vars ./build/cloud-service-broker.linux gcp-brokerpak/*.brokerpak
	GSB_BROKERPAK_BUILTIN_PATH=/broker/gcp-brokerpak \
	DB_HOST=host.docker.internal \
	docker run --rm -p 8080:8080 -v $(PWD):/broker \
	-e GSB_BROKERPAK_BUILTIN_PATH \
	-e DB_HOST \
	-e DB_USERNAME \
	-e DB_PASSWORD \
	-e PORT \
	-e SECURITY_USER_NAME \
	-e SECURITY_USER_PASSWORD \
	-e GOOGLE_CREDENTIALS \
	-e GOOGLE_PROJECT \
	alpine /broker/build/cloud-service-broker.linux serve	

# Azure broker

.PHONY:	test-brokerpak-azure
test-brokerpak-azure:
	docker run --rm -it -v $(PWD):/broker upstreamable/yamlint /usr/local/bin/yamllint -c /broker/yamllint.conf /broker/azure-brokerpak

.PHONY: run-broker-azure
run-broker-azure: check-azure-env-vars ./build/cloud-service-broker.$(OSFAMILY) azure-brokerpak/*.brokerpak test-brokerpak-azure
	GSB_BROKERPAK_BUILTIN_PATH=./azure-brokerpak ./build/cloud-service-broker.$(OSFAMILY) serve

azure-brokerpak/*.brokerpak: ./build/cloud-service-broker.$(OSFAMILY) ./azure-brokerpak/*.yml
	cd ./azure-brokerpak && ../build/cloud-service-broker.$(OSFAMILY) pak build

.PHONY: push-broker-azure
push-broker-azure: check-azure-env-vars ./build/cloud-service-broker.$(OSFAMILY) azure-brokerpak/*.brokerpak
	GSB_BROKERPAK_BUILTIN_PATH=./azure-brokerpak ./scripts/push-broker.sh

.PHONY: run-broker-azure-docker
run-broker-azure-docker: check-azure-env-vars ./build/cloud-service-broker.linux azure-brokerpak/*.brokerpak test-brokerpak-azure
	GSB_BROKERPAK_BUILTIN_PATH=/broker/azure-brokerpak \
	DB_HOST=host.docker.internal \
	docker run --rm -p 8080:8080 -v $(PWD):/broker \
	-e GSB_BROKERPAK_BUILTIN_PATH \
	-e DB_HOST \
	-e DB_USERNAME \
	-e DB_PASSWORD \
	-e PORT \
	-e SECURITY_USER_NAME \
	-e SECURITY_USER_PASSWORD \
	-e ARM_SUBSCRIPTION_ID \
	-e ARM_TENANT_ID \
	-e ARM_CLIENT_ID \
	-e ARM_CLIENT_SECRET \
	alpine /broker/build/cloud-service-broker.linux serve	

# AWS broker 
.PHONY: aws-brokerpak
aws-brokerpak: aws-brokerpak/*.brokerpak

aws-brokerpak/*.brokerpak: ./build/cloud-service-broker.$(OSFAMILY) ./aws-brokerpak/*.yml
	cd ./aws-brokerpak && ../build/cloud-service-broker.$(OSFAMILY) pak build

.PHONY: run-broker-aws 
run-broker-aws: check-aws-env-vars ./build/cloud-service-broker.$(OSFAMILY) aws-brokerpak/*.brokerpak
	GSB_BROKERPAK_BUILTIN_PATH=./aws-brokerpak ./build/cloud-service-broker.$(OSFAMILY) serve

.PHONY: run-broker-aws-docker
run-broker-aws-docker: check-aws-env-vars ./build/cloud-service-broker.linux aws-brokerpak/*.brokerpak
	GSB_BROKERPAK_BUILTIN_PATH=/broker/aws-brokerpak \
	DB_HOST=host.docker.internal \
	docker run --rm -p 8080:8080 -v $(PWD):/broker \
	-e GSB_BROKERPAK_BUILTIN_PATH \
	-e DB_HOST \
	-e DB_USERNAME \
	-e DB_PASSWORD \
	-e PORT \
	-e SECURITY_USER_NAME \
	-e SECURITY_USER_PASSWORD \
	-e AWS_ACCESS_KEY_ID \
    -e AWS_SECRET_ACCESS_KEY \
	alpine /broker/build/cloud-service-broker.linux serve

# env vars checks

.PHONY: google-credentials
google-credentials:
ifndef GOOGLE_CREDENTIALS
	$(error variable GOOGLE_CREDENTIALS not defined)
endif

.PHONY: google-project
google-project:
ifndef GOOGLE_PROJECT
	$(error variable GOOGLE_PROJECT not defined)
endif

.PHONY: security-user-name
security-user-name:
ifndef SECURITY_USER_NAME
	$(error variable SECURITY_USER_NAME not defined)
endif

.PHONY: security-user-password
security-user-password:
ifndef SECURITY_USER_PASSWORD
	$(error variable SECURITY_USER_PASSWORD not defined)
endif

.PHONY: db-host
db-host:
ifndef DB_HOST
	$(error variable DB_HOST not defined)
endif

.PHONY: db-username
db-username:
ifndef DB_USERNAME
	$(error variable DB_USERNAME not defined)
endif

.PHONY: db-password
db-password:
ifndef DB_PASSWORD
	$(error variable DB_PASSWORD not defined)
endif

.PHONY: arm-subscription-id
arm-subscription-id:
ifndef ARM_SUBSCRIPTION_ID
	$(error variable ARM_SUBSCRIPTION_ID not defined)
endif

.PHONY: arm-tenant-id
arm-tenant-id:
ifndef ARM_TENANT_ID
	$(error variable ARM_TENANT_ID not defined)
endif

.PHONY: arm-client-id
arm-client-id:
ifndef ARM_CLIENT_ID
	$(error variable ARM_CLIENT_ID not defined)
endif

.PHONY: arm-client-secret
arm-client-secret:
ifndef ARM_CLIENT_SECRET
	$(error variable ARM_CLIENT_SECRET not defined)
endif

.PHONY: aws_access_key_id
aws_access_key_id:
ifndef AWS_ACCESS_KEY_ID
	$(error variable AWS_ACCESS_KEY_ID not defined)
endif

.PHONY: aws_secret_access_key
aws_secret_access_key:
ifndef AWS_SECRET_ACCESS_KEY
	$(error variable AWS_SECRET_ACCESS_KEY not defined)
endif

.PHONY: check-gcp-env-vars
check-gcp-env-vars: google-credentials google-project security-user-name security-user-password db-host db-username db-password

.PHONY: check-azure-env-vars
check-azure-env-vars: arm-subscription-id arm-tenant-id arm-client-id arm-client-secret security-user-name security-user-password db-host db-username db-password

.PHONY: check-aws-env-vars
check-aws-env-vars: aws_access_key_id aws_secret_access_key security-user-password db-host db-username db-password
