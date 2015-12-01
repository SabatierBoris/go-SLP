package packet

import (
	"encoding/binary"
	"fmt"
	"io"
)

type PacketVersion uint8

const (
	_          PacketVersion = iota
	V1         PacketVersion = iota
	V2         PacketVersion = iota
	NB_VERSION PacketVersion = iota
)

var Encoding = binary.BigEndian

type Packet struct {
	Version PacketVersion
	Header  Header
}

func GetPacket(data io.Reader) (p Packet, err error) {
	var v PacketVersion

	if err = binary.Read(data, Encoding, &v); err != nil {
		err = fmt.Errorf("Error during parsing packet version : ", err)
		return
	}
	p.Version = v

	var h Header
	if h, err = GetHeader(v); err != nil {
		err = fmt.Errorf("Error during getting header type : ", err)
		return
	}
	h.Read(data)
	p.Header = h
	return
}
