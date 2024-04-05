package main

// главная причина почему го лучше в бэке
// многопоточность
// gorutines takes stack place
import "fmt"

func treadFunc() {
	fmt.Println("in thread")
}

//func main() {
//	go treadFunc()
//	fmt.Println("main thread")
//
//	fmt.Scanln() // we need this to hold main thread, and let goruitine finish
//	// if main is terminated other threads "dies"
//	// all gr has shared memo
//}
