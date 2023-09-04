// Read or write text data to or from the clipboard.
//
// Check out a usage example at https://github.com/sven-seyfert/gomisc/blob/main/examples/clipboard/main.go
package clipboard

// Read returns the current text data of the clipboard. In case a file was
// copied to the clipboard, the file path will be returned.
// If there is an error, the error will be returned.
func Read() (string, error) {
	return read()
}

// Write sets the given text data to the clipboard. If there is an error,
// the error will be returned.
func Write(text string) error {
	return write(text)
}
