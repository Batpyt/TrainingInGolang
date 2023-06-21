package main

//func main() {
//
//	person := &Person{26}
//
//	//前两种defer都属于有入参
//	defer fmt.Println("defer1", person.age)
//
//	defer func(p *Person) {
//		fmt.Println("defer2", p.age)
//	}(person)
//
//	//这种defer无入参，所以捕获外部变量
//	defer func() {
//		fmt.Println("defer3", person.age)
//	}()
//
//	person = &Person{27}
//	/*
//	打印结果
//	defer3 27
//	defer2 26
//	defer1 26
//	*/
//}

type Person struct {
	age int
}
