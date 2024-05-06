package imageon

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/google/uuid"
)

type RawRequest struct {
	Method string
	Path   string
	Body   string
}

type RawResponse struct {
	StatusCode int
	Body       string
}

func HandleRequest(ctx context.Context, request RawRequest) (*RawResponse, error) {
	switch request.Path {
	case "/":
		return &RawResponse{
			StatusCode: 200,
			Body:       "ok",
		}, nil

	case "/new":
		if request.Method != "POST" {
			return &RawResponse{
				StatusCode: 405,
				Body:       "method not allowed",
			}, nil
		}
		return handleNew(ctx, request)

	default:
		return &RawResponse{
			StatusCode: 404,
			Body:       "not found",
		}, nil
	}
}

func handleNew(ctx context.Context, request RawRequest) (*RawResponse, error) {
	options, err := url.ParseQuery(request.Body)
	if err != nil {
		return &RawResponse{
			StatusCode: 400,
			Body:       "bad request",
		}, nil
	}

	name := fmt.Sprintf("%s.jpg", uuid.New())

	size := options.Get("size")          // 1280 or 720,1080
	quality := options.Get("quality")    // 75
	format := options.Get("format")      // jpeg or png
	fit := options.Get("fit")            // cover or none
	cropArea := options.Get("crop_area") // 0,0,1280,1280

	bucketName := os.Getenv("UNPROCESSED_BUCKET")

	url, err := PresignUploadUrl(bucketName, name, map[string]string{
		"size":      size,
		"quality":   quality,
		"format":    format,
		"fit":       fit,
		"crop_area": cropArea,
	})

	if err != nil {
		return &RawResponse{
			StatusCode: 500,
			Body:       "internal server error",
		}, nil
	}

	return &RawResponse{
		StatusCode: 200,
		Body:       url,
	}, nil
}
