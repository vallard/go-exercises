package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// create an output file.
	fparts := strings.SplitAfter(url, "://")
	fn := fparts[1] + "." + fmt.Sprintf("%d", time.Now().Unix()) + ".out"
	fmt.Println("Creating: " + fn)
	in, err := os.Create(fn)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	nbytes, err := io.Copy(in, resp.Body)
	defer in.Close()
	in.Sync()
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
