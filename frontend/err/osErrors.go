package errors

import (
	"fmt"
)

type OperatingSystemErrorType struct {
	Type      int
	NestedErr error
}

// constants to be used
const (
// <ErrorName> = iota
)

func OperatingSystemError(typeof int, err error) *OperatingSystemErrorType {
	operatingSystemError := &OperatingSystemErrorType{
		Type:      typeof,
		NestedErr: err,
	}

	return operatingSystemError
}

func (operatingSystemError OperatingSystemErrorType) Error() string {

	errorField := []string{
		//actual thing to display goes here.
	}

	if operatingSystemError.NestedErr != nil {
		return fmt.Sprintf("server error: %s : %s", errorField[operatingSystemError.Type], operatingSystemError.NestedErr.Error())
	}

	return fmt.Sprintf("server error: %s", errorField[operatingSystemError.Type])

}
