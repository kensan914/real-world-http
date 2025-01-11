package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// NOTE: HTTP/1.0 のクエリパラメータを付与したGETクライアントの実装 (P.85)

func main() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
	log.Println("Content-Type:", resp.Header.Get("Content-Type"))
	log.Println("Date", resp.Header.Get("Date"))
}
