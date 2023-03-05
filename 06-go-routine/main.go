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
	existChannelFromMain := make(chan bool)
	for i:=0;i<workerCount;i++{
		workerID := fmt.Sprintf("worker-%v",i)
		go worker3(workerID, existChannelFromMain)
	}
	time.Sleep(10* time.Second)
	for i:=0;i< workerCount;i++{
		existChannelFromMain <- true
	}
	close(existChannelFromMain)
	time.Sleep(2*time.Second)
	fmt.Println("Main channel routine is exited")


}

func worker3(workerID string, existChannelFromMain chan bool){
	i := 0
	for  {
		i++
		select {
		case <- existChannelFromMain :
				fmt.Println(fmt.Sprintf(" worker-%v has exited",workerID))
				return
		default:
			fmt.Println(fmt.Sprintf("woker= %v, couter=%v", workerID,i))
			time.Sleep(1*time.Second)	
		}
	}
}

func worker1(workerID string, responseChannel chan string) {
	//simulate request latency
	time.Sleep(1 * time.Second)
	//send data to channel
	responseChannel <- (workerID + "Response")
}
