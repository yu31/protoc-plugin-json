package importpkg

import "google.golang.org/protobuf/compiler/protogen"

var (
	// TODO: rmeove it.
	//Base64    = protogen.GoImportPath("encoding/base64")
	//JSON      = protogen.GoImportPath("encoding/json")
	//Strconv   = protogen.GoImportPath("strconv")

	FMT       = protogen.GoImportPath("fmt")
	Errors    = protogen.GoImportPath("errors")
	PJEncoder = protogen.GoImportPath("github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder")
	PJDecoder = protogen.GoImportPath("github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder")
)
