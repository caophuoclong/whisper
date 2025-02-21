package httpErrors

import (
	"fmt"
	"net/http"
)

type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

// Causes implements RestErr.
func (r RestError) Causes() interface{} {
	return r.ErrCauses
}

// Error implements RestErr.
func (r RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", r.ErrStatus, r.ErrError, r.ErrCauses)
}

// Status implements RestErr.
func (r RestError) Status() int {
	return r.ErrStatus
}

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func ParseErrors(err error) RestErr {
	if restErr, ok := err.(RestErr); ok {
		return restErr
	}
	return RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  "Internal server errors",
		ErrCauses: err,
	}
}

func ErrorResponse(err error) (int, interface{}) {
	fmt.Println(err)
	return ParseErrors(err).Status(), ParseErrors(err).Causes()
}
