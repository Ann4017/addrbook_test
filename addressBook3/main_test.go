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

func Test__input(t *testing.T) {

	// array 구조체 정의
	arrpt_address := []address{}

	// for문으로 array 구조체 안에 반복적으로 데이터 입력
	for i := 0; i <= 10; i++ {
		pt_address := address{}
		pt_address.age = 10
		pt_address.number = 123
		pt_address.name = "123"

		// 반복 데이터를 append 해서 array 구조체 저장
		arrpt_address = append(arrpt_address, pt_address)
	}

	// array 구조체 출력
	fmt.Println(arrpt_address)

	// array 구조체를 for range 후 랜덤 출력
	// for range 문제
	// 순차적으로 출력하지 않음.
	for _, pt_address := range arrpt_address {
		fmt.Println(pt_address.name, pt_address.age, pt_address.number)
	}
}

// 함수 작성하는 기본 기조.
func new_address(name string, age, number int) (res string) {

	return res
}
