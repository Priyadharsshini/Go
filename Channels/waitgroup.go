package main

/* WAIT GROUP
Wait group is used to block a go routine until all the other go routines have completed execution.
Wait group is a strcut which has a counter which will keep track of the number of go routines spawned and finished.
When the counter reached 0, all the go routines have completed job.
Wait() method is used to block the current go routine and that will unblock only when the counter reaches 0
Add(int) method accepts an int and we should increment it depending on how many go routines we are spawing. That does not happen automatically
Done() is used to decrement and it does not accept any value. So decrement by 1 only. This also should be done by us.

Also need to pass the wg to the function, so that it can decrement. Need to pass the address of ot, since go is not pass by value.
*/

import (
	"fmt"
	"sync"
)

func read(wg *sync.WaitGroup, c chan int) {
	fmt.Println("M in the read")
	wg.Done()
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go read(&wg, c)
	}
	wg.Wait()
	fmt.Println("Done with main")

}
