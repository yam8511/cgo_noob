package main

/*
#include <stdio.h>  // <--- puts
#include <stdlib.h> // <--- free
#include "hello.h"

char *zuolar = "magic";
*/
import "C"
import (
	"mini"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(time.Second)
		// call C in goroutine
		C.hello()
		wg.Done()
	}()

	// call C *char from Go
	C.SayHello(C.zuolar)
	// Go string to C *char
	mini.PrintCString("Go String to C Char")
	wg.Wait()
}
