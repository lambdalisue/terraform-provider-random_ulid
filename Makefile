.PHONY: build install test fmt clean

HOSTNAME=lambdalisue
NAMESPACE=terraform
NAME=random_ulid
BINARY=terraform-provider-${NAME}
VERSION=0.1.0
OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -v ./...

fmt:
	go fmt ./...
	terraform fmt -recursive .

clean:
	rm -f ${BINARY}
	rm -rf examples/.terraform
	rm -f examples/.terraform.lock.hcl
	rm -f examples/terraform.tfstate*