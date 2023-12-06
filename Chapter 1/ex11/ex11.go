package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := []string{"bing.com", "google.com", "bad.bing.com", "baidu.com", "gopl.io", "golang.org", "godoc.org", "rfc-editor.org/rfc/rfc2616", "bilibili.com"}
	for _, url := range urls {
		go fetch(url, ch)
	}
	for index, _ := range urls {
		fmt.Printf("%d: try to receive info from channel\n", index+1)
		fmt.Println(<-ch)
		fmt.Printf("%d: finish receiving info from channel\n\n", index+1)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%s: try to send info to channel\n", url)
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	fmt.Printf("%s: finish sending info\n", url)
}
