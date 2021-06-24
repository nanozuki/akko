package desc

import (
	"errors"
	"regexp"

	"github.com/nanozuki/akko/strutil"
)

var nameRe = regexp.MustCompile("^[[:alpha:]][[:word:]]*$")

// Name presents valid name of golang.
type Name string

// NewName check string and convert to a valid golang name. the name must be alphanumeric or underscore,
// and must start with letter. If name string has invalid characters, it will panic.
func NewName(name string) Name {
	if !nameRe.MatchString(name) {
		panic(errors.New("the name must be alphanumeric, and must start with letter"))
	}
	return Name(strutil.Pascal(name))
}
