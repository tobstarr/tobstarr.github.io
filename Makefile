default: test vet build

build:
	@go get -v ./...

vet:
	go vet ./...

test:
	go test ./...

release: build
	ts-gen-tobstarr

preview: build
	ts-gen-tobstarr --do-not-push

godep:
	cd cmd/ts-gen-tobstarr && godep save -r ./... && godep save -r ./...
