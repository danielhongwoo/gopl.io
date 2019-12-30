// Write a function CountingWriter with the signature below that,
// given an io.Writer, returns a new Writer that wraps the original,
// and a pointer to an int64 variable that at any moment contains
// the number of bytes written to the new Writer.
package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	w     io.Writer
	count int64
}

func (c *ByteCounter) Write(p []byte) (count int, err error) {
	count, err = c.w.Write(p)
	c.count = int64(count)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	new_writer := ByteCounter{w, 0}
	return &new_writer, &new_writer.count
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "Hello!\n")
	fmt.Println(*count)

	fmt.Fprint(writer, "Nice to see you\n")
	fmt.Println(*count)
}
