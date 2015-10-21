# install with go get github.com/alecthomas/gometalinter
LINT_CMD := gometalinter \
        --deadline=10s \
        --exclude=".*comment.*" \
        --exclude=".*defer.*errcheck.*" \
        --exclude=".*unused struct field.*" \
        --exclude="scripts/goassets.*" \
        --exclude="assets/assets.go" \
        --severity=errcheck:error \
        --severity=varcheck:error \
        --severity=defercheck:error \
        --severity=go-nyet:error \
        ./...

lint_errors: build
    $(LINT_CMD) --disable vet --disable deadcode --disable gocyclo --enable go-nyet --enable golint --enable defercheck --disable varcheck --disable structcheck --disable gotype --enable errcheck --disable dupl

lint: build
    $(LINT_CMD) --disable vet --enable deadcode --disable gocyclo --enable go-nyet --enable golint --enable defercheck --enable varcheck --enable structcheck --disable gotype --enable errcheck --disable dupl
