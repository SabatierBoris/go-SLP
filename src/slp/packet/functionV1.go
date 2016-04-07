package packet

import (
	"encoding/binary"
	"fmt"
	"io"
)

type ErrorCode uint16

const (
	V1Success              = iota
	V1ProtocolParseError   = iota
	V1InvalidRegistration  = iota
	V1ScopeNotSupported    = iota
	V1CharsetNotUnderstood = iota
	V1AuthenticationAbsent = iota
	V1AuthenticationFalied = iota
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
