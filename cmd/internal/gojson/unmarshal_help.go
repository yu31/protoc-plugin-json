package gojson

//func (p *Plugin) unmarshalReadObjectKeyBefore(loopLabel string) {
//	p.g.P("if decoder.ReadObjectKeyBefore() { // before read object key")
//	p.g.P("    break ", loopLabel)
//	p.g.P("}")
//}
//
//func (p *Plugin) unmarshalReadObjectValueBefore() {
//	p.g.P("decoder.ReadObjectValueBefore() // Before read object value")
//}

func (p *Plugin) unmarshalReadObjectValueAfter(loopLabel string) {
	p.g.P("if decoder.ReadObjectValueAfter() { // After read object value")
	p.g.P("    break ", loopLabel)
	p.g.P("}")
}

//func (p *Plugin) unmarshalReadMapValueAfter(loopLabel string) {
//	p.g.P("if decoder.afterReadMapValue() { // After read map value")
//	p.g.P("    break ", loopLabel)
//	p.g.P("}")
//}

// FIXME: remove it.
//func (p *Plugin) unmarshalReadArrayElemBefore(loopLabel string) {
//	p.g.P("if decoder.ReadArrayElemBefore() { // Before read array value.")
//	p.g.P("    break ", loopLabel)
//	p.g.P("}")
//}

//func (p *Plugin) unmarshalReadArrayElemAfter(loopLabel string) {
//	p.g.P("if decoder.afterReadArrayElem() { // After read array value.")
//	p.g.P("    break ", loopLabel)
//	p.g.P("}")
//}
