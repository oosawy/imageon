package main

import (
	"image"
	"image/jpeg"
	"os"

	"golang.org/x/image/draw"
)

func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func SaveImage(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = jpeg.Encode(f, img, nil)
	if err != nil {
		return err
	}

	return nil
}

func ResizeImage(img image.Image, width, height int) image.Image {
	// do not resize if the image is smaller than the target size
	bounds := img.Bounds()
	if bounds.Dx() < width && bounds.Dy() < height {
		return img
	}

	// calculate the new size
	var w, h int
	if bounds.Dx() > bounds.Dy() {
		w = width
		h = bounds.Dy() * width / bounds.Dx()
	} else {
		w = bounds.Dx() * height / bounds.Dy()
		h = height
	}

	// resize the image
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.BiLinear.Scale(dst, dst.Bounds(), img, bounds, draw.Over, nil)

	return dst
}
