package main

import (
	"fmt"
	"time"
)

func main() {
	workerCount := 10

	// #send response from channel to main line routine
	//1. use channel to receive data from go routine
	// receive data from go routine back to main routine

	//step 1 make channel by make(chan type-DatasendinChannel)
	responseChannel := make(chan string)
	//step2 make worker 10 and return response each worker id
	for i := 0; i < workerCount; i++ {
		workerID := fmt.Sprintf("worker-%v", i)
		go worker1(workerID, responseChannel)
	}

	//step3 read response receive from channel worker id
	for i := 0; i < workerCount; i++ {
		readResponse := <-responseChannel
		fmt.Println(readResponse)
	}

	//step4 close channel when not use
	close(responseChannel)
	fmt.Println("All response from channel returned")


	// 2# send data from main routine to go channel all


}

func worker1(workerID string, responseChannel chan string) {
	//simulate request latency
	time.Sleep(1 * time.Second)
	//send data to channel
	responseChannel <- (workerID + "Response")
}
