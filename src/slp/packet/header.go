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
	m [NB_VERSION]HeaderContructor
	sync.RWMutex
}{}

func RegisterHeader(version PacketVersion, constructor HeaderContructor) {
	headers.Lock()
	headers.m[version] = constructor
	headers.Unlock()
	fmt.Printf("Header V%d is registered\n", version)
}

func GetHeader(id PacketVersion) (h Header, err error) {
	err = nil
	if id > NB_VERSION {
		err = fmt.Errorf("SLP V%d is not supported", id)
		return
	}
	headers.RLock()
	ctor := headers.m[id]
	headers.RUnlock()
	if ctor == nil {
		err = fmt.Errorf("SLP V%d is not supported", id)
		return
	}
	h = ctor()
	return
}
