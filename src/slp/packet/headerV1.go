package packet

import (
	"encoding/binary"
	"fmt"
	"io"
	"regexp"
)

type HeaderFlags uint8

const (
	_         HeaderFlags = 1 << iota //0x01
	_         HeaderFlags = 1 << iota //0x02
	_         HeaderFlags = 1 << iota //0x04
	V1_FlagsF HeaderFlags = 1 << iota //0x08
	V1_FlagsA HeaderFlags = 1 << iota //0x10
	V1_FlagsU HeaderFlags = 1 << iota //0x20
	V1_FlagsM HeaderFlags = 1 << iota //0x40
	V1_FlagsO HeaderFlags = 1 << iota //0x80
)

type HeaderV1 struct {
	Function      Function
	Length        uint16
	Flags         HeaderFlags
	Dialect       uint8
	Language_code [2]byte
	Char_encoding uint16
	Xid           uint16
}

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

func (h *HeaderV1) Validate() (err error) {
	if h.Flags&0x7 != 0 {
		err = fmt.Errorf("Error rsvd should be to 0 and is %d", h.Flags&0x7)
		return
	}
	if h.Dialect != 0 {
		err = fmt.Errorf("Error dialect should be to 0 and is %d", h.Flags&0x7)
		return
	}
	lang, err := h.GetLanguageCode()
	if err != nil {
		return
	}
	_, ok := SupportedLanguages[lang]
	if !ok {
		err = fmt.Errorf("Error LanguageCode %s isn't supported", lang)
		return
	}
	return
}

func (h *HeaderV1) HasFlags(f HeaderFlags) (r bool, err error) {
	r = ((h.Flags & f) == f)
	return
}

func (h *HeaderV1) GetFlags() (f HeaderFlags, err error) {
	f = h.Flags
	return
}

func (h *HeaderV1) GetLanguageCode() (r string, err error) {
	r = string(h.Language_code[:2])
	re := regexp.MustCompile("^[a-zA-Z0-9_]{2}$")
	if !re.MatchString(r) {
		err = fmt.Errorf("Error LanguageCode ins't set")
		return
	}
	return
}

func (h *HeaderV1) Read(data io.Reader) (err error) {
	if err = binary.Read(data, Encoding, h); err != nil {
		err = fmt.Errorf("Error during parsing HeaderV1 : %s", err)
		return
	}
	return
}

func (h *HeaderV1) GetFunction() (f Function) {
	f = h.Function
	return
}

func HeaderV1Constructor() Header {
	return &HeaderV1{}
}

func init() {
	RegisterHeader(V1, HeaderV1Constructor)
}

//func (h *SLPHeader) Print() (s string) {
//	s = fmt.Sprintf("Version : %d", h.Version)
//	s = fmt.Sprintf("%s\nFunction : %s", s, h.Function.Print())
//	s = fmt.Sprintf("%s\nLength : %d", s, h.Length)
//	s = fmt.Sprintf("%s\nFlags : %s", s, h.Flags.Print())
//	s = fmt.Sprintf("%s\nDialect : %d", s, h.Dialect)
//	s = fmt.Sprintf("%s\nLanguage_code : %d", s, h.Language_code)
//	s = fmt.Sprintf("%s\nChar_encoding : %d", s, h.Char_encoding)
//	s = fmt.Sprintf("%s\nXid : %d", s, h.Xid)
//	return
//}
//
//func (f *SLPFunction) Print() (s string) {
//	switch *f {
//	case SLPSrvReq:
//		s = "SLPSrvReq"
//	case SLPSrvRply:
//		s = "SLPSrvRply"
//	case SLPSrvReg:
//		s = "SLPSrvReg"
//	case SLPSrvDereg:
//		s = "SLPSrvDereg"
//	case SLPSrvAck:
//		s = "SLPSrvAck"
//	case SLPAttrRqst:
//		s = "SLPAttrRqst"
//	case SLPAttrRply:
//		s = "SLPAttrRply"
//	case SLPDAAdvert:
//		s = "SLPDAAdvert"
//	case SLPSrvTypeRqst:
//		s = "SLPSrvTypeRqst"
//	case SLPSrvTypeRply:
//		s = "SLPSrvTypeRply"
//	}
//	return
//}
//
//func (f *SLPHeaderFlags) Print() (s string) {
//	s = ""
//	if (*f & SLPFlagsO) != 0 {
//		s = fmt.Sprintf("%s O", s)
//	}
//	if (*f & SLPFlagsM) != 0 {
//		s = fmt.Sprintf("%s M", s)
//	}
//	if (*f & SLPFlagsU) != 0 {
//		s = fmt.Sprintf("%s U", s)
//	}
//	if (*f & SLPFlagsA) != 0 {
//		s = fmt.Sprintf("%s A", s)
//	}
//	if (*f & SLPFlagsF) != 0 {
//		s = fmt.Sprintf("%s F", s)
//	}
//	return
//}
