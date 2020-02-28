package nicrebooter

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

// worker is a function that you want to repeating when panic
// intervalSecond is a time in second that you want to repeat per second when panic
func Run(worker interface{}, intervalSecond time.Duration) {
	c := make(chan int, 100)
	v := reflect.ValueOf(worker)
	if v.Kind() != reflect.Func {
		return
	}
	c <- 1
	for {
		<-c
		runner(v, c)
		time.Sleep(intervalSecond * time.Second)
	}
}

func runner(v reflect.Value, c chan int) {
	fmt.Println("............Starting............")
	defer func() {
		if exception := recover(); exception != nil {
			log.Printf("Panic: %v\n", exception)
			c <- 1
		}
	}()
	v.Call(nil)
}
