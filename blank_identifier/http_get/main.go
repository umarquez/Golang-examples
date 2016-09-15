package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	/*
	  blank identifier "_" is used to throw away a unused value like the error that
	  could be returned if something goes wrong with the next functions calls.
	*/
	res, _ := http.Get("http://www.google.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
