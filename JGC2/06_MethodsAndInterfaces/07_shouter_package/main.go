package shouter

import (
	"bufio"
	"io"
	"strings"
)

func Shout(s string) string {
	return strings.ToUpper(s)
}

func ShoutReallyLoud(s string) string {
	return Shout(s) + "!"
}

/*
Here we have replaced the concrete type "file *os.File" with io.Reader interface
$ godoc io Reader
type Reader interface {
	Read(p []byte) (n int, err error)
}
*/
func ReadAndShout(r io.Reader) (string, error) {
	gather := ""
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		gather += Shout(scanner.Text()) + "\n"
	}

	return gather, scanner.Err()
}
