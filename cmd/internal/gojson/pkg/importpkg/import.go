package importpkg

import "google.golang.org/protobuf/compiler/protogen"

var (
	Errors    = protogen.GoImportPath("errors")
	PJEncoder = protogen.GoImportPath("github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder")
	PJDecoder = protogen.GoImportPath("github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder")
)
