package shouter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Shout(s string) string {
	return strings.ToUpper(s)
}

func ShoutReallyLoud(s string) string {
	return Shout(s) + "!"
}

func ReadAndShout(file *os.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(Shout(scanner.Text()))
	}

	return scanner.Err()
}
