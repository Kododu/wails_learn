package exception

import "github.com/pkg/errors"

var (
	Err5001 = errors.New("not found conf")
)

var (
	ErrConnNotFoundInConf = errors.New("connection not found in conf")
)
