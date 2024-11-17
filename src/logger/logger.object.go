package logger

const (
	FATAL = iota
	ERROR
	WARN
	INFO
	DEBUG
)

var Levels = map[string]int{
	"Fatal": FATAL,
	"Error": ERROR,
	"Warn":  WARN,
	"Info":  INFO,
	"Debug": DEBUG,
}
