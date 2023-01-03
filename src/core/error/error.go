package error

type BadRequest struct {
	message      string
	statusRecord string
	trace        string
	error        error
}

type NotFound struct {
	message      string
	statusRecord string
	trace        string
	error        error
}

type InternalServer struct {
	message      string
	statusRecord string
	trace        string
	error        error
}

func NewError() error {
	return nil
}
