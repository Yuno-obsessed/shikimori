// Logs package is responsible for providing logs
// to keep bot working the right way and to keep track of messages
package logs

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type BotError string

var (
	ErrWrongCommand          BotError = "Wrong command syntax"
	ErrInvalidUser           BotError = "The user isn't in database yet"
	ErrUnableToSendMessage   BotError = "Inability to send a message"
	ErrNotEnoughPermissions  BotError = "Not enough permission to perform an action"
	ErrFuncUnavailable       BotError = "Function is not implemented yet"
	ErrSendingImage          BotError = "Image can not be send"
	ErrStatusUpdate          BotError = "Status can not be updated"
	ErrSessionOpening        BotError = "Session can not be opened"
	ErrYetAnotherStupidError BotError = "How the heck did that break"
	ErrTooMuchFlags          BotError = "Too much flags specified to the command"
	ErrInvalidCommandUsage   BotError = "The command wasn't used in a proper way"
	ErrPublishingCommands    BotError = "Can not publish commands"
)

type ErrLogs struct {
	Err   BotError  `json:"error"`
	Msg   string    `json:"msg,omitempty"`
	Ltime time.Time `json:"time"`
}

type MsgLogs struct {
	Msg   string    `json:"msg"`
	Ltime time.Time `json:"time"`
}

func Log(msgs string, v ...any) {
	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	loggingMessages := MsgLogs{
		Ltime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC),
		Msg:   msgs,
	}

	byteArray, err := json.Marshal(loggingMessages)
	if err != nil {
		log.Println(err)
	}

	n, err := f.Write(byteArray)
	if err != nil {
		log.Println(n, err)
	}
	if n, err = f.WriteString("\n"); err != nil {
		log.Println(n, err)
	}
}

func LogErr(errortype BotError, msgs string, v ...any) {
	errf, err := os.OpenFile("errlogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	defer errf.Close()
	loggingErrors := ErrLogs{
		Err:   errortype,
		Msg:   msgs,
		Ltime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC),
	}
	byteArray, err := json.Marshal(loggingErrors)
	if err != nil {
		log.Println(err)
	}
	n, err := errf.Write(byteArray)
	if err != nil {
		log.Println(n, err)
	}
	if n, err = errf.WriteString("\n"); err != nil {
		log.Println(n, err)
	}
}
