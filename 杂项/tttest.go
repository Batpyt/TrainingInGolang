package main

import "fmt"

func main() {
	//total, sum := 0, 0
	//for i := 1; i <= 10; i++ {
	//	sum += i
	//	go func() {
	//		total += i
	//	}()
	//}
	//fmt.Printf("total: %v, sum: %v", total, sum)

	person := &Person{26}
	defer fmt.Println(person.age)

	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	defer func() {
		fmt.Println(person.age)
	}()

}

type Person struct {
	age int
}
