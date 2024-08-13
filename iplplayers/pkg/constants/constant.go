package constants

const (
	CONST_PLAYER_PREFIX = "indpl"
	CONST_NUMBER_FIVE   = 5
)

const (
	DB_STATUS_REGISTERED = 0
	DB_STATUS_ACTIVE     = 1
	DB_STATUS_INACTIVE   = 3
	DB_STATUS_DEACTIVE   = 9
)

var (
	StatusMap = map[uint8]string{
		0: "REGISTERED",
		1: "ACTIVE",
		2: "INACTIVE",
		3: "DEACTIVE",
	}
)
