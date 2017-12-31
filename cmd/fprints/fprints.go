package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {

	// Create a file and use bufio.NewWriter.
	f, _ := os.Create("cmd/fprints/file.txt")
	w := bufio.NewWriter(f)

	// Use Fprint to write things to the file.
	// ... No trailing newline is inserted.
	fmt.Fprint(w, "Hello", "\n")
	fmt.Fprint(w, 123, "\n")
	fmt.Fprint(w, "...\n")

	// Use Fprintf to write formatted data to the file.
	value1 := "cat"
	value2 := 900
	fmt.Fprintf(w, "%v %d...\n", value1, value2)

	fmt.Fprintln(w, "DONE...")

	// Done.
	w.Flush()

}
