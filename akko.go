package akko

type Akko struct{}

func New() *Akko { return &Akko{} }

func (a *Akko) Attach(adHoc AdHoc) *Akko {
	return a
}

type Route interface{} // must be function

func (a *Akko) Mount(base string, routes []Route) *Akko {
	return a
}
