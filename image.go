package main

import (
	"bytes"
	"errors"
	"image"
	"io"

	"github.com/disintegration/imaging"
)

func ResizeImage(r io.Reader, width, height int) (io.Reader, error) {
	img, err := imaging.Decode(r, imaging.AutoOrientation(true))
	if err != nil {
		return nil, errors.Join(errors.New("Failed to decode image"), err)
	}

	img = processImage(img, width, height)

	w := new(bytes.Buffer)

	err = imaging.Encode(w, img, imaging.JPEG, imaging.JPEGQuality(75))
	if err != nil {
		return nil, errors.Join(errors.New("Failed to encode image"), err)
	}

	return w, nil

}

func processImage(img image.Image, width, height int) image.Image {
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
