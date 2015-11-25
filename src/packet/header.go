package packet

type SPLFunction int8

const (
	SrvReq      SPLFunction = 1
	SrvRply     SPLFunction = 2
	SrvReq      SPLFunction = 3
	SrvDereg    SPLFunction = 4
	SrvAck      SPLFunction = 5
	AttrRqst    SPLFunction = 6
	AttrRply    SPLFunction = 7
	DAAdvert    SPLFunction = 8
	SrvTypeRqst SPLFunction = 9
	SrvTypeRply SPLFunction = 10
)

type SPLHeaderFlags struct {
	flag_o bool
	flag_m bool
	flag_u bool
	flag_a bool
	flag_f bool
}

type SPLHeader struct {
	version       int8
	function      SPLFunction
	length        int16
	flags         SPLHeaderFlags
	rvds          [4]bool
	dialect       int8
	language_code int16
	char_encoding int16
	xid           int16
}
