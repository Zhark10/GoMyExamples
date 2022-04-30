package ctx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func CancelRequest() {
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)

	go func() {
		if err := cancelMethod(ctx); err != nil {
			cancel()
		}
	}()

	doRequest(ctx, "https://ya.ru/")
}

func cancelMethod(ctx context.Context) error {
	const requestDeadline time.Duration = 5000
	time.Sleep(requestDeadline * time.Millisecond)
	return fmt.Errorf("some error")
}
func doRequest(ctx context.Context, url string) {
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
