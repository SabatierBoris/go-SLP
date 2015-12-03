package packet

import (
	"encoding/binary"
	"fmt"
	"io"
)

type ErrorCode uint16

const (
	V1_SUCCESS                = iota
	V1_PROTOCOL_PARSE_ERROR   = iota
	V1_INVALID_REGISTRATION   = iota
	V1_SCOPE_NOT_SUPPORTED    = iota
	V1_CHARSET_NOT_UNDERSTOOD = iota
	V1_AUTHENTICATION_ABSENT  = iota
	V1_AUTHENTICATION_FALIED  = iota
)

type SrvAckV1 struct {
	ErrorCode ErrorCode
}

func (f *SrvAckV1) Read(data io.Reader) (err error) {
	if err = binary.Read(data, Encoding, f); err != nil {
		err = fmt.Errorf("Error during parsing SrvAckV1 : %s", err)
		return
	}
	return
}

func SrvAckV1Constructor() SLPFunction {
	return &SrvAckV1{}
}

func init() {
	RegisterFunction(V1, SrvAck, SrvAckV1Constructor)
}
