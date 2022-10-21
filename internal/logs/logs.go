// Logs package is responsible for providing logs
// to keep bot working the right way and to keep track of messages
package logs

import (
	"errors"
	"log"
	"os"

	"github.com/yuno-obsessed/shikimori/internal/settings"
)

var (
	ErrWrongCommand         = errors.New("Wrong command syntax")
	ErrInvalidUser          = errors.New("The user isn't in database yet")
	ErrUnableToSendMessage  = errors.New("Inability to send a message")
	ErrNotEnoughPermissions = errors.New("Not enough permission to perform an action")
)

type logs struct {
}

func CreateLogFile(PathToLogFile string) *os.File {
	if _, err := os.Stat(PathToLogFile); !os.IsNotExist(err) {
		log.Println(err)
		return nil
	}
	logFile, err := os.Create(PathToLogFile)
	if err != nil {
		log.Println(err)
	}
	return logFile
}

func CreateLogFiles() {
	CreateLogFile(settings.PathToLogFile)
	CreateLogFile(settings.PathToErrLogFile)
}

const (
	logNone = iota
	logInfo
	logWarning
	logError
	logVerbose
	logDebug
)

type myFileLogger struct {
	logger   *log.Logger
	logFile  *os.File
	logLevel int
}

func newFileLogger() *myFileLogger {
	return &myFileLogger{
		logger:   nil,
		logFile:  nil,
		logLevel: logNone,
	}
}

func (myLogger *myFileLogger) startLog(level int, file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	myLogger.logger = log.New(f, "", 0)
	myLogger.logLevel = level
	myLogger.logFile = f
	return nil
}

func (myLogger *myFileLogger) stopLog() error {
	if myLogger.logFile != nil {
		return myLogger.logFile.Close()
	}
	return nil
}

// You can add a log of auxiliary functions here to make the log more easier
func (myLogger *myFileLogger) log(level int, msg string) error {
	if myLogger.logger == nil {
		return errors.New("myFileLogger is not initialized correctly")
	}
	if level >= myLogger.logLevel {
		myLogger.logger.Print(msg) // maybe you want to include the loglevel here, modify it as you want
	}
	return nil
}

// func main() {
//
//     logger := newFileLogger()
//     if err := logger.startLog(logError, "myLogFile.log"); err != nil {
//         panic(err.Error())
//     }
//
//     defer func() {
//         logger.stopLog()
//     }()
//
//     logger.log(logInfo, "Info level log msg\n") // this will be ignored
//     logger.log(logError, "Error: error message\n") // this should included in the log file
//
// }
