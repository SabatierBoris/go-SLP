package packet

import "fmt"

type ReadError struct{}

func (e *ReadError) Error() string {
	return "Cannot read data"
}

type VersionError struct {
	version Version
}

func (e *VersionError) Error() string {
	return fmt.Sprintf("SLP V%d isn't supported", e.version)
}
