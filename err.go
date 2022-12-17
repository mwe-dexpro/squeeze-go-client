package squeeze_go_client

import (
	"net/http"
)

type Error struct {
	err error
	res *http.Response
}

func newApiErr(err error, res *http.Response) *Error {
	return &Error{
		err: err,
		res: res,
	}
}

func newErr(err error) *Error {
	return &Error{
		err: err,
	}
}

func (e *Error) Error() string {
	return e.err.Error()
}
