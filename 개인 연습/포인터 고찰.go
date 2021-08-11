package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("%s<%s>\n", u.name, u.email)
}

func main() {

	var kim user
	fmt.Println(kim)       // { }
	fmt.Println(kim.name)  // ""
	fmt.Println(kim.email) // ""

	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify() // 실제 동작 : (*lisa).notify()

	fmt.Println(bill)  // {Bill bill@email.com}
	fmt.Println(lisa)  // &{Lisa lisa@email.com}
	fmt.Println(*lisa) // {Lisa lisa@email.com}

	fmt.Println("--------------------------------------")

	var a int
	var p *int
	p = &a
	*p = 20
	fmt.Println(a, &a, p) // a = 20,  &a와 p는 같은 값이 출력된다.
	p2 := &a

	*p2 = 30
	fmt.Println(a, &a, p, p2) // a = 30, &a와 p와 p2는 같은 주소값이 출력된다.
}
