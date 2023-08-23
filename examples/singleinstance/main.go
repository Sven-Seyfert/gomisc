package main

import (
	"fmt"
	"time"

	"github.com/sven-seyfert/gomisc/singleinstance"
)

func main() {
	// CreateInstanceFile creates a instance indicator file for this instance of
	// the program. To ensure this current instance of the program is the only one,
	// the function checks for a existing instance indicator file. In case there
	// is already a instance indicator file, this program will be exited.
	singleinstance.CreateInstanceFile()

	// RemoveInstanceFile removes the previously created instance indicator file.
	// Always set this to ensure the next single instance of the program can be run.
	// If the instance indicator file still exists, you will not be able to run
	// your program because the exit procedure in CreateInstanceFile.
	defer singleinstance.RemoveInstanceFile()

	// Do something
	for i := 1; i < 6; i++ {
		fmt.Print(i, ", ")
	}

	time.Sleep(time.Second * 3)

	// Do something more
	for i := 6; i < 11; i++ {
		fmt.Print(i, ", ")
	}

	fmt.Println("Finished!")
}
