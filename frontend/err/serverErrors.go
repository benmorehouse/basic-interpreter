package errors

import (
	"fmt"
)

//########################################################
//################### Server #############################

type ServerErrorType struct {
	Type      int
	NestedErr error
}

const (
	CreateAppFailed      = iota
	ScriptDirNotFound    = iota
	ListenAndServeFailed = iota
)

func ServerError(typeof int, err error) *ServerErrorType {
	serverError := &ServerErrorType{
		Type:      typeof,
		NestedErr: err,
	}

	return serverError
}

func (serverError ServerErrorType) Error() string {

	errorField := []string{
		"The app failed to initialize",
		"Directory for javascript files not found",
		"Listen and serve suddenly stopped",
	}

	if serverError.NestedErr != nil {
		return fmt.Sprintf("server error: %s : %s", errorField[serverError.Type], serverError.NestedErr.Error())
	}

	return fmt.Sprintf("server error: %s", errorField[serverError.Type])

}
