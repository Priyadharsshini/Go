package main

/* MUTEX - mutual exclusion
To avoid race condition.
each go routine have their own independent stack and hence they dont share data.
But it may happen like every go routine has to access and change a common variable. So when one go routine is accessing and before it completes, another go routines access and hence different value.
multiple goroutines are trying to manipulate data at the same memory location resulting in unexpected results

To avoid this, to make only one go routine to execute that particular step at a time. we lock that statement. so at a time only one can execute.
*/

import (
	"fmt"
	"sync"
)

var i int

func check(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() //acquire lock
	i = i + 1
	m.Unlock() //release lock
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go check(&wg, &m)
	}
	wg.Wait()
	fmt.Println(i)
}
