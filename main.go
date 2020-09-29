package main

import "fmt"

func main() {
	var day int8
	var month int8
	fmt.Scanln(&day)
	fmt.Scanln(&month)
	switch {
	case day >= 21 && month == 3 || day < 21 && month == 4:
		fmt.Println("aries")
	case day >= 21 && month == 4 || day < 21 && month == 5:
		fmt.Println("tauro")
	case day >= 21 && month == 5 || day <= 21 && month == 6:
		fmt.Println("geminis")
	case day >= 22 && month == 6 || day <= 22 && month == 7:
		fmt.Println("cancer")
	case day >= 23 && month == 7 || day <= 22 && month == 8:
		fmt.Println("leo")
	case day >= 23 && month == 8 || day <= 22 && month == 9:
		fmt.Println("virgo")
	case day >= 23 && month == 9 || day <= 22 && month == 10:
		fmt.Println("libra")
	case day >= 23 && month == 10 || day <= 22 && month == 11:
		fmt.Println("escorpio")
	case day >= 23 && month == 11 || day <= 21 && month == 12:
		fmt.Println("sagitario")
	case day >= 22 && month == 12 || day <= 20 && month == 1:
		fmt.Println("capricornio")
	case day >= 21 && month == 1 || day <= 18 && month == 2:
		fmt.Println("acuario")
	case day >= 19 && month == 2 || day <= 20 && month == 3:
		fmt.Println("picis")

	}
}
