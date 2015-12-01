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
