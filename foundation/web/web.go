package web

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}
