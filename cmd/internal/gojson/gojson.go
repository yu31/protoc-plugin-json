package gojson

import (
	"flag"

	"github.com/yu31/protoc-kit-go/pkgenerator"
	"github.com/yu31/protoc-kit-go/utils/pkmessage"
	"google.golang.org/protobuf/compiler/protogen"
)

const version = "0.0.1"

// Plugin for implements pgkgenerator.Plugin
type Plugin struct {
	flags flag.FlagSet
	g     *protogen.GeneratedFile

	pp   *protogen.Plugin
	file *protogen.File

	messages []*protogen.Message
}

func New() pkgenerator.Plugin {
	p := &Plugin{}
	return p
}

// Name identifies the Plugin.
func (p *Plugin) Name() string {
	return "json"
}

// Version identifies the Plugin version.
func (p *Plugin) Version() string {
	return version
}

// ParamFunc is used for accept parameters from the command line.
func (p *Plugin) ParamFunc() func(name, value string) error {
	return p.flags.Set
}

// Init is called once before code generated.
// The `file` will be ignored if return false.
func (p *Plugin) Init(pp *protogen.Plugin, file *protogen.File) bool {
	if len(file.Messages) == 0 {
		return false
	}
	p.pp = pp
	p.file = file
	p.messages = pkmessage.LoadValidMessages(file.Messages)
	return true
}

// Generate to generate codes for specified file.
// except for the imports, by calling the generator's methods P, In, and Out.
func (p *Plugin) Generate(g *protogen.GeneratedFile) {
	p.g = g
	for _, msg := range p.messages {
		p.generateForMessage(msg)
	}
}

func (p *Plugin) generateForMessage(msg *protogen.Message) {
	options := loadMessageOptions(msg)
	if options.Ignore {
		return
	}

	loader := &FieldLoader{}
	fieldSets := loader.Load(p.file, msg)

	marshal := &Marshal{
		g:       p.g,
		file:    p.file,
		message: msg,
		options: options,
	}
	marshal.GenerateCode(fieldSets)

	unmarshal := &Unmarshal{
		g:       p.g,
		file:    p.file,
		message: msg,
		options: options,
	}
	unmarshal.GenerateCode(fieldSets)
}
