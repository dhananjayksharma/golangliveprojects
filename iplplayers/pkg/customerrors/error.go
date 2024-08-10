package customerrors

// user interfaced error
var ERR_INVALID_PLAYER_CODE = "invalid player code found, player code:%s"
var ERR_PLAYER_CODE_GENERATION_FAILED = "error: unable to create player code"
var ERR_INVALID_PLAYER_ACTIVATION_CODE = "error: invalid player activation code"
var ERR_INTERNAL_SERVER_ERROR = "error: internal server error"

// DB error
var ERR_REDIS_DB_KEY_NOT_FOUND = "redis: nil"
var ERR_MYSQL_DB_UNKNOWN_COLUMN = "Error 1054 (42S22): Unknown column"
