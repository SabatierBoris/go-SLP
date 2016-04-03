package packet

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PacketVersion is use for get the version of a packet.
type PacketVersion uint8

// Values of supported SLP version.
const (
	_          PacketVersion = iota
	V1         PacketVersion = iota
	V2         PacketVersion = iota
	NB_VERSION PacketVersion = iota
)

var Encoding = binary.BigEndian

// Packet is the main struct for SLP packet.
// It's generique depending of the packetVersion
// and Function
type Packet struct {
	Version PacketVersion
	Header  Header
	Data    SLPFunction
}

// GetPacket create a Packet with data read in a io.Reader.
func GetPacket(data io.Reader) (p Packet, err error) {
	var v PacketVersion

	if err = binary.Read(data, Encoding, &v); err != nil {
		err = fmt.Errorf("Error during parsing packet version : %s", err)
		return
	}
	p.Version = v

	var h Header
	if h, err = GetHeader(p.Version); err != nil {
		err = fmt.Errorf("Error during getting header type : %s", err)
		return
	}
	if err = h.Read(data); err != nil {
		err = fmt.Errorf("Error during parsing header : %s", err)
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
