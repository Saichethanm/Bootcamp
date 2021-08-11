package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/home/saichm/Saichethan/fromWritingFilesProgram.txt", d1, 0644)
	check(err)

	f, err := os.Create("/home/saichm/Saichethan/fromWritingFilesProgram2.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("\nHello world this is writing from writingFiles of gobyexample\n")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("This is using buffered writer in writingFiles of gobyexample\n")
	fmt.Printf("Wrote %d bytes\n", n4)

	w.Flush()

}
