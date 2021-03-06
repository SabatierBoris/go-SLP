package slp

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
		{[]byte{0x01, 0xFF, 0x00, 0x0D, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00}, &FunctionError{}},
		{[]byte{0x01, 0x05, 0x00, 0x0D, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00}, &ReadError{}},
		{[]byte{0x01, 0x01, 0x00, 0x0D, 0x00, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00}, &FunctionError{}},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x02, 0x00, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}, &FlagError{}},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x01, 'e', 'n', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}, &DialectError{}},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}, &LanguageError{}},
		{[]byte{0x01, 0x05, 0x00, 0x0E, 0x00, 0x00, 'z', 'z', 0x00, 0x03, 0x00, 0x00, 0x00, 0x00}, &LanguageError{}},
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
