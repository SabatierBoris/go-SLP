package packet

import (
	"io"
	"log"
	"sync"
)

// Function is enum for SLP Function mapping
type Function uint8

// Values of supported SLP Function.
const (
	_           Function = iota
	SrvReq      Function = iota
	SrvRply     Function = iota
	SrvReg      Function = iota
	SrvDereg    Function = iota
	SrvAck      Function = iota
	AttrRqst    Function = iota
	AttrRply    Function = iota
	DAAdvert    Function = iota
	SrvTypeRqst Function = iota
	SrvTypeRply Function = iota
	NbFunction  Function = iota
)

// SLPFunction is the interface for all SLP Function version
type SLPFunction interface {
	Read(io.Reader) error
}

// FunctionContructor is the interface for all SLP Function constructor
type FunctionContructor func() SLPFunction

var functions = struct {
	m [NbVersion][NbFunction]FunctionContructor
	sync.RWMutex
}{}

// RegisterFunction permit to dynamically add supported SLP Function for an SLP version and Function id
func RegisterFunction(version Version, function Function, constructor FunctionContructor) {
	functions.Lock()
	functions.m[version][function] = constructor
	functions.Unlock()
	log.Printf("Function %d (V%d) is registered\n", function, version)
}

// GetFunction get the SLP Function type depending of the SLP Version and Function id
func GetFunction(id Version, function Function) (f SLPFunction, err error) {
	err = nil
	if function >= NbFunction {
		err = &FunctionError{function, nil}
		return
	}
	functions.RLock()
	ctor := functions.m[id][function]
	functions.RUnlock()
	if ctor == nil {
		err = &FunctionError{function, &id}
		return
	}
	f = ctor()
	return
}
