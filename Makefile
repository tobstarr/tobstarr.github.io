default: test vet build

build:
	@go get -v ./...

vet:
	go vet ./...

test:
	go test ./...

run: build
	ts-gen-tobstarr

godep:
	cd cmd/ts-gen-tobstarr && godep save -r ./... && godep save -r ./...
