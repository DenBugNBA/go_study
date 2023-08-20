package main

import (
	"errors"
	"fmt"
	"strings"
)

// некритичная ошибка валидации
type NonCriticalError struct{}

func (e NonCriticalError) Error() string {
	return "validation error"
}

// критичные ошибки
var (
	ErrBadConnection = errors.New("bad connection")
	ErrBadRequest    = errors.New("bad request")
)

const UnknownErrorMsg = "unknown error"

func GetErrorMsg(err error) string {
	if errors.Is(err, ErrBadConnection) || errors.Is(err, ErrBadRequest) {
		if strings.HasPrefix(err.Error(), "wrap: ") {
			return err.Error()[6:]
		}

		return err.Error()
	}

	if errors.As(err, &NonCriticalError{}) {
		return ""
	}

	return UnknownErrorMsg
}

var criticalErrs = []error{ErrBadRequest, ErrBadConnection}

func GetErrorMsgTeacher(err error) string {
	for _, crErr := range criticalErrs {
		if errors.Is(err, crErr) {
			return crErr.Error()
		}
	}

	if errors.As(err, &NonCriticalError{}) {
		return ""
	}

	return UnknownErrorMsg
}

func main() {
	fmt.Println(GetErrorMsg(ErrBadConnection))                      // "bad connection"
	fmt.Println(GetErrorMsg(ErrBadRequest))                         // "bad request"
	fmt.Println(GetErrorMsg(errors.New("random error")))            // "unknown error"
	fmt.Println(GetErrorMsg(NonCriticalError{}))                    // ""
	fmt.Println(GetErrorMsg(fmt.Errorf("wrap: %w", ErrBadRequest))) // ""
}
