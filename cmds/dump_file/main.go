package main

import (
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ppacher/nbt"
)

var fileName = flag.String("file", "", "The path to the NBT file")
var compressed = flag.Bool("gzip", true, "Whether or not the file is GZIP compressed")

func main() {
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Missing -file parameter")
		return
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Failed to open file")
		return
	}

	var reader io.Reader = file

	if *compressed {
		reader, err = gzip.NewReader(file)
		if err != nil {
			fmt.Println("Failed to create gzip reader")
			return
		}
	}

	tag, err := nbt.ReadNamedTag(reader)
	if err != nil {
		fmt.Printf("Failed to parse file: %s\n", err.Error())
		return
	}

	content, err := json.Marshal(tag)
	if err != nil {
		fmt.Println("Failed to convert to json")
		return
	}

	fmt.Println(string(content))
}
