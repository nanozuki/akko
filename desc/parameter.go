package desc

// ParamProp descript a property in parameter object.
type ParamProp struct {
	*Prop
	In Location

	paramterProp
}

type Location uint8

const (
	_ Location = iota // inbody
	InPath
	InQuery
)
