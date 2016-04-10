package packet

import (
	"io"
	"log"
	"sync"
)

// Header is the interface for all SLP Header version
type Header interface {
	Read(io.Reader) error
	GetFunction() Function
	HasFlags(HeaderFlags) (bool, error)
	GetFlags() (HeaderFlags, error)
	GetLanguageCode() (string, error)
	Validate() error
}

// HeaderContructor is the interface for all SLP Header constructor
type HeaderContructor func() Header

var headers = struct {
	m [NbVersion]HeaderContructor
	sync.RWMutex
}{}

// RegisterHeader permit to dynamically add supported SLP Header version
func RegisterHeader(version Version, constructor HeaderContructor) {
	headers.Lock()
	headers.m[version] = constructor
	headers.Unlock()
	log.Printf("Header V%d is registered\n", version)
}

// GetHeader get the SLP Header type depending of the SLP Version
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
