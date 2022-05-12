package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const port = 8900

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/input/fluentd", Capture)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8900", nil))
}

func Capture(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	base := "/home/app/.cache/"
	filename := randString(6)
	if err := os.WriteFile(base+filename, body, 0o755); err != nil {
		log.Fatal(err)
	}

	headers := r.Header
	if err := os.WriteFile(base+filename+"-headers", []byte(fmt.Sprint(headers)), 0o755); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Hello, there\n")
}
