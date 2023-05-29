package main

import (
	"github.com/yu31/protoc-kit-go/pkgenerator"

	"github.com/yu31/protoc-plugin-json/cmd/internal/gojson"
)

func main() {
	pkgenerator.Run(gojson.New())
}
