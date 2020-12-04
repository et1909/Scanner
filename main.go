package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/et1909/Scanner/port"
)

func main() {
	fmt.Println("Enter file path")       // Enter file name prompt
	var file_name string                 // variable file_name stores the user input
	fmt.Scanln(&file_name)               // input taken from the user and passed to the variable file_name and stored there temporarily.
	file_open, err := os.Open(file_name) // os.Open(variable_name) allows to open a file from the local machine. A path has to be given if the file is kept somewhere else.
	if err != nil {                      // a error check if the file doesn't open please log the error.
		log.Fatal(err) // error logging
	}
	defer file_open.Close() // file_open closed when the work is done.

	file_scan := bufio.NewScanner(file_open) // bufio or 'buffer input/output it takes the io.Reader or io.Writer object and wraps around it to create another object giving a buffer space for thr input/output'
	for file_scan.Scan() {                   // bufio.NewScanner(variable_name) scans/read lines from the file that was given by the user and for each line that is read by NewScanner(), .Scan() scans each line and if finds a input that it cannot recognises, it will send it to error and the error will be output except EOF.
		fmt.Println(file_scan.Text())
		portlist := []int{80, 443}
		for _, value := range portlist {

			port.GetPort(file_scan.Text(), value) // checking if the port 80 is open.
		}
	}

}
