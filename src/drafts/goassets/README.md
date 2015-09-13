# Goassets

## Installation

	mkdir -p scripts/goassets
	curl -SsfLO scripts/goassets/main.go https://raw.githubusercontent.com/dynport/dgtk/master/goassets/script/goassets.go

and then copy bits from Makefile

## .gitattributes

Add this to `.gitattributes`

	 assets/assets.go -diff

## in github.com/julienschmidt/httprouter

	r.ServeFiles("/assets/*filepath", assets.FileSystem(""))
