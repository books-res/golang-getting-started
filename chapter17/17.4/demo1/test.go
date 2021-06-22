package main

import "io/ioutil"

func main() {
	ioutil.WriteFile("demo.txt", []byte("ABCD"), 0644)
	ioutil.WriteFile("demo.txt", []byte("OPQRS"), 0644)
	ioutil.WriteFile("demo.txt", []byte("KLMN"), 0644)
}
