package logger

const (
	Fatal = iota
	Error
	Warn
	Info
	Debug
)

var Levels = map[string]int{
	"Fatal": Fatal,
	"Error": Error,
	"Warn":  Warn,
	"Info":  Info,
	"Debug": Debug,
}
