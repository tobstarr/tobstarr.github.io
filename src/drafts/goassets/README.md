# Goassets

## Installation

	mkdir -p scripts/goassets
	curl -SsfLO scripts/goassets/main.go https://raw.githubusercontent.com/dynport/dgtk/master/goassets/script/goassets.go

and then copy bits from Makefile

## in httprouter

	r.ServeFiles("/assets/*filepath", assets.FileSystem(""))
