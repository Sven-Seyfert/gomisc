package main

import (
	"fmt"

	"github.com/sven-seyfert/gomisc/clipboard"
)

func main() {
	// Write your text to the clipboard.
	if err := clipboard.Write("This is a text\nthat will be written to the clipboard."); err != nil {
		// Do your error handling
		fmt.Println(err)
	}

	// Read text from the clipboard.
	content, err := clipboard.Read()
	if err != nil {
		// Do your error handling
		fmt.Println(err)
	}

	fmt.Printf(`Clipboard content is: "%s"`, content)
}
