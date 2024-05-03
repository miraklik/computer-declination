package main

import "fmt"

func Mass(a []int) []int {
	var del []int
	for i := 1; i <= a[0]; i++ {
		if a[0]%i == 0 {
			iss := true
			for _, num := range a[1:] {
				if num%i != 0 {
					iss = false
					break
				}
			}
			if iss {
				del = append(del, i)
			}
		}
	}
	return del
}

func main() {
	numbers := []int{42, 12, 18}
	fmt.Println(Mass(numbers))
}
