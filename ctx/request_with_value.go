package ctx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type TContextValue struct {
	key   string
	value interface{}
}

var contextValueExample = TContextValue{
	key:   "message",
	value: "Hi",
}

func CancelRequestWithValue() {
	parentCtx := context.Background()
	parentCtx = context.WithValue(parentCtx, contextValueExample.key, contextValueExample.value)
	ctx, cancel := context.WithCancel(parentCtx)

	go func() {
		if err := cancelMethodWithValue(ctx); err != nil {
			cancel()
		}
	}()

	doRequestWithValue(ctx, "https://ya.ru/")
}

func cancelMethodWithValue(ctx context.Context) error {
	const requestDeadline time.Duration = 5000
	time.Sleep(requestDeadline * time.Millisecond)
	return fmt.Errorf("some error")
}
func doRequestWithValue(ctx context.Context, url string) {
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
		fmt.Printf("Message %s\n", ctx.Value(contextValueExample.key))
	case <-ctx.Done():
		fmt.Println("Request to long")
	}
}
