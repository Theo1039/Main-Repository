package main

import (
	"fmt"

	"github.com/theo1039/pointerPractice/birthDay"
	"github.com/theo1039/pointerPractice/learn"
	"github.com/theo1039/pointerPractice/print"
)

func main() {
	f := 56
	s := 4555
	learn.Add(&f)
	fmt.Println(f)
	learn.Add(&s)
	fmt.Println(s)

	print.Name()

	user := birthDay.Birth{
		Name: "Monday",
		Age:  34,
	}
	birthDay.AddOne(&user)
	fmt.Println(user.Name, "is", user.Age, "year old")
}
