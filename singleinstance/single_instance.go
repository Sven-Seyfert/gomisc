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
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// This second ensures a robust file handling
	time.Sleep(1 * time.Second)
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
	if err != nil {
		panic(err)
	}
}
