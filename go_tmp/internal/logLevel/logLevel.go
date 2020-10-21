package logLevel

type LogLevel int8

const (
	QUIET  LogLevel = 0
	NORMAL LogLevel = 1
	DEBUG  LogLevel = 2
)

var LOG_LEVEL = NORMAL
