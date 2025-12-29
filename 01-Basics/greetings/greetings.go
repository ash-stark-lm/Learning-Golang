/*
# This code is a library
# Other programs can import it
# It cannot be run directly
If you try:
go run greetings.go
❌ Go will refuse —package command-line-arguments is not a main package.
*/
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//%v means: “print the value in its default format.”
    return message
}