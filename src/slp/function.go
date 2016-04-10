package slp

import (
	"io"
	"log"
	"sync"
)

// FunctionID is enum for SLP Function mapping
type FunctionID uint8

// Values of supported SLP FunctionID.
const (
	_           FunctionID = iota
	SrvReq      FunctionID = iota
	SrvRply     FunctionID = iota
	SrvReg      FunctionID = iota
	SrvDereg    FunctionID = iota
	SrvAck      FunctionID = iota
	AttrRqst    FunctionID = iota
	AttrRply    FunctionID = iota
	DAAdvert    FunctionID = iota
	SrvTypeRqst FunctionID = iota
	SrvTypeRply FunctionID = iota
	NbFunction  FunctionID = iota
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

// RegisterFunction permit to dynamically add supported SLP Function for an SLP version and FunctionID
func RegisterFunction(version Version, functionID FunctionID, constructor FunctionContructor) {
	functions.Lock()
	functions.m[version][functionID] = constructor
	functions.Unlock()
	log.Printf("Function %d (V%d) is registered\n", functionID, version)
}

// GetFunction get the SLP Function type depending of the SLP Version and FunctionID
func GetFunction(id Version, functionID FunctionID) (f SLPFunction, err error) {
	err = nil
	if functionID >= NbFunction {
		err = &FunctionError{functionID, nil}
		return
	}
	functions.RLock()
	ctor := functions.m[id][functionID]
	functions.RUnlock()
	if ctor == nil {
		err = &FunctionError{functionID, &id}
		return
	}
	f = ctor()
	return
}
