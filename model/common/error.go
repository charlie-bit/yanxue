package common

type Error int

// successful business processing
const ErrNone Error = 0

// permission verification failure
const (
	ErrUnsupportedToken Error = iota + 30000
	ErrInvalidToken
)

// parameter checksum failure: include normal parameter and third party
const (
	ErrParams Error = iota + 40000
	ErrAuth
)

// errors within the system, include sql, redis ...etc
const (
	ErrSQL Error = iota + 50000
	ErrSystem
)

var ErrMsg = map[Error]string{
	ErrNone:             "Success",
	ErrParams:           "Invalid parameters",
	ErrSQL:              "Sql exec failed",
	ErrUnsupportedToken: "Unsupported token",
	ErrInvalidToken:     "Invalid token",
	ErrSystem:           "System error",
	ErrAuth:             "Invalid auth",
}
