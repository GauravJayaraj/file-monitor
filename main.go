package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Cannot read monitor from stdin pipe\nexpected usage: ./file-monitor [filename]")
	}

	fname := os.Args[1]
	f, ferr := os.Open(fname)
	if ferr != nil {
		log.Fatal("error opening file: ", ferr)
	}
	defer f.Close()

	freader := bufio.NewReader(f)

	buf := make([]byte, 1)
	for {
		n, err := freader.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal("error reading file: ", err)
		} else if err == io.EOF {
			// Don't exit as we are monitoring
		} else {
			fmt.Printf("%v", string(buf[:n]))
		}

		// some other process has deleted the file being monitored
		_, err = os.Stat(os.Args[1])
		if os.IsNotExist(err) {
			fmt.Println("\nFile deleted by someone else in the interim...")
			break
		}

	}
}
