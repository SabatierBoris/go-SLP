package packet

import (
	"bytes"
	"testing"
)

func TestV1SrvReq(t *testing.T) {
	// TODO
}

func TestV1SrvRply(t *testing.T) {
	// TODO
}

func TestV1SrvReg(t *testing.T) {
	// TODO
}

func TestV1SrvDereg(t *testing.T) {
	// TODO
}

func TestV1SrvAck(t *testing.T) {
	var datas = []struct {
		data   []byte
		result ErrorCode
	}{
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}, V1_SUCCESS},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x01}, V1_PROTOCOL_PARSE_ERROR},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x02}, V1_INVALID_REGISTRATION},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x03}, V1_SCOPE_NOT_SUPPORTED},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x04}, V1_CHARSET_NOT_UNDERSTOOD},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x05}, V1_AUTHENTICATION_ABSENT},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x06}, V1_AUTHENTICATION_FALIED},
	}

	for _, infos := range datas {
		buf := bytes.NewReader(infos.data)
		p, err := GetPacket(buf)
		if err != nil {
			t.Errorf("Test failed, expected no error, got:  '%s'", err)
		}
		if p.Version != V1 {
			t.Errorf("Test failed, expected packet version : '%d' , got:  '%d'", V1, p.Version)
		}
		if f := p.Header.GetFunction(); f != SrvAck {
			t.Errorf("Test failed, expected Function : '%d' , got:  '%d'", SrvAck, f)
		}
		data := p.Data.(*SrvAckV1)
		if data.ErrorCode != infos.result {
			t.Errorf("Test failed, expected Data : '%d' , got:  '%d'", infos.result, data.ErrorCode)
		}
	}
}

func TestV1AttrRqst(t *testing.T) {
	// TODO
}

func TestV1AttrRply(t *testing.T) {
	// TODO
}

func TestV1DAAdvert(t *testing.T) {
	// TODO
}

func TestV1SrvTypeRqst(t *testing.T) {
	// TODO
}

func TestV1SrvTypeRply(t *testing.T) {
	// TODO
}

func TestV1Flags(t *testing.T) {
	// TODO
}

func TestV1LanguageCode(t *testing.T) {
	// TODO
}

func TestV1CharacterEncoding(t *testing.T) {
	// TODO
}
