// Logs package is responsible for providing logs
// to keep bot working the right way and to keep track of messages
package logger

import (
	"errors"
	"go.uber.org/zap"
	"log"
)

var (
	ErrWrongCommand          = errors.New("wrong command syntax")
	ErrUserNotExist          = errors.New("there is no such user")
	ErrUnableToSendMessage   = errors.New("inability to send a message")
	ErrNotEnoughPermissions  = errors.New("not enough permission to perform an action")
	ErrFuncUnavailable       = errors.New("function is not implemented yet")
	ErrSendingImage          = errors.New("image can not be send")
	ErrStatusUpdate          = errors.New("status can not be updated")
	ErrSessionOpening        = errors.New("session can not be opened")
	ErrYetAnotherStupidError = errors.New("how the heck did that break")
	ErrTooMuchFlags          = errors.New("too much flags specified to the command")
	ErrInvalidCommandUsage   = errors.New("the command wasn't used in a proper way")
	ErrPublishingCommands    = errors.New("can not publish commands")
	ErrUnmarshalingJSON      = errors.New("can not unmarshal Json")
	ErrReachingAPI           = errors.New("can not get a response from API")
	ErrBuildingQuery         = errors.New("error occurred while building query")
)

type Logger struct {
	*zap.Logger
}

func loggerConfigInit() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevel(),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"../../logs/info.log"},
		ErrorOutputPaths: []string{"../../logs/error.log"},
	}
}

func NewLogger() Logger {
	logger, err := loggerConfigInit().Build()
	if err != nil {
		log.Fatalf("failed to build logger")
	}
	return Logger{logger}
}
