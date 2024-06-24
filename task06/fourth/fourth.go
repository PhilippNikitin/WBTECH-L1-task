// use of context to manage goroutine lifecycle

package fourth

import (
	"context"
	"fmt"
	"time"
)

var (
	Ctx    context.Context
	Cancel context.CancelFunc
)

func init() {
	Ctx, Cancel = context.WithCancel(context.Background())
}

func FourthProcess(ctx context.Context) {
	fmt.Println("Начало работы горутины 4.")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Горутина 4 продолжает работать")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
