// Package errmsg provides error message package.
package errmsg

type Errmsg struct {
	StatusCode int64		`json:"status_code"`
	StatusMsg string	`json:"status_msg"`
}

func (err Errmsg) Error() string {
	return err.StatusMsg
}

func DecodeErr(err error) (int64, string) {
	if err == nil {
		return OK.StatusCode, OK.StatusMsg
	}


	switch typed := err.(type) {
		case *Errmsg:
			return typed.StatusCode, typed.StatusMsg
		default:
			// TODO: implement
	}
	return InternalServerError.StatusCode, err.Error()
}