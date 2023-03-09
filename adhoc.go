package akko

type AdHoc interface{}

type adHoc struct{}

func OnIgnite(name string, process func(*Akko)) AdHoc {
	return &adHoc{}
}

func Provide[T any](func() (T, error)) AdHoc {
	return &adHoc{}
}
