package packet

import (
	"io"
	"log"
	"sync"
)

type Function uint8

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

type SLPFunction interface {
	Read(io.Reader) error
}

type FunctionContructor func() SLPFunction

var functions = struct {
	m [NbVersion][NbFunction]FunctionContructor
	sync.RWMutex
}{}

func RegisterFunction(version Version, function Function, constructor FunctionContructor) {
	functions.Lock()
	functions.m[version][function] = constructor
	functions.Unlock()
	log.Printf("Function %d (V%d) is registered\n", function, version)
}

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
