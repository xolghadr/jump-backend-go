// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

func addOneLocked(number *int) {

	defer wg.Done()
	mu.Lock()
	*number += 1
	time.Sleep(3 * time.Millisecond)
	mu.Unlock()
}
func addOneNoLock(number *int) {
	defer wg.Done()
	*number += 1
	time.Sleep(3 * time.Millisecond)
}

func main() {
	raceVariable := 1
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addOneLocked(&raceVariable)
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addOneNoLock(&raceVariable)
	}

	wg.Wait()
	fmt.Println(raceVariable)
}
