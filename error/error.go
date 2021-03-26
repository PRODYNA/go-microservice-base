package error

import (
	"fmt"
)


type ServiceError struct {
	Message string
	Status int
}

func (s ServiceError) Error() string {
	return fmt.Sprintf("message: %s, MinStatus %d", s.Message, s.Status)
}

func (s ServiceError) ErrorJson() string {
	return fmt.Sprintf("{ message: \"%s\", MinStatus: %d }", s.Message, s.Status)
}

func (s ServiceError) ErrorJsonBytes() []byte {
	return []byte(fmt.Sprintf("{ message: \"%s\", MinStatus: %d }", s.Message, s.Status))
}

func NewError(message string, status int) error {
	return ServiceError{
		Message: message,
		Status: status,
	}
}

func NotFound(message string) error {
	return ServiceError{
		Message: message,
		Status:  404,
	}
}

func BadRequest(message string) error {
	return ServiceError{
		Message: message,
		Status:  400,
	}
}

func Unauthorized(message string) error {
	return ServiceError{
		Message: message,
		Status:  401,
	}
}


func Forbidden(message string) error {
	return ServiceError{
		Message: message,
		Status:  403,
	}
}

func InternalServerError(message string) error {
	return ServiceError{
		Message: message,
		Status:  500,
	}
}

func NotImplemented(message string) error {
	return ServiceError{
		Message: message,
		Status:  501,
	}
}

func BadGateway(message string) error {
	return ServiceError{
		Message: message,
		Status:  502,
	}
}

func ServiceUnavailable(message string) error {
	return ServiceError{
		Message: message,
		Status:  503,
	}
}


