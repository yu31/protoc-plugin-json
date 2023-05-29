package gojson

func (p *Plugin) genVariableOneofIsStore(oneofName string) string {
	return "oneOfIsFill_" + oneofName
}

func (p *Plugin) unmarshalReadObjectKeyBefore(loopLabel string) {
	p.g.P("if decoder.ReadObjectKeyBefore() { // before read object key")
	p.g.P("    break ", loopLabel)
	p.g.P("}")
}

func (p *Plugin) unmarshalReadObjectValueBefore() {
	p.g.P("decoder.ReadObjectValueBefore() // Before read object value")
}

func (p *Plugin) unmarshalReadObjectValueAfter(loopLabel string) {
	p.g.P("if decoder.ReadObjectValueAfter() { // After read object value")
	p.g.P("    break ", loopLabel)
	p.g.P("}")
}

func (p *Plugin) unmarshalReadArrayValueBefore(loopLabel string) {
	p.g.P("if decoder.ReadArrayValueBefore() { // Before read array value.")
	p.g.P("    break ", loopLabel)
	p.g.P("}")
}

func (p *Plugin) unmarshalReadArrayValueAfter(loopLabel string) {
	p.g.P("if decoder.AfterReadArrayValueAfter() { // After read array value.")
	p.g.P("    break ", loopLabel)
	p.g.P("}")
}
