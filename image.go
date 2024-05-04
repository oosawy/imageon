package main

import (
	"image"
	"os"

	"github.com/disintegration/imaging"
)

func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := imaging.Decode(f, imaging.AutoOrientation(true))
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

	err = imaging.Encode(f, img, imaging.JPEG, imaging.JPEGQuality(75))
	if err != nil {
		return err
	}

	return nil
}

func ResizeImage(img image.Image, width, height int) image.Image {
	dst := resizeAndCrop(img, width, height, imaging.Center, imaging.Lanczos)

	return dst
}

// copied form https://github.com/disintegration/imaging/blob/d40f48ce0f098c53ab1fcd6e0e402da682262da5/resize.go#L315-L335
// Copyright (c) 2012 Grigory Dryapak
// resizeAndCrop resizes the image to the smallest possible size that will cover the specified dimensions,
// crops the resized image to the specified dimensions using the given anchor point and returns
// the transformed image.
func resizeAndCrop(img image.Image, width, height int, anchor imaging.Anchor, filter imaging.ResampleFilter) *image.NRGBA {
	dstW, dstH := width, height

	srcBounds := img.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()
	srcAspectRatio := float64(srcW) / float64(srcH)
	dstAspectRatio := float64(dstW) / float64(dstH)

	var tmp *image.NRGBA
	if srcAspectRatio < dstAspectRatio {
		tmp = imaging.Resize(img, dstW, 0, filter)
	} else {
		tmp = imaging.Resize(img, 0, dstH, filter)
	}

	return imaging.CropAnchor(tmp, dstW, dstH, anchor)
}
