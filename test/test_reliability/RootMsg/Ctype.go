package RootMsg

const (
	VERITIFY             = 0x00
	STAT_ONLINE          = 0x01
	CYCLE                = 0x02
	ONLINE_TIME          = 0x03
	P2R_GET_RANGE_TMP_ID = 0x4 // parent to root, get temp id
	L2P_GET_TMP_ID       = 0x5 // leaf <-> parent, get one temp id
	P2R_CONFIRM_TMP_ID   = 0x6 // parent to root, confirm reg temp id
	ADD_NODE             = 0x10
	SET_NEW_PARENT       = 0x11
	DEL_NODE             = 0x12
	CMD                  = 0x20
	DIR                  = 0x21
	PUT_TO_NODE_FILE     = 0x22
	GET_NODE_FILE        = 0x23
	DEL_FILE             = 0x24
)
