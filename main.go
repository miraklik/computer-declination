package main

import "fmt"

func Spen(a int) string {
	var word string
	switch a % 10 {
	case 1:
		return fmt.Sprintln(a, "компьютер")
	case 2, 3, 4:
		return fmt.Sprintln(a, "компьютера")
	default:
		return fmt.Sprintln(a, "компьютеров")
	}

	if a/10%10 == 1 {
		word = "компьютер"
	}

	return fmt.Sprintln(a, word)

}

func main() {
	fmt.Println(Spen(2301))
}
