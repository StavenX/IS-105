package main

import (
	"fmt"
	"os"
)

// Function that goes through the properties of a file, and prints them out to
// the user. This includes filesize (in bytes, kilobytes, megabytes and gigabytes).
// Also gives information about file being a directory or not, etc.
func info() {
	 args := os.Args
	 if len(args) == 0 {
	 	fmt.Println("Missing argument (File)")
	 	fmt.Println("Usage: go run fileinfo.go myfile.txt")
	 	return
	 }

	 filename := args[1]

	 file, err := os.Open(filename)
	 if err != nil {
	 	fmt.Println(err)
	 	return
	 }

	 filestat, err := file.Stat()
	 if err != nil {
	 	fmt.Println(err)
	 	return
	 }

	 // Filesize information
	 bytes := float64(filestat.Size())
	 fmt.Printf("Information about file %s:\n", filestat.Name())
	 fmt.Printf("Size: %.0f bytes\n", bytes)
	 fmt.Printf("Size: %f kilobytes\n", bytes / 1024)
	 fmt.Printf("Size: %f megabytes\n", bytes / (1024 * 1024))
	 fmt.Printf("Size: %f gigabytes\n", bytes / (1024 * 1024 * 1024))

	 // Check if directory or not
	 if filestat.Mode().IsDir() {
	 	fmt.Println(" Is a directory")
	 } else {
	 	fmt.Println(" Is not a directory")
	 }

	 // Check if regular file or not
	 if filestat.Mode().IsRegular() {
	 	fmt.Println(" Is a regular file")
	 } else {
	 	fmt.Println(" Is not a regular file")
	 }

	 // Check if file has UNIX permission bits
	 fmt.Printf(" Has UNIX permission bits: %s\n", filestat.Mode().Perm())

	 // Check if append only
	 if filestat.Mode() & os.ModeAppend != 0 {
	 	fmt.Println(" Is append only")
	 } else {
	 	fmt.Println(" Is not append only")
	 }

	 // Check if device file or not
	 if filestat.Mode() & os.ModeDevice != 0 {
	 	fmt.Println(" Is a device file")
	 } else {
	 	fmt.Println(" Is not a device file")
	 }

	 // Check if UNIX character device
	 //if filestat.Mode() & os.ModeCharDevice != 0 {
	 //	fmt.Println(" Is a UNIX character device")
	 //} else {
	 //	fmt.Println(" Is not a UNIX character device")
	 //}

	 // Check if symbolic link
	 if filestat.Mode() & os.ModeSymlink != 0 {
	 	fmt.Println(" Is a symbolic link")
	 } else {
	 	fmt.Println(" Is not a symbolic link")
	 }
}

// Info function is run here
func main() {
	info()
}