package server

import (
	"log"
	"os"
)

// logger type - should remain unchanged
type CustomLogger struct {
	infoLogger *log.Logger
	warnLogger *log.Logger
	errLogger  *log.Logger
}

func (l *CustomLogger) info(msg string) {
	l.infoLogger.Println(msg)
}

func (l *CustomLogger) warn(msg string) {
	l.warnLogger.Println(msg)
}

func (l *CustomLogger) error(msg string) {
	l.errLogger.Println(msg)
}

// logger implementation - can be editted to fit the needs
var SimpleLogger = CustomLogger{
	infoLogger: log.New(os.Stdout, "INFO: ", log.Ltime),
	warnLogger: log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
	errLogger:  log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
}
