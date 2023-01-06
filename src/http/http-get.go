package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	SimpleGet()
	fmt.Println("-----")
	CustomHttpGet()
}

func SimpleGet() {
	resp, err := http.Get("https://openlibrary.org/api/books?bibkeys=ISBN:0201558025,LCCN:93005405&format=json")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(data))
}

func CustomHttpGet() {
	req, err := http.NewRequest("GET", "https://openlibrary.org/api/books?bibkeys=ISBN:0201558025,LCCN:93005405&format=json", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Crawler")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)
}
