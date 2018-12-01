package main

import (
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ppacher/nbt"
)

var fileName = flag.String("file", "", "The path to the NBT file")

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

	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("Failed to create gzip reader")
		return
	}

	tag, err := nbt.ReadNamedTag(reader)
	if err != nil {
		fmt.Printf("Failed to parse file: %s\n", err.Error())
		return
	}

	fmt.Println("File parsed successfully \n\n")

	content, err := json.Marshal(tag)
	if err != nil {
		fmt.Println("Failed to convert to json")
		return
	}

	fmt.Println(content)
}
