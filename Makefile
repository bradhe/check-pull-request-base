SHELL            := /bin/bash
LOCAL_OS         := $(shell uname -s | tr A-Z a-z)
VERSION          := $(shell grep Version version.go | sed -e 's/\"//g' -e 's/const Version = //')

GO               := go
GOOS             := $(LOCAL_OS)

PACKAGE_BINARY	 := github-pull-request-base
PACKAGE_NAME	 := github-pull-request-base

AWS_PROFILE		  = personal

build:
	$(GO) build

setup:
	$(GO) get
	@if [ "`which fpm`" == "" ]; then gem install fpm; fi
	@if [ "`which aws`" == "" ]; then pip install --upgrade awscli; fi

clean:
	$(GO) clean

package:
	GOOS=linux GOARCH=amd64 $(GO) build -o bin/$(PACKAGE_BINARY)
	fpm -s dir -t deb --name $(PACKAGE_NAME) \
		--version $(VERSION) \
		--architecture amd64 \
		./bin/$(PACKAGE_BINARY)=/usr/bin/$(PACKAGE_BINARY)

release: package
				aws s3 cp $(PACKAGE_NAME)_$(VERSION)_$(GOARCH).deb s3://bradhe-packages/$(PACKAGE_NAME)-$(VERSION)-$(GOARCH).deb --acl public-read --region us-east-1 --profile $(AWS_PROFILE)
