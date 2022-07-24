package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test1(t *testing.T) {

	resp, err := http.Get("http://localhost:8080/temp.txt")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	if sb != "notessss" {
		t.Fail()
	}

}
