package main

import (
	"flag"
	"net/http"

	"fmt"
	"net/url"

	"log"

	"github.com/astranet/astranet"
)

var name = flag.String("name", "anonym", "recepient to send messages to")
var msg = flag.String("msg", "abra-kadabra", "message")

func main() {
	flag.Parse()

	var astraNet = astranet.New().Client().WithEnv(*name, "chat")
	astraNet.Join("tcp4", "0.0.0.0:20000")
	var httpClient = &http.Client{
		Transport: &http.Transport{
			Dial: astraNet.HttpDial,
		},
	}
	var q = url.Values{}
	q.Set("msg", *msg)

	var req = url.URL{}
	req.Scheme = "http"
	req.Host = fmt.Sprintf("http-api")
	req.Path = "/send"
	req.RawQuery = q.Encode()

	var resp, reqFail = httpClient.Get(req.String())
	if reqFail != nil {
		log.Panicln(reqFail)
	}
	fmt.Println(resp.Status)
}
