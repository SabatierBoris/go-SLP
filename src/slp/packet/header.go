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
	m [NbVersion]HeaderContructor
	sync.RWMutex
}{}

type VersionError struct {
	version Version
}

func (e *VersionError) Error() string {
	return fmt.Sprintf("SLP V%d isn't supported", e.version)
}

func RegisterHeader(version Version, constructor HeaderContructor) {
	headers.Lock()
	headers.m[version] = constructor
	headers.Unlock()
	log.Printf("Header V%d is registered\n", version)
}

func GetHeader(id Version) (h Header, err error) {
	err = nil
	if id >= NbVersion {
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
