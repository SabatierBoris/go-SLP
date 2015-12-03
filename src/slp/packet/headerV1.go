package packet

import (
	"encoding/binary"
	"fmt"
	"io"
)

type HeaderFlags uint8

const (
	_      HeaderFlags = 1 << iota
	_      HeaderFlags = 1 << iota
	_      HeaderFlags = 1 << iota
	FlagsF HeaderFlags = 1 << iota
	FlagsA HeaderFlags = 1 << iota
	FlagsU HeaderFlags = 1 << iota
	FlagsM HeaderFlags = 1 << iota
	FlagsO HeaderFlags = 1 << iota
)

type HeaderV1 struct {
	Function      Function
	Length        uint16
	Flags         HeaderFlags
	Dialect       uint8
	Language_code uint16
	Char_encoding uint16
	Xid           uint16
}

func (h *HeaderV1) Read(data io.Reader) (err error) {
	if err = binary.Read(data, Encoding, h); err != nil {
		err = fmt.Errorf("Error during parsing HeaderV1 : %s", err)
		return
	}
	return
}

func (h *HeaderV1) GetFunction() (f Function) {
	f = h.Function
	return
}

func HeaderV1Constructor() Header {
	return &HeaderV1{}
}

func init() {
	RegisterHeader(V1, HeaderV1Constructor)
}

//func (h *SLPHeader) Print() (s string) {
//	s = fmt.Sprintf("Version : %d", h.Version)
//	s = fmt.Sprintf("%s\nFunction : %s", s, h.Function.Print())
//	s = fmt.Sprintf("%s\nLength : %d", s, h.Length)
//	s = fmt.Sprintf("%s\nFlags : %s", s, h.Flags.Print())
//	s = fmt.Sprintf("%s\nDialect : %d", s, h.Dialect)
//	s = fmt.Sprintf("%s\nLanguage_code : %d", s, h.Language_code)
//	s = fmt.Sprintf("%s\nChar_encoding : %d", s, h.Char_encoding)
//	s = fmt.Sprintf("%s\nXid : %d", s, h.Xid)
//	return
//}
//
//func (f *SLPFunction) Print() (s string) {
//	switch *f {
//	case SLPSrvReq:
//		s = "SLPSrvReq"
//	case SLPSrvRply:
//		s = "SLPSrvRply"
//	case SLPSrvReg:
//		s = "SLPSrvReg"
//	case SLPSrvDereg:
//		s = "SLPSrvDereg"
//	case SLPSrvAck:
//		s = "SLPSrvAck"
//	case SLPAttrRqst:
//		s = "SLPAttrRqst"
//	case SLPAttrRply:
//		s = "SLPAttrRply"
//	case SLPDAAdvert:
//		s = "SLPDAAdvert"
//	case SLPSrvTypeRqst:
//		s = "SLPSrvTypeRqst"
//	case SLPSrvTypeRply:
//		s = "SLPSrvTypeRply"
//	}
//	return
//}
//
//func (f *SLPHeaderFlags) Print() (s string) {
//	s = ""
//	if (*f & SLPFlagsO) != 0 {
//		s = fmt.Sprintf("%s O", s)
//	}
//	if (*f & SLPFlagsM) != 0 {
//		s = fmt.Sprintf("%s M", s)
//	}
//	if (*f & SLPFlagsU) != 0 {
//		s = fmt.Sprintf("%s U", s)
//	}
//	if (*f & SLPFlagsA) != 0 {
//		s = fmt.Sprintf("%s A", s)
//	}
//	if (*f & SLPFlagsF) != 0 {
//		s = fmt.Sprintf("%s F", s)
//	}
//	return
//}
