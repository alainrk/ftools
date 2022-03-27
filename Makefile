.EXPORT_ALL_VARIABLES:
GO111MODULE = on

run: # export GO111MODULE = on
	@go run cmd/main.go

test:
	go test -v -coverprofile /tmp/c.out ./...
