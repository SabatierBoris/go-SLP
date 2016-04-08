package packet

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Version is use for get the version of a packet.
type Version uint8

// Values of supported SLP version.
const (
	_         Version = iota
	V1        Version = iota
	V2        Version = iota
	NbVersion Version = iota
)

var Encoding = binary.BigEndian

// Packet is the main struct for SLP packet.
// It's generique depending of the Version
// and Function
type Packet struct {
	Version Version
	Header  Header
	Data    SLPFunction
}

// GetPacket create a Packet with data read in a io.Reader.
func GetPacket(data io.Reader) (p Packet, err error) {
	var v Version

	if err = binary.Read(data, Encoding, &v); err != nil {
		err = &ReadError{}
		return
	}
	p.Version = v

	var h Header
	if h, err = GetHeader(p.Version); err != nil {
		return
	}
	if err = h.Read(data); err != nil {
		return
	}
	if err = h.Validate(); err != nil {
		err = fmt.Errorf("Error during validation of the header : %s", err)
		return
	}
	p.Header = h

	var f SLPFunction
	if f, err = GetFunction(p.Version, p.Header.GetFunction()); err != nil {
		err = fmt.Errorf("Error during getting function type : %s", err)
		return
	}
	if err = f.Read(data); err != nil {
		err = fmt.Errorf("Error during parsing function : %s", err)
		return
	}
	p.Data = f
	return
}
