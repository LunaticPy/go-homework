package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

var (
	timeout time.Duration
	url     string
)

func init() {
	flag.DurationVar(&timeout, "timeout", time.Second*10, "connection timeout")
	flag.StringVar(&url, "url", "google.com:80", "url")
}

func main() {
	flag.Parse()

	c := NewClient(url, timeout, ioutil.NopCloser(os.Stdin), os.Stdout)
	log.Printf("Connected to %s\n", url)
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := c.Close()
		log.Fatal(err)
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			if err := c.Send(); err != nil {
				log.Fatal(err)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			if err := c.Receive(); err != nil {
				log.Fatal(err)
				return
			}
		}
	}()

	wg.Wait()
}
