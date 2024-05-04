package imageon

import "context"

func HandleRequest(ctx context.Context) (*string, error) {
	message := "Hello World!"
	return &message, nil
}
