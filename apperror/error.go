package apperror

import "fmt"

type AppError struct {
	Code    int
	Message string
}

func (ae AppError) Error() string {
	return fmt.Sprintf("Code %d | Error: %s", ae.Code, ae.Message)
}
