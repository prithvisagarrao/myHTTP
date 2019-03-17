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
	log.Printf("myHTTP start -- %v", time.Now())

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
		h := md5.New()
		err := resp.Write(h)
		if err != nil {
			log.Printf("ERROR: %v\n", err.Error())
			return
		}

		fmt.Printf("%v %x\n", urlStr, h.Sum(nil))
	}
}

func checkURLs(urls []string) (retURLList []string) {
	for _, v := range urls {

		urlParsed, err := url.Parse(v)

		if err != nil {
			log.Printf("ERROR: %v\n", err.Error())
			break
		}

		if urlParsed.Scheme == ""{
			v = "http://" + v
		}

		retURLList = append(retURLList, v)
	}

	return
}
