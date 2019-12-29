// Enhance comma so that it deals correctly with floating-point numbers and an optional sign
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	n := len(s)
	period_index := strings.Index(s, ".")
	if period_index == -1 {
		period_index = n
	}
	for i := 0; i < n; i++ {
		if i < period_index {
			if (i != 0) && (n != 0) && ((period_index-i)%3) == 0 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(byte(s[i]))
	}
	return buf.String()
}
