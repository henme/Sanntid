// Go 1.2
// go run helloworld_go.go

package main

import (
    "fmt"
    "runtime"
    //"time"
    "sync"
)

var i int

func main() {
    //message := make(chan int) //Messages between processes
    var wg sync.WaitGroup       //Synchronize a group of processes
    runtime.GOMAXPROCS(1)    // I guess this is a hint to what GOMAXPROCS does...
    // limit the ammount of threads to one, simple but slow implementation
    wg.Add(2)                                        // Try doing the exercise both with and without it!
    go func() {
        defer wg.Done()
        for x := 0; x < 100000; x++ {
            i++
        } 
        //message <- 1
    }()                 // This spawns someGoroutine() as a goroutine
    go func() {
        defer wg.Done()
        for x := 0; x < 100000; x++ {
            i--
        } 
        //message <- 2
    }() 
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    //time.Sleep(100*time.Millisecond)
    wg.Wait()
    fmt.Printf("%v\n",i)
}