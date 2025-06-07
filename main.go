package main

import (
	"flag"
	"fmt"
	"log/slog"
	"time"
)

var (
	size  = flag.Int("size", 16, "size[MB] of the heap memory")
	delay = flag.Int("delay", 10, "delay[seconds] before triggering the panic")
)

func main() {
	flag.Parse()
	if *size <= 0 {
		slog.Error("size must be greater than 0")
		return
	}

	l := *size * 1024 * 1024 / 8
	slog.Info("Allocating memory", "size", l)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = i
	}

	<-time.After(time.Duration(*delay) * time.Second)
	function10()
	// slog.Info("first", "arr[0]", arr[0])
	// slog.Info("last", fmt.Sprintf("arr[%d]", l-1), arr[l-1])
}

func function1() {
	panic(fmt.Errorf("make a shit"))
}

func function2() {
	function1()
}

func function3() {
	function2()
}

func function4() {
	function3()
}

func function5() {
	function4()
}

func function6() {
	function5()
}

func function7() {
	function6()
}

func function8() {
	function7()
}

func function9() {
	function8()
}

func function10() {
	function9()
}
