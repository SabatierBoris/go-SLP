package packet

import (
	"encoding/binary"
	"fmt"
	"io"
)

type SLPFunction uint8

const (
	SLPSrvReq      SLPFunction = 1
	SLPSrvRply     SLPFunction = 2
	SLPSrvReg      SLPFunction = 3
	SLPSrvDereg    SLPFunction = 4
	SLPSrvAck      SLPFunction = 5
	SLPAttrRqst    SLPFunction = 6
	SLPAttrRply    SLPFunction = 7
	SLPDAAdvert    SLPFunction = 8
	SLPSrvTypeRqst SLPFunction = 9
	SLPSrvTypeRply SLPFunction = 10
)

type SLPHeaderFlags uint8

const (
	_         SLPHeaderFlags = 1 << iota
	_         SLPHeaderFlags = 1 << iota
	_         SLPHeaderFlags = 1 << iota
	SLPFlagsF SLPHeaderFlags = 1 << iota
	SLPFlagsA SLPHeaderFlags = 1 << iota
	SLPFlagsU SLPHeaderFlags = 1 << iota
	SLPFlagsM SLPHeaderFlags = 1 << iota
	SLPFlagsO SLPHeaderFlags = 1 << iota
)

type SLPHeader struct {
	Version       uint8
	Function      SLPFunction
	Length        uint16
	Flags         SLPHeaderFlags
	Dialect       uint8
	Language_code uint16
	Char_encoding uint16
	Xid           uint16
}

func (h *SLPHeader) Print() (s string) {
	s = fmt.Sprintf("Version : %d", h.Version)
	s = fmt.Sprintf("%s\nFunction : %s", s, h.Function.Print())
	s = fmt.Sprintf("%s\nLength : %d", s, h.Length)
	s = fmt.Sprintf("%s\nFlags : %s", s, h.Flags.Print())
	s = fmt.Sprintf("%s\nDialect : %d", s, h.Dialect)
	s = fmt.Sprintf("%s\nLanguage_code : %d", s, h.Language_code)
	s = fmt.Sprintf("%s\nChar_encoding : %d", s, h.Char_encoding)
	s = fmt.Sprintf("%s\nXid : %d", s, h.Xid)
	return
}

func (f *SLPFunction) Print() (s string) {
	switch *f {
	case SLPSrvReq:
		s = "SLPSrvReq"
	case SLPSrvRply:
		s = "SLPSrvRply"
	case SLPSrvReg:
		s = "SLPSrvReg"
	case SLPSrvDereg:
		s = "SLPSrvDereg"
	case SLPSrvAck:
		s = "SLPSrvAck"
	case SLPAttrRqst:
		s = "SLPAttrRqst"
	case SLPAttrRply:
		s = "SLPAttrRply"
	case SLPDAAdvert:
		s = "SLPDAAdvert"
	case SLPSrvTypeRqst:
		s = "SLPSrvTypeRqst"
	case SLPSrvTypeRply:
		s = "SLPSrvTypeRply"
	}
	return
}

func (f *SLPHeaderFlags) Print() (s string) {
	s = ""
	if (*f & SLPFlagsO) != 0 {
		s = fmt.Sprintf("%s O", s)
	}
	if (*f & SLPFlagsM) != 0 {
		s = fmt.Sprintf("%s M", s)
	}
	if (*f & SLPFlagsU) != 0 {
		s = fmt.Sprintf("%s U", s)
	}
	if (*f & SLPFlagsA) != 0 {
		s = fmt.Sprintf("%s A", s)
	}
	if (*f & SLPFlagsF) != 0 {
		s = fmt.Sprintf("%s F", s)
	}
	return
}
