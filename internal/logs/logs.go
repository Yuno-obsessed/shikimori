// Logs package is responsible for providing logs
// to keep bot working the right way and to keep track of messages
package logs

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"
)

var (
	ErrWrongCommand         = errors.New("Wrong command syntax")
	ErrInvalidUser          = errors.New("The user isn't in database yet")
	ErrUnableToSendMessage  = errors.New("Inability to send a message")
	ErrNotEnoughPermissions = errors.New("Not enough permission to perform an action")
	ErrFuncUnavailable      = errors.New("Function is not implemented yet")
)

type errLogs struct {
	Err   error     `json:"error"`
	Msg   string    `json:"msg"`
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

func LogErr(errortype error, msgs string, v ...any) {
	errf, err := os.OpenFile("errlogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	defer errf.Close()
	loggingErrors := errLogs{
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
