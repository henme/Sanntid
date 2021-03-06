// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var global_var i

func someGoroutine() {
    for x := 0; x < 100000; x++ {
        i++
    } 
}

func someGoroutine2() {
    for x := 0; x < 100000; x++ {
        i--
    } 
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!
    go someGoroutine()                      // This spawns someGoroutine() as a goroutine
    go someGoroutine2() 
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(100*time.Millisecond)
    Println(i)
}