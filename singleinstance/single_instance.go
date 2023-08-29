// Ensures only a single instance of the program runs at the same time.
//
// Check out a usage example at https://github.com/sven-seyfert/gomisc/blob/main/examples/singleinstance/main.go
package singleinstance

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const instanceFile = "instance_indicator.lock"

// CreateInstanceFile creates a instance indicator file for this instance of
// the program. To ensure this current instance of the program is the only one,
// the function checks for a existing instance indicator file. If there is an
// error, the error will be returned.
func CreateInstanceFile() error {
	if existsInstanceFile() {
		errMessage := fmt.Errorf("instance indicator file \"%s\" could not be created because it already exist", instanceFile) //nolint:goerr113,lll

		return errMessage
	}

	file, err := os.Create(instanceFile)
	if err != nil {
		return err
	}

	defer file.Close()

	// This 100 milliseconds delay ensures a robust file handling
	time.Sleep(time.Millisecond * 100) //nolint:gomnd

	return nil
}

func existsInstanceFile() bool {
	_, err := os.Stat(instanceFile)

	return !errors.Is(err, os.ErrNotExist)
}

// RemoveInstanceFile removes the previously created instance indicator file.
// Always set this to ensure the next single instance of the program can be run.
// If there is an error, the error will be returned.
func RemoveInstanceFile() error {
	err := os.Remove(instanceFile)
	if err != nil {
		return err
	}

	return nil
}
