package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	// ----- Example 1: Open a file and get its info -----
	f, err := os.Open("new.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	// Should always close the file after we open it
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Is the file folder:", fileInfo.IsDir())
	fmt.Println("File size (Bytes):", fileInfo.Size())
	fmt.Println("File permissions:", fileInfo.Mode())
	fmt.Println("File modified at:", fileInfo.ModTime())

	// ----- Example 2: Open a file and read it's contents -----
	fr, err1 := os.Open("new.txt")
	if err1 != nil {
		panic(err1)
	}

	defer fr.Close()

	// We need to store the data we read from the file into a buffer
	buf := make([]byte, 29) // Note: we can use the fileinfo to get the size

	d, errR := fr.Read(buf)
	if errR != nil {
		panic(errR)
	}

	// Read the contents at once
	// Note: If we don't convert to string
	// the ascii/utf-8 value of the chars is shown
	fmt.Println("data", d, string(buf))

	// Better yet, we can use a loop to go over each byte
	for i := 0; i < len(buf); i++ {
		fmt.Println("data", string(buf[i]))
	}

	// ----- Example 3: Read a file using ReadFile() -----
	data, err2 := os.ReadFile("new.txt")
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(string(data))
	// Why not always use this?
	// Ans: This method loads the complete data of the file
	// into memory, and for large files, that would lead to
	// just a single file read hogging a big chunk of memory

	// ----- Example 4: Read a directory -----
	frD, err3 := os.Open("..")
	if err3 != nil {
		panic(err3)
	}

	defer frD.Close()

	// here if we pass:
	// n >= 0: returns at most n directory records
	// n < 0: returns all directory records
	dirInfo, errRD := frD.ReadDir(-1) // returns a slice
	if errRD != nil {
		panic(errRD)
	}

	fmt.Println("Files in the current workdir:")
	for _, fi := range dirInfo {
		fmt.Println(fi.Name(), fi.IsDir())
	}

	// ----- Example 5: Create & Write a file -----
	frW, err4 := os.Create("example.txt")
	// Note: This under the hood uses the os.OpenFile("name", O_RDWR|O_CREATE|O_TRUNC, 0666)
	// We can also pass permissions when creating:
	// a file using: os.WriteFile(err := os.WriteFile("name", []byte, 0600))
	// a directory using: os.mkDir("name", 0644)
	if err4 != nil {
		panic(err4)
	}

	defer frW.Close()

	frW.WriteString("Hello, World!")
	frW.WriteString("Added new text.")     //  Appends the new text to the file
	frW.Write([]byte("Another new text.")) //  Appends the new text to the file
	// Note: This syntax converts string to a slice of bytes: []byte("any string")

	// Difference between: os.Write and os.WriteFile
	// WriteFile() is a method in itself which is self sufficient
	// Write() can only be called on an open file object

	// ----- Example 6: read and write to another file (used for big files) -----
	sourceFile, errS := os.Open("example.txt")
	if errS != nil {
		panic(errS)
	}

	defer sourceFile.Close()

	destFile, errD := os.Create("example_new.txt")
	if errD != nil {
		panic(errD)
	}

	defer destFile.Close()

	// Here we will use the bufio package (below are some reasons why).
	// The bufio package is used in Go to implement buffered I/O operations.
	// By wrapping underlying readers and writers
	// (like files or network connections), it groups many small, expensive
	// system calls into fewer, larger memory operations. This drastically reduces
	// latency and improves overall application performance

	// Unbuffered I/O (like reading or writing one byte at a time) requires the
	// operating system to process a separate system call for every single byte.
	// bufio maintains an in-memory buffer (default 4KB) to batch these operations,
	// meaning the program only interacts with the OS when the buffer is full or empty.

	reader := bufio.NewReader(sourceFile) // default buffer size is 4096 bytes (4KB)
	writer := bufio.NewWriter(destFile)

	for {
		readByte, errRead := reader.ReadByte()
		if errRead != nil {
			if errRead.Error() == "EOF" {
				break
			}
			panic(errRead)
		}

		errWrite := writer.WriteByte(readByte)
		if errWrite != nil {
			panic(errWrite)
		}
	}

	// We always do this as a good practice
	writer.Flush()

	fmt.Println("Data written to new file successfully.")

	// ----- Example 7: Deleting a file -----
	fC, errC := os.CreateTemp(".", "sample-*.txt")
	if errC != nil {
		panic(errC)
	}

	time.Sleep(time.Second * 2)

	errDel := os.Remove(fC.Name())
	if errDel != nil {
		panic(errDel)
	}

	// Note: There is also a createTemp() method to create temporary files
	// dir: The directory path where you want to store the file.
	// Passing an empty string "" tells Go to automatically use the
	// operating system's default temporary folder
	// (such as /tmp on Linux/macOS or AppData\Local\Temp on Windows)

	// pattern: The naming template for your file. Go appends a random
	// unique string to ensure no naming conflicts occur, making it
	// completely concurrency-safe.Example: "logs-*.txt" will generate
	// something like logs-3849102.txt.Example: "preview" will append
	// a random string right to the end, resulting in preview483140887
}
