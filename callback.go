package akko

import "github.com/getkin/kin-openapi/openapi3"

type CallbackBuilder struct {
	expression string
	callback   *openapi3.Callback
}

func Callback(expression string, path *PathBuilder) *CallbackBuilder {
	b := &CallbackBuilder{expression: expression}
	callback := openapi3.Callback{}
	callback[path.path] = path.item
	b.callback = &callback
	return b
}

func (b *CallbackBuilder) AttachToOp(op *OperationBuilder) *CallbackBuilder {
	op.operation.Callbacks[b.expression] = &openapi3.CallbackRef{Value: b.callback}
	return b
}
