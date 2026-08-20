package main

import(
	"fmt"
)

func main7() {
	const baseRate=5.50
	const taxRate=0.12
	const distanceRate=2.0
	const fragileFee=0.2

	var a string
	fmt.Print("Введите имя: ")
	fmt.Scan(&a)
	var b float64
	fmt.Print("вес груза(кг): ")
	fmt.Scan(&b)
	var c int
	fmt.Print("Дистанция (км): ")
	fmt.Scan(&c)
	var d int
	result:= b+taxRate
	fmt.Print("Количество хрупких упаковок: ")
	fmt.Scan(&d)
	fmt.Println(a)
	fmt.Println(float64(b) * baseRate)
	fmt.Println(float64(c) * distanceRate)
	fmt.Println(float64(d) * fragileFee + 1)
	fmt.Println("Отчет о доставке: ")
	fmt.Println("Отправитель: ", a)
	fmt.Println("Итоговая стоимость: ", result)
}