package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// NOTE: HTTP/1.0 の単純なGETクライアントの実装 (P.82)

func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
	log.Println("Content-Type:", resp.Header.Get("Content-Type"))
	log.Println("Date", resp.Header.Get("Date"))
}
