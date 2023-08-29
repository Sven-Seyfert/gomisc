package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sven-seyfert/gomisc/singleinstance"
)

func main() {
	// CreateInstanceFile creates a instance indicator file for this instance of
	// the program. To ensure this current instance of the program is the only one,
	// the function checks for a existing instance indicator file.
	if err := singleinstance.CreateInstanceFile(); err != nil {
		// Do your error handling
		fmt.Println(err)

		// For this case the current program instance will be exits.
		os.Exit(1)
	}

	// RemoveInstanceFile removes the previously created instance indicator file.
	// Always set this to ensure the next single instance of the program can be run.
	defer func() {
		if err := singleinstance.RemoveInstanceFile(); err != nil {
			// Do your error handling
			fmt.Println(err)
		}
	}()

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
