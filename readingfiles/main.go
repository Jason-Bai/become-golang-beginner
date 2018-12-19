package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readingFileUsingAbsoluteFilePath() {
	data, err := ioutil.ReadFile("/Users/jason-bai/go/src/github.com/Jason-Bai/golang_codes/readingfiles/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func readingFileUsingCommandLineFlag() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func readingAFileInSmallChunks() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}

func readingAFileLineByLine() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("1. reading entire file in memory using absolute file path")
	readingFileUsingAbsoluteFilePath()
	fmt.Println("2. reading entire file in memory using command line flag")
	// readingFileUsingCommandLineFlag()
	fmt.Println("3. reading a file in small chunks")
	// readingAFileInSmallChunks()
	fmt.Println("4. reading a file line by line")
	readingAFileLineByLine()
}
