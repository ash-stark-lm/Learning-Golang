package handleError

import (
	"errors"
	"fmt"
)

// Hello takes a string called name and returns a string and an error
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("Please provide a name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//%v means: “print the value in its default format.”
	return message, nil
}
