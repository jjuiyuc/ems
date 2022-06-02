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
	@go build -o dist/weather-worker cmd/daemon/weather-worker/weather-worker.go

swag:
	@$(shell if [ ! -f $(GOPATH)/bin/swag ] ; then go get -u github.com/swaggo/swag/cmd/swag; fi)
	$(AT_RUN)$(GOPATH)/bin/swag init -g routers/router.go

systemd:
	@./scripts/dist.sh

go-test:
	@go test -v -count=1 -coverpkg=$(COVER_PKG) -coverprofile cover.out $(GO_PKGS)

go-cover: go-test
	@go tool cover -html=cover.out -o cover.html
	@go tool cover -func=cover.out