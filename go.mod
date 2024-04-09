module github.com/yu31/protoc-plugin-json

go 1.18

require (
	github.com/golang/protobuf v1.5.0
	github.com/json-iterator/go v1.1.12
	github.com/stretchr/testify v1.8.3
	github.com/yu31/protoc-kit-go v0.0.0-20230701152537-70155099b8ab
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace github.com/yu31/protoc-kit-go => ../protoc-kit-go
