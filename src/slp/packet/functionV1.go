package packet

import (
	"encoding/binary"
	"io"
)

// ErrorCode is use for get the ErrorCode of an Ack packet
type ErrorCode uint16

// Values of supported Error Code.
const (
	V1Success              ErrorCode = iota
	V1ProtocolParseError   ErrorCode = iota
	V1InvalidRegistration  ErrorCode = iota
	V1ScopeNotSupported    ErrorCode = iota
	V1CharsetNotUnderstood ErrorCode = iota
	V1AuthenticationAbsent ErrorCode = iota
	V1AuthenticationFalied ErrorCode = iota
)

// SrvAckV1 is the structure of an SLP Ack packet for V1
type SrvAckV1 struct {
	ErrorCode ErrorCode
}

// Read parse the data for extract packet information
func (f *SrvAckV1) Read(data io.Reader) (err error) {
	if err = binary.Read(data, Encoding, f); err != nil {
		err = &ReadError{}
		return
	}
	return
}

// SrvAckV1Constructor is the constructor for SrvAckV1 packet
func SrvAckV1Constructor() SLPFunction {
	return &SrvAckV1{}
}

func init() {
	RegisterFunction(V1, SrvAck, SrvAckV1Constructor)
}
