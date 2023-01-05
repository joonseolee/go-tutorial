package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 절대경로
	// path, _ := os.Getwd()
	fi, err := os.Open("./sample.json")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fo, err := os.Create("./sample_copy.json")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if cnt == 0 {
			break
		}

		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

	// 파일이 크지않을경우 유틸사용
	bytes, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("bytes ->", bytes)

	err = ioutil.WriteFile("./writeFile.json", bytes, 0)
	if err != nil {
		panic(err)
	}
}
