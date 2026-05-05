package errs

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func (err *Error) Encode() ([]byte, string, error) {
	data, e := json.Marshal(err)
	return data, "application/json", e
}

func (err *Error) Error() string {
	return err.Message
}

// ErrCode represents an error code in the system.
type ErrCode struct {
	value int
}

// String returns the string representation of the error code.
func (ec ErrCode) String() string {
	return codeNames[ec]
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (ec *ErrCode) UnmarshalText(data []byte) error {
	errName := string(data)

	v, exists := codeNumbers[errName]
	if !exists {
		return fmt.Errorf("err code %q does not exist", errName)
	}

	*ec = v

	return nil
}

type Error struct {
	Code     ErrCode `json:"code"`
	Message  string  `json:"message"`
	FuncName string  `json:"-"`
	FileName string  `json:"-"`
}

func New(code ErrCode, err error) *Error {
	pc, filename, line, _ := runtime.Caller(1)

	return &Error{
		Code:     code,
		Message:  err.Error(),
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fmt.Sprintf("%s:%d", filename, line),
	}
}

func Newf(code ErrCode, format string, v ...any) *Error {
	pc, filename, line, _ := runtime.Caller(1)

	return &Error{
		Code:     code,
		Message:  fmt.Sprintf(format, v...),
		FuncName: runtime.FuncForPC(pc).Name(),
		FileName: fmt.Sprintf("%s:%d", filename, line),
	}
}
