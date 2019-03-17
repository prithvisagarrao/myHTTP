package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

func main() {

	file, err := os.OpenFile("/root/go/src/github.com/prithvisagarrao/myHTTP/myHttp.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
		fmt.Println(err.Error())
	}

	defer file.Close()
	log.SetOutput(file)
	startTime := time.Now()
	log.Printf("myHTTP start -- %v", startTime.Format("2006-01-02 15:04:05" ))

	limit := flag.Int("limit", 10, "limit value for number of parallel requests")
	timeoutInt := flag.Int("timeout", 10, "timeout value for http requests")
	timeoutSec := time.Duration(*timeoutInt) * time.Second

	flag.Parse()

	client := http.Client{
		Timeout: timeoutSec,
	}

	urls := flag.Args()

	n := len(urls)

	if n == 0 {
		log.Printf("ERROR: No arguments passed")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, *limit)

	checkedURLs := checkURLs(urls)

	for i := 0; i < n; i++ {

		wg.Add(1)

		go makeRequest(checkedURLs[i], &wg, &semaphore, client)
	}

	wg.Wait()
	endTime := time.Now()
	log.Printf("myHTTP end -- %v", endTime.Format("2006-01-02 15:04:05" ))

}

func makeRequest(urlStr string, wg *sync.WaitGroup, semaphore *chan struct{}, client http.Client) {

	defer wg.Done()

	*semaphore <- struct{}{}

	defer func() {
		<-*semaphore
	}()

	resp, err := client.Get(urlStr)

	if err != nil {
		log.Printf("ERROR: %v\n", err.Error())

	} else {
		hash := md5.New()
		err := resp.Write(hash)
		if err != nil {
			log.Printf("ERROR: %v\n", err.Error())
			return
		}

		fmt.Printf("%v %x\n", urlStr, hash.Sum(nil))
	}
}

func checkURLs(urls []string) (retURLList []string) {
	for _, val := range urls {

		urlParsed, err := url.Parse(val)

		if err != nil {
			log.Printf("ERROR: %v\n", err.Error())
			break
		}

		if urlParsed.Scheme == ""{
			val = "http://" + val
		}

		retURLList = append(retURLList, val)
	}

	return
}
