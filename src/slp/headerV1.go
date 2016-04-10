package slp

import (
	"encoding/binary"
	"io"
	"regexp"
)

// HeaderFlags is use for get flags store in Header V1
type HeaderFlags uint8

// Values of V1 Flags
const (
	_        HeaderFlags = 1 << iota //0x01
	_        HeaderFlags = 1 << iota //0x02
	_        HeaderFlags = 1 << iota //0x04
	V1FlagsF HeaderFlags = 1 << iota //0x08
	V1FlagsA HeaderFlags = 1 << iota //0x10
	V1FlagsU HeaderFlags = 1 << iota //0x20
	V1FlagsM HeaderFlags = 1 << iota //0x40
	V1FlagsO HeaderFlags = 1 << iota //0x80
)

// HeaderV1 is the structure of an SLP Header for V1
type HeaderV1 struct {
	FunctionID   FunctionID
	Length       uint16
	Flags        HeaderFlags
	Dialect      uint8
	LanguageCode [2]byte
	CharEncoding uint16
	Xid          uint16
}

// SupportedLanguages are the list of supportedLanguage
var SupportedLanguages = map[string]string{
	"aa": "Afar",
	"ab": "Abkhazian",
	"af": "Afrikaans",
	"am": "Amharic",
	"ar": "Arabic",
	"as": "Assamese",
	"ay": "Aymara",
	"az": "Azerbaijani",
	"ba": "Bashkir",
	"be": "Byelorussian",
	"bg": "Bulgarian",
	"bh": "Bihari",
	"bi": "Bislama",
	"bn": "Bengali",
	"bo": "Tibetan",
	"br": "Breton",
	"ca": "Catalan",
	"co": "Corsican",
	"cs": "Czech",
	"cy": "Welsh",
	"da": "Danish",
	"de": "German",
	"dz": "Bhutani",
	"el": "Greek",
	"en": "English",
	"eo": "Esperanto",
	"es": "Spanish",
	"et": "Estonian",
	"eu": "Basque",
	"fa": "Persian",
	"fi": "Finnish",
	"fj": "Fiji",
	"fo": "Faeroese",
	"fr": "French",
	"fy": "Frisiansa",
	"ga": "Irish",
	"gd": "Scots Gaelic",
	"gl": "Galician",
	"gn": "Guarani",
	"gu": "Gujarati",
	"ha": "Hausa",
	"he": "Hebrew",
	"hi": "Hindi",
	"hr": "Croatian",
	"hu": "Hungarian",
	"hy": "Armenian",
	"ia": "Interlingua",
	"in": "Indonesian",
	"ie": "Interlingue",
	"ik": "Inupiak",
	"is": "Icelandic",
	"it": "Italian",
	"ja": "Japanese",
	"jw": "Javanese",
	"ka": "Georgian",
	"kk": "Kazakh",
	"kl": "Greenlandic",
	"km": "Cambodian",
	"kn": "Kannada",
	"ko": "Korean",
	"ks": "Kashmiri",
	"ku": "Kurdish",
	"ky": "Kirghiz",
	"la": "Latin",
	"ln": "Lingala",
	"lo": "Laothian",
	"lt": "Lithuanian",
	"lv": "Latvian, Lettish",
	"mg": "Malagasy",
	"mi": "Maori",
	"mk": "Macedonian",
	"ml": "Malayalam",
	"mn": "Mongolian",
	"mo": "Moldavian",
	"mr": "Marathi",
	"ms": "Malay",
	"mt": "Maltese",
	"my": "Burmese",
	"na": "Nauru",
	"ne": "Nepali",
	"nl": "Dutch",
	"no": "Norwegian",
	"oc": "Occitan",
	"om": "(Afan) Oromo",
	"or": "Oriya",
	"pa": "Punjabi",
	"pl": "Polish",
	"ps": "Pashto, Pushto",
	"pt": "Portuguese",
	"qu": "Quechua",
	"rm": "Rhaeto-Romance",
	"rn": "Kirundi",
	"ro": "Romanian",
	"ru": "Russian",
	"rw": "Kinyarwanda",
	"sd": "Sindhi",
	"sg": "Sangro",
	"sh": "Serbo",
	"si": "Singhalese",
	"sk": "Slovak",
	"sl": "Slovenian",
	"sm": "Samoan",
	"sn": "Shona",
	"so": "Somali",
	"sq": "Albanian",
	"sr": "Serbian",
	"ss": "Siswati",
	"st": "Sesotho",
	"su": "Sundanese",
	"sv": "Swedish",
	"sw": "Swahili",
	"ta": "Tamil",
	"te": "Telugu",
	"tg": "Tajik",
	"th": "Thai",
	"ti": "Tigrinya",
	"tk": "Turkmen",
	"tl": "Tagalog",
	"tn": "Setswana",
	"to": "Tonga",
	"tr": "Turkish",
	"ts": "Tsonga",
	"tt": "Tatar",
	"tw": "Twi",
	"ug": "Uigar",
	"uk": "Ukrainian",
	"ur": "Urdu",
	"uz": "Uzbek",
	"vi": "Vietnamese",
	"vo": "Volapuk",
	"wo": "Wolof",
	"xh": "Xhosa",
	"yi": "Yiddish",
	"yo": "Yoruba",
	"za": "Zhuang",
	"zh": "Chinese",
	"zu": "Zulu",
}

// Validate check if the Header is conforme
func (h *HeaderV1) Validate() (err error) {
	if h.Flags&0x7 != 0 {
		err = &FlagError{"rsvd", 0, h.Flags & 0x7}
		return
	}
	if h.Dialect != 0 {
		err = &DialectError{h.Dialect}
		return
	}
	lang, err := h.GetLanguageCode()
	if err != nil {
		return
	}
	_, ok := SupportedLanguages[lang]
	if !ok {
		err = &LanguageError{&lang}
		return
	}
	return
}

// HasFlags check if the flag f are rised in the header
func (h *HeaderV1) HasFlags(f HeaderFlags) (r bool, err error) {
	r = ((h.Flags & f) == f)
	return
}

// GetFlags return all flag in the header
func (h *HeaderV1) GetFlags() (f HeaderFlags, err error) {
	f = h.Flags
	return
}

// GetLanguageCode return the full name of 2 char language Code
func (h *HeaderV1) GetLanguageCode() (r string, err error) {
	r = string(h.LanguageCode[:2])
	re := regexp.MustCompile("^[a-zA-Z0-9_]{2}$")
	if !re.MatchString(r) {
		err = &LanguageError{nil}
		return
	}
	return
}

// Read parse the data for extract packet information
func (h *HeaderV1) Read(data io.Reader) (err error) {
	if err = binary.Read(data, Encoding, h); err != nil {
		err = &ReadError{}
		return
	}
	return
}

// GetFunction return the functionID
func (h *HeaderV1) GetFunction() (f FunctionID) {
	f = h.FunctionID
	return
}

// HeaderV1Constructor is the constructor for HeaderV1 packet
func HeaderV1Constructor() Header {
	return &HeaderV1{}
}

func init() {
	RegisterHeader(V1, HeaderV1Constructor)
}
