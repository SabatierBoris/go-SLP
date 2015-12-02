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
		result V1AckCode
	}{
		{[]byte{0x01, 0x01, 0x00, 0x0E, TODO_FLAG, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}, V1_ACK_OK},
		{[]byte{0x01, 0x01, 0x00, 0x0E, TODO_FLAG, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x01}, V1_ACK_PROTOCOL_PARSE_ERROR},
		{[]byte{0x01, 0x01, 0x00, 0x0E, TODO_FLAG, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x02}, V1_ACK_INVALID_REGISTRATION},
		{[]byte{0x01, 0x01, 0x00, 0x0E, TODO_FLAG, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x03}, V1_ACK_SCOPE_NOT_SUPPORTED},
		{[]byte{0x01, 0x01, 0x00, 0x0E, TODO_FLAG, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x04}, V1_ACK_AUTHENTICATION_ABSENT},
		{[]byte{0x01, 0x01, 0x00, 0x0E, TODO_FLAG, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x05}, V1_ACK_AUTHENTICATION_FAILED},
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
		if p.Header.Function != V1SrvAck {
			t.Errorf("Test failed, expected Function : '%d' , got:  '%d'", V1SrvAck, p.Header.Function)
		}
		if p.Data != infos.result {
			t.Errorf("Test failed, expected Data : '%d' , got:  '%d'", infos.result, p.Data)
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
