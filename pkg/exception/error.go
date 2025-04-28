package exception

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("status %d %s", c.Code, c.Message)
}

func NewError(code int, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
