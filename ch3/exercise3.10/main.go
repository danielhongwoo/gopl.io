// Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	var buf bytes.Buffer
	for i, v := range s {
		if (i != 0) && (n != 0) && ((n-i)%3) == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(byte(v))
	}
	return buf.String()
}
