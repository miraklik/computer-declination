package main

import "fmt"

func Spen(a int) string {
	if a == 1 {
		return string(a) + " компьютер"
	}
	if a > 1 && a < 5 {
		return string(a) + "компьютера"
	}
	if a > 4 && a < 41 {
		return string(a) + "компьютеров"
	}
	if a >= 41 && a < 42 {
		return string(a) + "компьютер"
	}
	return fmt.Sprintln(a, "компьютеров")
}

func main() {
	fmt.Println(Spen(1024))
}
