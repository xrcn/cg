//

package gcmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xrcn/cg/text/gstr"
)

// Scan prints `info` to stdout, reads and returns user input, which stops by '\n'.
func Scan(info ...interface{}) string {
	fmt.Print(info...)
	return readline()
}

// Scanf prints `info` to stdout with `format`, reads and returns user input, which stops by '\n'.
func Scanf(format string, info ...interface{}) string {
	fmt.Printf(format, info...)
	return readline()
}

func readline() string {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = gstr.Trim(s)
	return s
}
