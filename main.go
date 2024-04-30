package main

import (
	"os"
	"strconv"
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

	img, err := LoadImage(inputPath)
	if err != nil {
		panic(err)
	}

	img = ResizeImage(img, width, height)

	err = SaveImage(outputPath, img)
	if err != nil {
		panic(err)
	}
}
