package main

import "fmt"

func homework() {
	var schooling int
	schooling=10
	fmt.Println(schooling)
	schooling=11
	fmt.Println(schooling)

	var name string
	name="Vladislav"
	fmt.Println(name)
	name="Adlet"
	fmt.Println(name)

	steps:=0
	fmt.Println(steps)
	steps=2000
	fmt.Println(steps)
	fmt.Println("Хорошая работа! Вы уже на пути к своей ежедневной цели")

	largeNumber:=6000000
	fmt.Println(largeNumber)

	const breakTime=15
	fmt.Println(breakTime)
	breakTime=20
	fmt.Println(breakTime)
	//вылезла ошибка потому что константу нельзя изменить//
}