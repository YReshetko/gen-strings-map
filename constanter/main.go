package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	baseDir = flag.String("baseDir", "./", "the base dir where stored go packages by patter; must be set; example ./test")
	pattern = flag.String("pattern", "", "the directory/packages pattern to generate a constant map; must be set; example: /schema/*/*/schema")
	output  = flag.String("output", "constanter.go", "output file name; default ./constanter.go")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of constanter:\n")
	fmt.Fprintf(os.Stderr, "\tconstanter -baseDir [directory] -pattern [pattern] -output [file_name]\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	// Init flags
	log.SetFlags(0)
	log.SetPrefix("constanter: ")
	flag.Usage = Usage
	flag.Parse()
	if *baseDir == "" || *pattern == ""{
		flag.Usage()
		os.Exit(2)
	}
	generator := NewGenerator(*baseDir, *pattern, *output)
	generator.Process()
	generator.WriteFile()
}
