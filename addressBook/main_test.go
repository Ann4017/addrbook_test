package main

import (
	"fmt"
	"testing"
)

var m = map[string]int{
	"ann":  27,
	"kim":  30,
	"pack": 34,
}

func Test(t *testing.T) {

	fmt.Println("address")
	for k, v := range m {
		fmt.Printf("name : %s, age : %d\n", k, v)
	}
}

func addAddress(name string, age int) {
	m[name] = age
}

func deleteAddress(name string) {
	delete(m, name)
}

func deleteAllAddress() {
	if m != nil {
		for k := range m {
			delete(m, k)
		}
	}
}
