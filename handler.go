package imageon

import (
	"context"
)

type RawRequest struct {
	Path string
	Body string
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

	default:
		return &RawResponse{
			StatusCode: 404,
			Body:       "not found",
		}, nil
	}
}
