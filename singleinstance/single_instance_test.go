package singleinstance_test

import (
	"errors"
	"os"
	"testing"

	"github.com/sven-seyfert/gomisc/singleinstance" //nolint:depguard
)

func TestCreateInstanceFile(t *testing.T) {
	tests := []struct {
		name               string
		expectError        bool
		expectFileCreation bool
	}{
		{name: "file created", expectError: false, expectFileCreation: true},
	}

	const instanceFile = "instance_indicator.lock"

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := singleinstance.CreateInstanceFile()

			defer func() {
				os.Remove(instanceFile)
			}()

			if !test.expectError && err != nil {
				t.Fatalf("unexpected error for testcase \"%s\"", test.name)
			}

			_, err = os.Stat(instanceFile)
			fileExists := !errors.Is(err, os.ErrNotExist)

			if test.expectFileCreation && !fileExists {
				t.Fatalf("instance indicator file should be created for testcase \"%s\"", test.name)
			}
		})
	}
}
