package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {
	start := time.Now().UTC().UnixMilli()
	ch := make(chan string)
	urls := []string {"https://www.rfc-editor.org/rfc/rfc2616", "https://www.rfc-editor.org/rfc/rfc2616"}
	for index, url := range urls {
		go fetch(strconv.Itoa(index+1)+".txt", url, ch)
	}
	for index, _ := range urls {
		fmt.Printf("%d: try to receive info from channel\n", index+1)
		fmt.Println(<-ch)
		fmt.Printf("%d: finish receiving info from channel\n\n", index+1)
	}
	fmt.Printf("%dms elapsed\n", time.Now().UTC().UnixMilli()-start)
	if check("1.txt", "2.txt") {
		fmt.Println("The content of both documents is identical.")
	} else {
		fmt.Println("The content of both documents is not identical.")
	}
}

func fetch(filename string, url string, ch chan<- string) {
	start := time.Now().UTC().UnixMilli()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	buffer := bytes.Buffer{}
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	err = ioutil.WriteFile(filename, buffer.Bytes(), 0666)
	if err != nil {
		ch <- fmt.Sprintf("while writing %s into %s: %v", url, filename, err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	milliSecs := time.Now().UTC().UnixMilli()-start
	fmt.Printf("%s: try to send info to channel\n", url)
	ch <- fmt.Sprintf("%dms %7d %s", milliSecs, nbytes, url)
	fmt.Printf("%s: finish sending info\n", url)
}

func check(filename1 string, filename2 string) bool {
	f1, err := ioutil.ReadFile(filename1)
	if err != nil {
		fmt.Print("reading file %s: %v\n", filename1, err)
		return false
	}
	f2, err := ioutil.ReadFile(filename2)
	if err != nil {
		fmt.Printf("reading file %s: %v\n", filename2, err)
		return false
	}
	for i := 0; i < len(f1); i++ {
		if f1[i] != f2[i] {
			return false
		}
	}
	return true
}
