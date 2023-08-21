package singleinstance

import (
	"os"
	"time"
)

const instanceFile = "instance_indicator"

func CreateInstanceFile() {
	if existsInstanceFile() {
		os.Exit(0)
	}

	file, err := os.Create(instanceFile)
	check(err)

	defer file.Close()

	// This 400 milliseconds delay ensures a robust file handling
	time.Sleep(time.Millisecond * 400)
}

func existsInstanceFile() bool {
	fileInfo, err := os.Stat(instanceFile)
	if os.IsNotExist(err) {
		return false
	}

	return !fileInfo.IsDir()
}

func RemoveInstanceFile() {
	err := os.Remove(instanceFile)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
