package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UTC().UnixNano())
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		params := Params{5, 0.001, 100, 64, 8}
		for k := range r.Form {
			var err error
			switch k {
			case "cycles":
				params.cycles, err = strconv.Atoi(r.FormValue(k))
			case "delay":
				params.delay, err = strconv.Atoi(r.FormValue(k))
			case "nframes":
				params.nframes, err = strconv.Atoi(r.FormValue(k))
			case "res":
				params.res, err = strconv.ParseFloat(r.FormValue(k), 64)
			case "size":
				params.size, err = strconv.Atoi(r.FormValue(k))
			default:
				log.Printf("server: get unexpected param: %q=%s\n", k, r.FormValue(k))
			}
			if err != nil {
				fmt.Fprintf(w, "server: %q\n", err)
			}
		}
		lissajous(w, params)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
