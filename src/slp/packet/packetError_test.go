package packet

import (
	"bytes"
	"reflect"
	"testing"
)

func TestVersionToto(t *testing.T) {
	var datas = []struct {
		data   []byte
		result error
	}{
		{[]byte{0xFF}, &VersionError{}},
		{[]byte{0x02}, &VersionError{}},
		{[]byte{}, &ReadError{}},
		{[]byte{0x01}, &ReadError{}},
	}

	for _, infos := range datas {
		buf := bytes.NewReader(infos.data)
		p, err := GetPacket(buf)
		if err == nil {
			t.Errorf("Test failed, expected an error, got:  '%v'", p)
		}
		if reflect.TypeOf(err) != reflect.TypeOf(infos.result) {
			t.Errorf("Test failed, expected an %s, got:  '%s'", reflect.TypeOf(infos.result), err)
		}
		t.Logf("Got the error %s", err)
	}
}

func TestFunctionUnknow(t *testing.T) {
	data := []byte{0x01, 0xFF, 0x00, 0x0D, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf) //TODO Check if the err is the right one
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%v'", p)
	}
}

func TestErrorFunctionParsing(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0D, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf) //TODO Check if the err is the right one
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%v'", p)
	}
}

func TestErrorFunctionUnsupported(t *testing.T) {
	data := []byte{0x01, 0x01, 0x00, 0x0D, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf) //TODO Check if the err is the right one
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%v'", p)
	}
}

func TestErrorV1Flags(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0E, 0x02, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf) //TODO Check if the err is the right one
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%v'", p)
	}
}

func TestErrorV1Dialect(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x01, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf) //TODO Check if the err is the right one
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%v'", p)
	}
}

func TestErrorV1NoLanguageCode(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf) //TODO Check if the err is the right one
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%v'", p)
	}
}

func TestErrorV1UnknownLanguageCode(t *testing.T) {
	data := []byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'z', 'z', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}
	buf := bytes.NewReader(data)

	p, err := GetPacket(buf) //TODO Check if the err is the right one
	if err == nil {
		t.Errorf("Test failed, expected an error, got:  '%v'", p)
	}
}
