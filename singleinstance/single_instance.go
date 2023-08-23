package singleinstance

import (
	"os"
	"time"
)

const instanceFile = "instance_indicator.lock"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// CreateInstanceFile creates a instance indicator file for this instance of
// the program. To ensure this current instance of the program is the only one,
// the function checks for a existing instance indicator file. In case there
// is already a instance indicator file, this program will be exited.
func CreateInstanceFile() {
	if existsInstanceFile() {
		os.Exit(0)
	}

	file, err := os.Create(instanceFile)
	check(err)

	defer file.Close()

	// This 150 milliseconds delay ensures a robust file handling
	time.Sleep(time.Millisecond * 150)
}

func existsInstanceFile() bool {
	fileInfo, err := os.Stat(instanceFile)
	if os.IsNotExist(err) {
		return false
	}

	return !fileInfo.IsDir()
}

// RemoveInstanceFile removes the previously created instance indicator file.
// Always set this to ensure the next single instance of the program can be run.
// If the instance indicator file still exists, you will not be able to run
// your program because the exit procedure in CreateInstanceFile.
func RemoveInstanceFile() {
	err := os.Remove(instanceFile)
	check(err)
}
