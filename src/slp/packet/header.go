package packet

import (
	"fmt"
	"io"
	"sync"
)

type Header interface {
	Read(io.Reader) error
}

type HeaderContructor func() Header

var headers = struct {
	m map[PacketVersion]HeaderContructor
	sync.RWMutex
}{m: make(map[PacketVersion]HeaderContructor)}

func RegisterHeader(version PacketVersion, constructor HeaderContructor) {
	headers.Lock()
	headers.m[version] = constructor
	headers.Unlock()
	fmt.Println("Header V", version, " is registered")
}

func GetHeader(id PacketVersion) (h Header, err error) {
	err = nil
	headers.RLock()
	ctor, ok := headers.m[id]
	headers.RUnlock()
	if !ok {
		err = fmt.Errorf("SLP V%d is not supported", id)
		return
	}
	h = ctor()
	return
}
