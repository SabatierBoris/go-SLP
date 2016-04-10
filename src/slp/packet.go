package slp

import (
	"encoding/binary"
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

// Encoding is the current encoding way for SLP message
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
		return
	}
	p.Header = h

	var f SLPFunction
	if f, err = GetFunction(p.Version, p.Header.GetFunction()); err != nil {
		return
	}
	if err = f.Read(data); err != nil {
		return
	}
	p.Data = f
	return
}
