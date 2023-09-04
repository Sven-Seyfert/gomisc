package clipboard_test

import (
	"testing"

	"github.com/sven-seyfert/gomisc/clipboard" //nolint:depguard
)

func TestWriteAndRead(t *testing.T) {
	tests := []struct {
		name         string
		expectError  bool
		expectedText string
	}{
		{name: "happy path 1", expectError: false, expectedText: "this is a lower case text"},
		{name: "happy path 2", expectError: false, expectedText: "THIS IS A UPPER CASE TEXT"},
		{name: "happy path 3", expectError: false, expectedText: "This is a multi line text.\n- second line\n- third line"},
		{name: "happy path 4", expectError: false, expectedText: "901003d73e3efd4d950160c07dd5e1418df75394d4319825c99ecca1ce7fac892e4adaf838d4a12dff771dfc"}, //nolint:lll
		{name: "happy path 5", expectError: false, expectedText: `special character like " in the text`},
		{name: "happy path 6", expectError: false, expectedText: `or like ' in the text`},
		{name: "happy path 7", expectError: false, expectedText: `or ""`},
		{name: "happy path 8", expectError: false, expectedText: `or ''`},
		{name: "happy path 9", expectError: false, expectedText: `or even a pipe | ...`},
		{name: "happy path 10", expectError: false, expectedText: `or !"#$%&()*+,-./:;<=>?@[\]^_{|}~`},
		{name: "happy path 11", expectError: false, expectedText: " "},
		{name: "sad path empty string", expectError: true, expectedText: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// arrange

			// act
			err := clipboard.Write(test.expectedText)

			// assert
			if !test.expectError && err != nil {
				t.Fatalf(`unexpected write error for testcase "%s"`, test.name)
			}

			if test.expectError && err == nil {
				t.Fatalf(`a error is expected for testcase "%s"`, test.name)
			}

			// act
			content, err := clipboard.Read()

			// assert
			if !test.expectError && err != nil {
				t.Fatalf(`unexpected read error for testcase "%s"`, test.name)
			}

			if test.expectedText != content {
				message := `written text was "%s" but read text is "%s" for testcase "%s"`
				t.Fatalf(message, test.expectedText, content, test.name)
			}
		})
	}
}
