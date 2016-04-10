package slp

import "fmt"

// ReadError is the error send when there is an issue during reading data
type ReadError struct{}

func (e *ReadError) Error() string {
	return "Cannot read data"
}

// VersionError is the error send when the SLP version isn't knowed or supported
type VersionError struct {
	version Version
}

func (e *VersionError) Error() string {
	return fmt.Sprintf("SLP V%d isn't supported", e.version)
}

// FunctionError is the error send when the SLP function isn't knowed or not supported for the SLP version
type FunctionError struct {
	function Function
	version  *Version
}

func (e *FunctionError) Error() string {
	var s string
	if e.version != nil {
		s = fmt.Sprintf("SLP function %d for V%d isn't supported", e.function, *e.version)
	} else {
		s = fmt.Sprintf("SLP function %d isn't supported", e.function)
	}
	return s
}

// FlagError is the error send when a Flag isn't set with the attended value
type FlagError struct {
	name   string
	target HeaderFlags
	value  HeaderFlags
}

func (e *FlagError) Error() string {
	return fmt.Sprintf("SLP flag error. %s is %d and should be %d", e.name, e.value, e.target)
}

// DialectError is the error send when the Dialect data is set to something else than 0
// For the moment, SLP RFCs say : Dialect should be to 0, it's might be use in futur
type DialectError struct {
	value uint8
}

func (e *DialectError) Error() string {
	return fmt.Sprintf("SLP dialect is %d and should be 0", e.value)
}

// LanguageError is the error send when the Language isn't set or not supported
type LanguageError struct {
	languageCode *string
}

func (e *LanguageError) Error() string {
	var s string
	if e.languageCode != nil {
		s = fmt.Sprintf("SLP LanguageCode '%s' isn't supported", *e.languageCode)
	} else {
		s = fmt.Sprintf("SLP LanguageCode isn't set")
	}
	return s
}
