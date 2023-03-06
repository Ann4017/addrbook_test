package main

import (
	"fmt"
	"testing"
)

type address struct {
	name   string
	age    int
	number int
}

var m = map[int]address{
	1: {"ann", 29, 10 - 4519 - 6551},
	2: {"kim", 30, 10 - 4640 - 4432},
	3: {"pack", 34, 10 - 6877 - 8445},
}

func Test(t *testing.T) {
	fmt.Println(m)
	add := address{}
	add.addAddress(4, address{"sss", 24, 01 - 4821 - 2245})
	fmt.Println(m)
	del := address{}
	del.deleteAddress(2)
	fmt.Println(m)
	delAll := address{}
	delAll.deleteAllAddress()
	fmt.Println(m)
}

func (add *address) addAddress(key int, data address) {
	m[key] = data
}

func (del *address) deleteAddress(key int) {
	for k := range m {
		if key == k {
			delete(m, key)
		}
	}
}

func (delAll *address) deleteAllAddress() {
	if m != nil {
		for k := range m {
			delete(m, k)
		}
	}
}
