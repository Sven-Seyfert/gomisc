package singleinstance_test

import (
	"errors"
	"os"
	"testing"

	"github.com/sven-seyfert/gomisc/singleinstance" //nolint:depguard
)

var fileHandle *os.File //nolint:gochecknoglobals

func TestCreateInstanceFile(t *testing.T) {
	const instanceFile = "instance_indicator.lock"

	tests := []struct {
		name                string
		expectError         bool
		expectFileIsCreated bool
	}{
		{name: "happy path file is created", expectError: false, expectFileIsCreated: true},
		{name: "sad path file already exists", expectError: true, expectFileIsCreated: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// arrange
			if test.expectError {
				// This locks the file because there is no .Close() until
				// the defer function will be executed.
				fileHandle = createFile(t, test.name, instanceFile)
			}

			// act
			err := singleinstance.CreateInstanceFile()

			defer func() {
				fileHandle.Close()
				fileHandle = nil

				os.Remove(instanceFile)
			}()

			// assert
			if !test.expectError && err != nil {
				t.Fatalf(`unexpected error for testcase "%s"`, test.name)
			}

			if test.expectError && err == nil {
				t.Fatalf(`a error is expected for testcase "%s"`, test.name)
			}

			_, err = os.Stat(instanceFile)
			fileExists := !errors.Is(err, os.ErrNotExist)

			if test.expectFileIsCreated && !fileExists {
				t.Fatalf(`instance indicator file should be created for testcase "%s"`, test.name)
			}
		})
	}
}

func createFile(t *testing.T, testcase string, filename string) *os.File {
	t.Helper()

	file, err := os.Create(filename)
	if err != nil {
		t.Logf(`arrange failed for testcase "%s"`, testcase)
	}

	return file
}

func TestRemoveInstanceFile(t *testing.T) {
	const instanceFile = "instance_indicator.lock"

	tests := []struct {
		name                string
		expectError         bool
		expectFileIsRemoved bool
	}{
		{name: "happy path file is removed", expectError: false, expectFileIsRemoved: true},
		{name: "sad path file is not removed", expectError: true, expectFileIsRemoved: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// arrange
			if test.expectError {
				// This locks the file because there is no .Close() until
				// the defer function will be executed.
				fileHandle = createFile(t, test.name, instanceFile)
			} else {
				file := createFile(t, test.name, instanceFile)
				file.Close()
			}

			// act
			err := singleinstance.RemoveInstanceFile()

			defer func() {
				fileHandle.Close()
				fileHandle = nil

				os.Remove(instanceFile)
			}()

			// assert
			if !test.expectError && err != nil {
				t.Fatalf(`unexpected error for testcase "%s"`, test.name)
			}

			if test.expectError && err == nil {
				t.Fatalf(`a error is expected for testcase "%s"`, test.name)
			}

			_, err = os.Stat(instanceFile)
			fileExists := !errors.Is(err, os.ErrNotExist)

			if test.expectFileIsRemoved && fileExists {
				t.Fatalf(`instance indicator file should be removed for testcase "%s"`, test.name)
			}
		})
	}
}
