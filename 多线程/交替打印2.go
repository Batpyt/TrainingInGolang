package main

//var wg sync.WaitGroup
//var flag = make(chan struct{})

//func main() {
//	var arr1 = []int{1, 2, 3, 4, 5, 6, 7}
//	var arr2 = []string{"A", "B", "C", "D", "E", "F", "G"}
//	wg.Add(2)
//	go func(str []int) {
//		defer wg.Done()
//		for _, v := range str {
//			fmt.Print(v)
//			flag <- struct{}{}
//			<-flag
//		}
//	}(arr1)
//	go func(str []string) {
//		defer wg.Done()
//		for _, v := range str {
//			<-flag
//			fmt.Print(v)
//			flag <- struct{}{}
//		}
//	}(arr2)
//	wg.Wait()
//}
