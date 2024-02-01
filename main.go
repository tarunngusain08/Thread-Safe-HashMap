package main

import (
	"fmt"
	"sync"
)

func testConcurrentMap(waitGroup *sync.WaitGroup) {
	concurrentMap := NewConcurrentMap()
	defer waitGroup.Done()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			concurrentMap.Set(fmt.Sprintf("key%d", i), i)
			fmt.Printf("Key = %d is set\n", i)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val, ok := concurrentMap.Get(fmt.Sprintf("key%d", i))
			if !ok {
				fmt.Println("Key is not present")
			} else {
				fmt.Printf("Key = %d has the value = %d\n", i, val)
			}
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			concurrentMap.Delete(fmt.Sprintf("key%d", i))
			fmt.Printf("Key = %d deleted \n", i)
		}(i)
	}
	wg.Wait()
}

func testGenericConcurrentMap(waitGroup *sync.WaitGroup) {
	genericConcurrentMap := NewGenericConcurrentMap[string, int]()
	defer waitGroup.Done()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			genericConcurrentMap.Set(fmt.Sprintf("key%d", i), i)
			fmt.Printf("Key = %d is set\n", i)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val, ok := genericConcurrentMap.Get(fmt.Sprintf("key%d", i))
			if !ok {
				fmt.Println("Key is not present")
			} else {
				fmt.Printf("Key = %d has the value = %d\n", i, val)
			}
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			genericConcurrentMap.Delete(fmt.Sprintf("key%d", i))
			fmt.Printf("Key = %d deleted \n", i)
		}(i)
	}
	wg.Wait()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go testConcurrentMap(&wg)
	wg.Add(1)
	go testGenericConcurrentMap(&wg)
	wg.Wait()
}
