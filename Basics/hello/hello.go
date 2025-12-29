// If a Go program has package main, Go treats it as a standalone executable program.
package main

/* A program will not compile if there are missing imports or if there are unnecessary ones.
This strict requirement prevents references to unused packages from accumulating as programs evolve. */
import (
	"fmt"
	"log"

	"example.com/greetings"
	"example.com/handleError"
	"rsc.io/quote"
)

// Execution of the program always begins at main()
func main() {

	// Configure logger
	log.SetPrefix("app: ")
	log.SetFlags(0)

	fmt.Println(quote.Hello())

	message := greetings.Hello("Ashish")
	fmt.Println(message)

	/*
		In Go:
		- NO try/catch
		- NO exceptions
		- Functions return (value, error)
	*/

	result, err := handleError.Hello("") // passing empty string → error
	if err != nil {
		log.Fatal(err) // logs error and exits program
	}

	fmt.Println(result)
	fmt.Println("Hello, 世界")
}
