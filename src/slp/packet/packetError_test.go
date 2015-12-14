package packet

import (
	"bytes"
	"testing"
)

func TestVersionUnknow(t *testing.T) {
	data := []byte{0xFF}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestVersionUnsupported(t *testing.T) {
	data := []byte{0x02}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestNoData(t *testing.T) {
	data := []byte{}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestMissingError(t *testing.T) {
	data := []byte{0x01}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestFunctionUnknow(t *testing.T) {
	data := []byte{0x01, 0xFF, 0x00, 0x0D, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestErrorFunctionParsing(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0D, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestErrorFunctionUnsupported(t *testing.T) {
	data := []byte{0x01, 0x01, 0x00, 0x0D, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestErrorV1Flags(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0E, 0x02, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestErrorV1Dialect(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x01, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestErrorV1NoLanguageCode(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}

func TestErrorV1UnknownLanguageCode(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'z', 'z', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf)
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%s'", p)
	}
}
