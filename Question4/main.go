// 		M readers, N writers
// Solve for M = 8 and N = 2
// Solve for M = 8 and N = 8
// Solve for M = 8 and N = 16
// Solve for M = 2 and N = 8

// sharedBuffer=8 byte M=8 N=8
// we can read and write in each byte on parallel without locks (uncomment below)
// if the writer need to write 8bytes then only one writer (one writer multiple readers) can do it using Mutual Exclusion: sync.Mutex
// and the readers can use sync.RWMutex

// one writer multiple readers
package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 8

func main() {
	sharedBuffer := make([]byte, bufferSize)
	copy(sharedBuffer, []byte("Hello, World!"))

	var sharedCounter int
	var mu sync.Mutex
	var rwMutex sync.RWMutex

	readChannel := make(chan struct{})
	writeChannel := make(chan struct{})

	for i := 0; i < 8; i++ {
		go writer(i, &sharedCounter, sharedBuffer, writeChannel, &mu)
	}

	for i := 0; i < 8; i++ {
		go reader(i, &sharedCounter, sharedBuffer, readChannel, &rwMutex)
	}

	for {
		writeChannel <- struct{}{}
		time.Sleep(5 * time.Second)
		data := <-readChannel
		fmt.Printf("Read data: %s, sharedcounter: %d\n", data, sharedCounter)
		time.Sleep(5 * time.Second)
	}
}

func reader(id int, sharedCounter *int, sharedBuffer []byte, readChannel chan struct{}, rwMutex *sync.RWMutex) {
	for {

		rwMutex.RLock()
		data := sharedBuffer
		counterValue := *sharedCounter
		fmt.Printf("Reader %d read: %s balance: %d\n", id, data, counterValue)
		time.Sleep(1 * time.Second)
		rwMutex.RUnlock()
	}
}

func writer(id int, sharedCounter *int, sharedBuffer []byte, writeChannel chan struct{}, mu *sync.Mutex) {
	for {
		mu.Lock()
		copy(sharedBuffer, []byte("Hello, World!"))
		*sharedCounter++
		fmt.Printf("Writer %d wrote: %s balance: %d \n", id, sharedBuffer, *sharedCounter)
		time.Sleep(2 * time.Second)
		mu.Unlock()
	}
}

// parallel reading and writing
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// const bufferSize = 8

// func main() {
// 	// Shared buffer (byte slice)
// 	sharedBuffer := make([]byte, bufferSize)

// 	// Create an RWMutex
// 	var rwMutex sync.RWMutex
// 	var mu sync.Mutex

// 	// Channels for signaling readers and writers
// 	readChannel := make(chan struct{})
// 	writeChannel := make(chan struct{})

// 	// Start writer goroutines
// 	for i := 0; i < 8; i++ {
// 		go writer(i, sharedBuffer, writeChannel, &mu)
// 	}

// 	// Start reader goroutines
// 	for i := 0; i < 8; i++ {
// 		go reader(i, sharedBuffer, readChannel, &rwMutex)
// 	}

// 	// Simulate continuous writing and reading
// 	for {
// 		writeChannel <- struct{}{}
// 		// time.Sleep(500 * time.Millisecond) //delay
// 		<-readChannel
// 		fmt.Printf("Read buffer: %v\n", sharedBuffer)
// 		// time.Sleep(500 * time.Millisecond) //delay
// 	}
// }

// func reader(id int, sharedBuffer []byte, readChannel chan struct{}, rwMutex *sync.RWMutex) {
// 	for {

// 		// rwMutex.RLock() // Acquire a read lock
// 		// Simulate some reading logic
// 		data := make([]byte, len(sharedBuffer))
// 		copy(data, sharedBuffer)
// 		fmt.Printf("Reader %d read: %v\n", id, data)
// 		// time.Sleep(100 * time.Millisecond)
// 		// rwMutex.RUnlock()
// 	}
// }

// func writer(id int, sharedBuffer []byte, writeChannel chan struct{}, rwMutex *sync.Mutex) {
// 	for {

// 		// rwMutex.Lock()
// 		// incrementing one byte
// 		sharedBuffer[id] = byte(sharedBuffer[id] + 1)
// 		fmt.Printf("Writer %d wrote: %v\n", id, sharedBuffer)
// 		// time.Sleep(200 * time.Millisecond)
// 		// rwMutex.Unlock()
// 	}
// }
