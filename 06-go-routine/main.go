package main

import (
	"fmt"
	"time"
)

const NumberWorker = 10

func main() {

	//1 worker 10 num work with routine and send response to channelpipe
	//then read data from channel pipe in main routine
	responsechannelForWorkpipe := make(chan string)
	for i := 0; i < NumberWorker; i++ {
		workerId := fmt.Sprintf("worker ID=%v \n", i+1)

		go workerCall(workerId, responsechannelForWorkpipe)
	}
	for i := 0; i < NumberWorker; i++ {
		responseReturntoMain := <-responsechannelForWorkpipe
		fmt.Println(responseReturntoMain)
	}

	fmt.Println("=========")
	//2 send data from main to channel
	//if receive signal then exit work
	for i := 0; i < NumberWorker; i++ {
		exitSignal := fmt.Sprintf("worker exist ID=%v", i+1)
		mainRoutineWork(exitSignal, responsechannelForWorkpipe)
	}

}
func mainRoutineWork(signalToexist string, channelpipe chan string) {

}

func workerCall(workerId string, responsechannelForWorkpipe chan string) {
	time.Sleep(1 + time.Second)
	responsechannelForWorkpipe <- workerId + "response"

}
