package common

import (
	"errors"
	"fmt"
)

var (
	ErrNotFoundEndpoint = errors.New("Not found endponit")
	ErrLoadConfig       = errors.New("failed to Load config")
	ErrConnectDB        = errors.New("failed to Connect database")
	ErrLoadTLSFiles     = errors.New("failed to load cert file or cert key file")
	ErrListenServer     = errors.New("failed to listen server")
	ErrCreateDir        = errors.New("failed to make dir")
)

func NewError(err error, cause interface{}) error {
	return fmt.Errorf("%s: %s", err, cause)
}

func NewErrLoadConfig(cause interface{}) error {
	return NewError(ErrLoadConfig, cause)
}

func NewErrConnectDB(cause interface{}) error {
	return NewError(ErrConnectDB, cause)
}

func NewErrNotFoundEndpoint(cause interface{}) error {
	return NewError(ErrNotFoundEndpoint, cause)
}

func NewErrLoadTLSFiles(cause interface{}) error {
	return NewError(ErrLoadTLSFiles, cause)
}

func NewErrListenServer(cause interface{}) error {
	return NewError(ErrListenServer, cause)
}

func NewErrCreateDir(cause interface{}) error {
	return NewError(ErrCreateDir, cause)
}
