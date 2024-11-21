package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/dykoffi/forexauto/src/config"
)

type Interface interface {
	writeInFile(message string, level string)
	Fatal(message string)
	Error(message string)
	Warning(message string)
	Info(message string)
	Debug(message string)
}

type Service struct {
	level  int
	folder string
}

var (
	iLoggerService Service
	once           sync.Once
)

func New(config config.Interface) *Service {

	once.Do(func() {
		level, exist := Levels[config.GetOrDefault("LOG_LEVEL", "Debug")]

		if !exist {
			level = DEBUG
		}

		iLoggerService = Service{
			level:  level,
			folder: config.GetOrDefault("LOG_FOLDER", "logs"),
		}
	})

	return &iLoggerService

}

func (ls *Service) writeInFile(message string, level string) {

	if levelVal, exist := Levels[level]; !exist || levelVal > ls.level {
		return
	}

	dateTime := time.Now().Format("02-01-2006 15:04:05")

	logFormat := fmt.Sprintf("[%s] %s %s \n", dateTime, level, message)

	if level == "Error" {
		pc, fileName, line, _ := runtime.Caller(2)
		caller2 := runtime.FuncForPC(pc).Name()
		logFormat = fmt.Sprintf("[%s] %s %s - %s (%d) {%s}\n", dateTime, level, message, fileName, line, caller2)
	}

	date := time.Now().Format("02-01-2006")
	filePath := fmt.Sprintf("%s/%s.log", ls.folder, date)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Panic(err)
	}

	defer file.Close()
	data := []byte(logFormat)

	fmt.Println(string(data))

	file.Write(data)

}

func (ls *Service) Fatal(message string) {
	ls.writeInFile(message, "Fatal")
}

func (ls *Service) Error(message string) {
	ls.writeInFile(message, "Error")
}

func (ls *Service) Warning(message string) {
	ls.writeInFile(message, "Warn")
}

func (ls *Service) Info(message string) {
	ls.writeInFile(message, "Info")
}

func (ls *Service) Debug(message string) {
	ls.writeInFile(message, "Debug")
}
