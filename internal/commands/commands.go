package commands

import (
	"log"
	"regexp"
)

var Commands []string = []string{
	"/tags",
	"/advice",
	"/avatar",
	"/info",
	"/tagadd",
}

// To invoke a function:
// if string(str[1]) != "/" && string(str[2]) != "" {
func NewCommand(str string) {

	if CountWords(str) == 2 {
		Commands = append(Commands, str)
	}
	log.Printf("Command %v was successfully added\n", str)
}

func CountWords(str string) int {
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	stringQuantity := re.FindAllString(str, -1)
	return len(stringQuantity)
}

func ListTags() string {
	result := "```"
	for _, val := range Commands {
		result += (val + "\n")
	}
	result += "```"
	return result
}
