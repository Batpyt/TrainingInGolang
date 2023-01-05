package main

//func main() {
//
//}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return pow(x, n)
	}
	return 1.0 / pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := pow(x, n/2)
	if n%2 == 0 {
		return y * y
	} else {
		return y * y * x
	}
}
