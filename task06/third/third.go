//  use channel closure as a termination signal

package third

import (
	"fmt"
	"time"
)

var StopChannel = make(chan struct{})

func ThirdProcess(stopChannel chan struct{}) {
	fmt.Println("Начало работы горутины 3.")
	for {
		select {
		case <-stopChannel:
			return
		default:
			fmt.Println("Горутина 3 продолжает работу.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
