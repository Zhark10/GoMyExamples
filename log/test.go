package log

import (
	"fmt"
	"time"
)

func SpeedTest(operation func()) {
	start := time.Now()
	operation()
	duration := time.Since(start)
	fmt.Println("\nTIME OF COMPLETION --", duration.Milliseconds(), "ms")
}
