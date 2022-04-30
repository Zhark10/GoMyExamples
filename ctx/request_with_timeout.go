package ctx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func CancelRequestWithTimeout() {
	parentCtx := context.Background()
	duration := 2000 * time.Millisecond
	ctx, cancel := context.WithTimeout(parentCtx, duration)
	defer cancel()

	doRequestWithTimeout(ctx, "https://yandex.ru/")
}

func doRequestWithTimeout(ctx context.Context, url string) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	req := request.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Response completed with status code %d\n", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("Request to long")
	}
}
