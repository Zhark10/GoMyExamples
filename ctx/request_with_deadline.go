package ctx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func CancelRequestWithDeadline() {
	parentCtx := context.Background()
	duration := 2000 * time.Millisecond
	timeout := time.Now().Add(duration)
	ctx, cancel := context.WithDeadline(parentCtx, timeout)
	defer cancel()

	doRequestWithDeadline(ctx, "https://yandex.ru/")
}

func doRequestWithDeadline(ctx context.Context, url string) {
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
