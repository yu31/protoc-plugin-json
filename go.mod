module github.com/yu31/protoc-plugin-json

go 1.15

require (
	github.com/golang/protobuf v1.5.0
	github.com/json-iterator/go v1.1.12
	github.com/stretchr/testify v1.8.3
	github.com/yu31/protoc-kit-go v0.0.0-20230701152537-70155099b8ab
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1
)

//replace github.com/yu31/protoc-kit-go => ../protoc-kit-go
