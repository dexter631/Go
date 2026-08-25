package main

import(
	"fmt"
)

func main9() {
	a:= 5 == 5 //true
	fmt.Println(a)
	b:= 10 != 3 //true
	fmt.Println(b)
	c:= 7 > 12 //false
	fmt.Println(c)
	d:= 15 < 20 //true
	fmt.Println(d)
	q:= 8 >= 8 //true
	fmt.Println(q)
	w:= 6 <= 4 //false
	fmt.Println(w)
	e:= (10 > 5) && (3 < 1) //false
	fmt.Println(e)
	r:= (10 > 5) || (3 < 1) //true
	fmt.Println(r)
	t:= !(5 == 5) //false
	fmt.Println(t)
	y:= !(7 < 3) //true
	fmt.Println(y)
	u:= true && false //false
	fmt.Println(u)
	i:= false || false //false
	fmt.Println(i)
	o := true || false //true
	fmt.Println(o)
	p := (4 + 6 == 10) && (9 > 2) //true
	fmt.Println(p)
	f:= (12 / 3 == 4) || (8 < 5) //true
	fmt.Println(f)

	age:= 18
	hasTicket:= true
	canEnter:= age >= 18 && hasTicket
	fmt.Println(canEnter)

	isLoggedIn:= true
	isAdmin:= true
	hasAccess:= isLoggedIn && isAdmin
	fmt.Println(hasAccess)
}