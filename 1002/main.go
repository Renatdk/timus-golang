package main

import "fmt"
import "io/ioutil"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := ioutil.ReadFile("input")
	check(err)
	fmt.Print(string(dat))

}
