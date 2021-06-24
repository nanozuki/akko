package akko

type Method string

const (
	MethodGET     Method = "get"
	MethodPUT     Method = "put"
	MethodPOST    Method = "post"
	MethodDELETE  Method = "delete"
	MethodOPTIONS Method = "options"
	MethodHEAD    Method = "head"
	MethodPATCH   Method = "patch"
	MethodTRACE   Method = "trace"
)
