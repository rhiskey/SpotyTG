package structures

import "fmt"

type ArgError struct {
	arg  int
	prob string
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
