package logger

const (
	Fatal = iota
	Error
	Warn
	Notice
	Info
	Debug
)

var Levels = map[string]int{
	"Fatal":  Fatal,
	"Error":  Error,
	"Warn":   Warn,
	"Notice": Notice,
	"Info":   Info,
	"Debug":  Debug,
}
