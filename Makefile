.EXPORT_ALL_VARIABLES:
GO111MODULE = on

test:
	go test -v -coverprofile /tmp/c.out ./...
