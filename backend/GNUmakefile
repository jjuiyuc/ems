GOPATH := $(shell go env GOPATH)

GO_PKGS := $(shell go list ./... | grep -v models\/ )

# saparated by comma
COVER_PKG := $(shell go list ./... | grep -v models\/ | paste -sd "," - )

clean:
	@rm -rf dist

gen-models:
	sqlboiler -c sqlboiler.der-ems.toml --struct-tag-casing camel mysql

go-lint:
	@$(shell if [ ! -f $(GOPATH)/bin/golint ] ; then go get -u golang.org/x/lint/golint; fi)
	@$(AT_RUN)$(GOPATH)/bin/golint $(GO_PKGS)

go-build: clean
	@./build/build.sh

swag:
	@$(shell if [ ! -f $(GOPATH)/bin/swag ] ; then go get -u github.com/swaggo/swag/cmd/swag; fi)
	$(AT_RUN)$(GOPATH)/bin/swag init -g routers/router.go

FOCUS ?= .
go-ginkgo:
	GIN_MODE=release ginkgo -r -v --coverpkg=$(COVER_PKG) --focus $(FOCUS) --coverprofile=cover.out --junit-report=report.xml ./...

go-cover: go-ginkgo
	@go tool cover -html=cover.out -o cover.html
	@go tool cover -func=cover.out
