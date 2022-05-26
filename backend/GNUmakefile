GOPATH := $(shell go env GOPATH)

GO_PKGS := $(shell go list ./... | grep -v models\/ )

clean:
	@rm -rf dist

gen-models:
	sqlboiler -c sqlboiler.der-ems.toml --struct-tag-casing camel mysql

go-lint:
	@$(shell if [ ! -f $(GOPATH)/bin/golint ] ; then go get -u golang.org/x/lint/golint; fi)
	@$(AT_RUN)$(GOPATH)/bin/golint $(GO_PKGS)

go-build: clean
	@go build -o dist/weather-worker cmd/daemon/weather-worker/weather-worker.go

systemd:
	@./scripts/dist.sh
