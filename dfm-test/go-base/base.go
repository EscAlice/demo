package go_base

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// chan用于goroutine通信,select用于轮询与阻塞
func checkOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			once.Do(onceBody)
			done <- i
		}
		quit <- 1
	}()

	for i := 13; i < 33; i++ {
		go func(x int) {
			once.Do(onceBody)
			done <- x
		}(i)
	}

	for {
		select {
		case x := <-done:
			fmt.Println("x:", x)
		case <-quit:
			return
		case <-time.After(time.Second * 5):
			fmt.Println("timeout 5")
			return
		}
	}
}

// struct比较
type M struct {
	X string
	I int
	F float64
	//V map[string]string
	A []byte
	C chan int
}

func CreatePointer() bool {
	//return new(int)
	qc := make(chan int, 1)
	qc <- 1
	ac := make(chan int, 1)
	ac <- 1
	//var q = M{X: "as", I: 1, F: 0.7, V: map[string]string{"ss": "ss"}}
	var q = M{X: "as", I: 1, F: 0.7, A: []byte{2}, C: qc}
	//var a = M{X: "as", I: 1, F: 0.7, V: map[string]string{"ss": "ss"}}
	var a = M{X: "as", I: 1, F: 0.7, A: []byte{2}, C: ac}
	return reflect.DeepEqual(q, a)
}

func goodnight(s string) {
	fmt.Println("GoodNight, " + s)
}

func goodbye(s string) {
	fmt.Println("Goodbye, " + s)
}

//defer
func LazyDefer() {
	defer goodbye("1")
	defer goodnight("1")
	fmt.Println("Hello world.")
	return
	defer goodbye("2")
	defer goodnight("2")
}
