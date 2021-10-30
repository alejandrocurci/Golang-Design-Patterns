package barrier

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// it is used when we have to wait for more than one response from different goroutines before letting the program continue
// purpose -> put up a barrier so that nobody passes until we have all the results we need
// use cases -> application where one service needs to compose its response by merging the responses of another several microservices

// example -> an HTTP GET aggregator

const timeout = 5000 // milliseconds

type barrierResponse struct {
	Err  error
	Resp string
}

func barrier(endpoints ...string) {
	// number of requests to make
	requests := len(endpoints)

	// buffered channel to receive each goroutine response
	in := make(chan barrierResponse, requests)
	defer close(in)

	// slice of responses
	responses := make([]barrierResponse, requests)

	// make each request in a different goroutine and send the response to the channel
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	// receive each response and check whether there has been an error
	// if no error has occurred, print each response
	var hasError bool
	for i := 0; i < requests; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}
	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

// makeRequest performs the request and sends the result (error or successful resp) to the output channel
func makeRequest(out chan<- barrierResponse, url string) {
	res := barrierResponse{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeout) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}
