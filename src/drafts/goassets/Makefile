ASSETS = $(shell find assets -type f | grep -v "^assets/assets.go$$")

# see README.md for install instructions

assets/assets.go: $(ASSETS)
	rm -f $@
	go run ./scripts/goassets/main.go -file $@ assets
	cd assets && go get .
