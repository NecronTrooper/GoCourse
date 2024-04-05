package main

import (
	"context"
	"fmt"
	"time"
)

func 2main() {
	ctx := context.Background()
	//ctxVal := context.WithValue(ctx,"deadline", 5*time.Second)
	//ctxVal := context.WithValue(ctxVal , "id", 123)

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()


	helper(ctxTimeout, 5)
}

func helper(ctx context.Context, req int) {
	val := ctx.Value("deadline")
	if val != nil {
		fmt.Println(val)
	}
	id := ctx.Value("id")
	if id != nil {
		fmt.Println(id)
	}
	time.Sleep(1 * time.Second)

	inputCh := make(chan int)
	outputCh := make(chan int)

	go execute(ctx, inputCh, outputCh)

	inputCh<- req


	select {
		case <-ctx.Done():
			fmt.Println("Context done")
			return

		default:
			fmt.Println("Everything is fine")
	}


	func execute(context.Context() ,input chan int, output chan<- int){
		for req :=  range inputCh{
			fmt.Println("process")

		}

	}

}
