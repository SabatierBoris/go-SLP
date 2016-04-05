package packet

import (
	"fmt"
	"io"
	"log"
	"sync"
)

type Header interface {
	Read(io.Reader) error
	GetFunction() Function
	HasFlags(HeaderFlags) (bool, error)
	GetFlags() (HeaderFlags, error)
	GetLanguageCode() (string, error)
	Validate() error
}

type HeaderContructor func() Header

var headers = struct {
	m [NB_VERSION]HeaderContructor
	sync.RWMutex
}{}

type VersionError struct {
	version PacketVersion
}

func (e *VersionError) Error() string {
	return fmt.Sprintf("SLP V%d isn't supported", e.version)
}

func RegisterHeader(version PacketVersion, constructor HeaderContructor) {
	headers.Lock()
	headers.m[version] = constructor
	headers.Unlock()
	log.Printf("Header V%d is registered\n", version)
}

func GetHeader(id PacketVersion) (h Header, err error) {
	err = nil
	if id >= NB_VERSION {
		err = &VersionError{id}
		return
	}
	headers.RLock()
	ctor := headers.m[id]
	headers.RUnlock()
	if ctor == nil {
		err = &VersionError{id}
		return
	}
	h = ctor()
	return
}
