package context

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

// example waits for 2 seconds and then prints the message
// you can cancel the process by writing anything to the console
func example() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	mySleepAndTalk(ctx, 2*time.Second, "Hello context!")
}

// mySleepAndTalk prints a message after a specified duration or exits with error if it receives the cancellation signal
func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
