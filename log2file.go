package main

import (
	"log"
	"os"
)

func main() {
	if !FileExists("logfile") {
		CreateFile("logfile")
	}
	f, _ := os.Create("logfile")

	// attempt #1
	// log.SetOutput(io.MultiWriter(os.Stderr, f))
	// log.Println("hello, logfile")
	//
	// // attempt #2
	// log.SetOutput(io.Writer(f))
	// log.Println("hello, logfile")

	// attempt #3
	log.SetOutput(f)
	log.Println("hello, logfile")
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}
