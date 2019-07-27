package main

import (
	"debug/elf"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./elfReader elf_file")
		os.Exit(1)
	}
	elfFile, err := elf.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	for _, section := range elfFile.Sections {
		fmt.Println(section)
	}

	fmt.Println("Machine Type:\t", elfFile.FileHeader.Machine)
	fmt.Println("Starting Memory Address:\t", elfFile.FileHeader.Entry)
}
