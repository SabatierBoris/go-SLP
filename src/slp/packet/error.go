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

type FunctionError struct {
	function Function
	version  *Version
}

func (e *FunctionError) Error() string {
	if e.version != nil {
		return fmt.Sprintf("SLP function %d for V%d isn't supported", e.function, *e.version)
	} else {
		return fmt.Sprintf("SLP function %d isn't supported", e.function)
	}
}

type FlagError struct {
	name   string
	target HeaderFlags
	value  HeaderFlags
}

func (e *FlagError) Error() string {
	return fmt.Sprintf("SLP flag error. %s is %d and should be %d", e.name, e.value, e.target)
}

type DialectError struct {
	value uint8
}

func (e *DialectError) Error() string {
	return fmt.Sprintf("SLP dialect is %d and should be 0", e.value)
}

type LanguageError struct {
	languageCode *string
}

func (e *LanguageError) Error() string {
	if e.languageCode != nil {
		return fmt.Sprintf("SLP LanguageCode '%s' isn't supported", *e.languageCode)
	} else {
		return fmt.Sprintf("SLP LanguageCode isn't set")
	}
}
