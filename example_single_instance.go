package main

import (
	"fmt"

	"github.com/sven-seyfert/gomisc/singleinstance"
)

func main() {
	// Create instance indicator file for this instance of the software.
	// In case a instance already exists (because the indicator file
	// perhaps already exists), this current new instance will be exited.
	singleinstance.CreateInstanceFile()

	// Remove the previously created instance file. Always set this
	// to ensure the next single instance of the software can be run.
	// If the instance indicator file still exists, you will not
	// be able to run your program (because the exit procedure in CreateInstanceFile()).
	defer singleinstance.RemoveInstanceFile()

	// Do something
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Finished!")
}
