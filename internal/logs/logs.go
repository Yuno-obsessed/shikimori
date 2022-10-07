// Logs package is responsible for providing logs
// to keep bot working the right way and to keep track of messages
package logs

import (
	"log"
)

func LogProviding(message string) {
	log.Println(message)
}
