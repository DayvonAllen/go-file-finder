package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// get command line args
	args := os.Args[1:]

	// if no command line arguments, run this
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	// ioutil package provides functions for working with files
	// uses command line arg for directory
	files, err := ioutil.ReadDir(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}
	var total int

	for _, file := range  files {
		if file.Size() == 0 {
			// the extra one is for the newline character
			total += len(file.Name()) + 1
		}
	}
	// for optimization use make function to have an initial capacity
	names := make([]byte, 0 , len(files) * 256)

	for _, file := range files {
		// determines whether a file is empty or not
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	// name of the created file, byte slice of data to write, permissions of the file
	err = ioutil.WriteFile("out.txt", names, 0644); if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Total size of files written: %d bytes\n", total)

}
