package main

import (
	"io"
	"os"
	"strconv"

	"github.com/oosawy/imageon"
)

func main() {
	if len(os.Args) < 5 {
		panic("Not enough arguments")
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]
	width, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(os.Args[4])
	if err != nil {
		panic(err)
	}

	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := imageon.ResizeImage(f, width, height)
	if err != nil {
		panic(err)
	}

	f, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Copy(f, img)
	if err != nil {
		panic(err)
	}
}
