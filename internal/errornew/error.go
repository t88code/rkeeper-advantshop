package errornew

type Error struct {
	Technical bool
	Err       error
}

func (e *Error) Error() string {
	return e.Err.Error()
}
