package main

import (
	"sync/atomic"
	"time"
)

func main() {
	//bitmap := make(map[string]int)
	//bitmap["a"] = 1
	//bitmap["b"] = 2
	//go func() {
	//	for {
	//		bitmap["a"] = 2
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		fmt.Println(bitmap["b"])
	//
	//	}
	//}()
	//time.Sleep(time.Second)
	//fmt.Println(bitmap["a"])

}

func T() int32 {
	var count int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&count, 1)
		}()
	}
	time.Sleep(time.Second)
	return count
}