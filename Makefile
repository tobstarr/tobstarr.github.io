default: run

build: vet test
	@go get -v ./...

vet:
	go vet ./...

test:
	go test ./...

run: build
	ts-gen-tobstarr
